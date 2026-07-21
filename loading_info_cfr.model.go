package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (LoadingInfoCfr) TableName() string {
	return "t_loading_info_cfr"
}

type LoadingInfoCfr struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPjbbId             *uuid.UUID        `gorm:"type:uuid" json:"pjbbId"`
	Pjbb                *Pjbb             `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId    *uuid.UUID        `gorm:"type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen       *PjbbAmandemen    `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TPelabuhanMuatId    *uuid.UUID        `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat       *PelabuhanMuat    `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId             *uuid.UUID        `gorm:"type:uuid" json:"spekId"`
	Spek                *Spek             `gorm:"foreignKey:TSpekId" json:"spek"`
	// TSumberTambangId    *uuid.UUID        `gorm:"type:uuid" json:"sumberTambangId"`
	// SumberTambang       *SumberTambang    `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`

	// Pejabat
	TBastPejabatId             *uuid.UUID           `gorm:"type:uuid" json:"bastPejabatId"`
	BastPejabat                *Pejabat_surat_kuasa `gorm:"foreignKey:TBastPejabatId" json:"bastPejabat"`
	TPenyesuaianHargaPejabatId *uuid.UUID           `gorm:"type:uuid" json:"penyesuaianHargaPejabatId"`
	PenyesuaianHargaPejabat    *Pejabat_surat_kuasa `gorm:"foreignKey:TPenyesuaianHargaPejabatId" json:"penyesuaianHargaPejabat"`

	SumberTambang []LoadingInfoCfrSumberTambang `gorm:"foreignKey:TLoadingInfoCfrId" json:"sumberTambang"`
}

func (LoadingInfoCfrSumberTambang) TableName() string {
	return "t_loading_info_cfr_sumber_tambang"
}

type LoadingInfoCfrSumberTambang struct {
	baseapp.BaseModel
	TLoadingInfoCfrId   *uuid.UUID      `gorm:"index;type:uuid" json:"loadingInfoCfrId"`
	LoadingInfoCfr      *LoadingInfoCfr `gorm:"foreignKey:TLoadingInfoCfrId" json:"loadingInfoCfr"`
	TSumberTambangId    *uuid.UUID      `gorm:"index;type:uuid" json:"sumberTambangId"`
	SumberTambang       *SumberTambang  `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	VolumeSumberTambang float64         `gorm:"type:float" json:"volumeSumberTambang"`
}
