package platform

import (
	"time"

	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/model"
)

type InstallmentRecord struct {
	model.InstallmentRecords
	InstallmentID uuid.UUID `json:"installment_id"`
	PaidAt        time.Time `json:"paid_at"`
}

type TransactionRecord struct {
	model.TransactionRecords
	ContractID  string    `json:"contract_id"`
	LoanID      uuid.UUID `json:"loan_id"`
	Otr         *int64    `json:"otr"`
	AdminFee    *int64    `json:"admin_fee"`
	Installment *int64    `json:"installment"`
	Interest    *int64    `json:"interest"`
	AssetName   string    `json:"asset_name"`
	Amount      int64     `json:"amount"`
	StatusName  string    `json:"status"`
	CatalogID   uuid.UUID `json:"catalog_id"`

	InstallmentRecords []InstallmentRecord `json:"installments"`
}

type Loan struct {
	model.Loans
	ConsumerID        uuid.UUID  `json:"consumer_id"`
	Amount            int64      `json:"amount"`
	Tenor             int16      `json:"tenor"`
	InstallmentLength int16      `json:"installment_length"`
	ConsumerLimit     int64      `json:"consumer_limit"`
	ConsumerSalary    int64      `json:"consumer_salary"`
	LoanedAt          *time.Time `json:"loaned_at"`
	StatusName        string     `json:"status"`

	TransactionRecords []TransactionRecord `json:"transactions"`
}
