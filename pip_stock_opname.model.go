package types
import (
	"time"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipStockOpname) TableName() string {
	return "t_pip_stock_opname"
}

type PipStockOpname struct {
	baseapp.BaseModel
	OperatingUnit       *string    `json:"operatingUnit" gorm:"type:varchar(255)"`
	TransactionId       *string    `json:"transactionId" gorm:"type:varchar(255)"`
	TransactionDate     *time.Time `json:"transactionDate" gorm:"type:timestamp"` // DD/MM/YYYY
	TipeMiscTransaction *string    `json:"tipeMiscTransaction" gorm:"type:varchar(255)"`
	ItemNumber          *string    `json:"itemNumber" gorm:"type:varchar(255)"`
	Storeroom           *string    `json:"storeroom" gorm:"type:varchar(255)"`
	Satuan              *string    `json:"satuan" gorm:"type:varchar(255)"`
	StokSebelum         *float64   `json:"stokSebelum" gorm:"type:float"`
	Quantity            *float64   `json:"quantity" gorm:"type:float"`
	StokSesudah         *float64   `json:"stokSesudah" gorm:"type:float"`
	ItemCost            *float64   `json:"itemCost" gorm:"type:float"`
	AmountLine          *float64   `json:"amountLine" gorm:"type:float"`
	CreatedBy           *string    `json:"createdBy" gorm:"type:varchar(255)"`
	Tipe                *string    `json:"tipe" gorm:"type:varchar(255)"`
	ResponseHash        *string    `gorm:"type:varchar(255)" json:"-"`
}
