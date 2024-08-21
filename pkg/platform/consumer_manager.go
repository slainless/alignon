package platform

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/slainless/my-alignon/pkg/internal/query"
)

var (
	ErrConsumerNotFound = errors.New("consumer not found")
)

type ConsumerManager struct {
	db *sql.DB

	authManager  *AuthManager
	errorTracker ErrorTracker
}

func NewConsumerManager(db *sql.DB, auth *AuthManager, tracker ErrorTracker) *ConsumerManager {
	return &ConsumerManager{
		authManager: auth,
		db:          db,

		errorTracker: tracker,
	}
}

func (m *ConsumerManager) GetByEmail(ctx context.Context, email string) (*Consumer, error) {
	consumer, err := query.GetConsumerByEmail(ctx, m.db, email)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrConsumerNotFound
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return &Consumer{
		Consumers: *consumer,
	}, nil
}
