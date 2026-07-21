package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PengalihanPasokan) TableName() string {
	return "t_pengalihan_pasokan"
}

type PengalihanPasokan struct {
	baseapp.BaseModel
	NoPengiriman        string          `gorm:"type:varchar(255)" json:"noPengiriman"`
	TPembangkitAsalId   uuid.UUID       `gorm:"type:uuid" json:"TPembangkitAsalId"`
	PembangkitAsal      Organization    `gorm:"foreignKey:TPembangkitAsalId" json:"PembangkitAsal"`
	TPembangkitTujuanId uuid.UUID       `gorm:"type:uuid" json:"TPembangkitTujuanId"`
	PembangkitTujuan    Organization    `gorm:"foreignKey:TPembangkitTujuanId" json:"PembangkitTujuan"`
	TglTaConfirm        DateField `gorm:"type:date" json:"tglTaConfirm"`
	VolumeConfirm       float64         `gorm:"type:float" json:"volumeConfirm"`
}
