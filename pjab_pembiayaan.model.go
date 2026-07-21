package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PjabPembiayaan) TableName() string {
	return "t_pjab_pembiayaan"
}

type PjabPembiayaan struct {
	baseapp.BaseModel
	TTransportirId   *uuid.UUID     `gorm:"type:uuid" json:"transportirId"`
	Transportir      *Organization  `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TPjabId          *uuid.UUID     `gorm:"type:uuid" json:"pjabId"`
	Pjab             *Pjab          `gorm:"foreignKey:TPjabId" json:"pjab"`
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
	// TrfHsd           *float64       `gorm:"type:double precision" json:"trfHsd"`
}
