package types
import (
	"time"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipReturnStock) TableName() string {
	return "t_pip_return_stock"
}

type PipReturnStock struct {
	baseapp.BaseModel
	OperatingUnit    *string    `gorm:"type:varchar(255)" json:"operatingUnit"`
	NoKontrakAdendum *string    `gorm:"type:varchar(255)" json:"noKontrakAdendum"`
	PoNumber         *string    `gorm:"type:varchar(255)" json:"poNumber"`
	RevisionPoNumber *string    `gorm:"type:varchar(255)" json:"revisionPoNumber"`
	SupplierId       *string    `gorm:"type:varchar(255)" json:"supplierId"`
	Supplier         *string    `gorm:"type:varchar(255)" json:"supplier"`
	ReturnId         *string    `gorm:"type:varchar(255)" json:"returnId"`
	TransactionDate  *time.Time `gorm:"type:timestamp" json:"transactionDate"` // DD/MM/YYYY
	ReceiptId        *string    `gorm:"type:varchar(255)" json:"receiptId"`
	ReturnBy         *string    `gorm:"type:varchar(255)" json:"returnBy"`
	ItemNumber       *string    `gorm:"type:varchar(255)" json:"itemNumber"`
	Storeroom        *string    `gorm:"type:varchar(255)" json:"storeroom"`
	Satuan           *string    `gorm:"type:varchar(255)" json:"satuan"`
	Quantity         *float64   `gorm:"type:float" json:"quantity"`
	ItemCost         *float64   `gorm:"type:float" json:"itemCost"`
	AmountLine       *float64   `gorm:"type:float" json:"amountLine"`
	TrxId            *string    `gorm:"type:varchar(255)" json:"trxId"`
	ResponseHash     *string    `gorm:"type:varchar(255)" json:"-"`
}
