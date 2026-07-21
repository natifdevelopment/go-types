package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (PjabAmdPembiayaan) TableName() string {
	return "t_pjab_amd_pembiayaan"
}

type PjabAmdPembiayaan struct {
	baseapp.BaseModel
	TTransportirId   *uuid.UUID     `gorm:"type:uuid" json:"transportirId"`
	Transportir      *Organization  `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TPjabAmandemenId *uuid.UUID     `gorm:"type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen    *PjabAmandemen `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TPembangkitId    *uuid.UUID     `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit       *Organization  `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId       *uuid.UUID     `gorm:"type:uuid" json:"pemasokId"`
	Pemasok          *Organization  `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPelabuhanMuatId *uuid.UUID     `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat    *PelabuhanMuat `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TModaId          *uuid.UUID     `gorm:"type:uuid" json:"modaId"`
	Moda             *Moda          `gorm:"foreignKey:TModaId" json:"moda"`
	TarifSewa        *float64       `gorm:"type:double precision" json:"tarifSewa"`
	HsdBase          *float64       `gorm:"type:double precision" json:"hsdBase"`
	FaktorX          *float64       `gorm:"type:double precision" json:"faktorX"`
	FaktorY          *float64       `gorm:"type:double precision" json:"faktorY"`
	FaktorY2         *float64       `gorm:"type:double precision" json:"faktorY2"`
	HargaTrb         *float64       `gorm:"type:double precision" json:"hargaTrb"`
	MfoBase          *float64       `gorm:"type:double precision" json:"mfoBase"`
}
