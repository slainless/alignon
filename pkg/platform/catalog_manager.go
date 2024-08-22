package platform

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/query"
)

type MissingProductError struct {
	Expected []uuid.UUID
	Got      []uuid.UUID
}

func (e *MissingProductError) Error() string {
	return fmt.Sprintf("expected products %v but only got %v", e.Expected, e.Got)
}

type CatalogManager struct {
	db *sql.DB

	errorTracker ErrorTracker
}

func NewCatalogManager(db *sql.DB, tracker ErrorTracker) *CatalogManager {
	return &CatalogManager{
		db: db,

		errorTracker: tracker,
	}
}

func (m *CatalogManager) GetItems(ctx context.Context, itemIDs []uuid.UUID) ([]Product, error) {
	var products []Product
	err := query.GetProducts(ctx, m.db, itemIDs, &products)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	ids := make([]uuid.UUID, 0, len(products))
	for i := range products {
		products[i].P_ID = products[i].ID.String()
		ids = append(ids, products[i].ID)
	}

	if len(ids) != len(itemIDs) {
		return nil, &MissingProductError{
			Expected: itemIDs,
			Got:      ids,
		}
	}

	return products, nil
}
