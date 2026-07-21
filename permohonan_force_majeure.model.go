package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PermohonanForceMajeure) TableName() string {
	return "t_permohonan_force_majeure"
}

type PermohonanForceMajeure struct {
	baseapp.BaseModel
	TPembangkitId      *uuid.UUID                         `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit         *Organization                      `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	Hari               float64                            `gorm:"type:float" json:"hari"`
	TDokForceMajeureId *uuid.UUID                         `gorm:"type:uuid" json:"dokForceMajeureId"`
	DokForceMajeure    *Media                             `gorm:"foreignKey:TDokForceMajeureId" json:"dokForceMajeure"`
	JadwalPengiriman   []PermohonanForceMajeurePengiriman `gorm:"foreignKey:TPermohonanForceMajeureId" json:"jadwalPengiriman"`
}

func (PermohonanForceMajeurePengiriman) TableName() string {
	return "t_permohonan_force_majeure_pengiriman"
}

type PermohonanForceMajeurePengiriman struct {
	baseapp.BaseModel
	TPermohonanForceMajeureId *uuid.UUID              `gorm:"type:uuid" json:"permohonanForceMajeureId"`
	PermohonanForceMajeure    *PermohonanForceMajeure `gorm:"foreignKey:TPermohonanForceMajeureId" json:"permohonanForceMajeure"`
	TJadwalPengirimanId       *uuid.UUID              `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman          *JadwalPengiriman       `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
}
