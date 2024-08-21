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
)

type LoanManager struct {
	db *sql.DB

	errorTracker ErrorTracker
}

func (m *LoanManager) GetCurrentLoanOfConsumer(ctx context.Context, consumerID uuid.UUID) (*Loan, error) {
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
