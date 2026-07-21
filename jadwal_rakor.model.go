package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"
"github.com/google/uuid"
)

func (JadwalRakor) TableName() string {
	return "t_jadwal_rakor"
}

type JadwalRakor struct {
	baseapp.BaseModel
	Judul               string                `gorm:"type:varchar(255)" json:"judul"`
	TJadwalPraRakorId   *uuid.UUID            `gorm:"type:uuid" json:"jadwalPraRakorId"`
	JadwalPraRakor      *JadwalPraRakor       `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalPraRakor"`
	TOrganizationId     *uuid.UUID            `gorm:"type:uuid" json:"organizationId"`
	Organization        *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TRegionalId         *uuid.UUID            `gorm:"type:uuid" json:"regionalId"`
	Regional            *Organization         `gorm:"foreignKey:TRegionalId" json:"regional"`
	Tanggal             DateField       `gorm:"type:date" json:"tanggal"`
	Periode             *YearMonthField `gorm:"type:date" json:"periode"`
	RakorDitutup        *time.Time            `gorm:"type:timestamp" json:"rakorDitutup"`
	TRakorDitutupOlehId *uuid.UUID            `gorm:"type:uuid" json:"rakorDitutupOlehId"`
	RakorDitutupOleh    *Account              `gorm:"foreignKey:TRakorDitutupOlehId" json:"rakorDitutupOleh"`
	JadwalRakorUndangan []JadwalRakorUndangan `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakorUndangan"`
	EntryJadwal         []EntryJadwal         `gorm:"foreignKey:TJadwalRakorId" json:"entryJadwal"`
	JadwalPengiriman    []JadwalPengiriman    `gorm:"foreignKey:TJadwalRakorId" json:"jadwalPengiriman"`
	PengirimanTerpenuhi bool                  `gorm:"type:boolean;default:false" json:"pengirimanTerpenuhi"`
}

type StatusUndanganRakor string

const (
	StatusUndanganRakorUnconfirmed StatusUndanganRakor = "unconfirmed"
	StatusUndanganRakorConfirmed   StatusUndanganRakor = "confirmed"
	StatusUndanganRakorOnBid       StatusUndanganRakor = "on-bid"
	StatusUndanganRakorApproved    StatusUndanganRakor = "approved"
	StatusUndanganRakorRejected    StatusUndanganRakor = "rejected"
)

func (o StatusUndanganRakor) IsValid() bool {
	return o == StatusUndanganRakorUnconfirmed || o == StatusUndanganRakorConfirmed ||
		o == StatusUndanganRakorOnBid || o == StatusUndanganRakorApproved || o == StatusUndanganRakorRejected
}

func (JadwalRakorUndangan) TableName() string {
	return "t_jadwal_rakor_undangan"
}

type JadwalRakorUndangan struct {
	baseapp.BaseModel
	TJadwalRakorId       *uuid.UUID          `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor          *JadwalRakor        `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TOrganizationId      *uuid.UUID          `gorm:"type:uuid" json:"organizationId"`
	Organization         *Organization       `gorm:"foreignKey:TOrganizationId" json:"organization"`
	StatusUndanganRakor  StatusUndanganRakor `gorm:"type:varchar(255);default:unconfirmed" json:"statusUndanganRakor"`
	KonfirmasiTutupRakor *time.Time          `gorm:"type:timestamp" json:"konfirmasiTutupRakor"`
}

func (TutupRakorKonfirmasi) TableName() string {
	return "t_tutup_rakor_konfirmasi"
}

type TutupRakorKonfirmasi struct {
	baseapp.BaseModel
	TJadwalRakorId   *uuid.UUID    `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor      *JadwalRakor  `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TOrganizationId  *uuid.UUID    `gorm:"type:uuid" json:"organizationId"`
	Organization     *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
	DikonfirmasiPada *time.Time    `gorm:"type:timestamp" json:"konfirmasiTutupRakor"`
	TAccountId       *uuid.UUID    `gorm:"type:uuid" json:"accountId"`
	Account          *AccountInfo  `gorm:"foreignKey:TAccountId;" json:"account"`
}

type TipeRiwayatRakor string

const (
	TipeRiwayatRakorRakor               TipeRiwayatRakor = "rakor"
	TipeRiwayatRakorPraRakor            TipeRiwayatRakor = "pra_rakor"
	TipeRiwayatRakorKemampuanPasok      TipeRiwayatRakor = "kemampuan_pasok"
	TipeRiwayatRakorRencanaPasokan      TipeRiwayatRakor = "rencana_pasokan"
	TipeRiwayatRakorVerifikasiRencana   TipeRiwayatRakor = "verifikasi_rencana_pra_rakor"
	TipeRiwayatRakorEntriJadwal         TipeRiwayatRakor = "entri_jadwal"
	TipeRiwayatRakorVerifikasiKebutuhan TipeRiwayatRakor = "verifikasi_kebutuhan_rakor"
)

func (o TipeRiwayatRakor) IsValid() bool {
	return o == TipeRiwayatRakorPraRakor || o == TipeRiwayatRakorKemampuanPasok ||
		o == TipeRiwayatRakorRencanaPasokan || o == TipeRiwayatRakorVerifikasiRencana ||
		o == TipeRiwayatRakorEntriJadwal || o == TipeRiwayatRakorVerifikasiKebutuhan
}

type AksiRiwayatRakor string

const (
	AksiRiwayatRakorMoveJadwal AksiRiwayatRakor = "pindah_jadwal"
	AksiRiwayatRakorSetAlokasi AksiRiwayatRakor = "alokasi_pasokan"
	AksiRiwayatRakorConfirm    AksiRiwayatRakor = "konfirmasi"
	AksiRiwayatRakorApprove    AksiRiwayatRakor = "setujui"
	AksiRiwayatRakorReject     AksiRiwayatRakor = "tolak"
	AksiRiwayatRakorCreate     AksiRiwayatRakor = "membuat"
	AksiRiwayatRakorEdit       AksiRiwayatRakor = "edit"
)

func (o AksiRiwayatRakor) IsValid() bool {
	return o == AksiRiwayatRakorConfirm || o == AksiRiwayatRakorApprove ||
		o == AksiRiwayatRakorReject || o == AksiRiwayatRakorCreate || o == AksiRiwayatRakorEdit
}

func (RiwayatRakor) TableName() string {
	return "t_jadwal_rakor_riwayat"
}

type RiwayatRakor struct {
	baseapp.BaseModel
	TJadwalRakorId    *uuid.UUID       `gorm:"index" json:"jadwalRakorId"`
	JadwalRakor       *JadwalRakor     `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TJadwalPraRakorId *uuid.UUID       `gorm:"index" json:"jadwalPraRakorId"`
	JadwalPraRakor    *JadwalPraRakor  `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalPraRakor"`
	TRencanaPasokanId *uuid.UUID       `gorm:"index" json:"rencanaPasokanId"`
	RencanaPasokan    *RencanaPasokan  `gorm:"foreignKey:TRencanaPasokanId" json:"rencanaPasokan"`
	TKemampuanPasokId *uuid.UUID       `gorm:"index" json:"kemampuanPasokId"`
	KemampuanPasok    *KemampuanPasok  `gorm:"foreignKey:TKemampuanPasokId" json:"kemampuanPasok"`
	TOrganizationId   *uuid.UUID       `gorm:"index" json:"organizationId"`
	Organization      *Organization    `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId    *uuid.UUID       `gorm:"index" json:"accessLevelId"`
	AccessLevel       *AccessLevel     `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TipeRiwayat       TipeRiwayatRakor `gorm:"type:varchar(255)" json:"tipeRiwayat"`
	AksiRiwayat       AksiRiwayatRakor `gorm:"type:varchar(255)" json:"aksiRiwayat"`
	Catatan           *string          `gorm:"type:text" json:"catatan"`
}
