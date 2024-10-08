//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PurchaseCatalog = newPurchaseCatalogTable("public", "purchase_catalog", "")

type purchaseCatalogTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnString
	ItemID     postgres.ColumnString
	ProviderID postgres.ColumnString
	Name       postgres.ColumnString
	Price      postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PurchaseCatalogTable struct {
	purchaseCatalogTable

	EXCLUDED purchaseCatalogTable
}

// AS creates new PurchaseCatalogTable with assigned alias
func (a PurchaseCatalogTable) AS(alias string) *PurchaseCatalogTable {
	return newPurchaseCatalogTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PurchaseCatalogTable with assigned schema name
func (a PurchaseCatalogTable) FromSchema(schemaName string) *PurchaseCatalogTable {
	return newPurchaseCatalogTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PurchaseCatalogTable with assigned table prefix
func (a PurchaseCatalogTable) WithPrefix(prefix string) *PurchaseCatalogTable {
	return newPurchaseCatalogTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PurchaseCatalogTable with assigned table suffix
func (a PurchaseCatalogTable) WithSuffix(suffix string) *PurchaseCatalogTable {
	return newPurchaseCatalogTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPurchaseCatalogTable(schemaName, tableName, alias string) *PurchaseCatalogTable {
	return &PurchaseCatalogTable{
		purchaseCatalogTable: newPurchaseCatalogTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newPurchaseCatalogTableImpl("", "excluded", ""),
	}
}

func newPurchaseCatalogTableImpl(schemaName, tableName, alias string) purchaseCatalogTable {
	var (
		IDColumn         = postgres.StringColumn("id")
		ItemIDColumn     = postgres.StringColumn("item_id")
		ProviderIDColumn = postgres.StringColumn("provider_id")
		NameColumn       = postgres.StringColumn("name")
		PriceColumn      = postgres.IntegerColumn("price")
		allColumns       = postgres.ColumnList{IDColumn, ItemIDColumn, ProviderIDColumn, NameColumn, PriceColumn}
		mutableColumns   = postgres.ColumnList{ItemIDColumn, ProviderIDColumn, NameColumn, PriceColumn}
	)

	return purchaseCatalogTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		ItemID:     ItemIDColumn,
		ProviderID: ProviderIDColumn,
		Name:       NameColumn,
		Price:      PriceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
