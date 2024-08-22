package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/table"
)

func GetProducts(ctx context.Context, db *sql.DB, itemIDs []uuid.UUID, ToProducts any) error {
	stmt := SELECT(table.PurchaseCatalog.AllColumns).
		FROM(table.PurchaseCatalog).
		WHERE(table.PurchaseCatalog.ItemID.IN(util.ToPostgresUUIDs(itemIDs)...))

	err := stmt.QueryContext(ctx, db, &ToProducts)
	if err != nil {
		return err
	}

	return nil
}
