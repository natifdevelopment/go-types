package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (SuratPengantarTagihan) TableName() string {
	return "t_surat_pengantar_tagihan"
}

type SuratPengantarTagihan struct {
	baseapp.BaseModel
	TOrganizationId         *uuid.UUID            `gorm:"type:uuid" json:"organizationId"`
	Organization            *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPermohonanPembayaranId *uuid.UUID            `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
	TDokSptId               *uuid.UUID            `gorm:"type:uuid" json:"dokSptId"`
	DokSpt                  *Media                `gorm:"foreignKey:TDokSptId" json:"dokSpt"`
	MembuatDokSpt           bool                  `gorm:"type:bool;default:true" json:"membuatDokSpt"`
	DokSptError             bool                  `gorm:"type:bool;default:false" json:"dokSptError"`
	TDokSptUploadId         *uuid.UUID            `gorm:"type:uuid" json:"dokSptUploadId"`
	DokSptUpload            *Media                `gorm:"foreignKey:TDokSptUploadId" json:"dokSptUpload"`
	TotalTagihan            float64               `gorm:"type:float" json:"totalTagihan"`
	JmlInvoice              int32                 `gorm:"type:integer" json:"JmlInvoice"`
	Spb                     *string               `gorm:"type:varchar(255)" json:"spb"`
	TglTerimaDokTagihan     *DateField      `gorm:"type:date" json:"tglTerimaDokTagihan"`
	NoSpt                   *string               `gorm:"type:varchar(255)" json:"noSpt"`
	TglSpt                  *DateField      `gorm:"type:date" json:"tglSpt"`
	TPejabatAmsId           *uuid.UUID            `gorm:"type:uuid" json:"pejabatAmsId"`
	PejabatAms              *PejabatAms           `gorm:"foreignKey:TPejabatAmsId" json:"pejabatAms"`

	// SAP
	TIntegrasiReqId  *uuid.UUID       `gorm:"type:uuid" json:"integrasiReqId"`
	IntegrasiReq     *IntegrasiStatus `gorm:"foreignKey:TIntegrasiReqId" json:"integrasiReq"`
	TIntegrasiRespId *uuid.UUID       `gorm:"type:uuid" json:"integrasiRespId"`
	IntegrasiResp    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiRespId" json:"integrasiResp"`
	SapStatus        ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapStatus"`
	SapInvNo         *string          `gorm:"type:varchar(20)" json:"sapInvNo"`
	SapInvYear       *int             `gorm:"type:int" json:"sapInvYear"`
}
