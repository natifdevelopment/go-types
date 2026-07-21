package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipPembangkitMap) TableName() string {
	return "t_pip_pembangkit_map"
}

type PipPembangkitMap struct {
	baseapp.BaseModel
	TPipPembangkitId *uuid.UUID     `gorm:"type:uuid" json:"pipPembangkitId"`
	PipPembangkit    *PipPembangkit `gorm:"foreignKey:TPipPembangkitId" json:"pipPembangkit"`
	TPembangkitId    *uuid.UUID     `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit       *Pembangkit    `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
}
