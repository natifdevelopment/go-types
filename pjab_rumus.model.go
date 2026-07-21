package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"

	"github.com/google/uuid"
)

type TipePjab string

const (
	TipePjabAmandemen TipePjab = "AMANDEMEN"
	TipePjabUtama     TipePjab = "KONTRAK_UTAMA"
)

func (PjabRumus) TableName() string {
	return "t_pjab_rumus"
}

type PjabRumus struct {
	baseapp.BaseModel
	TPjabId               *uuid.UUID           `gorm:"index;type:uuid" json:"pjabId"`
	Pjab                  *Pjab                `gorm:"foreignKey:TPjabId" json:"pjab"`
	TPjabAmandemenId      *uuid.UUID           `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen         *PjabAmandemen       `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TipePjab              TipePjab             `gorm:"type:varchar(255);default:'KONTRAK_UTAMA'" json:"tipePjab"`
	TTransportirId        *uuid.UUID           `gorm:"type:uuid" json:"transportirId"`
	Transportir           *Organization        `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TRumusTemplateId      *uuid.UUID           `gorm:"index;type:uuid" json:"rumusTemplateId"`
	RumusTemplate         *RumusTemplate       `gorm:"foreignKey:TRumusTemplateId" json:"rumusTemplate"`
	JenisTransportasiCode *string              `json:"jenisTransportasiCode"`
	JenisTransportasi     *ConfigDataInfo      `gorm:"foreignKey:JenisTransportasiCode;references:Code;" json:"jenisTransportasi"`
	RumusBlok             []PjabRumusBlok      `gorm:"foreignKey:TPjabRumusId;constraint:OnDelete:CASCADE" json:"rumusBlocks"`
	Pembangkit            []PjabRumusPembankit `gorm:"foreignKey:TPjabRumusId;constraint:OnDelete:CASCADE" json:"pembangkit"`
}

func (PjabRumusPembankit) TableName() string {
	return "t_pjab_rumus_pembangkit"
}

type PjabRumusPembankit struct {
	baseapp.BaseModel
	TPjabRumusId  *uuid.UUID    `gorm:"index;type:uuid" json:"pjabRumusId"`
	PjabRumus     *PjabRumus    `gorm:"foreignKey:TPjabRumusId" json:"pjabRumus"`
	TPembangkitId *uuid.UUID    `gorm:"index;type:uuid" json:"pembangkitId"`
	Pembangkit    *Organization `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
}

func (PjabRumusBlok) TableName() string {
	return "t_pjab_rumus_blok"
}

type PjabRumusBlok struct {
	baseapp.BaseModel
	Type                    string                  `gorm:"type:varchar(255)" json:"type"`
	TPjabRumusId            *uuid.UUID              `gorm:"index;type:uuid" json:"pjabRumusId"`
	PjabRumus               *PjabRumus              `gorm:"foreignKey:TPjabRumusId" json:"pjabRumus"`
	NamaRumusBlockCode      *string                 `gorm:"type:varchar(255)" json:"namaRumusBlockCode"`
	NamaRumusBlock          *ConfigDataInfo         `gorm:"foreignKey:NamaRumusBlockCode;references:Code;" json:"namaRumusBlock"`
	Deskripsi               *string                 `gorm:"type:text" json:"description"`
	VariableItem            json.RawMessage         `gorm:"type:jsonb" json:"variables"`
	PjabRumusBlokItem       []PjabRumusBlokItem     `gorm:"foreignKey:TPjabRumusBlokId;constraint:OnDelete:CASCADE" json:"items"`
	TPjabKualitasTemplateId *uuid.UUID              `gorm:"index;type:uuid" json:"pjabKualitasTemplateId"`
	PjabKualitasTemplate    *Pjbb_kualitas_template `gorm:"foreignKey:TPjabKualitasTemplateId" json:"pjabKualitasTemplate"`
	OrderNumber             int                     `json:"orderNumber"`
}

func (PjabRumusBlokItem) TableName() string {
	return "t_pjab_rumus_blok_item"
}

type PjabRumusBlokItem struct {
	baseapp.BaseModel
	TPjabRumusId               *uuid.UUID                   `gorm:"index;type:uuid" json:"pjabRumusId"`
	PjabRumus                  *PjabRumus                   `gorm:"foreignKey:TPjabRumusId" json:"pjabRumus"`
	TPjabRumusBlokId           *uuid.UUID                   `gorm:"index;type:uuid" json:"pjabRumusBlokId"`
	PjabRumusBlok              *PjabRumusBlok               `gorm:"foreignKey:TPjabRumusBlokId" json:"pjabRumusBlok"`
	TUraianId                  *uuid.UUID                   `gorm:"type:uuid" json:"uraianId"`
	Uraian                     *Uraian                      `gorm:"foreignKey:TUraianId" json:"uraian"`
	NamaVariable               string                       `gorm:"type:varchar(255)" json:"variable"`
	NilaiAwal                  *string                      `gorm:"type:varchar(255)" json:"nilaiAwal"`
	Hasil                      *float64                     `gorm:"type:float" json:"hasil"`
	PjabRumusBlokItemCondition []PjabRumusBlokItemCondition `gorm:"foreignKey:TPjabRumusBlokItemId;constraint:OnDelete:CASCADE" json:"conditions"`
	TampilkanDiPropose         *bool                        `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	SatuanHasil                *string                      `gorm:"type:varchar(255)" json:"satuanHasil"`
	OrderNumber                int                          `json:"orderNumber"`
	Catatan                    *string                      `gorm:"type:varchar(255)" json:"catatan"`
}

func (PjabRumusBlokItemCondition) TableName() string {
	return "t_pjab_rumus_blok_item_condition"
}

type PjabRumusBlokItemCondition struct {
	baseapp.BaseModel
	TPjabRumusId         *uuid.UUID         `gorm:"index;type:uuid" json:"pjabRumusId"`
	PjabRumus            *PjabRumus         `gorm:"foreignKey:TPjabRumusId" json:"pjabRumus"`
	TPjabRumusBlokItemId *uuid.UUID         `gorm:"index;type:uuid" json:"pjabRumusBlokItemId"`
	PjabRumusBlokItem    *PjabRumusBlokItem `gorm:"foreignKey:TPjabRumusBlokItemId" json:"pjabRumusBlokItem"`
	Rumus                *string            `gorm:"type:text" json:"rumus"`
	VariableItemValue    json.RawMessage    `gorm:"type:jsonb" json:"variables"`
	OrderNumber          int                `json:"orderNumber"`
}
