package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"

	"github.com/google/uuid"
)

type PenggunaanRumus string

const (
	PenggunaanRumusPjbb PenggunaanRumus = "PJBB"
	PenggunaanRumusPjab PenggunaanRumus = "PJAB"
)

func (RumusTemplate) TableName() string {
	return "t_rumus_template"
}

type RumusTemplate struct {
	baseapp.BaseModel
	Nama              string              `gorm:"type:varchar(255)" json:"name"`
	Deskripsi         *string             `gorm:"type:text" json:"description"`
	PenggunaanRumus   PenggunaanRumus     `gorm:"type:varchar(255);default:'PJBB'" json:"penggunaanRumus"`
	RumusTemplateBlok []RumusTemplateBlok `gorm:"foreignKey:TRumusTemplateId;constraint:OnDelete:CASCADE" json:"blocks"`
}

func (RumusTemplateBlok) TableName() string {
	return "t_rumus_template_blok"
}

type RumusTemplateBlok struct {
	baseapp.BaseModel
	Type                          string                          `gorm:"type:varchar(255)" json:"type"`
	TRumusTemplateId              *uuid.UUID                      `gorm:"index;type:uuid" json:"rumusTemplateId"`
	RumusTemplate                 *RumusTemplate                  `gorm:"foreignKey:TRumusTemplateId" json:"rumusTemplate"`
	NamaRumusBlockCode            *string                         `gorm:"type:varchar(255)" json:"namaRumusBlockCode"`
	NamaRumusBlock                *ConfigDataInfo                 `gorm:"foreignKey:NamaRumusBlockCode;references:Code;" json:"namaRumusBlock"`
	Deskripsi                     *string                         `gorm:"type:text" json:"description"`
	VariableItem                  json.RawMessage                 `gorm:"type:jsonb" json:"variables"`
	RumusTemplateBlokItem         []RumusTemplateBlokItem         `gorm:"foreignKey:TRumusTemplateBlokId;constraint:OnDelete:CASCADE" json:"items"`
	TPjbbKualitasTemplateId       *uuid.UUID                      `gorm:"index;type:uuid" json:"pjbbKualitasTemplateId"`
	PjbbKualitasTemplate          *Pjbb_kualitas_template         `gorm:"foreignKey:TPjbbKualitasTemplateId" json:"pjbbKualitasTemplate"`
	RumusTemplateBlokKualitasItem []RumusTemplateBlokKualitasItem `gorm:"foreignKey:TRumusTemplateBlokId;constraint:OnDelete:CASCADE" json:"kualitasItems"`
	OrderNumber                   int                             `json:"orderNumber"`
}

func (RumusTemplateBlokItem) TableName() string {
	return "t_rumus_template_blok_item"
}

type RumusTemplateBlokItem struct {
	baseapp.BaseModel
	TRumusTemplateId               *uuid.UUID                       `gorm:"index;type:uuid" json:"rumusTemplateId"`
	RumusTemplate                  *RumusTemplate                   `gorm:"foreignKey:TRumusTemplateId" json:"rumusTemplate"`
	TRumusTemplateBlokId           *uuid.UUID                       `gorm:"index;type:uuid" json:"rumusTemplateBlokId"`
	RumusTemplateBlok              *RumusTemplateBlok               `gorm:"foreignKey:TRumusTemplateBlokId" json:"rumusTemplateBlok"`
	TUraianId                      *uuid.UUID                       `gorm:"type:uuid" json:"uraianId"`
	Uraian                         *Uraian                          `gorm:"foreignKey:TUraianId" json:"uraian"`
	NamaVariable                   string                           `gorm:"type:varchar(255)" json:"variable"`
	NilaiAwal                      *string                          `gorm:"type:varchar(255)" json:"nilaiAwal"`
	RumusTemplateBlokItemCondition []RumusTemplateBlokItemCondition `gorm:"foreignKey:TRumusTemplateBlokItemId;constraint:OnDelete:CASCADE" json:"conditions"`
	Hasil                          *float64                         `gorm:"type:float" json:"hasil"`
	TampilkanDiPropose             *bool                            `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	SatuanHasil                    *string                          `gorm:"type:varchar(255)" json:"satuanHasil"`
	OrderNumber                    int                              `json:"orderNumber"`
}

func (RumusTemplateBlokItemCondition) TableName() string {
	return "t_rumus_template_blok_item_condition"
}

type RumusTemplateBlokItemCondition struct {
	baseapp.BaseModel
	TRumusTemplateId         *uuid.UUID             `gorm:"index;type:uuid" json:"rumusTemplateId"`
	RumusTemplate            *RumusTemplate         `gorm:"foreignKey:TRumusTemplateId" json:"rumusTemplate"`
	TRumusTemplateBlokItemId *uuid.UUID             `gorm:"index;type:uuid" json:"rumusTemplateBlokItemId"`
	RumusTemplateBlokItem    *RumusTemplateBlokItem `gorm:"foreignKey:TRumusTemplateBlokItemId" json:"rumusTemplateBlokItem"`
	Rumus                    *string                `gorm:"type:text" json:"rumus"`
	VariableItemValue        json.RawMessage        `gorm:"type:jsonb" json:"variables"`
	OrderNumber              int                    `json:"orderNumber"`
}

func (RumusTemplateBlokKualitasItem) TableName() string {
	return "t_rumus_template_blok_kualitas_item"
}

type RumusTemplateBlokKualitasItem struct {
	baseapp.BaseModel
	TRumusTemplateId     *uuid.UUID                `gorm:"index;type:uuid" json:"rumusTemplateId"`
	RumusTemplate        *RumusTemplate            `gorm:"foreignKey:TRumusTemplateId" json:"rumusTemplate"`
	TRumusTemplateBlokId *uuid.UUID                `gorm:"index;type:uuid" json:"rumusTemplateBlokId"`
	RumusTemplateBlok    *RumusTemplateBlok        `gorm:"foreignKey:TRumusTemplateBlokId" json:"rumusTemplateBlok"`
	TipeKualitas         Pjbb_kualitasItemType     `gorm:"type:varchar(255);default:'none'" json:"tipeKualitas"`
	TUraianId            uuid.UUID                 `gorm:"type:uuid" json:"uraianId"`
	Uraian               *Uraian                   `gorm:"foreignKey:TUraianId" json:"uraian"`
	BasisCode            *string                   `gorm:"type:varchar(255)" json:"basisCode"`
	Basis                *ConfigDataInfo           `gorm:"foreignKey:BasisCode;references:Code;" json:"basis"`
	Tipikal              *string                   `gorm:"type:varchar(255)" json:"tipikal"`
	Penolakan            *string                   `gorm:"type:varchar(255)" json:"penolakan"`
	Simbol               *string                   `gorm:"type:varchar(255)" json:"simbol"`
	BatasBawah           *string                   `gorm:"type:varchar(255)" json:"batasBawah"`
	BatasAtas            *string                   `gorm:"type:varchar(255)" json:"batasAtas"`
	TipeData             Pjbb_kualitasItemDataType `gorm:"type:varchar(255);default:'Num'" json:"tipeData"`
	TurunHarga           *string                   `gorm:"type:varchar(255)" json:"turunHarga"`
	HasilUji             *string                   `gorm:"type:varchar(255)" json:"hasilUji"`
	Denda                *string                   `gorm:"type:varchar(255)" json:"denda"`
	Rumus                *string                   `gorm:"type:text" json:"rumus"`
	Keterangan           *string                   `gorm:"type:varchar(255)" json:"keterangan"`
	Variable             *string                   `gorm:"type:varchar(255)" json:"variable"`
	Hasil                *float64                  `gorm:"type:float" json:"hasil"`
	TampilkanDiPropose   *bool                     `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	OrderNumber          int                       `json:"orderNumber"`
}
