package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PermohonanPembayaran) TableName() string {
	return "t_permohonan_pembayaran"
}

type PermohonanPembayaran struct {
	baseapp.BaseModel
	NoPenagihan             string                      `gorm:"type:varchar(255);unique;" json:"noPenagihan"`
	TotalTagihan            float64                     `gorm:"type:float" json:"totalTagihan"`
	TglDokPenagihan         *DateField            `gorm:"type:date" json:"tglDokPenagihan"`
	TDokPenagihanId         *uuid.UUID                  `gorm:"type:uuid" json:"dokPenagihanId"`
	DokPenagihan            *Media                      `gorm:"foreignKey:TDokPenagihanId" json:"dokPenagihan"`
	TglKirimDokFisikPln     *DateField            `gorm:"type:date" json:"tglKirimDokFisikPln"`
	SuratPengantarTagihan   []SuratPengantarTagihan     `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"suratPengantarTagihan"`
	NotaDinasItem           []NotaDinasItem             `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"notaDinasItem"`
	UploadDokPembayaranItem []Upload_dok_pembayaranItem `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"uploadDokPembayaranItem"`
	Items                   []PermohonanPembayaranItem  `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"items"`
	PenagihanAkuntansi      []PenagihanAkuntansi        `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"penagihanAkuntansi"`
	PelunasanPenagihan      []PelunasanPenagihanItem    `gorm:"foreignKey:TPermohonanPembayaranId; constraint:OnDelete:CASCADE" json:"pelunasanPenagihan"`

	// Pertama dibuat akan menjadi belum_dikirim,
	// Ketika kirim permohonan pembayaran, maka akan menjadi dalam_proses,
	StatusPenagihan StatusPenagihan `gorm:"type:varchar(255);default:'BELUM_DIKIRIM'" json:"statusPenagihan"` // belum_dikirim/dalam_proses

	// This is new chapter, we need to disconnect the relation
	// So filter, search etc no need to get from jadwal pengiriman
	TOrganizationId  *uuid.UUID            `gorm:"type:uuid" json:"organizationId"` // Maker
	Organization     *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPembangkitId    *uuid.UUID            `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit       *Organization         `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId       *uuid.UUID            `gorm:"type:uuid" json:"pemasokId"`
	Pemasok          *Organization         `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TTransportirId   *uuid.UUID            `gorm:"type:uuid" json:"transportirId"`
	Transportir      *Organization         `gorm:"foreignKey:TTransportirId" json:"transportir"`
	Periode          *YearMonthField `gorm:"type:date" json:"periode"`
	SkemaKontrakCode *JenisPengiriman      `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak     *ConfigDataInfo       `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`

	TipeKontrak TipeKontrakPropose `gorm:"type:varchar(255);" json:"tipeKontrak"`

	TPjbbId          *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb          `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TipePjbb         TipePjbb       `gorm:"type:varchar(255);default:'KONTRAK_UTAMA'" json:"tipePjbb"`

	TPjabId          *uuid.UUID     `gorm:"index;type:uuid" json:"pjabId"`
	Pjab             *Pjab          `gorm:"foreignKey:TPjabId" json:"pjab"`
	TPjabAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen    *PjabAmandemen `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TipePjab         TipePjab       `gorm:"type:varchar(255);default:'KONTRAK_UTAMA'" json:"tipePjab"`

	// @deprecated
	TBundleInvoicePenagihanId *uuid.UUID               `gorm:"type:uuid" json:"bundleInvoicePenagihanId"`
	BundleInvoicePenagihan    *Bundle_invoice_pangihan `gorm:"foreignKey:TBundleInvoicePenagihanId" json:"bundleInvoicePenagihan"`
}

func (PermohonanPembayaranItem) TableName() string {
	return "t_permohonan_pembayaran_item"
}

type PermohonanPembayaranItem struct {
	baseapp.BaseModel
	TPermohonanPembayaranId *uuid.UUID            `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran `gorm:"foreignKey:TPermohonanPembayaranId;constraint:OnDelete:CASCADE" json:"permohonanPembayaran"`
	TProposePerhitunganId   *uuid.UUID            `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan      *ProposePerhitungan   `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
}
