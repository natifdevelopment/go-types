package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"

	"github.com/google/uuid"
)

type TipePjbb string

const (
	TipePjbbAmandemen TipePjbb = "AMANDEMEN"
	TipePjbbUtama     TipePjbb = "KONTRAK_UTAMA"
)

func (PjbbRumus) TableName() string {
	return "t_pjbb_rumus"
}

type PjbbRumus struct {
	baseapp.BaseModel
	TPjbbId               *uuid.UUID           `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb                  *Pjbb                `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId      *uuid.UUID           `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen         *PjbbAmandemen       `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TipePjbb              TipePjbb             `gorm:"type:varchar(255);default:'KONTRAK_UTAMA'" json:"tipePjbb"`
	TPemasokId            *uuid.UUID           `gorm:"index;type:uuid" json:"pemasokId"`
	Pemasok               *Organization        `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	SkemaKontrakCode      *JenisPengiriman     `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak          *ConfigDataInfo      `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	JenisTransportasiCode *string              `json:"jenisTransportasiCode"`
	JenisTransportasi     *ConfigDataInfo      `gorm:"foreignKey:JenisTransportasiCode;references:Code;" json:"jenisTransportasi"`
	RumusBlok             []PjbbRumusBlok      `gorm:"foreignKey:TPjbbRumusId;constraint:OnDelete:CASCADE" json:"rumusBlocks"`
	Pembangkit            []PjbbRumusPembankit `gorm:"foreignKey:TPjbbRumusId;constraint:OnDelete:CASCADE" json:"pembangkit"`
}

func (PjbbRumusPembankit) TableName() string {
	return "t_pjbb_rumus_pembangkit"
}

type PjbbRumusPembankit struct {
	baseapp.BaseModel
	TPjbbRumusId  *uuid.UUID    `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus     *PjbbRumus    `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	TPembangkitId *uuid.UUID    `gorm:"index;type:uuid" json:"pembangkitId"`
	Pembangkit    *Organization `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
}

// Rumus Block
// -------------------------
type TipeRumusBlok string

const (
	TipeRumusBlokRumus    TipeRumusBlok = "PERHITUNGAN"
	TipeRumusBlokKualitas TipeRumusBlok = "KUALITAS"
)

func (PjbbRumusBlok) TableName() string {
	return "t_pjbb_rumus_blok"
}

type PjbbRumusBlok struct {
	baseapp.BaseModel
	Type                      TipeRumusBlok               `gorm:"type:varchar(255)" json:"type"`
	TPjbbRumusId              *uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus                 *PjbbRumus                  `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	NamaRumusBlockCode        *string                     `gorm:"type:varchar(255)" json:"namaRumusBlockCode"`
	NamaRumusBlock            *ConfigDataInfo             `gorm:"foreignKey:NamaRumusBlockCode;references:Code;" json:"namaRumusBlock"`
	Deskripsi                 *string                     `gorm:"type:text" json:"description"`
	VariableItem              json.RawMessage             `gorm:"type:jsonb" json:"variables"`
	PjbbRumusBlokItem         []PjbbRumusBlokItem         `gorm:"foreignKey:TPjbbRumusBlokId;constraint:OnDelete:CASCADE" json:"items"`
	TPjbbKualitasTemplateId   *uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbKualitasTemplateId"`
	PjbbKualitasTemplate      *Pjbb_kualitas_template     `gorm:"foreignKey:TPjbbKualitasTemplateId" json:"pjbbKualitasTemplate"`
	PjbbRumusBlokPelMuatSpek  []PjbbRumusBlokPelMuatSpek  `gorm:"foreignKey:TPjbbRumusBlokId;constraint:OnDelete:CASCADE" json:"pelMuatSpeks"`
	PjbbRumusBlokKualitasItem []PjbbRumusBlokKualitasItem `gorm:"foreignKey:TPjbbRumusBlokId;constraint:OnDelete:CASCADE" json:"kualitasItems"`
	OrderNumber               int                         `json:"orderNumber"`
}

// Rumus Perhitungan
// -------------------------

func (PjbbRumusBlokItemCondition) TableName() string {
	return "t_pjbb_rumus_blok_item_condition"
}

type PjbbRumusBlokItemCondition struct {
	baseapp.BaseModel
	TPjbbRumusId         *uuid.UUID         `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus            *PjbbRumus         `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	TPjbbRumusBlokItemId *uuid.UUID         `gorm:"index;type:uuid" json:"pjbbRumusBlokItemId"`
	PjbbRumusBlokItem    *PjbbRumusBlokItem `gorm:"foreignKey:TPjbbRumusBlokItemId" json:"pjbbRumusBlokItem"`
	Rumus                *string            `gorm:"type:text" json:"rumus"`
	VariableItemValue    json.RawMessage    `gorm:"type:jsonb" json:"variables"`
	OrderNumber          int                `json:"orderNumber"`
}

func (PjbbRumusBlokItem) TableName() string {
	return "t_pjbb_rumus_blok_item"
}

type PjbbRumusBlokItem struct {
	baseapp.BaseModel
	TPjbbRumusId               *uuid.UUID                   `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus                  *PjbbRumus                   `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	TPjbbRumusBlokId           *uuid.UUID                   `gorm:"index;type:uuid" json:"pjbbRumusBlokId"`
	PjbbRumusBlok              *PjbbRumusBlok               `gorm:"foreignKey:TPjbbRumusBlokId" json:"pjbbRumusBlok"`
	TUraianId                  *uuid.UUID                   `gorm:"type:uuid" json:"uraianId"`
	Uraian                     *Uraian                      `gorm:"foreignKey:TUraianId" json:"uraian"`
	NamaVariable               string                       `gorm:"type:varchar(255)" json:"variable"`
	NilaiAwal                  *string                      `gorm:"type:varchar(255)" json:"nilaiAwal"`
	Hasil                      *float64                     `gorm:"type:float" json:"hasil"`
	PjbbRumusBlokItemCondition []PjbbRumusBlokItemCondition `gorm:"foreignKey:TPjbbRumusBlokItemId;constraint:OnDelete:CASCADE" json:"conditions"`
	TampilkanDiPropose         *bool                        `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	SatuanHasil                *string                      `gorm:"type:varchar(255)" json:"satuanHasil"`
	OrderNumber                int                          `json:"orderNumber"`
	Catatan                    *string                      `gorm:"type:varchar(255)" json:"catatan"` // Catatan perbaikan
}

// Rumus Kualitas
// -------------------------

func (PjbbRumusBlokPelMuatSpek) TableName() string {
	return "t_pjbb_rumus_blok_pelmuat_spek"
}

type PjbbRumusBlokPelMuatSpek struct {
	baseapp.BaseModel
	TPjbbRumusId              *uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus                 *PjbbRumus                  `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	TPjbbRumusBlokId          *uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbRumusBlokId"`
	PjbbRumusBlok             *PjbbRumusBlok              `gorm:"foreignKey:TPjbbRumusBlokId" json:"pjbbRumusBlok"`
	TPelabuhanMuatId          *uuid.UUID                  `gorm:"index;type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat             *PelabuhanMuat              `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId                   *uuid.UUID                  `gorm:"index;type:uuid" json:"spekId"`
	Spek                      *Spek                       `gorm:"foreignKey:TSpekId" json:"spek"`
	PjbbRumusBlokKualitasItem []PjbbRumusBlokKualitasItem `gorm:"foreignKey:TPjbbRumusBlokPelMuatSpekId;constraint:OnDelete:CASCADE" json:"kualitasItems"`
	OrderNumber               int                         `json:"orderNumber"`
}

func (PjbbRumusBlokKualitasItem) TableName() string {
	return "t_pjbb_rumus_blok_kualitas_item"
}

type PjbbRumusBlokKualitasItem struct {
	baseapp.BaseModel
	TPjbbRumusId                *uuid.UUID                `gorm:"index;type:uuid" json:"pjbbRumusId"`
	PjbbRumus                   *PjbbRumus                `gorm:"foreignKey:TPjbbRumusId" json:"pjbbRumus"`
	TPjbbRumusBlokId            *uuid.UUID                `gorm:"index;type:uuid" json:"pjbbRumusBlokId"`
	PjbbRumusBlok               *PjbbRumusBlok            `gorm:"foreignKey:TPjbbRumusBlokId" json:"pjbbRumusBlok"`
	TPjbbRumusBlokPelMuatSpekId *uuid.UUID                `gorm:"index;type:uuid" json:"pjbbRumusBlokPelMuatSpekId"`
	PjbbRumusBlokPelMuatSpek    *PjbbRumusBlokPelMuatSpek `gorm:"foreignKey:TPjbbRumusBlokPelMuatSpekId" json:"pjbbRumusBlokPelMuatSpek"`
	TipeKualitas                Pjbb_kualitasItemType     `gorm:"type:varchar(255);default:'none'" json:"tipeKualitas"`
	TUraianId                   uuid.UUID                 `gorm:"type:uuid" json:"uraianId"`
	Uraian                      *Uraian                   `gorm:"foreignKey:TUraianId" json:"uraian"`
	BasisCode                   *string                   `gorm:"type:varchar(255)" json:"basisCode"`
	Basis                       *ConfigDataInfo           `gorm:"foreignKey:BasisCode;references:Code;" json:"basis"`
	Tipikal                     *string                   `gorm:"type:varchar(255)" json:"tipikal"`
	Penolakan                   *string                   `gorm:"type:varchar(255)" json:"penolakan"`
	Simbol                      *string                   `gorm:"type:varchar(255)" json:"simbol"`
	BatasBawah                  *string                   `gorm:"type:varchar(255)" json:"batasBawah"`
	BatasAtas                   *string                   `gorm:"type:varchar(255)" json:"batasAtas"`
	TipeData                    Pjbb_kualitasItemDataType `gorm:"type:varchar(255);default:'Num'" json:"tipeData"`
	TurunHarga                  *string                   `gorm:"type:varchar(255)" json:"turunHarga"`
	HasilUji                    *string                   `gorm:"type:varchar(255)" json:"hasilUji"`
	Denda                       *string                   `gorm:"type:varchar(255)" json:"denda"`
	Rumus                       *string                   `gorm:"type:text" json:"rumus"`
	Keterangan                  *string                   `gorm:"type:varchar(255)" json:"keterangan"`
	Variable                    *string                   `gorm:"type:varchar(255)" json:"variable"`
	Hasil                       *float64                  `gorm:"type:float" json:"hasil"`
	TampilkanDiPropose          *bool                     `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	OrderNumber                 int                       `json:"orderNumber"`
	Catatan                     *string                   `gorm:"type:varchar(255)" json:"catatan"` // Catatan perbaikan

	// @deprecated, ganti jadi jangan relasi kesini
	// CoaCowItems                 []CoaCowCifItem           `gorm:"foreignKey:TKualitasId;constraint:OnDelete:CASCADE" json:"-"`
}
