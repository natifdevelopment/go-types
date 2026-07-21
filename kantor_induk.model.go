package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type KantorInduk struct {
	baseapp.BaseModel
	NamaInstansi     string           `gorm:"type:varchar(255)" json:"namaInstansi"`
	NamaSingkat      string           `gorm:"type:varchar(255)" json:"namaSingkat"`
	Kota             string           `gorm:"type:varchar(255)" json:"kota"`
	Alamat           string           `gorm:"type:text" json:"alamat"`
	MasaBerlakuAwal  *DateField `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir *DateField `gorm:"type:date" json:"masaBerlakuAkhir"`
	TRegionalId      *uuid.UUID       `gorm:"type:uuid" json:"regionalId"`
	Regional         *Regional        `gorm:"foreignKey:TRegionalId" json:"regional"`
	CompanyCode      string           `gorm:"type:varchar(45)" json:"companyCode"`
}

func (KantorInduk) TableName() string {
	return "t_kantor_induk"
}
