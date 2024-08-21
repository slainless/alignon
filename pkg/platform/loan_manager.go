package platform

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/query"
)

var (
	ErrConsumerNotBorrowingAnyNow = errors.New("consumer not borrowing any loan right now")
	ErrTransactionRecordNotFound  = errors.New("transaction record not found")
)

type LoanManager struct {
	db *sql.DB

	errorTracker ErrorTracker
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
