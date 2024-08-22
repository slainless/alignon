package platform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/query"
)

var (
	ErrConsumerNotBorrowingAnyNow = errors.New("consumer not borrowing any loan right now")
	ErrTransactionRecordNotFound  = errors.New("transaction record not found")
	ErrInvalidTenor               = errors.New("invalid tenor")
	ErrInvalidProduct             = errors.New("invalid product")
)

type LoanManager struct {
	db *sql.DB

	consumer *ConsumerManager
	catalog  *CatalogManager

	errorTracker ErrorTracker
}

func NewLoanManager(db *sql.DB, consumer *ConsumerManager, catalog *CatalogManager, tracker ErrorTracker) *LoanManager {
	return &LoanManager{
		db:       db,
		consumer: consumer,
		catalog:  catalog,

		errorTracker: tracker,
	}
}

func (m *LoanManager) GetCurrentLoan(ctx context.Context, consumerID uuid.UUID) (*Loan, error) {
	var loan Loan
	err := query.GetCurrentLoanOfConsumer(ctx, m.db, consumerID, &loan)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrConsumerNotBorrowingAnyNow
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return &loan, nil
}

func (m *LoanManager) GetTransactionOfCurrentLoan(ctx context.Context, consumerID uuid.UUID, transactionID string) (*TransactionRecord, error) {
	var transaction TransactionRecord
	err := query.GetTransactionOfCurrentLoan(ctx, m.db, consumerID, transactionID, &transaction)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrTransactionRecordNotFound
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return &transaction, nil
}

func (m *LoanManager) GetLoans(ctx context.Context, consumerID uuid.UUID) ([]Loan, error) {
	loans := make([]Loan, 0)
	err := query.GetLoansOfConsumer(ctx, m.db, consumerID, loans)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return loans, nil
}

func (m *LoanManager) CreateTransactionRecords(loan *Loan, products []Product) error {
	loan.TransactionRecords = make([]TransactionRecord, 0, len(products))
	for product := range products {
		loan.TransactionRecords = append(loan.TransactionRecords, TransactionRecord{
			AssetName: products[product].Name,
			Amount:    products[product].Price,
			CatalogID: products[product].ID,
			LoanID:    loan.ID,
		})
	}

	return nil
}

func (m *LoanManager) Apply(ctx context.Context, consumer *Consumer, tenor int16, catalogs []uuid.UUID) (*Loan, error) {
	limit, err := m.consumer.GetLimit(ctx, consumer.ID)
	if err != nil {
		return nil, err
	}

	var limitAmount int64
	var installmentLength int16
	switch tenor {
	case 1:
		limitAmount = limit.Tenor1
		installmentLength = 3
	case 2:
		limitAmount = limit.Tenor2
		installmentLength = 6
	case 3:
		limitAmount = limit.Tenor3
		installmentLength = 12
	case 4:
		limitAmount = limit.Tenor4
		installmentLength = 24
	default:
		return nil, ErrInvalidTenor
	}

	items, err := m.catalog.GetItems(ctx, catalogs)
	if err != nil {
		return nil, err
	}

	total := int64(0)
	for item := range items {
		total += items[item].Price
	}

	if total > limitAmount {
		return nil, &LoanLimitExceededError{
			Limit: limitAmount,
			Tenor: tenor,
			Month: installmentLength,

			Got:      total,
			Products: items,
		}
	}

	loan := &Loan{
		ConsumerID:        consumer.ID,
		Amount:            total,
		ConsumerLimit:     limitAmount,
		ConsumerSalary:    consumer.Salary,
		Tenor:             tenor,
		InstallmentLength: installmentLength,
		LoanedAt:          nil,
	}

	err = m.CreateTransactionRecords(loan, items)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	err = query.ApplyLoan(ctx, m.db, loan)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return loan, nil
}

type LoanLimitExceededError struct {
	Limit int64
	Tenor int16
	Month int16

	Got      int64
	Products []Product

	str string
}

func (e *LoanLimitExceededError) Error() string {
	if e.str != "" {
		return e.str
	}

	items := make([]string, 0, len(e.Products))
	for i := range e.Products {
		items = append(items, fmt.Sprintf("%s (%d)", e.Products[i].Name, e.Products[i].Price))
	}
	e.str = fmt.Sprintf("Products total (%d) exceeded the limit of loan (%d) for tenor %d (%d-months):\n%s", e.Got, e.Limit, e.Tenor, e.Month, strings.Join(items, "\n"))
	return e.str
}
