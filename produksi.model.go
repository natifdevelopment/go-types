package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

type Produksi struct {
	baseapp.BaseModel
	BulanTahun YearMonthField `gorm:"type:date" json:"BulanTahun"`
	SFC        string               `gorm:"type:varchar(255)" json:"SFC"`
	CF         string               `gorm:"type:varchar(255)" json:"CF"`
	Bedding    float64              `gorm:"type:float" json:"Bedding"` // In ton
}

func (Produksi) TableName() string {
	return "t_produksi"
}
