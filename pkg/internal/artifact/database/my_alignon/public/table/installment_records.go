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

var InstallmentRecords = newInstallmentRecordsTable("public", "installment_records", "")

type installmentRecordsTable struct {
	postgres.Table

	// Columns
	InstallmentID postgres.ColumnString
	TransactionID postgres.ColumnString
	PaidAt        postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type InstallmentRecordsTable struct {
	installmentRecordsTable

	EXCLUDED installmentRecordsTable
}

// AS creates new InstallmentRecordsTable with assigned alias
func (a InstallmentRecordsTable) AS(alias string) *InstallmentRecordsTable {
	return newInstallmentRecordsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new InstallmentRecordsTable with assigned schema name
func (a InstallmentRecordsTable) FromSchema(schemaName string) *InstallmentRecordsTable {
	return newInstallmentRecordsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new InstallmentRecordsTable with assigned table prefix
func (a InstallmentRecordsTable) WithPrefix(prefix string) *InstallmentRecordsTable {
	return newInstallmentRecordsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new InstallmentRecordsTable with assigned table suffix
func (a InstallmentRecordsTable) WithSuffix(suffix string) *InstallmentRecordsTable {
	return newInstallmentRecordsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newInstallmentRecordsTable(schemaName, tableName, alias string) *InstallmentRecordsTable {
	return &InstallmentRecordsTable{
		installmentRecordsTable: newInstallmentRecordsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newInstallmentRecordsTableImpl("", "excluded", ""),
	}
}

func newInstallmentRecordsTableImpl(schemaName, tableName, alias string) installmentRecordsTable {
	var (
		InstallmentIDColumn = postgres.StringColumn("installment_id")
		TransactionIDColumn = postgres.StringColumn("transaction_id")
		PaidAtColumn        = postgres.TimestampColumn("paid_at")
		allColumns          = postgres.ColumnList{InstallmentIDColumn, TransactionIDColumn, PaidAtColumn}
		mutableColumns      = postgres.ColumnList{TransactionIDColumn, PaidAtColumn}
	)

	return installmentRecordsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		InstallmentID: InstallmentIDColumn,
		TransactionID: TransactionIDColumn,
		PaidAt:        PaidAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
