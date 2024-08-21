package platform

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/pkg/internal/query"
)

var (
	ErrConsumerNotFound          = errors.New("consumer not found")
	ErrConsumerAlreadyRegistered = errors.New("consumer already registered")
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

func (m *ConsumerManager) Register(ctx context.Context, email string) error {
	err := query.InsertFreshCustomer(ctx, m.db, email)
	if err != nil {
		if err := util.PQError(err); err != nil {
			switch err.Code.Name() {
			case "unique_violation":
				return ErrConsumerAlreadyRegistered
			}
		}
		m.errorTracker.Report(ctx, err)
		return err
	}

	return nil
}
