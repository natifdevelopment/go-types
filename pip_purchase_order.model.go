package types
import (
	"time"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipPurchaseOrder) TableName() string {
	return "t_pip_purchase_order"
}

type PipPurchaseOrder struct {
	baseapp.BaseModel
	OperatingUnit         *string    `json:"operatingUnit" gorm:"type:varchar(255)"`
	NoKontrakAdendum      *string    `json:"noKontrakAdendum" gorm:"type:varchar(255)"`
	PoNumber              *string    `json:"poNumber" gorm:"type:varchar(255)"`
	TanggalCreatePo       *time.Time `json:"tanggalCreatePo" gorm:"type:timestamp"` // DD/MM/YYYY
	TipeKontrak           *string    `json:"tipeKontrak" gorm:"type:varchar(255)"`
	RevisionPoNumber      *string    `json:"revisionPoNumber" gorm:"type:varchar(255)"`
	Description           *string    `json:"description" gorm:"type:varchar(255)"`
	Supplier              *string    `json:"supplier" gorm:"type:varchar(255)"`
	PoStatus              *string    `json:"poStatus" gorm:"type:varchar(255)"`
	Currency              *string    `json:"currency" gorm:"type:varchar(255)"`
	TotalAmount           *float64   `json:"totalAmount,string" gorm:"type:float"`
	ItemNumber            *string    `json:"itemNumber" gorm:"type:varchar(255)"`
	ItemDescription       *string    `json:"itemDescription" gorm:"type:varchar(255)"`
	Satuan                *string    `json:"satuan" gorm:"type:varchar(255)"`
	Quantity              *float64   `json:"quantity,string" gorm:"type:float"`
	ItemCost              *float64   `json:"itemCost,string" gorm:"type:float"`
	AmountLine            *float64   `json:"amountLine,string" gorm:"type:float"`
	Storeroom             *string    `json:"storeroom" gorm:"type:varchar(255)"`
	TanggalJatuhTempo     *time.Time `json:"tanggalJatuhTempo" gorm:"type:timestamp"`
	TanggalKontrakAdendum *time.Time `json:"tanggalKontrakAdendum" gorm:"type:timestamp"`
	NpwpSupplier          *string    `json:"npwpSupplier" gorm:"type:varchar(255)"`
	NpwpOperatingUnit     *string    `json:"npwpOperatingUnit" gorm:"type:varchar(255)"`
	NoKontrakPemasokPjbb  *string    `json:"noKontrakPemasokPjbb" gorm:"type:varchar(255)"`
	SupplierId            *string    `json:"supplierId" gorm:"type:varchar(255)"`
	PoHeaderId            *string    `json:"poHeaderId" gorm:"type:varchar(255)"`
	ResponseHash          *string    `gorm:"type:varchar(255)" json:"-"`
}
