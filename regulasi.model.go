package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (Regulasi) TableName() string {
	return "t_regulasi"
}

type Regulasi struct {
	baseapp.BaseModel
	RegulasiCode string          `gorm:"type:varchar(255)" json:"regulasiCode"`
	Regulasi     *ConfigData     `gorm:"foreignKey:RegulasiCode;references:Code;" json:"regulasi"`
	Nomor        string          `gorm:"type:varchar(255)" json:"nomor"`
	Tanggal      DateField `gorm:"type:date" json:"tanggal"`
}
