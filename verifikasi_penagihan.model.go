package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

// Ada status belum diproses, dan lengkap
// Belum di proses jika belum input, jika sudah jadi lengkap,
// Sama persis kaya invoice
// Setelah lengkap, akan buat report, rpt_lampVerifikasi.pdf
// Kualitas ambil dari denda_penolakan, jika ga ada penolakan
// Ambil dari penyesuaian, kaya coa cow

type StatusVerifikasiPenagihan string

const (
	StatusVerifikasiPenagihanBelumDiperiksa StatusVerifikasiPenagihan = "BELUM_DIPERIKSA"
	StatusVerifikasiPenagihanSudahLengkap   StatusVerifikasiPenagihan = "SUDAH_LENGKAP"
	StatusVerifikasiPenagihanBelumLengkap   StatusVerifikasiPenagihan = "BELUM_LENGKAP"
	StatusVerifikasiPenagihanDalamProses    StatusVerifikasiPenagihan = "DALAM_PROSES"
)

func (o StatusVerifikasiPenagihan) IsValid() bool {
	return o == StatusVerifikasiPenagihanSudahLengkap || o == StatusVerifikasiPenagihanBelumLengkap ||
		o == StatusVerifikasiPenagihanDalamProses
}

func (VerifikasiPenagihan) TableName() string {
	return "t_verifikasi_penagihan"
}

type VerifikasiPenagihan struct {
	baseapp.BaseModel
	TProposePerhitunganId         *uuid.UUID                `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan            *ProposePerhitungan       `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	JenisPengirimanPembangkitName string                    `gorm:"type:varchar(255)" json:"pembangkitName"`
	TglTerimaDokFisik             *DateField          `gorm:"type:date" json:"tglTerimaDokFisik"`
	TglValidasiDokFisik           *DateField          `gorm:"type:date" json:"tglValidasiDokFisik"`
	StatusVerifikasiPenagihan     StatusVerifikasiPenagihan `gorm:"type:text;default:'BELUM_LENGKAP'" json:"statusVerifikasiPenagihan"`
	Disetujui                     bool                      `gorm:"type:boolean" json:"disetujui"`
	SubmitUlang                   bool                      `gorm:"type:boolean;default:false" json:"submitUlang"`
}

// func (VerifikasiPenagihanTidakSetujuiDokumen) TableName() string {
// 	return "t_verifikasi_penagihan_tidak_setujui_dokumen"
// }

// type VerifikasiPenagihanTidakSetujuiDokumen struct {
// 	baseapp.BaseModel
// 	TVerifikasiPenagihanId *uuid.UUID           `gorm:"type:uuid" json:"verifikasiPenagihanId"`
// 	VerifikasiPenagihan    *VerifikasiPenagihan `gorm:"foreignKey:TVerifikasiPenagihanId" json:"verifikasiPenagihan"`
// 	TProposePerhitunganId  *uuid.UUID           `gorm:"type:uuid" json:"proposePerhitunganId"`
// 	ProposePerhitungan     *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
// 	TMediaId               *uuid.UUID           `gorm:"type:uuid" json:"mediaId"`
// 	Media                  *Media               `gorm:"foreignKey:TMediaId" json:"media"`
// 	TipeDokumen            string               `gorm:"type:varchar(255)" json:"tipeDokumen"`
// 	Keterangan             *string              `gorm:"type:varchar(255)" json:"keterangan"`
// 	SudahDiperbaiki        bool                 `gorm:"type:boolean;default:false" json:"sudahDiperbaiki"`
// 	Diperiksa              bool                 `gorm:"type:boolean;default:false" json:"diperiksa"`
// }

// func (VerifikasiPenagihanSetujuiInvoice) TableName() string {
// 	return "t_verifikasi_penagihan_setujui_invoice"
// }

// type VerifikasiPenagihanSetujuiInvoice struct {
// 	baseapp.BaseModel
// 	TVerifikasiPenagihanId    *uuid.UUID                `gorm:"type:uuid" json:"verifikasiPenagihanId"`
// 	VerifikasiPenagihan       *VerifikasiPenagihan      `gorm:"foreignKey:TVerifikasiPenagihanId" json:"verifikasiPenagihan"`
// 	TProposePerhitunganId     *uuid.UUID                `gorm:"type:uuid" json:"proposePerhitunganId"`
// 	ProposePerhitungan        *ProposePerhitungan      `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
// 	TUploadInvoicePenagihanId *uuid.UUID                `gorm:"type:uuid" json:"uploadInvoicePenagihanId"`
// 	UploadInvoicePenagihan    *UploadInvoicePenagihan `gorm:"foreignKey:TUploadInvoicePenagihanId" json:"uploadInvoicePenagihan"`
// 	TSuratPengantarTagihanId  *uuid.UUID                `gorm:"type:uuid" json:"suratPengantarTagihanId"`
// 	SuratPengantarTagihan     *SuratPengantarTagihan    `gorm:"foreignKey:TSuratPengantarTagihanId" json:"suratPengantarTagihan"`
// 	NoFakturPajak             *string                   `gorm:"type:varchar(255)" json:"noFakturPajak"`
// 	NoInvoiceSap              *string                   `gorm:"type:varchar(255)" json:"noInvoiceSap"`
// 	NoDendaSap                *string                   `gorm:"type:varchar(255)" json:"noDendaSap"`
// 	TglDokDikirimOlehPemasok  *DateField          `gorm:"type:date" json:"tglDokDikirimOlehPemasok"`
// 	TotalInvoice              float64                   `gorm:"type:float" json:"totalInvoice"`
// 	TglTerimaDokFisikPemasok  *DateField          `gorm:"type:date" json:"tglTerimaDokFisikPemasok"`
// 	TglValidasiDokFisik       *DateField          `gorm:"type:date" json:"tglValidasiDokFisik"`
// }
