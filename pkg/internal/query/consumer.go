package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/model"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/table"
)

func GetConsumerByEmail(ctx context.Context, db *sql.DB, email string) (*model.Consumers, error) {
	stmt := SELECT(table.Consumers.AllColumns).
		FROM(table.Consumers).
		WHERE(table.Consumers.Email.EQ(String(email)))

	var consumer model.Consumers
	err := stmt.QueryContext(ctx, db, &consumer)
	if err != nil {
		return nil, err
	}

	return &consumer, err
}

func GetConsumerLimit(ctx context.Context, db *sql.DB, id uuid.UUID) (*model.Limits, error) {
	stmt := SELECT(table.Limits.AllColumns).
		FROM(table.Limits).
		WHERE(table.Limits.ConsumerID.EQ(UUID(id)))

	var limit model.Limits
	err := stmt.QueryContext(ctx, db, &limit)
	if err != nil {
		return nil, err
	}

	return &limit, err
}

// func InsertFreshCustomer(ctx context.Context, db *sql.DB, email string) error {
// 	stmt := table.
// 		Consumers.INSERT(table.Consumers.MutableColumns).
// 		MODELS(&model.Consumers{Email: email})

// 	_, err := stmt.ExecContext(ctx, db)
// 	return err
// }
