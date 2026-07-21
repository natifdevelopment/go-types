package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"
"github.com/google/uuid"
)

type TipePropose string

const (
	TipeProposeManual      TipePropose = "MANUAL_PROPOSE"
	TipeProposeTermin      TipePropose = "PROPOSE_TERMIN"
	TipeProposePerhitungan TipePropose = "PROPOSE_PERHITUNGAN"
)

type StatusPerhitungan string

const (
	StatusPerhitunganBelumDihitung       StatusPerhitungan = "BELUM_DIHITUNG"
	StatusPerhitunganPerhitunganDitolak  StatusPerhitungan = "PERHITUNGAN_DITOLAK"
	StatusPerhitunganSudahDihitung       StatusPerhitungan = "SUDAH_DIHITUNG"
	StatusPerhitunganMenungguPersetujuan StatusPerhitungan = "MENUNGGU_PERSETUJUAN"
	StatusPerhitunganDisetujui           StatusPerhitungan = "DISETUJUI"
)

func (o StatusPerhitungan) IsValid() bool {
	return o == StatusPerhitunganBelumDihitung || o == StatusPerhitunganSudahDihitung ||
		o == StatusPerhitunganMenungguPersetujuan || o == StatusPerhitunganDisetujui
}

type StatusPenagihan string

const (
	StatusPenagihanBelumDikirim StatusPenagihan = "BELUM_DIKIRIM"
	StatusPenagihanDalamProses  StatusPenagihan = "DALAM_PROSES"
	StatusPenagihanDalamSelesai StatusPenagihan = "PENAGIHAN_SELESAI"
)

func (o StatusPenagihan) IsValid() bool {
	return o == StatusPenagihanBelumDikirim || o == StatusPenagihanDalamProses || o == StatusPenagihanDalamSelesai
}

type StatusInvoice string

const (
	StatusInvoiceBelumLengkap StatusInvoice = "BELUM_LENGKAP"
	StatusInvoiceSudahLengkap StatusInvoice = "SUDAH_LENGKAP"
)

func (o StatusInvoice) IsValid() bool {
	return o == StatusInvoiceBelumLengkap || o == StatusInvoiceSudahLengkap
}

func (ProposePerhitungan) TableName() string {
	return "t_propose_perhitungan"
}

type ProposePerhitungan struct {
	baseapp.BaseModel
	TJadwalPengirimanId       *uuid.UUID                  `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman          *JadwalPengiriman           `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TipePropose               TipePropose                 `gorm:"type:varchar(255)" json:"tipePropose"`
	SkemaKontrakCode          *JenisPengiriman            `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak              *ConfigDataInfo             `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TBastCifId                *uuid.UUID                  `gorm:"type:uuid" json:"bastCifId"`
	BastCif                   *BastCif                    `gorm:"foreignKey:TBastCifId" json:"bastCif"`
	HasilPerhitungan          float64                     `gorm:"type:float;default:0;" json:"hasilPerhitungan"`
	PersentaseTermin          float64                     `gorm:"type:float;default:0;" json:"persentaseTermin"`
	TotalTagihan              float64                     `gorm:"type:float;default:0;" json:"totalTagihan"`
	NilaiDenda                float64                     `gorm:"type:float;default:0;" json:"nilaiDenda"`
	AdaTermin                 bool                        `gorm:"type:bool" json:"adaTermin"`
	AdaManualPropose          bool                        `gorm:"type:bool" json:"adaManualPropose"`
	MatDocNo                  string                      `gorm:"type:varchar(255)" json:"matDocNo"`
	MatDocNoTransSelisih      *string                     `gorm:"type:varchar(255)" json:"matDocNoTransSelisih"`
	MatDocYear                int                         `gorm:"type:integer" json:"matDocYear"`
	UploadInvoicePenagihan    []UploadInvoicePenagihan    `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"uploadInvoicePenagihan"`
	BundleInvoicePangihanItem []BundleInvoicePangihanItem `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"bundleInvoicePangihanItem"`
	VerifikasiPenagihan       []VerifikasiPenagihan       `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"verifikasiPenagihan"`
	PenagihanAkuntansi        []PenagihanAkuntansiItem    `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"penagihanAkuntansi"`
	ProposeRumus              []ProposeRumus              `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"proposeRumus"`
	PelunasanPenagihan        []PelunasanPenagihanItem    `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"pelunasanPenagihan"`

	// Next data connection
	PermohonanPembayaran []PermohonanPembayaranItem `gorm:"foreignKey:TProposePerhitunganId;constraint:OnDelete:CASCADE" json:"permohonanPembayaran"`

	// BAPH Document
	NoBaph        string     `gorm:"type:varchar(255)" json:"noBaph"`
	DokumenDibuat bool       `gorm:"type:bool;default:false" json:"dokumenDibuat"`
	DokumenError  *string    `gorm:"type:text" json:"dokumenError"`
	TDokBaphId    *uuid.UUID `gorm:"type:uuid" json:"dokBaphId"`
	DokBaph       *Media     `gorm:"foreignKey:TDokBaphId" json:"dokBaph"`

	// This is new chapter, we need to disconnect the relation
	// So filter, search etc no need to get from jadwal pengiriman
	TOrganizationId *uuid.UUID            `gorm:"type:uuid" json:"organizationId"` // Maker
	Organization    *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPembangkitId   *uuid.UUID            `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit      *Organization         `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId      *uuid.UUID            `gorm:"type:uuid" json:"pemasokId"`
	Pemasok         *Organization         `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TTransportirId  *uuid.UUID            `gorm:"type:uuid" json:"transportirId"`
	Transportir     *Organization         `gorm:"foreignKey:TTransportirId" json:"transportir"`
	Periode         *YearMonthField `gorm:"type:date" json:"periode"`

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

	// Jika propose perhitungan tidak disetujui, maka akan menjadi dibutuhkan_perbaikan
	DibutuhkanPerbaikan                     bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan                        *string `gorm:"text" json:"catatanPerbaikan"`
	DibutuhkanPerbaikanPermohonanPembayaran bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanPermohonanPembayaran"`
	CatatanPerbaikanPermohonanPembayaran    *string `gorm:"text" json:"catatanPerbaikanPermohonanPembayaran"`

	// Setelah dihitung, pemasok akan propose dan mengisi data ini
	TglDihitung                *DateTimeField `gorm:"type:timestamp" json:"tglDihitung"`
	TDihitungOlehId            *uuid.UUID           `gorm:"type:uuid" json:"dihitungOlehId"`
	DihitungOleh               *Account             `gorm:"foreignKey:TDihitungOlehId" json:"dihitungOleh"`
	TDihitungOlehOrgId         *uuid.UUID           `gorm:"type:uuid" json:"dihitungOlehOrgId"`
	DihitungOlehOrg            *Organization        `gorm:"foreignKey:TDihitungOlehOrgId" json:"dihitungOlehOrg"`
	TDihitungOlehAccessLevelId *uuid.UUID           `gorm:"type:uuid" json:"dihitungOlehAccessLevelId"`
	DihitungOlehAccessLevel    *AccessLevel         `gorm:"foreignKey:TDihitungOlehAccessLevelId" json:"dihitungOlehAccessLevel"`

	// Setelah dihitung, pemasok akan propose dan mengisi data ini
	// Data ini juga akan tersimpan pada propose approval data
	// termasuk jika ditolak
	TglDiajukan                *DateTimeField `gorm:"type:timestamp" json:"tglDiajukan"`
	TDiajukanOlehId            *uuid.UUID           `gorm:"type:uuid" json:"diajukanOlehId"`
	DiajukanOleh               *Account             `gorm:"foreignKey:TDiajukanOlehId" json:"diajukanOleh"`
	TDiajukanOlehOrgId         *uuid.UUID           `gorm:"type:uuid" json:"diajukanOlehOrgId"`
	DiajukanOlehOrg            *Organization        `gorm:"foreignKey:TDiajukanOlehOrgId" json:"diajukanOlehOrg"`
	TDiajukanOlehAccessLevelId *uuid.UUID           `gorm:"type:uuid" json:"diajukanOlehAccessLevelId"`
	DiajukanOlehAccessLevel    *AccessLevel         `gorm:"foreignKey:TDiajukanOlehAccessLevelId" json:"diajukanOlehAccessLevel"`

	// Setelah catatan perbaikan diperbaiki, submit_ulang akan menjadi true
	SubmitUlang bool `gorm:"type:bool;default:false" json:"submitUlang"`

	// Used for grouping on fe
	PembangkitName string `gorm:"type:varchar(255)" json:"pembangkitName"`

	// Setelah upload invoice dan user klik buat penagihan
	// Maka no penagihan akan terisi
	NoPenagihan *string `gorm:"type:varchar(255)" json:"noPenagihan"`

	// pertama kali dibuat itu ketika sudah melakukan perhitungan
	// kemudian pemasok dapat klik propose dan status menjadi menunggu_persetujuan,
	// setelah bast disetujui oleh semua pihak, maka akan menjadi disetujui,
	StatusPerhitungan StatusPerhitungan `gorm:"type:varchar(255);default:'SUDAH_DIHITUNG'" json:"statusPerhitungan"`

	// Tanggal di isi ketika semua approval sudah approve
	TglVerifikasi *DateTimeField `gorm:"type:timestamp" json:"tglVerifikasi"`

	// Pertama dibuat akan menjadi belum_dikirim,
	// Ketika kirim permohonan pembayaran, maka akan menjadi dalam_proses,
	StatusPenagihan StatusPenagihan `gorm:"type:varchar(255);default:'BELUM_DIKIRIM'" json:"statusPenagihan"` // belum_dikirim/dalam_proses

	// Pertama dibuat status invoice akan menjadi belum_lengkap,
	// Jika semua document invoice sudah diupload, maka akan menjadi sudah_lengkap,
	StatusInvoice StatusInvoice `gorm:"type:varchar(255);default:'BELUM_LENGKAP'" json:"statusInvoice"` // Sudah_lengkap/Belum_lengkap

	// SAP Integration
	TIntegrasiReqId  *uuid.UUID       `gorm:"type:uuid" json:"integrasiReqId"`
	IntegrasiReq     *IntegrasiStatus `gorm:"foreignKey:TIntegrasiReqId" json:"integrasiReq"`
	TIntegrasiRespId *uuid.UUID       `gorm:"type:uuid" json:"integrasiRespId"`
	IntegrasiResp    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiRespId" json:"integrasiResp"`
	SapStatus        ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapStatus"`
	SapMatDocNo      *string          `gorm:"type:varchar(255);" json:"sapMatDocNo"`
	SapDocYear       *int             `gorm:"type:int;" json:"sapDocYear"`
	SapLineNo        *string          `gorm:"type:varchar(255);" json:"sapLineNo"`

	// SAP SPT
	TIntegrasiSptReqId  *uuid.UUID       `gorm:"type:uuid" json:"integrasiSptReqId"`
	IntegrasiSptReq     *IntegrasiStatus `gorm:"foreignKey:TIntegrasiSptReqId" json:"integrasiSptReq"`
	TIntegrasiSptRespId *uuid.UUID       `gorm:"type:uuid" json:"integrasiSptRespId"`
	IntegrasiSptResp    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiSptRespId" json:"integrasiSptResp"`
	SapSptStatus        ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapSptStatus"` // @deprecated, moved to spt table
	SapInvNo            *string          `gorm:"type:varchar(20)" json:"sapInvNo"`
	SapInvYear          *int             `gorm:"type:int" json:"sapInvYear"`
	SapSptUpdateCount   *int             `gorm:"type:int" json:"sapSptUpdateCount"`

	// SAP Nota Dinas
	TIntegrasiNotaDinasReqId  *uuid.UUID       `gorm:"type:uuid" json:"integrasiNotaDinasReqId"`
	IntegrasiNotaDinasReq     *IntegrasiStatus `gorm:"foreignKey:TIntegrasiReqId" json:"integrasiNotaDinasReq"`
	TIntegrasiNotaDinasRespId *uuid.UUID       `gorm:"type:uuid" json:"integrasiNotaDinasRespId"`
	IntegrasiNotaDinasResp    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiRespId" json:"integrasiNotaDinasResp"`
	SapNotaDinasStatus        ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapNotaDinasStatus"`
	SapPostingDate            *DateField `gorm:"type:date" json:"sapPostingDate"`
}

func (ProposeSimple) TableName() string {
	return "t_propose_perhitungan"
}

type ProposeSimple struct {
	ID                  uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	TJadwalPengirimanId *uuid.UUID              `gorm:"type:uuid" json:"-"`
	Pengiriman          *JadwalPengirimanSimple `gorm:"foreignKey:TJadwalPengirimanId" json:"pengiriman"`
}

func (ProposePerhitunganHistory) TableName() string {
	return "t_propose_perhitungan_history"
}

type ProposePerhitunganHistory struct {
	baseapp.BaseModel
	TProposePerhitunganId *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	TOrganizationId       *uuid.UUID          `gorm:"type:uuid" json:"organizationId"`
	Organization          *Organization       `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId        *uuid.UUID          `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel           *AccessLevel        `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId            *uuid.UUID          `gorm:"type:uuid" json:"accountId"`
	Account               *Account            `gorm:"foreignKey:TAccountId" json:"account"`
	Tipe                  BastCifHistoryType  `gorm:"type:varchar(255)" json:"tipe"`
	Keterangan            *string             `gorm:"type:varchar(255)" json:"keterangan"`
	IdDokumen             *string             `gorm:"type:varchar(255)" json:"idDokumen"`
	NamaDokumen           *string             `gorm:"type:varchar(255)" json:"namaDokumen"`
}

// Ketika pemasok/transportir klik propose
// maka akan merubah status propose perhitungan dan mengisi
// data pada approval ini berdasarkan pelimpahan
// Setelah itu hanya terdaftar pada approval ini
// yang dapat melakukan approval
func (ProposePerhitunganApproval) TableName() string {
	return "t_propose_perhitungan_approval"
}

type ProposePerhitunganApproval struct {
	baseapp.BaseModel
	TProposePerhitunganId *uuid.UUID                          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan                 `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	TJadwalPengirimanId   *uuid.UUID                          `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman      *JadwalPengiriman                   `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TOrganizationId       *uuid.UUID                          `gorm:"type:uuid" json:"organizationId"`
	Organization          *Organization                       `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId        *uuid.UUID                          `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel           *AccessLevel                        `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId            *uuid.UUID                          `gorm:"type:uuid" json:"accountId"`
	Account               *Account                            `gorm:"foreignKey:TAccountId" json:"account"`
	Disetujui             bool                                `gorm:"type:boolean" json:"disetujui"`
	SubmitUlang           bool                                `gorm:"type:boolean;default:false" json:"submitUlang"`
	Dokumen               []ProposePerhitunganApprovalDokumen `gorm:"foreignKey:TProposePerhitunganApprovalId;constraint:OnDelete:CASCADE" json:"dokumen"`
}

func (ProposePerhitunganApprovalDokumen) TableName() string {
	return "t_propose_perhitungan_approval_dokumen"
}

type ProposePerhitunganApprovalDokumen struct {
	baseapp.BaseModel
	TProposePerhitunganApprovalId *uuid.UUID                  `gorm:"type:uuid" json:"proposePerhitunganApprovalId"`
	ProposePerhitunganApproval    *ProposePerhitunganApproval `gorm:"foreignKey:TProposePerhitunganApprovalId" json:"proposePerhitunganApproval"`
	TJadwalPengirimanId           *uuid.UUID                  `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman              *JadwalPengiriman           `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TMediaId                      *uuid.UUID                  `gorm:"type:uuid" json:"mediaId"`
	Media                         *Media                      `gorm:"foreignKey:TMediaId" json:"media"`
	TipeDokumen                   string                      `gorm:"type:varchar(255)" json:"tipeDokumen"`
	Keterangan                    *string                     `gorm:"type:varchar(255)" json:"keterangan"`
	SudahDiperbaiki               bool                        `gorm:"type:boolean;default:false" json:"sudahDiperbaiki"`
	Diperiksa                     bool                        `gorm:"type:boolean;default:false" json:"diperiksa"`
}

// Rumus
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// @deprecated
// hasil perhitungan rumus sekarang dipindah dari propose
// ke jadwal pengiriman, digunakan juga untuk propose termin dan
// manual propose agar dapat digunakan re-usable oleh ketiga module
type TipeKontrakPropose string

const (
	TipeKontrakProposePjbb TipeKontrakPropose = "PJBB"
	TipeKontrakProposePjab TipeKontrakPropose = "PJAB"
)

func (ProposeRumus) TableName() string {
	return "t_propose_perhitungan_rumus"
}

type ProposeRumus struct {
	baseapp.BaseModel
	TProposePerhitunganId *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	TipeKontrak           TipeKontrakPropose  `gorm:"type:varchar(255);" json:"tipeKontrak"`

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

	TPemasokId            *uuid.UUID         `gorm:"index;type:uuid" json:"pemasokId"`
	Pemasok               *Organization      `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TTransportirId        *uuid.UUID         `gorm:"index;type:uuid" json:"transportirId"`
	Transportir           *Organization      `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TPembangkitId         *uuid.UUID         `gorm:"index;type:uuid" json:"pembangkitId"`
	Pembangkit            *Organization      `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TipePengiriman        *JenisPengiriman   `json:"tipePengiriman"`
	JenisTransportasiCode *string            `json:"jenisTransportasiCode"`
	JenisTransportasi     *ConfigDataInfo    `gorm:"foreignKey:JenisTransportasiCode;references:Code;" json:"jenisTransportasi"`
	RumusBlok             []ProposeRumusBlok `gorm:"foreignKey:TProposeRumusId;constraint:OnDelete:CASCADE" json:"rumusBlocks"`
}

// Rumus Block
// -------------------------
func (ProposeRumusBlok) TableName() string {
	return "t_propose_perhitungan_rumus_blok"
}

type ProposeRumusBlok struct {
	baseapp.BaseModel
	TRefId                  *uuid.UUID                     `gorm:"index;type:uuid" json:"-"`
	Type                    TipeRumusBlok                  `gorm:"type:varchar(255)" json:"type"`
	TProposeRumusId         *uuid.UUID                     `gorm:"index;type:uuid" json:"ProposeRumusId"`
	ProposeRumus            *ProposeRumus                  `gorm:"foreignKey:TProposeRumusId" json:"ProposeRumus"`
	NamaRumusBlockCode      *string                        `gorm:"type:varchar(255)" json:"namaRumusBlockCode"`
	NamaRumusBlock          *ConfigDataInfo                `gorm:"foreignKey:NamaRumusBlockCode;references:Code;" json:"namaRumusBlock"`
	Deskripsi               *string                        `gorm:"type:text" json:"description"`
	VariableItem            json.RawMessage                `gorm:"type:jsonb" json:"variables"`
	RumusBlokItem           []ProposeRumusBlokItem         `gorm:"foreignKey:TProposeRumusBlokId;constraint:OnDelete:CASCADE" json:"items"`
	TPjbbKualitasTemplateId *uuid.UUID                     `gorm:"index;type:uuid" json:"pjbbKualitasTemplateId"`
	PjbbKualitasTemplate    *Pjbb_kualitas_template        `gorm:"foreignKey:TPjbbKualitasTemplateId" json:"pjbbKualitasTemplate"`
	RumusBlokPelMuatSpek    []ProposeRumusBlokPelMuatSpek  `gorm:"foreignKey:TProposeRumusBlokId;constraint:OnDelete:CASCADE" json:"pelMuatSpeks"`
	RumusBlokKualitasItem   []ProposeRumusBlokKualitasItem `gorm:"foreignKey:TProposeRumusBlokId;constraint:OnDelete:CASCADE" json:"kualitasItems"`
	OrderNumber             int                            `json:"orderNumber"`
}

// Rumus Perhitungan
// -------------------------
func (ProposeRumusBlokItem) TableName() string {
	return "t_propose_perhitungan_blok_rumus_item"
}

type ProposeRumusBlokItem struct {
	baseapp.BaseModel
	TProposeRumusId        *uuid.UUID                      `gorm:"index;type:uuid" json:"proposeRumusId"`
	ProposeRumus           *ProposeRumus                   `gorm:"foreignKey:TProposeRumusId" json:"proposeRumus"`
	TProposeRumusBlokId    *uuid.UUID                      `gorm:"index;type:uuid" json:"proposeRumusBlokId"`
	ProposeRumusBlok       *ProposeRumusBlok               `gorm:"foreignKey:TProposeRumusBlokId" json:"proposeRumusBlok"`
	TUraianId              *uuid.UUID                      `gorm:"type:uuid" json:"uraianId"`
	Uraian                 *Uraian                         `gorm:"foreignKey:TUraianId" json:"uraian"`
	NamaVariable           string                          `gorm:"type:varchar(255)" json:"variable"`
	NilaiAwal              *string                         `gorm:"type:varchar(255)" json:"nilaiAwal"`
	Hasil                  *float64                        `gorm:"type:float" json:"hasil"`
	RumusBlokItemCondition []ProposeRumusBlokItemCondition `gorm:"foreignKey:TProposeRumusBlokItemId;constraint:OnDelete:CASCADE" json:"conditions"`
	TampilkanDiPropose     *bool                           `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	SatuanHasil            *string                         `gorm:"type:varchar(255)" json:"satuanHasil"`
	OrderNumber            int                             `json:"orderNumber"`

	// Khusus propose
	TRefId  *uuid.UUID `gorm:"index;type:uuid" json:"-"`
	Catatan *string    `gorm:"type:varchar(255)" json:"catatan"`
}

func (ProposeRumusBlokItemCondition) TableName() string {
	return "t_propose_perhitungan_rumus_blok_item_condition"
}

type ProposeRumusBlokItemCondition struct {
	baseapp.BaseModel
	TRefId                  *uuid.UUID            `gorm:"index;type:uuid" json:"-"`
	TProposeRumusId         *uuid.UUID            `gorm:"index;type:uuid" json:"proposeRumusId"`
	ProposeRumus            *ProposeRumus         `gorm:"foreignKey:TProposeRumusId" json:"proposeRumus"`
	TProposeRumusBlokItemId *uuid.UUID            `gorm:"index;type:uuid" json:"proposeRumusBlokItemId"`
	ProposeRumusBlokItem    *ProposeRumusBlokItem `gorm:"foreignKey:TProposeRumusBlokItemId" json:"proposeRumusBlokItem"`
	Rumus                   *string               `gorm:"type:text" json:"rumus"`
	VariableItemValue       json.RawMessage       `gorm:"type:jsonb" json:"variables"`
	OrderNumber             int                   `json:"orderNumber"`
}

// Rumus Kualitas
// -------------------------
func (ProposeRumusBlokPelMuatSpek) TableName() string {
	return "t_propose_perhitungan_rumus_blok_pelmuat_spek"
}

type ProposeRumusBlokPelMuatSpek struct {
	baseapp.BaseModel
	TRefId                *uuid.UUID                     `gorm:"index;type:uuid" json:"-"`
	TProposeRumusId       *uuid.UUID                     `gorm:"index;type:uuid" json:"proposeRumusId"`
	ProposeRumus          *ProposeRumus                  `gorm:"foreignKey:TProposeRumusId" json:"proposeRumus"`
	TProposeRumusBlokId   *uuid.UUID                     `gorm:"index;type:uuid" json:"proposeRumusBlokId"`
	ProposeRumusBlok      *ProposeRumusBlok              `gorm:"foreignKey:TProposeRumusBlokId" json:"proposeRumusBlok"`
	TPelabuhanMuatId      *uuid.UUID                     `gorm:"index;type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat         *PelabuhanMuat                 `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId               *uuid.UUID                     `gorm:"index;type:uuid" json:"spekId"`
	Spek                  *Spek                          `gorm:"foreignKey:TSpekId" json:"spek"`
	RumusBlokKualitasItem []ProposeRumusBlokKualitasItem `gorm:"foreignKey:TProposeRumusBlokPelMuatSpekId;constraint:OnDelete:CASCADE" json:"kualitasItems"`
	OrderNumber           int                            `json:"orderNumber"`
}

func (ProposeRumusBlokKualitasItem) TableName() string {
	return "t_propose_perhitungan_rumus_blok_kualitas_item"
}

type ProposeRumusBlokKualitasItem struct {
	baseapp.BaseModel
	TProposeRumusId                *uuid.UUID                   `gorm:"index;type:uuid" json:"proposeRumusId"`
	ProposeRumus                   *ProposeRumus                `gorm:"foreignKey:TProposeRumusId" json:"proposeRumus"`
	TProposeRumusBlokId            *uuid.UUID                   `gorm:"index;type:uuid" json:"proposeRumusBlokId"`
	ProposeRumusBlok               *ProposeRumusBlok            `gorm:"foreignKey:TProposeRumusBlokId" json:"proposeRumusBlok"`
	TProposeRumusBlokPelMuatSpekId *uuid.UUID                   `gorm:"index;type:uuid" json:"proposeRumusBlokPelMuatSpekId"`
	ProposeRumusBlokPelMuatSpek    *ProposeRumusBlokPelMuatSpek `gorm:"foreignKey:TProposeRumusBlokPelMuatSpekId" json:"proposeRumusBlokPelMuatSpek"`
	TipeKualitas                   Pjbb_kualitasItemType        `gorm:"type:varchar(255);default:'none'" json:"tipeKualitas"`
	TUraianId                      uuid.UUID                    `gorm:"type:uuid" json:"uraianId"`
	Uraian                         *Uraian                      `gorm:"foreignKey:TUraianId" json:"uraian"`
	BasisCode                      *string                      `gorm:"type:varchar(255)" json:"basisCode"`
	Basis                          *ConfigDataInfo              `gorm:"foreignKey:BasisCode;references:Code;" json:"basis"`
	Tipikal                        *string                      `gorm:"type:varchar(255)" json:"tipikal"`
	Penolakan                      *string                      `gorm:"type:varchar(255)" json:"penolakan"`
	Simbol                         *string                      `gorm:"type:varchar(255)" json:"simbol"`
	BatasBawah                     *string                      `gorm:"type:varchar(255)" json:"batasBawah"`
	BatasAtas                      *string                      `gorm:"type:varchar(255)" json:"batasAtas"`
	TipeData                       Pjbb_kualitasItemDataType    `gorm:"type:varchar(255);default:'Num'" json:"tipeData"`
	TurunHarga                     *string                      `gorm:"type:varchar(255)" json:"turunHarga"`
	HasilUji                       *string                      `gorm:"type:varchar(255)" json:"hasilUji"`
	Denda                          *string                      `gorm:"type:varchar(255)" json:"denda"`
	Rumus                          *string                      `gorm:"type:text" json:"rumus"`
	Keterangan                     *string                      `gorm:"type:varchar(255)" json:"keterangan"`
	Variable                       *string                      `gorm:"type:varchar(255)" json:"variable"`
	Hasil                          *float64                     `gorm:"type:float" json:"hasil"`
	TampilkanDiPropose             *bool                        `gorm:"type:boolean;default:true" json:"tampilkanDiPropose"`
	OrderNumber                    int                          `json:"orderNumber"`

	// Khusus propose
	TRefId  *uuid.UUID `gorm:"index;type:uuid" json:"-"`
	Catatan *string    `gorm:"type:varchar(255)" json:"catatan"`
}
