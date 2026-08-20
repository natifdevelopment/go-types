package types

import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (UploadMasterInvoice) TableName() string {
	return "t_upload_master_invoice"
}

type UploadMasterInvoice struct {
	baseapp.BaseModel
	TOrganizationId  *uuid.UUID    `gorm:"type:uuid" json:"organizationId"`
	Organization     *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
	NoNpwp           string        `gorm:"type:varchar(255)" json:"noNpwp"`
	TglTerdaftarNpwp *DateField    `gorm:"type:date" json:"tglTerdaftarNpwp"`
	TDokNpwpId       *uuid.UUID    `gorm:"type:uuid" json:"dokNpwpId"`
	DokNpwp          *Media        `gorm:"foreignKey:TDokNpwpId" json:"dokNpwp"`
	NoPkp            string        `gorm:"type:varchar(255)" json:"noPkp"`
	TDokPkpId        *uuid.UUID    `gorm:"type:uuid" json:"dokPkpId"`
	DokPkp           *Media        `gorm:"foreignKey:TDokPkpId" json:"dokPkp"`
	TglSkbAwal       *DateField    `gorm:"type:date" json:"tglSkbAwal"`
	TglSkbAkhir      *DateField    `gorm:"type:date" json:"tglSkbAkhir"`
	TDokSkbId        *uuid.UUID    `gorm:"type:uuid" json:"dokSkbId"`
	DokSkb           *Media        `gorm:"foreignKey:TDokSkbId" json:"dokSkb"`
	TglSptpAwal      *DateField    `gorm:"type:date" json:"tglSptpAwal"`
	TglSptpAkhir     *DateField    `gorm:"type:date" json:"tglSptpAkhir"`
	TDokSptpId       *uuid.UUID    `gorm:"type:uuid" json:"dokSptpId"`
	DokSptp          *Media        `gorm:"foreignKey:TDokSptpId" json:"dokSptp"`
	StatusInvoice    bool          `gorm:"type:bool" json:"statusInvoice"`
}
