package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (KemampuanPasok) TableName() string {
	return "t_kemampuan_pasok"
}

type KemampuanPasok struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID            `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	Periode         *YearMonthField `gorm:"type:date" json:"periode"`
	KemampuanPasok  []KemampuanPasokItem  `gorm:"foreignKey:TKemampuanPasokId; constraint:OnDelete:CASCADE" json:"kemampuanPasok"`
}

func (KemampuanPasokItem) TableName() string {
	return "t_kemampuan_pasok_item"
}

type KemampuanPasokItem struct {
	baseapp.BaseModel
	TKemampuanPasokId *uuid.UUID       `gorm:"type:uuid" json:"kemampuanPasokId"`
	KemampuanPasok    *KemampuanPasok  `gorm:"foreignKey:TKemampuanPasokId" json:"kemampuanPasok"`
	SkemaKontrakCode  *JenisPengiriman `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak      *ConfigDataInfo  `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	VolumeLrc         *float64         `gorm:"type:float" json:"volumeLrc"`
	VolumeMrc         *float64         `gorm:"type:float" json:"volumeMrc"`
	VolumeHrc         *float64         `gorm:"type:float" json:"volumeHrc"`
}
