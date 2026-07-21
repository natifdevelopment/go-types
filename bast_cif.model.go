package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"
"github.com/google/uuid"
)

type JenisBast string

const (
	JenisBastCif   JenisBast = "CIF"
	JenisBastFob   JenisBast = "FOB"
	JenisBastTrans JenisBast = "TRANS"
	JenisBastCfr   JenisBast = "CFR"
)

func (o JenisBast) IsValid() bool {
	return o == JenisBastCif || o == JenisBastFob || o == JenisBastTrans || o == JenisBastCfr
}

type TipeDendaDesdem string

const (
	TipeDendaDesdemDes TipeDendaDesdem = "DES"
	TipeDendaDesdemDem TipeDendaDesdem = "DEM"
)

func (o TipeDendaDesdem) IsValid() bool {
	return o == TipeDendaDesdemDem || o == TipeDendaDesdemDes
}

func (BastCif) TableName() string {
	return "t_bast_cif"
}

type BastCif struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	SkemaKontrakCode    *JenisPengiriman  `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak        *ConfigDataInfo   `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	OrderCode           int               `gorm:"autoIncrement"`
	NoBast              *string           `gorm:"type:varchar(255)" json:"noBast"`
	TanggalBast         DateField   `gorm:"type:date" json:"tanggalBast"`

	// Revisi
	DibutuhkanPerbaikanBast  bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanBast"`
	CatatanPerbaikanBast     *string `gorm:"text" json:"catatanPerbaikanBast"`
	DibutuhkanPerbaikanDenda bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanDenda"`
	CatatanPerbaikanDenda    *string `gorm:"text" json:"catatanPerbaikanDenda"`

	// BAST Loading
	TPelabuhanMuatId *uuid.UUID       `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat    *PelabuhanMuat   `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	VolumeBl         *float32         `gorm:"type:float" json:"volumeBl"`
	TanggalBl        *DateField `gorm:"type:date" json:"tanggalBl"`

	// Loading
	MembuatDokBast  bool `gorm:"type:bool;default:true" json:"membuatDokBast"`
	MembuatDokDenda bool `gorm:"type:bool;default:true" json:"membuatDokDenda"`

	// Error to display on fe
	// so user can retry if error
	// error message saved on activity log
	DokBastError  bool `gorm:"type:bool;default:false" json:"dokBastError"`
	DokDendaError bool `gorm:"type:bool;default:false" json:"dokDendaError"`

	// Dokumen BAST
	NoDokBast  *string    `gorm:"type:varchar(255)" json:"noDokBast"`
	TDokBastId *uuid.UUID `gorm:"type:uuid" json:"dokBastId"`
	DokBast    *Media     `gorm:"foreignKey:TDokBastId" json:"dokBast"`

	// Dokumen Denda
	NoDokDenda         *string    `gorm:"type:varchar(255)" json:"noDokDenda"`
	TDokDendaId        *uuid.UUID `gorm:"type:uuid" json:"dokDendaId"`
	DokDenda           *Media     `gorm:"foreignKey:TDokDendaId" json:"dokDenda"`
	TDokRincianDendaId *uuid.UUID `gorm:"type:uuid" json:"dokRincianDendaId"`
	DokRincianDenda    *Media     `gorm:"foreignKey:TDokRincianDendaId" json:"dokRincianDenda"`

	// Pejabat
	TPejabatAmsId        *uuid.UUID           `gorm:"type:uuid" json:"pejabatAmsId"`
	PejabatAms           *PejabatAms          `gorm:"foreignKey:TPejabatAmsId" json:"pejabatAms"`
	TPejabatSuratKuasaId *uuid.UUID           `gorm:"type:uuid" json:"pejabatSuratKuasaId"`
	PejabatSuratKuasa    *Pejabat_surat_kuasa `gorm:"foreignKey:TPejabatSuratKuasaId" json:"pejabatSuratKuasa"`

	// BAST Loading
	TipePejabat          *string              `gorm:"type:varchar(255);default:'PEJABAT_INTERNAL'" json:"tipePejabat"`
	TPejabatPelimpahanId *uuid.UUID           `gorm:"type:uuid" json:"pejabatPelimpahanId"`
	PejabatPelimpahan    *Pejabat_surat_kuasa `gorm:"foreignKey:TPejabatPelimpahanId" json:"pejabatPelimpahan"`

	// Approval
	BastCifApproval     []BastCifApproval `gorm:"foreignKey:TBastCifId; constraint:OnDelete:CASCADE" json:"bastCifApproval"`
	BastCifHistory      []BastCifHistory  `gorm:"foreignKey:TBastCifId; constraint:OnDelete:CASCADE" json:"bastCifHistory"`
	DibutuhkanPerbaikan bool              `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan    *string           `gorm:"text" json:"catatanPerbaikan"`

	// Jika true maka input data sebelum ini sudah tidak bisa lagi
	// kecuali ada pengajuan rollback
	IsFinished bool       `gorm:"type:bool;default:false" json:"sudahSelesai"`
	FinishedAt *time.Time `gorm:"type:timestamp" json:"finishedAt"`
}

// @deprecated
// Dipindah ke t_dok_approval
type BastCifHistoryType string

const (
	BastCifHistoryTypeBuat         BastCifHistoryType = "BUAT"
	BastCifHistoryTypeSetujui      BastCifHistoryType = "SETUJUI"
	BastCifHistoryTypeTolak        BastCifHistoryType = "TOLAK"
	BastCifHistoryTypeEdit         BastCifHistoryType = "EDIT"
	BastCifHistoryTypePeriksa      BastCifHistoryType = "PERIKSA"
	BastCifHistoryTypeCatatan      BastCifHistoryType = "CATATAN"
	BastCifHistoryTypeCatatanRumus BastCifHistoryType = "CATATAN_RUMUS"
	BastCifHistoryTypePerbaikan    BastCifHistoryType = "PERBAIKAN"
)

func (o BastCifHistoryType) IsValid() bool {
	return o == BastCifHistoryTypeBuat || o == BastCifHistoryTypeSetujui || o == BastCifHistoryTypeTolak || o == BastCifHistoryTypeEdit
}

func (BastCifHistory) TableName() string {
	return "t_bast_cif_approval_history"
}

type BastCifHistory struct {
	baseapp.BaseModel
	TBastCifId      *uuid.UUID         `gorm:"type:uuid" json:"bastCifId"`
	BastCif         *BastCif           `gorm:"foreignKey:TBastCifId" json:"bastCif"`
	TOrganizationId *uuid.UUID         `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization      `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId  *uuid.UUID         `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevel       `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId      *uuid.UUID         `gorm:"type:uuid" json:"accountId"`
	Account         *Account           `gorm:"foreignKey:TAccountId" json:"account"`
	Tipe            BastCifHistoryType `gorm:"type:varchar(255)" json:"tipe"`
	Keterangan      *string            `gorm:"type:varchar(255)" json:"keterangan"`
	IdDokumen       *string            `gorm:"type:varchar(255)" json:"idDokumen"`
	NamaDokumen     *string            `gorm:"type:varchar(255)" json:"namaDokumen"`
}

func (BastCifApproval) TableName() string {
	return "t_bast_cif_approval"
}

type BastCifApproval struct {
	baseapp.BaseModel
	TBastCifId          *uuid.UUID               `gorm:"type:uuid" json:"bastCifId"`
	BastCif             *BastCif                 `gorm:"foreignKey:TBastCifId" json:"bastCif"`
	TJadwalPengirimanId *uuid.UUID               `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman        `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TOrganizationId     *uuid.UUID               `gorm:"type:uuid" json:"organizationId"`
	Organization        *Organization            `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId      *uuid.UUID               `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel         *AccessLevel             `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId          *uuid.UUID               `gorm:"type:uuid" json:"accountId"`
	Account             *Account                 `gorm:"foreignKey:TAccountId" json:"account"`
	Disetujui           bool                     `gorm:"type:boolean" json:"disetujui"`
	SubmitUlang         bool                     `gorm:"type:boolean;default:false" json:"submitUlang"`
	Dokumen             []BastCifApprovalDokumen `gorm:"foreignKey:TBastCifApprovalId; constraint:OnDelete:CASCADE" json:"dokumen"`
}

func (BastCifApprovalDokumen) TableName() string {
	return "t_bast_cif_approval_dokumen"
}

type BastCifApprovalDokumen struct {
	baseapp.BaseModel
	TBastCifApprovalId  *uuid.UUID        `gorm:"type:uuid" json:"bastCifApprovalId"`
	BastCifApproval     *BastCifApproval  `gorm:"foreignKey:TBastCifApprovalId" json:"bastCifApproval"`
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TMediaId            *uuid.UUID        `gorm:"type:uuid" json:"mediaId"`
	Media               *Media            `gorm:"foreignKey:TMediaId" json:"media"`
	TipeDokumen         string            `gorm:"type:varchar(255)" json:"tipeDokumen"`
	Keterangan          *string           `gorm:"type:varchar(255)" json:"keterangan"`
	SudahDiperbaiki     bool              `gorm:"type:boolean;default:false" json:"sudahDiperbaiki"`
	Diperiksa           bool              `gorm:"type:boolean;default:false" json:"diperiksa"`
}

// ---------------------------------------------------------------------
// Denda
// 1. Desdem // CIF/TRANS
// 2. Keterlambatan CIF/CFR
// 3. Denda muat FOB
// 4. Senda susut TRANS
// ---------------------------------------------------------------------
func (DendaDesdem) TableName() string {
	return "t_denda_desdem"
}

type DendaDesdem struct {
	baseapp.BaseModel

	// Metadata
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`

	// Informasi Denda
	TipeDesdem    TipeDendaDesdem `gorm:"type:varchar(255)" json:"tipeDesdem"` // DES/DEM
	HariDespatch  float32         `gorm:"type:float" json:"hariDespatch"`
	HargaDespatch float32         `gorm:"type:float" json:"hargaDespatch"`
	HariDemurrage float32         `gorm:"type:float" json:"hariDemurrage"`
	HargaDemurage float32         `gorm:"type:float" json:"hargaDemurage"`
}

func (DendaMuat) TableName() string {
	return "t_denda_muat"
}

type DendaMuat struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`

	TglAwal         time.Time
	TglAkhir        time.Time
	MaksHariDenda   float64
	SelisihHari     float64
	HariTerlambat   float64
	VolumeBl        float64
	ToleransiDenda  float64 `gorm:"type:float" json:"toleransiDenda"`
	PersentaseDenda float64 `gorm:"type:float" json:"persentaseDenda"`

	// Nilai
	HargaFob  float64 // Ambil dari rumus "Harga Dasar Batubara per Ton + PPN"
	DendaMuat float64
}

func (DendaSusut) TableName() string {
	return "t_denda_susut"
}

type DendaSusut struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`

	// Informasi Denda
	VolumeBl       float64 `gorm:"type:float" json:"volumeBl"`
	VolumeDs       float64 `gorm:"type:float" json:"volumeDs"`
	ToleransiSusut float64 `gorm:"type:float" json:"toleransiSusut"`
	SusutVolume    float64 `gorm:"type:float" json:"susutVolume"`
	SusutPercent   float64 `gorm:"type:float" json:"susutPercent"`

	// Nilai
	HargaFob   float64
	DendaSusut float64 `gorm:"type:float" json:"dendaSusut"`
}
