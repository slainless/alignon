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

var Loans = newLoansTable("public", "loans", "")

type loansTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnString
	ConsumerID        postgres.ColumnString
	Amount            postgres.ColumnInteger
	Tenor             postgres.ColumnInteger
	InstallmentLength postgres.ColumnInteger
	ConsumerLimit     postgres.ColumnInteger
	ConsumerSalary    postgres.ColumnInteger
	LoanedAt          postgres.ColumnTimestamp
	Status            postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LoansTable struct {
	loansTable

	EXCLUDED loansTable
}

// AS creates new LoansTable with assigned alias
func (a LoansTable) AS(alias string) *LoansTable {
	return newLoansTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LoansTable with assigned schema name
func (a LoansTable) FromSchema(schemaName string) *LoansTable {
	return newLoansTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LoansTable with assigned table prefix
func (a LoansTable) WithPrefix(prefix string) *LoansTable {
	return newLoansTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LoansTable with assigned table suffix
func (a LoansTable) WithSuffix(suffix string) *LoansTable {
	return newLoansTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLoansTable(schemaName, tableName, alias string) *LoansTable {
	return &LoansTable{
		loansTable: newLoansTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newLoansTableImpl("", "excluded", ""),
	}
}

func newLoansTableImpl(schemaName, tableName, alias string) loansTable {
	var (
		IDColumn                = postgres.StringColumn("id")
		ConsumerIDColumn        = postgres.StringColumn("consumer_id")
		AmountColumn            = postgres.IntegerColumn("amount")
		TenorColumn             = postgres.IntegerColumn("tenor")
		InstallmentLengthColumn = postgres.IntegerColumn("installment_length")
		ConsumerLimitColumn     = postgres.IntegerColumn("consumer_limit")
		ConsumerSalaryColumn    = postgres.IntegerColumn("consumer_salary")
		LoanedAtColumn          = postgres.TimestampColumn("loaned_at")
		StatusColumn            = postgres.IntegerColumn("status")
		allColumns              = postgres.ColumnList{IDColumn, ConsumerIDColumn, AmountColumn, TenorColumn, InstallmentLengthColumn, ConsumerLimitColumn, ConsumerSalaryColumn, LoanedAtColumn, StatusColumn}
		mutableColumns          = postgres.ColumnList{ConsumerIDColumn, AmountColumn, TenorColumn, InstallmentLengthColumn, ConsumerLimitColumn, ConsumerSalaryColumn, LoanedAtColumn, StatusColumn}
	)

	return loansTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		ConsumerID:        ConsumerIDColumn,
		Amount:            AmountColumn,
		Tenor:             TenorColumn,
		InstallmentLength: InstallmentLengthColumn,
		ConsumerLimit:     ConsumerLimitColumn,
		ConsumerSalary:    ConsumerSalaryColumn,
		LoanedAt:          LoanedAtColumn,
		Status:            StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
