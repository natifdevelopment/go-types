package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PindahPasokan) TableName() string {
	return "t_pindah_pasokan"
}

type PindahPasokan struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TTujuanPembangkitId *uuid.UUID        `gorm:"type:uuid" json:"tujuanPembangkitId"`
	TujuanPembangkit    *Organization     `gorm:"foreignKey:TTujuanPembangkitId" json:"tujuanPembangkit"`
	TanggalEta          *DateField  `gorm:"type:date" json:"tanggalEta"`
	VolumeConfirm       float64           `gorm:"type:float" json:"volumeConfirm"`
	TDokPindahPasokanId *uuid.UUID        `gorm:"type:uuid" json:"dokPindahPasokanId"`
	DokPindahPasokan    *Media            `gorm:"foreignKey:TDokPindahPasokanId" json:"dokPindahPasokan"`
	Keterangan          *string           `gorm:"type:text" json:"keterangan"`
}

type PindahPasokanTipeHistory string

const (
	PindahPasokanTipeHistoryNew      PindahPasokanTipeHistory = "new"
	PindahPasokanTipeHistoryRollback PindahPasokanTipeHistory = "rollback"
)

func (o PindahPasokanTipeHistory) IsValid() bool {
	switch o {
	case PindahPasokanTipeHistoryNew, PindahPasokanTipeHistoryRollback:
		return true
	default:
		return false
	}
}

func (PindahPasokanHistory) TableName() string {
	return "t_pindah_pasokan_history"
}

type PindahPasokanHistory struct {
	baseapp.BaseModel
	TPindahPasokanId    *uuid.UUID               `gorm:"type:uuid" json:"pindahPasokanId"`
	PindahPasokan       *PindahPasokan           `gorm:"foreignKey:TPindahPasokanId" json:"pindahPasokan"`
	TJadwalPengirimanId *uuid.UUID               `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman        `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TTujuanPembangkitId *uuid.UUID               `gorm:"type:uuid" json:"tujuanPembangkitId"`
	TujuanPembangkit    *Organization            `gorm:"foreignKey:TTujuanPembangkitId" json:"tujuanPembangkit"`
	TanggalEta          *DateField         `gorm:"type:date" json:"tanggalEta"`
	VolumeConfirm       float64                  `gorm:"type:float" json:"volumeConfirm"`
	TDokPindahPasokanId *uuid.UUID               `gorm:"type:uuid" json:"dokPindahPasokanId"`
	DokPindahPasokan    *Media                   `gorm:"foreignKey:TDokPindahPasokanId" json:"dokPindahPasokan"`
	Keterangan          *string                  `gorm:"type:text" json:"keterangan"`
	TipeHistory         PindahPasokanTipeHistory `gorm:"type:varchar(255);default:'new'" json:"tipeHistory"`
}
