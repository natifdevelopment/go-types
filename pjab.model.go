package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Pjab) TableName() string {
	return "t_pjab"
}

type Pjab struct {
	baseapp.BaseModel
	JudulKontrakCode   *string          `gorm:"index;type:uuid" json:"judulKontrakCode"`
	JudulKontrak       *ConfigDataInfo  `gorm:"foreignKey:JudulKontrakCode;references:Code;" json:"judulKontrak"`
	TTransportirId     *uuid.UUID       `gorm:"type:uuid" json:"transportirId"`
	Transportir        *Organization    `gorm:"foreignKey:TTransportirId" json:"transportir"`
	NoPjab             string           `gorm:"type:varchar(255)" json:"noPjab"`
	NoPenyediaJasa     string           `gorm:"type:varchar(255)" json:"noPenyediaJasa"`
	TglPjab            *DateField `gorm:"type:date" json:"tglPjab"`
	Jaminan            string           `gorm:"type:varchar(255)" json:"jaminan"`
	MasaBerlakuAwal    *DateField `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir   *DateField `gorm:"type:date" json:"masaBerlakuAkhir"`
	VolumePjab         float64          `gorm:"type:float" json:"volumePjab"`
	VolumePjabPerTahun float64          `gorm:"type:float" json:"volumePjabPerTahun"`
	LrDrp              *float64         `gorm:"type:float" json:"lrDrp"`
	AdaDendaSusut      bool             `gorm:"type:bool" json:"adaDendaSusut"`
	PengaliDenda       string           `gorm:"type:varchar(255)" json:"pengaliDenda"`
	PenagihanSusut     string           `gorm:"type:varchar(255)" json:"penagihanSusut"`
	SusutPersen        float64          `gorm:"type:float" json:"susutPersen"`
	AdaDesdem          bool             `gorm:"type:bool" json:"adaDesdem"`
	PenentuanNorDesdem string           `gorm:"type:varchar(255)" json:"penentuanNorDesdem"`
	Y2VesselSpb        bool             `gorm:"type:boolean" json:"y2VesselSpb"`

	TipePeriodeHargaHsd string        `gorm:"type:varchar(20)" json:"tipePeriodeHargaHsd"`
	BulanPeriodeHsd     BulanPeriode  `gorm:"type:varchar(255);default:'BULAN_N'" json:"bulanPeriodeHsd"`
	TanggalAcuanHsd     TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHsd"`
	PeriodeHargaHsdCode *string       `gorm:"type:varchar(255)" json:"periodeHargaHsdCode"`
	PeriodeHargaHsd     *ConfigData   `gorm:"foreignKey:PeriodeHargaHsdCode;references:Code;" json:"periodeHargaHsd"`
	TanggalAcuanHarga   TglAcuanHarga `gorm:"type:varchar(20)" json:"tanggalAcuanHarga"`

	// Termin
	PemilihanTermin bool    `gorm:"type:bool" json:"pemilihanTermin"`
	Termin1         float64 `gorm:"type:float" json:"termin1"`

	PemilihanLsfoHsfo bool             `gorm:"type:boolean" json:"pemilihanLsfoHsfo"`
	MenggunakanFc     bool             `gorm:"type:boolean" json:"menggunakanFc"`
	Fc                *float64         `gorm:"type:float" json:"fc"`
	TDokumenKontrakId *uuid.UUID       `gorm:"type:uuid" json:"dokumenKontrakId"`
	DokumenKontrak    *Media           `gorm:"foreignKey:TDokumenKontrakId" json:"dokumenKontrak"`
	Pembangkits       []PjabPembangkit `gorm:"foreignKey:TPjabId" json:"pembangkit"`
	Desdems           []PjabDesdem     `gorm:"foreignKey:TPjabId" json:"desdems"`
	PjabRumus         []PjabRumus      `gorm:"foreignKey:TPjabId" json:"pjabRumus"`
	PjabPembiayaan    []PjabPembiayaan `gorm:"foreignKey:TPjabId" json:"pjabPembiayaan"`
	PjabAmandemen     []PjabAmandemen  `gorm:"foreignKey:TPjabId" json:"pjabAmandemen"`
	NoRekening        []PjabNoRekening `gorm:"foreignKey:TPjabId" json:"noRekening"`

	// Bank Garansi
	AdaBankGaransi bool `gorm:"type:boolean" json:"adaBankGaransi"`
}

func (PjabPembangkit) TableName() string {
	return "t_pjab_pembangkit"
}

type PjabPembangkit struct {
	baseapp.BaseModel
	TPjabId         *uuid.UUID    `gorm:"index;type:uuid" json:"pjabId"`
	Pjab            *Pjab         `gorm:"foreignKey:TPjabId" json:"pjab"`
	TOrganizationId *uuid.UUID    `gorm:"index;type:uuid" json:"organizationId"`
	Organization    *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjabDesdem) TableName() string {
	return "t_pjab_desdem"
}

type PjabDesdem struct {
	baseapp.BaseModel
	TPjabId               uuid.UUID              `gorm:"index;type:uuid" json:"pjabId"`
	Pjab                  *Pjab                  `gorm:"foreignKey:TPjabId" json:"pjab"`
	HargaDemurageVessel   float64                `gorm:"type:decimal(15,2)" json:"hargaDemurageVessel"`
	HargaDespatchVessel   float64                `gorm:"type:decimal(15,2)" json:"hargaDespatchVessel"`
	BatasHariVessel       int                    `gorm:"type:int" json:"batasHariVessel"`
	HargaDemurageSpb      float64                `gorm:"type:decimal(15,2)" json:"hargaDemurageSpb"`
	HargaDespatchSpb      float64                `gorm:"type:decimal(15,2)" json:"hargaDespatchSpb"`
	BatasHariSpb          int                    `gorm:"type:int" json:"batasHariSpb"`
	HargaDemurageTongkang float64                `gorm:"type:decimal(15,2)" json:"hargaDemurageTongkang"`
	HargaDespatchTongkang float64                `gorm:"type:decimal(15,2)" json:"hargaDespatchTongkang"`
	BatasHariTongkang     int                    `gorm:"type:int" json:"batasHariTongkang"`
	TIsianMulaiId         *uuid.UUID             `gorm:"type:uuid" json:"isianMulaiId"`
	IsianMulai            *ConfigData            `gorm:"foreignKey:TIsianMulaiId" json:"isianMulai"`
	TIsianSelesaiId       *uuid.UUID             `gorm:"type:uuid" json:"isianSelesaiId"`
	IsianSelesai          *ConfigData            `gorm:"foreignKey:TIsianSelesaiId" json:"isianSelesai"`
	Pembangkits           []PjabDesdemPembangkit `gorm:"foreignKey:TPjabDesdemId" json:"pembangkits"`
}

func (PjabDesdemPembangkit) TableName() string {
	return "t_pjbb_desdem_pembangkit"
}

type PjabDesdemPembangkit struct {
	baseapp.BaseModel
	TPjabId         uuid.UUID     `gorm:"index;type:uuid" json:"pjabId"`
	TPjabDesdemId   uuid.UUID     `gorm:"index;type:uuid" json:"pjabDesdemId"`
	PjabDesdem      *PjabDesdem   `gorm:"foreignKey:TPjabDesdemId" json:"pjabDesdem"`
	TOrganizationId uuid.UUID     `gorm:"index;type:uuid" json:"organizationId"`
	Organization    *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjabNoRekening) TableName() string {
	return "t_pjab_no_rekening"
}

type PjabNoRekening struct {
	baseapp.BaseModel
	TPjabId             *uuid.UUID        `gorm:"index;type:uuid" json:"pjabId"`
	Pjab                *Pjab             `gorm:"foreignKey:TPjabId" json:"pjab"`
	TMasterNoRekeningId uuid.UUID         `gorm:"index;type:uuid" json:"masterNoRekeningId"`
	MasterNoRekening    *MasterNoRekening `gorm:"foreignKey:TMasterNoRekeningId" json:"masterNoRekening"`

	TPjabAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen    *PjabAmandemen `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
}
