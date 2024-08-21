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

var TransactionRecords = newTransactionRecordsTable("public", "transaction_records", "")

type transactionRecordsTable struct {
	postgres.Table

	// Columns
	ContractID  postgres.ColumnString
	LoanID      postgres.ColumnString
	Otr         postgres.ColumnInteger
	AdminFee    postgres.ColumnInteger
	Installment postgres.ColumnInteger
	Interest    postgres.ColumnInteger
	AssetName   postgres.ColumnString
	Total       postgres.ColumnInteger
	Status      postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TransactionRecordsTable struct {
	transactionRecordsTable

	EXCLUDED transactionRecordsTable
}

// AS creates new TransactionRecordsTable with assigned alias
func (a TransactionRecordsTable) AS(alias string) *TransactionRecordsTable {
	return newTransactionRecordsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransactionRecordsTable with assigned schema name
func (a TransactionRecordsTable) FromSchema(schemaName string) *TransactionRecordsTable {
	return newTransactionRecordsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransactionRecordsTable with assigned table prefix
func (a TransactionRecordsTable) WithPrefix(prefix string) *TransactionRecordsTable {
	return newTransactionRecordsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransactionRecordsTable with assigned table suffix
func (a TransactionRecordsTable) WithSuffix(suffix string) *TransactionRecordsTable {
	return newTransactionRecordsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransactionRecordsTable(schemaName, tableName, alias string) *TransactionRecordsTable {
	return &TransactionRecordsTable{
		transactionRecordsTable: newTransactionRecordsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newTransactionRecordsTableImpl("", "excluded", ""),
	}
}

func newTransactionRecordsTableImpl(schemaName, tableName, alias string) transactionRecordsTable {
	var (
		ContractIDColumn  = postgres.StringColumn("contract_id")
		LoanIDColumn      = postgres.StringColumn("loan_id")
		OtrColumn         = postgres.IntegerColumn("otr")
		AdminFeeColumn    = postgres.IntegerColumn("admin_fee")
		InstallmentColumn = postgres.IntegerColumn("installment")
		InterestColumn    = postgres.IntegerColumn("interest")
		AssetNameColumn   = postgres.StringColumn("asset_name")
		TotalColumn       = postgres.IntegerColumn("total")
		StatusColumn      = postgres.IntegerColumn("status")
		allColumns        = postgres.ColumnList{ContractIDColumn, LoanIDColumn, OtrColumn, AdminFeeColumn, InstallmentColumn, InterestColumn, AssetNameColumn, TotalColumn, StatusColumn}
		mutableColumns    = postgres.ColumnList{LoanIDColumn, OtrColumn, AdminFeeColumn, InstallmentColumn, InterestColumn, AssetNameColumn, TotalColumn, StatusColumn}
	)

	return transactionRecordsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ContractID:  ContractIDColumn,
		LoanID:      LoanIDColumn,
		Otr:         OtrColumn,
		AdminFee:    AdminFeeColumn,
		Installment: InstallmentColumn,
		Interest:    InterestColumn,
		AssetName:   AssetNameColumn,
		Total:       TotalColumn,
		Status:      StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
