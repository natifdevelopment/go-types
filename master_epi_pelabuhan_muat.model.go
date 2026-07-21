package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (MasterEpiPelabuhanMuat) TableName() string {
	return "t_master_epi_pelabuhan_muat"
}

type MasterEpiPelabuhanMuat struct {
	baseapp.BaseModel
	Nama          string          `gorm:"type:varchar(255)" json:"nama"`
	Latitude      float64         `gorm:"type:float" json:"latitude"`
	Longitude     float64         `gorm:"type:float" json:"longitude"`
	TWilayahId    uuid.UUID       `gorm:"type:uuid" json:"wilayahId"`
	Wilayah       *Wilayah        `gorm:"foreignKey:TWilayahId" json:"wilayah"`
	UnLocode      string          `gorm:"type:varchar(255)" json:"unLocode"`
	ZonaWaktuCode string          `gorm:"type:varchar(255)" json:"zonaWaktuCode"`
	ZonaWaktu     *ConfigDataInfo `gorm:"foreignKey:ZonaWaktuCode;references:Code;" json:"zonaWaktu"`
}
