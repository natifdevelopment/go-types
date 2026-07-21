package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"
)

func (PipPemakaian) TableName() string {
	return "t_pip_pemakaian"
}

type PipPemakaian struct {
	baseapp.BaseModel
	OperatingUnit   *string    `json:"operatingUnit" gorm:"type:varchar(255)"`
	TransactionId   *string    `json:"transactionId" gorm:"type:varchar(255)"`
	TransactionDate *time.Time `json:"transactionDate" gorm:"type:timestamp"` // DD/MM/YYYY
	ItemNumber      *string    `json:"itemNumber" gorm:"type:varchar(255)"`
	Storeroom       *string    `json:"storeroom" gorm:"type:varchar(255)"`
	Satuan          *string    `json:"satuan" gorm:"type:varchar(255)"`
	Quantity        *float64   `json:"quantity,string" gorm:"type:float"`
	ItemCost        *float64   `json:"itemCost,string" gorm:"type:float"`
	AmountLine      *float64   `json:"amountLine,string" gorm:"type:float"`
	IssueBy         *string    `json:"issueBy" gorm:"type:varchar(255)"`
	NamaMesin       *string    `json:"namaMesin" gorm:"type:varchar(255)"`
	TrxIdMmt        *string    `json:"trxIdMmt" gorm:"type:varchar(255)"`
	ResponseHash    *string    `gorm:"type:varchar(255)" json:"-"`
}
