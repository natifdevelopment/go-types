package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (NotaDinas) TableName() string {
	return "t_nota_dinas"
}

type NotaDinas struct {
	baseapp.BaseModel
	NoNotaDinas     *string          `gorm:"type:varchar(255)" json:"noNotaDinas"`
	TPejabatAmsId   *uuid.UUID       `gorm:"type:uuid" json:"pejabatAmsId"`
	PejabatAms      *PejabatAms      `gorm:"foreignKey:TPejabatAmsId" json:"pejabatAms"`
	NotaDinasItem   []NotaDinasItem  `gorm:"foreignKey:TNotaDinasId" json:"items"`
	TDokNotaDinasId *uuid.UUID       `gorm:"type:uuid" json:"dokNotaDinasId"`
	DokNotaDinas    *Media           `gorm:"foreignKey:TDokNotaDinasId" json:"dokNotaDinas"`
	TglNotaDinas    *DateField `gorm:"type:date" json:"tglNotaDinas"`

	// Dokumen hasil generated
	TLembarVerifikasiId *uuid.UUID `gorm:"type:uuid" json:"lembarVerifikasiId"`
	LembarVerifikasi    *Media     `gorm:"foreignKey:TLembarVerifikasiId" json:"lembarVerifikasi"`
	TDokRptVendorId     *uuid.UUID `gorm:"type:uuid" json:"dokRptVendorId"`
	DokRptVendor        *Media     `gorm:"foreignKey:TDokRptVendorId" json:"dokRptVendor"`
	TDokRptId           *uuid.UUID `gorm:"type:uuid" json:"dokRptId"`
	DokRpt              *Media     `gorm:"foreignKey:TDokRptId" json:"dokRpt"`

	// Tanggal di isi ketika semua approval sudah approve
	// Akan null jika tidak ada proses approval
	TglVerifikasi *DateTimeField `gorm:"type:timestamp" json:"tglVerifikasi"`
	MyApproval    *DokApproval         `gorm:"-" json:"myApproval"`

	// SAP Integration
	TIntegrasiReqId      *uuid.UUID       `gorm:"type:uuid" json:"integrasiReqId"`
	IntegrasiReq         *IntegrasiStatus `gorm:"foreignKey:TIntegrasiReqId" json:"integrasiReq"`
	TIntegrasiRespId     *uuid.UUID       `gorm:"type:uuid" json:"integrasiRespId"`
	IntegrasiResp        *IntegrasiStatus `gorm:"foreignKey:TIntegrasiRespId" json:"integrasiResp"`
	SapStatus            ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapStatus"`
	SapInvDocumentNumber *string          `gorm:"type:varchar(255)" json:"SapInvDocumentNumber"`
	SapPostingDate       *string          `gorm:"type:varchar(255)" json:"SapPostingDate"`

	BboInvNo *string `gorm:"type:varchar(255)" json:"BboInvNo"`

	PaymentDate        *string `gorm:"type:varchar(255)" json:"PaymentDate"`
	PaymentPostingDate *string `gorm:"type:varchar(255)" json:"PaymentPostingDate"`
	PaymentIndicator   *string `gorm:"type:varchar(255)" json:"PaymentIndicator"`
	DocClearing        *string `gorm:"type:varchar(255)" json:"DocClearing"`
	PaidBy             *string `gorm:"type:varchar(255)" json:"PaidBy"`
}

func (NotaDinasItem) TableName() string {
	return "t_nota_dinas_item"
}

type NotaDinasItem struct {
	baseapp.BaseModel
	TNotaDinasId            *uuid.UUID            `gorm:"type:uuid" json:"notaDinasId"`
	NotaDinas               *NotaDinas            `gorm:"foreignKey:TNotaDinasId;constraint:OnDelete:CASCADE" json:"notaDinas"`
	TPermohonanPembayaranId *uuid.UUID            `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
}
