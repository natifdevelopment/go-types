package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PjbbAmandemen) TableName() string {
	return "t_pjbb_amandemen"
}

type PjbbAmandemen struct {
	baseapp.BaseModel
	// Amandemen Info Data
	TPjbbId                  *uuid.UUID       `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb                     *Pjbb            `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TKantorIndukId           *uuid.UUID       `gorm:"index;type:uuid" json:"kantorIndukId"`
	KantorInduk              *Organization    `gorm:"foreignKey:TKantorIndukId" json:"kantorInduk"`
	NoAmandemen              string           `gorm:"type:varchar(255)" json:"noAmandemen"`
	NoPemasokAmandemen       string           `gorm:"type:varchar(255)" json:"noPemasokAmandemen"`
	TglAmandemen             *DateField `gorm:"type:date" json:"tglAmandemen"`
	AmandemenKe              int              `gorm:"type:int" json:"amandemenKe"`
	VolumeAmandemen          float64          `gorm:"type:decimal(12,2)" json:"volumeAmandemen"`
	VolumeAmandemenPerTahun  float64          `gorm:"type:decimal(12,2)" json:"volumeAmandemenPerTahun"`
	TglBerlakuAmandemenAwal  *DateField `gorm:"type:date" json:"tglBerlakuAmandemenAwal"`
	TglBerlakuAmandemenAkhir *DateField `gorm:"type:date" json:"tglBerlakuAmandemenAkhir"`

	// Whats changed from base PJBB
	JudulKontrakId     *uuid.UUID           `gorm:"index;type:uuid" json:"judulKontrakId"`
	JudulKontrak       *ConfigData          `gorm:"foreignKey:JudulKontrakId" json:"judulKontrak"`
	JenisKontrakCode   *string              `gorm:"type:varchar(255)" json:"jenisKontrakCode"`
	JenisKontrak       *ConfigDataInfo      `gorm:"foreignKey:JenisKontrakCode;references:Code;" json:"jenisKontrak"`
	Kurs               PjbbKurs             `gorm:"type:varchar(255)" json:"kurs"`
	TipeKurs           *PjbbTipeKurs        `gorm:"type:varchar(255)" json:"tipeKurs"`
	Trunc              uint                 `gorm:"type:int" json:"trunc"`
	SkemaKontrakSkBbo1 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo1"`
	SkemaKontrakSkBbo2 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo2"`
	SkemaKontrakSkBbo3 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo3"`
	MenggunakanFc      bool                 `gorm:"type:boolean" json:"menggunakanFc"`
	Fc                 *float64             `gorm:"type:float" json:"fc"`
	PjbbAmdRumus       []PjbbRumus          `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmdRumus"`
	SumberTambang      []PjbbSumberTambang  `gorm:"foreignKey:TPjbbAmandemenId" json:"sumberTambang"`
	JenisBatuBara      []PjbbJenisBatubara  `gorm:"foreignKey:TPjbbAmandemenId" json:"jenisBatuBara"`
	NoRekening         []PjbbNoRekening     `gorm:"foreignKey:TPjbbAmandemenId" json:"noRekening"`
	VolumePerTahun     []PjbbVolumePerTahun `gorm:"foreignKey:TPjbbAmandemenId" json:"volumePerTahun"`
	Items              []PjbbAmandemenItem  `gorm:"foreignKey:TPjbbAmandemenId" json:"items"`

	// Bank Garansi
	AdaBankGaransi    bool       `gorm:"type:boolean" json:"adaBankGaransi"`
	NoBankGaransi     *string    `gorm:"type:varchar(255)" json:"noBankGaransi"`
	TDokBankGaransiId *uuid.UUID `gorm:"type:uuid" json:"dokBankGaransiId"`
	DokBankGaransi    *Media     `gorm:"foreignKey:TDokBankGaransiId" json:"dokBankGaransi"`
}

func (PjbbAmandemenItem) TableName() string {
	return "t_pjbb_amandemen_item"
}

type PjbbAmandemenItem struct {
	baseapp.BaseModel
	TPjbbAmandemenId uuid.UUID        `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen   `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	SkemaKontrakCode *JenisPengiriman `gorm:"type:varchar(255)" json:"skemaKontrakCode"`
	SkemaKontrak     *ConfigDataInfo  `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`

	// Common fields
	TDokumenKontrakId  *uuid.UUID    `gorm:"type:uuid" json:"dokumenKontrakId"`
	DokumenKontrak     *Media        `gorm:"foreignKey:TDokumenKontrakId" json:"dokumenKontrak"`
	AdaCsr             bool          `gorm:"type:bool" json:"adaCsr"`
	DokSkabSyaratBayar bool          `gorm:"type:bool" json:"dokSkabSyaratBayar"`
	TanggalAcuanHarga  TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHarga"`

	// Common bank garansi
	AdaBankGaransi              bool             `gorm:"type:boolean" json:"adaBankGaransi"`
	TDokumenBankGaransiId       *uuid.UUID       `gorm:"type:uuid" json:"dokumenBankGaransiId"`
	DokumenBankGaransi          *Media           `gorm:"foreignKey:TDokumenBankGaransiId" json:"dokumenBankGaransi"`
	NoBankGaransi               *string          `gorm:"type:varchar(255)" json:"noBankGaransi"`
	TanggalBankGaransi          *DateField `gorm:"type:date" json:"tanggalBankGaransi"`
	JangkaWaktuBankGaransiAwal  *DateField `gorm:"type:date" json:"jangkaWaktuBankGaransiAwal"`
	JangkaWaktuBankGaransiAkhir *DateField `gorm:"type:date" json:"jangkaWaktuBankGaransiAkhir"`

	// CIF and CFR common fields
	DendaKeterlambatan       bool    `gorm:"type:bool" json:"dendaKeterlambatan"`
	ToleransiVolumeRakor     float64 `gorm:"type:float" json:"toleransiVolumeRakor"`
	ToleransiVolumeTerlambat float64 `gorm:"type:float" json:"toleransiVolumeTerlambat"`
	BatasMaksHariDenda       float64 `gorm:"type:float" json:"batasMaksHariDenda"`
	PemilihanHargaPertamina  bool    `gorm:"type:bool" json:"pemilihanHargaPertamina"`
	BiayaTrucking            bool    `gorm:"type:bool" json:"biayaTrucking"`

	// FOB speCIFic fields
	DendaMuat             bool    `gorm:"type:bool" json:"dendaMuat"`
	Lhv                   bool    `gorm:"type:bool" json:"lhv"`
	MaksHariDendaVessel   float64 `gorm:"type:float" json:"maksHariDendaVessel"`
	MaksHariDendaSpb      float64 `gorm:"type:float" json:"maksHariDendaSpb"`
	MaksHariDendaTongkang float64 `gorm:"type:float" json:"maksHariDendaTongkang"`
	ToleransiDenda        float64 `gorm:"type:float" json:"toleransiDenda"`
	PersentaseDenda       float64 `gorm:"type:float" json:"persentaseDenda"`
	AdaRepro              bool    `gorm:"type:bool" json:"adaRepro"`

	// CIF speCIFic fields
	TipePeriodeHargaHsd string      `gorm:"type:varchar(20)" json:"tipePeriodeHargaHsd"`
	PeriodeHargaHsdCode *string     `gorm:"type:varchar(255)" json:"periodeHargaHsdCode"`
	PeriodeHargaHsd     *ConfigData `gorm:"foreignKey:PeriodeHargaHsdCode;references:Code;" json:"periodeHargaHsd"`
	PemilihanTermin     bool        `gorm:"type:bool" json:"pemilihanTermin"`
	Termin1             float64     `gorm:"type:float" json:"termin1"`
	AdaDesdem           bool        `gorm:"type:bool" json:"adaDesdem"`

	// New CIF specific fields
	BulanPeriodeHsd BulanPeriode  `gorm:"type:varchar(255);default:'BULAN_N'" json:"bulanPeriodeHsd"`
	TanggalAcuanHsd TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHsd"`

	// CFR speCIFic fields
	PerLot       bool    `gorm:"type:bool" json:"perLot"`
	VolumePerLot float64 `gorm:"type:float" json:"volumePerLot"`

	Pembangkits []PjbbAmandemenItemPembangkit `gorm:"foreignKey:TPjbbAmandemenItemId" json:"pembangkit"`
	Desdems     []PjbbAmandemenItemDesdem     `gorm:"foreignKey:TPjbbAmandemenItemId" json:"desdems"`
}

func (PjbbAmandemenItemPembangkit) TableName() string {
	return "t_pjbb_amandemen_item_pembangkit"
}

type PjbbAmandemenItemPembangkit struct {
	baseapp.BaseModel
	TPjbbAmandemenId     uuid.UUID          `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	TPjbbAmandemenItemId uuid.UUID          `gorm:"index;type:uuid" json:"pjbbAmandemenItemId"`
	PjbbAmandemenItem    *PjbbAmandemenItem `gorm:"foreignKey:TPjbbAmandemenItemId" json:"pjbbAmandemenItem"`
	TOrganizationId      uuid.UUID          `gorm:"index" json:"organizationId"`
	Organization         *Organization      `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjbbAmandemenItemDesdem) TableName() string {
	return "t_pjbb_amandemen_item_desdem"
}

type PjbbAmandemenItemDesdem struct {
	baseapp.BaseModel
	TPjbbAmandemenId      uuid.UUID                           `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	TPjbbAmandemenItemId  uuid.UUID                           `gorm:"index;type:uuid" json:"pjbbAmandemenItemId"`
	PjbbAmandemenItem     *PjbbAmandemenItem                  `gorm:"foreignKey:TPjbbAmandemenItemId" json:"pjbbAmandemenItem"`
	HargaDemurageVessel   float64                             `gorm:"type:decimal(15,2)" json:"hargaDemurageVessel"`
	HargaDespatchVessel   float64                             `gorm:"type:decimal(15,2)" json:"hargaDespatchVessel"`
	BatasHariVessel       int                                 `gorm:"type:int" json:"batasHariVessel"`
	HargaDemurageSpb      float64                             `gorm:"type:decimal(15,2)" json:"hargaDemurageSpb"`
	HargaDespatchSpb      float64                             `gorm:"type:decimal(15,2)" json:"hargaDespatchSpb"`
	BatasHariSpb          int                                 `gorm:"type:int" json:"batasHariSpb"`
	HargaDemurageTongkang float64                             `gorm:"type:decimal(15,2)" json:"hargaDemurageTongkang"`
	HargaDespatchTongkang float64                             `gorm:"type:decimal(15,2)" json:"hargaDespatchTongkang"`
	BatasHariTongkang     int                                 `gorm:"type:int" json:"batasHariTongkang"`
	TIsianMulaiId         *uuid.UUID                          `gorm:"type:uuid" json:"isianMulaiId"`
	IsianMulai            *ConfigData                         `gorm:"foreignKey:TIsianMulaiId" json:"isianMulai"`
	TIsianSelesaiId       *uuid.UUID                          `gorm:"type:uuid" json:"isianSelesaiId"`
	IsianSelesai          *ConfigData                         `gorm:"foreignKey:TIsianSelesaiId" json:"isianSelesai"`
	Pembangkits           []PjbbAmandemenItemDesdemPembangkit `gorm:"foreignKey:TPjbbAmandemenItemDesdemId" json:"pembangkits"`
}

func (PjbbAmandemenItemDesdemPembangkit) TableName() string {
	return "t_pjbb_amandemen_item_desdem_pembangkit"
}

type PjbbAmandemenItemDesdemPembangkit struct {
	baseapp.BaseModel
	TPjbbAmandemenId           uuid.UUID                `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	TPjbbAmandemenItemDesdemId uuid.UUID                `gorm:"index;type:uuid" json:"pjbbAmandemenItemDesdemId"`
	PjbbAmandemenItemDesdem    *PjbbAmandemenItemDesdem `gorm:"foreignKey:TPjbbAmandemenItemDesdemId" json:"pjbbAmandemenItemDesdem"`
	TOrganizationId            uuid.UUID                `gorm:"index;type:uuid" json:"organizationId"`
	Organization               *Organization            `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjbbAmandemenSumberTambang) TableName() string {
	return "t_pjbb_amandemen_sumber_tambang"
}

type PjbbAmandemenSumberTambang struct {
	baseapp.BaseModel
	TPjbbAmandemenId uuid.UUID      `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TSumberTambangId uuid.UUID      `gorm:"index;type:uuid" json:"sumberTambangId"`
	SumberTambang    *SumberTambang `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
}
