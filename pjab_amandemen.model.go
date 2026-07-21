package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PjabAmandemen) TableName() string {
	return "t_pjab_amandemen"
}

type PjabAmandemen struct {
	baseapp.BaseModel
	// Amandemen
	TPjabId                  uuid.UUID       `gorm:"index;type:uuid" json:"pjabId"`
	Pjab                     *Pjab           `gorm:"foreignKey:TPjabId" json:"pjab"`
	JudulKontrakCode         *string         `gorm:"index;type:uuid" json:"judulKontrakCode"`
	JudulKontrak             *ConfigData     `gorm:"foreignKey:JudulKontrakCode;references:Code;" json:"judulKontrak"`
	NoAmandemen              string          `gorm:"type:varchar(255)" json:"noAmandemen"`
	NoPenyediaJasaAmd        string          `gorm:"type:varchar(255)" json:"noPenyediaJasaAmd"`
	TglAmandemen             DateField `gorm:"type:date" json:"tglAmandemen"`
	AmandemenKe              int             `gorm:"type:int" json:"amandemenKe"`
	TglBerlakuAmandemenAwal  DateField `gorm:"type:date" json:"tglBerlakuAmandemenAwal"`
	TglBerlakuAmandemenAkhir DateField `gorm:"type:date" json:"tglBerlakuAmandemenAkhir"`
	VolumeAmandemen          float64         `gorm:"type:float" json:"volumeAmandemen"`
	VolumeAmandemenPerTahun  float64         `gorm:"type:float" json:"volumeAmandemenPerTahun"`

	// Perubahan Kontrak
	TDokumenKontrakId *uuid.UUID `gorm:"type:uuid" json:"dokumenKontrakId"`
	DokumenKontrak    *Media     `gorm:"foreignKey:TDokumenKontrakId" json:"dokumenKontrak"`
	AdaDendaSusut     bool       `gorm:"type:bool" json:"adaDendaSusut"`
	PengaliDenda      string     `gorm:"type:varchar(255)" json:"pengaliDenda"`
	PenagihanSusut    string     `gorm:"type:varchar(255)" json:"penagihanSusut"`
	SusutPersen       float64    `gorm:"type:float" json:"susutPersen"`

	TipePeriodeHargaHsd string        `gorm:"type:varchar(20)" json:"tipePeriodeHargaHsd"`
	BulanPeriodeHsd     BulanPeriode  `gorm:"type:varchar(255);default:'BULAN_N'" json:"bulanPeriodeHsd"`
	TanggalAcuanHsd     TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHsd"`
	PeriodeHargaHsdCode *string       `gorm:"type:varchar(255)" json:"periodeHargaHsdCode"`
	PeriodeHargaHsd     *ConfigData   `gorm:"foreignKey:PeriodeHargaHsdCode;references:Code;" json:"periodeHargaHsd"`
	TanggalAcuanHarga   TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHarga"`

	// Termin
	PemilihanTermin bool    `gorm:"type:bool" json:"pemilihanTermin"`
	Termin1         float64 `gorm:"type:float" json:"termin1"`

	Y2VesselSpb        bool                      `gorm:"type:boolean" json:"y2VesselSpb"`
	AdaDesdem          bool                      `gorm:"type:bool" json:"adaDesdem"`
	PenentuanNorDesdem string                    `gorm:"type:varchar(255)" json:"penentuanNorDesdem"`
	PemilihanLsfoHsfo  bool                      `gorm:"type:boolean" json:"pemilihanLsfoHsfo"`
	MenggunakanFc      bool                      `gorm:"type:boolean" json:"menggunakanFc"`
	Fc                 *float64                  `gorm:"type:float" json:"fc"`
	Pembangkits        []PjabAmandemenPembangkit `gorm:"foreignKey:TPjabAmandemenId" json:"pembangkit"`
	Desdems            []PjabAmandemenDesdem     `gorm:"foreignKey:TPjabAmandemenId" json:"desdems"`
	PjabAmdPembiayaan  []PjabAmdPembiayaan       `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmdPembiayaan"`
	NoRekening         []PjabNoRekening          `gorm:"foreignKey:TPjabAmandemenId" json:"noRekening"`

	// Bank Garansi
	AdaBankGaransi bool `gorm:"type:boolean" json:"adaBankGaransi"`
}

func (PjabAmandemenPembangkit) TableName() string {
	return "t_pjab_amandemen_pembangkit"
}

type PjabAmandemenPembangkit struct {
	baseapp.BaseModel
	TPjabAmandemenId uuid.UUID      `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen    *PjabAmandemen `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TOrganizationId  uuid.UUID      `gorm:"index;type:uuid" json:"organizationId"`
	Organization     *Organization  `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjabAmandemenDesdem) TableName() string {
	return "t_pjab_amandemen_desdem"
}

type PjabAmandemenDesdem struct {
	baseapp.BaseModel
	TPjabAmandemenId      uuid.UUID                       `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen         *PjabAmandemen                  `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	HargaDemurageVessel   float64                         `gorm:"type:decimal(15,2)" json:"hargaDemurageVessel"`
	HargaDespatchVessel   float64                         `gorm:"type:decimal(15,2)" json:"hargaDespatchVessel"`
	BatasHariVessel       int                             `gorm:"type:int" json:"batasHariVessel"`
	HargaDemurageSpb      float64                         `gorm:"type:decimal(15,2)" json:"hargaDemurageSpb"`
	HargaDespatchSpb      float64                         `gorm:"type:decimal(15,2)" json:"hargaDespatchSpb"`
	BatasHariSpb          int                             `gorm:"type:int" json:"batasHariSpb"`
	HargaDemurageTongkang float64                         `gorm:"type:decimal(15,2)" json:"hargaDemurageTongkang"`
	HargaDespatchTongkang float64                         `gorm:"type:decimal(15,2)" json:"hargaDespatchTongkang"`
	BatasHariTongkang     int                             `gorm:"type:int" json:"batasHariTongkang"`
	TIsianMulaiId         *uuid.UUID                      `gorm:"type:uuid" json:"isianMulaiId"`
	IsianMulai            *ConfigData                     `gorm:"foreignKey:TIsianMulaiId" json:"isianMulai"`
	TIsianSelesaiId       *uuid.UUID                      `gorm:"type:uuid" json:"isianSelesaiId"`
	IsianSelesai          *ConfigData                     `gorm:"foreignKey:TIsianSelesaiId" json:"isianSelesai"`
	Pembangkits           []PjabAmandemenDesdemPembangkit `gorm:"foreignKey:TPjabAmandemenDesdemId" json:"pembangkits"`
}

func (PjabAmandemenDesdemPembangkit) TableName() string {
	return "t_pjbb_amandemen_desdem_pembangkit"
}

type PjabAmandemenDesdemPembangkit struct {
	baseapp.BaseModel
	TPjabAmandemenId       uuid.UUID            `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	TPjabAmandemenDesdemId uuid.UUID            `gorm:"index;type:uuid" json:"pjabAmandemenDesdemId"`
	PjabAmandemenDesdem    *PjabAmandemenDesdem `gorm:"foreignKey:TPjabAmandemenDesdemId" json:"pjabAmandemenDesdem"`
	TOrganizationId        uuid.UUID            `gorm:"index;type:uuid" json:"organizationId"`
	Organization           *Organization        `gorm:"foreignKey:TOrganizationId" json:"organization"`
}
