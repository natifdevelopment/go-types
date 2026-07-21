package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type LaycanType string

const (
	TipeHapus      LaycanType = "hapus"
	TipeRevisi     LaycanType = "revisi"
	TipePengalihan LaycanType = "pengalihan"
)

func (o LaycanType) IsValid() bool {
	switch o {
	case TipeHapus, TipeRevisi, TipePengalihan:
		return true
	default:
		return false
	}
}

func (RevisiLaycan) TableName() string {
	return "t_revisi_laycan"
}

type RevisiLaycan struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID            `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman     `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId          *uuid.UUID            `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization         `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId       *uuid.UUID            `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization         `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TipeLaycan          LaycanType            `gorm:"type:varchar(255)" json:"tipeLaycan"`
	TTargetPembangkitId *uuid.UUID            `gorm:"type:uuid" json:"targetPembangkitId"`
	TargetPembangkit    *Organization         `gorm:"foreignKey:TTargetPembangkitId" json:"targetPembangkit"`
	TargetLoading       *DateField      `gorm:"type:date" json:"targetLoading"`
	VolumeConfirm       *float64              `gorm:"type:float" json:"volumeConfirm"`
	Keterangan          *string               `gorm:"type:text" json:"keterangan"`
	RevisiLaycanHistory []RevisiLaycanHistory `gorm:"foreignKey:TRevisiLaycanId; constraint:OnDelete:CASCADE" json:"revisiLaycanHistory"`
}

type LaycanHistoryType string

const (
	TipeHistoryNew      LaycanHistoryType = "new"
	TipeHistoryRollback LaycanHistoryType = "rollback"
)

func (o LaycanHistoryType) IsValid() bool {
	switch o {
	case TipeHistoryNew, TipeHistoryRollback:
		return true
	default:
		return false
	}
}

func (RevisiLaycanHistory) TableName() string {
	return "t_revisi_laycan_history"
}

type RevisiLaycanHistory struct {
	baseapp.BaseModel
	TRevisiLaycanId     *uuid.UUID        `gorm:"type:uuid" json:"revisiLaycanId"`
	RevisiLaycan        *RevisiLaycan     `gorm:"foreignKey:TRevisiLaycanId" json:"revisiLaycan"`
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"Pembangkit"`
	TipeLaycan          LaycanType        `gorm:"type:varchar(255)" json:"tipeLaycan"`
	TTargetPembangkitId *uuid.UUID        `gorm:"type:uuid" json:"targetPembangkitId"`
	TargetPembangkit    *Organization     `gorm:"foreignKey:TTargetPembangkitId" json:"targetPembangkit"`
	TargetLoading       *DateField  `gorm:"type:date" json:"targetLoading"`
	VolumeConfirm       *float64          `gorm:"type:float" json:"volumeConfirm"`
	Keterangan          *string           `gorm:"type:text" json:"keterangan"`
	TipeHistory         LaycanHistoryType `gorm:"type:varchar(255);default:'new'" json:"tipeHistory"`
}
