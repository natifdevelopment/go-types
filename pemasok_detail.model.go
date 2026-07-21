package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PemasokDetail) TableName() string {
	return "t_pemasok_detail"
}

type PemasokDetail struct {
	baseapp.BaseModel
	TraderId         uuid.UUID       `gorm:"type:uuid" json:"traderId"`
	NamaTambang      string          `gorm:"type:varchar(255)" json:"namaTambang"`
	ModiId           uuid.UUID       `gorm:"type:uuid" json:"modiId"`
	ProduksiId       uuid.UUID       `gorm:"type:uuid" json:"ProduksiId"`
	Produksi         Produksi        `gorm:"foreignKey:ProduksiId" json:"Produksi"`
	LandingPort      string          `gorm:"type:varchar(25)" json:"landingPort"`
	FasilitasLoading string          `gorm:"type:varchar(25)" json:"fasilitasLoading"`
	JarakHauling     float64         `gorm:"type:float" json:"jarakHauling"`
	MasaBerlakuAwal  DateField `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir DateField `gorm:"type:date" json:"masaBerlakuAkhir"`
}
