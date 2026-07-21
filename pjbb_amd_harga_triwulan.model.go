package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PjbbAmdHargaTriwulan) TableName() string {
	return "t_pjbb_amd_harga_triwulan"
}

type PjbbAmdHargaTriwulan struct {
	baseapp.BaseModel
	TPemasokId            *uuid.UUID       `gorm:"index;type:uuid" json:"pemasokId"`
	Pemasok               *Organization    `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPjbbAmandemenId      *uuid.UUID       `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen         *PjbbAmandemen   `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TipePengiriman        JenisPengiriman  `gorm:"type:varchar(255)" json:"tipePengiriman"`
	SkemaKontrakCode      *JenisPengiriman `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak          *ConfigDataInfo  `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TSpekId               *uuid.UUID       `gorm:"index;type:uuid" json:"spekId"`
	Spek                  *Spek            `gorm:"foreignKey:TSpekId" json:"spek"`
	JenisTransportasiCode *string          `gorm:"type:varchar(255)" json:"jenisTransportasiCode"`
	JenisTransportasi     *ConfigDataInfo  `gorm:"foreignKey:JenisTransportasiCode;references:Code;" json:"jenisTransportasi"`

	Currency string `gorm:"type:varchar(255)" json:"currency"` // USD/IDR

	HargaFob float64 `gorm:"type:float" json:"hargaFob"`

	// Triwulan
	NoSurat                   string           `gorm:"type:varchar(255)" json:"noSurat"`
	TglSurat                  *DateField `gorm:"type:date" json:"tglSurat"`
	PeriodeBerlakuAwal        DateField  `gorm:"type:date" json:"periodeBerlakuAwal"`
	PeriodeBerlakuAkhir       DateField  `gorm:"type:date" json:"periodeBerlakuAkhir"`
	TTriwulanId               *uuid.UUID       `gorm:"index;type:uuid" json:"triwulanId"`
	Triwulan                  *Triwulan        `gorm:"foreignKey:TTriwulanId" json:"triwulan"`
	BaTriwulanBaBankGaransiId *uuid.UUID       `gorm:"type:index;type:uuid" json:"baTriwulanBaBankGaransiId"`
	BaTriwulanBaBankGaransi   *Media           `gorm:"foreignKey:BaTriwulanBaBankGaransiId" json:"baTriwulanBaBankGaransi"`
}
