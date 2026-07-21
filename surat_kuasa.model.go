package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type TipeSuratKuasa string

const (
	TipeSuratKuasaSk         TipeSuratKuasa = "SURAT_KUASA"
	TipeSuratKuasaPelimpahan TipeSuratKuasa = "PELIMPAHAN_WEWENANG"
)

func (SuratKuasa) TableName() string {
	return "t_surat_kuasa"
}

type SuratKuasa struct {
	baseapp.BaseModel
	TOrganizationId  *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	Organization     *Organization          `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TTransportirId   *uuid.UUID             `gorm:"type:uuid" json:"transportirId"`
	Transportir      *Organization          `gorm:"foreignKey:TTransportirId" json:"transportir"`
	NoSuratKuasa     string                 `gorm:"type:varchar(255);unique;" json:"noSuratKuasa"`
	TglSuratKuasa    *DateField       `gorm:"type:date" json:"tglSuratKuasa"`
	FileSuratKuasaId *uuid.UUID             `gorm:"type:uuid" json:"fileSuratKuasaId"`
	FileSuratKuasa   *Media                 `gorm:"foreignKey:FileSuratKuasaId" json:"fileSuratKuasa"`
	TipeSuratKuasa   TipeSuratKuasa         `gorm:"type:varchar(255)" json:"tipeSuratKuasa"`
	Pembangkit       []SuratKuasaPembangkit `gorm:"foreignKey:TSuratKuasaId; constraint:OnDelete:CASCADE" json:"pembangkit"`
}

func (SuratKuasaPembangkit) TableName() string {
	return "t_surat_kuasa_pembangkit"
}

type SuratKuasaPembangkit struct {
	baseapp.BaseModel
	TSuratKuasaId   uuid.UUID     `gorm:"type:uuid" json:"suratKuasaId"`
	SuratKuasa      SuratKuasa    `gorm:"foreignKey:TSuratKuasaId;" json:"suratKuasa"`
	TOrganizationId *uuid.UUID    `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization `gorm:"foreignKey:TOrganizationId;" json:"organization"`
}
