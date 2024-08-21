package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/table"
)

func GetConsumerByEmail(ctx context.Context, db *sql.DB, email string, to any) error {
	stmt := SELECT(table.Consumers.AllColumns).
		FROM(table.Consumers).
		WHERE(table.Consumers.Email.EQ(String(email)))

	err := stmt.QueryContext(ctx, db, &to)
	if err != nil {
		return err
	}

	return err
}

func GetConsumerLimit(ctx context.Context, db *sql.DB, id uuid.UUID, to any) error {
	stmt := SELECT(table.Limits.AllColumns).
		FROM(table.Limits).
		WHERE(table.Limits.ConsumerID.EQ(UUID(id)))

	err := stmt.QueryContext(ctx, db, &to)
	if err != nil {
		return err
	}

	return nil
}

func GetCurrentLoanOfConsumer(ctx context.Context, db *sql.DB, id uuid.UUID, to any) error {
	stmt := SELECT(
		table.Loans.AllColumns,
		table.TransactionRecords.AllColumns,
		table.InstallmentRecords.AllColumns,
	).
		FROM(
			table.Loans.
				INNER_JOIN(table.TransactionRecords, table.TransactionRecords.LoanID.EQ(table.Loans.ID)).
				INNER_JOIN(table.InstallmentRecords, table.InstallmentRecords.ContractID.EQ(table.TransactionRecords.ContractID)),
		).
		WHERE(table.Loans.ConsumerID.EQ(UUID(id)))

	err := stmt.QueryContext(ctx, db, to)
	if err != nil {
		return err
	}

	return nil
}

// func InsertFreshCustomer(ctx context.Context, db *sql.DB, email string) error {
// 	stmt := table.
// 		Consumers.INSERT(table.Consumers.MutableColumns).
// 		MODELS(&model.Consumers{Email: email})

// 	_, err := stmt.ExecContext(ctx, db)
// 	return err
// }
