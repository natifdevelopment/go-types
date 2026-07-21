package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (UploadInvoicePenagihan) TableName() string {
	return "t_upload_invoice_penagihan"
}

type UploadInvoicePenagihan struct {
	baseapp.BaseModel
	TProposePerhitunganId      *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan         *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	TPembangkitId              *uuid.UUID          `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                 *Organization       `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterNoRekeningId        *uuid.UUID          `gorm:"type:uuid" json:"masterNoRekeningId"`
	MasterNoRekening           *MasterNoRekening   `gorm:"foreignKey:TMasterNoRekeningId" json:"masterNoRekening"`
	NoPengiriman               string              `gorm:"type:varchar(255)" json:"noPengiriman"`
	TanggalVerifikasiKit       *DateField    `gorm:"type:date" json:"tanggalVerifikasiKit"`
	TotalInvoice               *float64            `gorm:"type:float" json:"totalInvoice"`
	NoInvoice                  string              `gorm:"type:varchar(255)" json:"noInvoice"`    // Invoice input by user
	NoInvoiceBbo               *string             `gorm:"type:varchar(255)" json:"noInvoiceBbo"` // Invoice input by user + no pengiriman
	TanggalInvoice             *DateField    `gorm:"type:date" json:"tanggalInvoice"`
	TDokInvoiceId              *uuid.UUID          `gorm:"type:uuid" json:"dokInvoiceId"`
	DokInvoice                 *Media              `gorm:"foreignKey:TDokInvoiceId" json:"dokInvoice"`
	DibutuhkanPerbaikanInvoice bool                `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanInvoice"`
	CatatanPerbaikanInvoice    *string             `gorm:"text" json:"catatanPerbaikanInvoice"`
	NoBuktiRoyalti             string              `gorm:"type:varchar(255)" json:"noBuktiRoyalti"` // BuktiRoyalti
	TanggalBuktiRoyalti        *DateField    `gorm:"type:date" json:"tanggalBuktiRoyalti"`
	// TDokBuktiRoyaltiId               *uuid.UUID           `gorm:"type:uuid" json:"dokBuktiRoyaltiId"`
	// DokBuktiRoyalti                  *Media               `gorm:"foreignKey:TDokBuktiRoyaltiId" json:"dokBuktiRoyalti"`
	DibutuhkanPerbaikanRoyalti       bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanRoyalti"`
	CatatanPerbaikanRoyalti          *string          `gorm:"text" json:"catatanPerbaikanRoyalti"`
	NoKwitansi                       string           `gorm:"type:varchar(255)" json:"noKwitansi"` // Kwitansi
	TanggalKwitansi                  *DateField `gorm:"type:date" json:"tanggalKwitansi"`
	TDokKwitansiId                   *uuid.UUID       `gorm:"type:uuid" json:"dokKwitansiId"`
	DokKwitansi                      *Media           `gorm:"foreignKey:TDokKwitansiId" json:"dokKwitansi"`
	DibutuhkanPerbaikanKwitansi      bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanKwitansi"`
	CatatanPerbaikanKwitansi         *string          `gorm:"text" json:"catatanPerbaikanKwitansi"`
	NoRefPembayaran                  string           `gorm:"type:varchar(255)" json:"noRefPembayaran"` // RefPembayaran
	TanggalRefPembayaran             *DateField `gorm:"type:date" json:"tanggalRefPembayaran"`
	TDokRefPembayaranId              *uuid.UUID       `gorm:"type:uuid" json:"dokRefPembayaranId"`
	DokRefPembayaran                 *Media           `gorm:"foreignKey:TDokRefPembayaranId" json:"dokRefPembayaran"`
	DibutuhkanPerbaikanRefPembayaran bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanRefPembayaran"`
	CatatanPerbaikanRefPembayaran    *string          `gorm:"text" json:"catatanPerbaikanRefPembayaran"`
	NoFakturPajak                    string           `gorm:"type:varchar(255)" json:"noFakturPajak"` // FakturPajak
	TanggalFakturPajak               *DateField `gorm:"type:date" json:"tanggalFakturPajak"`
	TDokFakturPajakId                *uuid.UUID       `gorm:"type:uuid" json:"dokFakturPajakId"`
	DokFakturPajak                   *Media           `gorm:"foreignKey:TDokFakturPajakId" json:"dokFakturPajak"`
	DibutuhkanPerbaikanFakturPajak   bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFakturPajak"`
	CatatanPerbaikanFakturPajak      *string          `gorm:"text" json:"catatanPerbaikanFakturPajak"`
	Consent                          *bool            `gorm:"type:boolean" json:"consent"`

	// Simple relation
	Propose *ProposeSimple `gorm:"foreignKey:TProposePerhitunganId" json:"propose"`

	SumberTambang []UploadInvoicePenagihanSumberTambang `gorm:"foreignKey:TUploadInvoicePenagihanId" json:"sumberTambang"`
}

func (UploadInvoicePenagihanSumberTambang) TableName() string {
	return "t_upload_invoice_penagihan_sumber_tambang"
}

type UploadInvoicePenagihanSumberTambang struct {
	baseapp.BaseModel
	TUploadInvoicePenagihanId *uuid.UUID             `gorm:"index;type:uuid" json:"uploadInvoicePenagihanId"`
	UploadInvoicePenagihan    UploadInvoicePenagihan `gorm:"foreignKey:TUploadInvoicePenagihanId" json:"uploadInvoicePenagihan"`
	TSumberTambangId          *uuid.UUID             `gorm:"index;type:uuid" json:"sumberTambangId"`
	SumberTambang             SumberTambang          `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	TDokBuktiRoyaltiId        *uuid.UUID             `gorm:"index;type:uuid" json:"dokBuktiRoyaltiId"`
	DokBuktiRoyalti           *Media                 `gorm:"foreignKey:TDokBuktiRoyaltiId" json:"dokBuktiRoyalti"`
}
