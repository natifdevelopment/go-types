package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type TipeBa string

const (
	TipeBaBast         TipeBa = "BAST"
	TipeBaDenda        TipeBa = "DENDA"
	TipeBaRincianDenda TipeBa = "RINCIAN_DENDA"
	TipeBaBaph         TipeBa = "BAPH"
	TipeBaSpt          TipeBa = "SPT"
)

func (t TipeBa) Label() string {
	switch t {
	case TipeBaBast:
		return "Berita Acara Serah Terima"
	case TipeBaDenda:
		return "Berita Acara Serah Terima Denda"
	case TipeBaRincianDenda:
		return "Berita Acara Serah Terima Rincian Denda"
	case TipeBaBaph:
		return "Berita Acara Perhitungan Harga"
	case TipeBaSpt:
		return "Surat Pengantar Tagihan"
	default:
		return string(t)
	}
}

type Narasi struct {
	baseapp.BaseModel
	OrgAccessLevelCode      string                `gorm:"type:varchar(255)" json:"orgAccessLevelCode"`
	ContractType            TipeKontrakPropose    `gorm:"type:varchar(255)" json:"contractType"`
	TOrganizationId         *uuid.UUID            `gorm:"index;type:uuid" json:"organizationId"`
	Organization            *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPjbbId                 *uuid.UUID            `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb                    *Pjbb                 `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId        *uuid.UUID            `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen           *PjbbAmandemen        `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TPjabId                 *uuid.UUID            `gorm:"index;type:uuid" json:"pjabId"`
	Pjab                    *Pjab                 `gorm:"foreignKey:TPjabId" json:"pjab"`
	TPjabAmandemenId        *uuid.UUID            `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen           *PjabAmandemen        `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TPjbbHargaTriwulanId    *uuid.UUID            `gorm:"index;type:uuid" json:"pjbbHargaTriwulanId"`
	PjbbHargaTriwulan       *PjbbHargaTriwulan    `gorm:"foreignKey:TPjbbHargaTriwulanId" json:"pjbbHargaTriwulan"`
	TPjbbAmdHargaTriwulanId *uuid.UUID            `gorm:"index;type:uuid" json:"pjbbAmdHargaTriwulanId"`
	PjbbAmdHargaTriwulan    *PjbbAmdHargaTriwulan `gorm:"foreignKey:TPjbbAmdHargaTriwulanId" json:"pjbbAmdHargaTriwulan"`
	TipeBa                  TipeBa                `gorm:"type:varchar(255)" json:"tipeBa"`
	MasaBerlakuAwal         DateField       `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir        DateField       `gorm:"type:date" json:"masaBerlakuAkhir"`
	TDokNarasiId            *uuid.UUID            `gorm:"index;type:uuid" json:"dokNarasiId"`
	DokNarasi               *Media                `gorm:"foreignKey:TDokNarasiId" json:"dokNarasi"`
	Narasi                  string                `gorm:"type:text" json:"narasi"`
	NarasiPembangkit        []NarasiPembangkit    `gorm:"foreignKey:TNarasiId;constraint:OnDelete:CASCADE;" json:"pembangkit"`
}

func (Narasi) TableName() string {
	return "t_narasi"
}

type NarasiPembangkit struct {
	baseapp.BaseModel
	TNarasiId       *uuid.UUID    `gorm:"index;type:uuid" json:"narasiId"`
	Narasi          *Narasi       `gorm:"foreignKey:TNarasiId" json:"narasi"`
	TOrganizationId *uuid.UUID    `gorm:"index;type:uuid" json:"organizationId"`
	Organization    *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (NarasiPembangkit) TableName() string {
	return "t_narasi_pembangkit"
}
