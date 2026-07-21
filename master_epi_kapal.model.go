package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (MasterEpiKapal) TableName() string {
	return "t_master_epi_kapal"
}

type MasterEpiKapal struct {
	baseapp.BaseModel
	Nama           string      `gorm:"type:varchar(255)" json:"nama"`
	IMO            *string     `gorm:"type:varchar(255)" json:"imo"`
	MMSI           *string     `gorm:"type:varchar(255)" json:"mmsi"`
	JenisKapalCode *string     `gorm:"type:varchar(255)" json:"jenisKapalCode"`
	JenisKapal     *ConfigData `gorm:"foreignKey:JenisKapalCode;references:Code;" json:"jenisKapal"`
	TModaId        *uuid.UUID  `gorm:"type:uuid" json:"modaId"`
	Moda           *Moda       `gorm:"foreignKey:TModaId" json:"moda"`
	Kapasitas      *float64    `gorm:"type:decimal(20,3)" json:"kapasitas"`
	StatusKapal    bool        `gorm:"type:boolean" json:"statusKapal"`
}
