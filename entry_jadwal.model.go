package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type StatusEntryJadwal string

const (
	StatusEntryJadwalUnconfirmed StatusEntryJadwal = "unconfirmed"
	StatusEntryJadwalConfirmed   StatusEntryJadwal = "confirmed"
	StatusEntryJadwalOnBid       StatusEntryJadwal = "on-bid"
	StatusEntryJadwalApproved    StatusEntryJadwal = "approved"
	StatusEntryJadwalRejected    StatusEntryJadwal = "rejected"
)

func (o StatusEntryJadwal) IsValid() bool {
	return o == StatusEntryJadwalUnconfirmed || o == StatusEntryJadwalConfirmed ||
		o == StatusEntryJadwalOnBid || o == StatusEntryJadwalApproved || o == StatusEntryJadwalRejected
}

func (EntryJadwal) TableName() string {
	return "t_entry_jadwal"
}

type EntryJadwal struct {
	baseapp.BaseModel
	TJadwalRakorId          *uuid.UUID            `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor             *JadwalRakor          `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TPemasokId              *uuid.UUID            `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                 *Organization         `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId           *uuid.UUID            `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit              *Organization         `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TTransportirId          *uuid.UUID            `gorm:"type:uuid" json:"transportirId"`
	Transportir             *Organization         `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TglTargetLoading        *DateField      `gorm:"type:date" json:"tglTargetLoading"`
	TglEta                  *DateField      `gorm:"type:date" json:"tglEta"`
	TRencanaPasokanJadwalId *uuid.UUID            `gorm:"type:uuid" json:"rencanaPasokanJadwalId"`
	SkemaKontrakCode        *JenisPengiriman      `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak            *ConfigDataInfo       `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	RencanaPasokanJadwal    *RencanaPasokanJadwal `gorm:"foreignKey:TRencanaPasokanJadwalId" json:"rencanaPasokanJadwal"`
	StatusEntryJadwal       StatusEntryJadwal     `gorm:"type:varchar(45);default:'unconfirmed'" json:"statusEntryJadwal"`
	TMasterKapalId          *uuid.UUID            `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal             *MasterKapal          `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	Gcv                     float64               `gorm:"type:float" json:"gcv"`
	TSpekId                 *uuid.UUID            `gorm:"type:uuid" json:"spekId"`
	Spek                    *Spek                 `gorm:"foreignKey:TSpekId" json:"spek"`
	TJenisBatuBaraId        *uuid.UUID            `gorm:"type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara           *JenisBatuBara        `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
	BongkarHariKe           int                   `gorm:"type:int" json:"bongkarHariKe"`
	TotalHariBongkar        int                   `gorm:"type:int" json:"totalHariBongkar"`
	BongkarPengirimanKe     int                   `gorm:"type:int" json:"bongkarPengirimanKe"`
	TotalPengiriman         int                   `gorm:"type:int" json:"totalPengiriman"`
	VolumeHarian            float64               `gorm:"type:float" json:"volumeHarian"`
	TotalVolume             float64               `gorm:"type:float" json:"totalVolume"`
	KapasitasBongkar        float64               `gorm:"type:float" json:"kapasitasBongkar"` // Kapasitas bongkar pembangkit ketika data ini dibuat
}

type TipeEntryJadwalAksi string

const (
	EntryJadwalAksiSend    TipeEntryJadwalAksi = "send"
	EntryJadwalAksiApprove TipeEntryJadwalAksi = "approve"
	EntryJadwalAksiReject  TipeEntryJadwalAksi = "reject"
)

func (o TipeEntryJadwalAksi) IsValid() bool {
	return o == EntryJadwalAksiSend || o == EntryJadwalAksiApprove || o == EntryJadwalAksiReject
}

func (EntryJadwalAksi) TableName() string {
	return "t_entry_jadwal_aksi"
}

type EntryJadwalAksi struct {
	baseapp.BaseModel
	Tipe        TipeEntryJadwalAksi   `gorm:"type:varchar(255)" json:"tipe"`
	Catatan     string                `gorm:"type:text" json:"catatan"`
	EntryJadwal []EntryJadwalAksiItem `gorm:"foreignKey:TEntryJadwalAksiId" json:"EntryJadwal"`
}

func (EntryJadwalAksiItem) TableName() string {
	return "t_entry_jadwal_aksi_item"
}

type EntryJadwalAksiItem struct {
	baseapp.BaseModel
	TEntryJadwalAksiId *uuid.UUID       `gorm:"index" json:"entryJadwalAksiId"`
	EntryJadwalAksi    *EntryJadwalAksi `gorm:"foreignKey:TEntryJadwalAksiId" json:"entryJadwalAksi"`
	TEntryJadwalId     *uuid.UUID       `gorm:"index" json:"entryJadwalId"`
	EntryJadwal        *EntryJadwal     `gorm:"foreignKey:TEntryJadwalId" json:"entryJadwal"`
}
