package types
import (
	"time"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipReturnPemakaian) TableName() string {
	return "t_pip_return_pemakaian"
}

type PipReturnPemakaian struct {
	baseapp.BaseModel
	OperatingUnit    *string    `json:"operatingUnit" gorm:"type:varchar(255)"`
	NoKontrakAdendum *string    `json:"noKontrakAdendum" gorm:"type:varchar(255)"`
	PoNumber         *string    `json:"poNumber" gorm:"type:varchar(255)"`
	RevisionPoNumber *string    `json:"revisionPoNumber" gorm:"type:varchar(255)"`
	SupplierId       *string    `json:"supplierId" gorm:"type:varchar(255)"`
	Supplier         *string    `json:"supplier" gorm:"type:varchar(255)"`
	ReturnId         *string    `json:"returnId" gorm:"type:varchar(255)"`
	TransactionDate  *time.Time `json:"transactionDate" gorm:"type:timestamp"` // DD/MM/YYYY
	ReceiptId        *string    `json:"receiptId" gorm:"type:varchar(255)"`
	ReturnBy         *string    `json:"returnBy" gorm:"type:varchar(255)"`
	ItemNumber       *string    `json:"itemNumber" gorm:"type:varchar(255)"`
	Storeroom        *string    `json:"storeroom" gorm:"type:varchar(255)"`
	Satuan           *string    `json:"satuan" gorm:"type:varchar(255)"`
	Quantity         *float64   `json:"quantity,string" gorm:"type:float"`
	ItemCost         *float64   `json:"itemCost,string" gorm:"type:float"`
	AmountLine       *float64   `json:"amountLine,string" gorm:"type:float"`
	TrxId            *string    `json:"trxId" gorm:"type:varchar(255)"`
	ResponseHash     *string    `gorm:"type:varchar(255)" json:"-"`
}
