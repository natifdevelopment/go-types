package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PosisiKapal) TableName() string {
	return "t_posisi_kapal"
}

type PosisiKapal struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	Latitude            float64           `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude           float64           `gorm:"type:decimal(11,8)" json:"longitude"`
	StatusKapal         string            `gorm:"type:varchar(255)" json:"statusKapal"`
}
