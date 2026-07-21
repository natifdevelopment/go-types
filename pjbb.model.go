package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Pjbb) TableName() string {
	return "t_pjbb"
}

type TglAcuanHarga string

const (
	TglAcuanHargaTaKonfirmasi   TglAcuanHarga = "TA_Konfirmasi"
	TglAcuanHargaBl             TglAcuanHarga = "BL"
	TglAcuanHargaTa             TglAcuanHarga = "TA"
	TglAcuanHargaSandar         TglAcuanHarga = "Sandar"
	TglAcuanHargaMulaiBongkar   TglAcuanHarga = "Mulai_Bongkar"
	TglAcuanHargaSelesaiBongkar TglAcuanHarga = "Selesai_Bongkar"
	TglAcuanHargaMulaiMuat      TglAcuanHarga = "Mulai_Muat"
	TglAcuanHargaSelesaiMuat    TglAcuanHarga = "Selesai_Muat"
)

type BulanPeriode string

const (
	PeriodeBulanN     BulanPeriode = "BULAN_N"
	PeriodeBulanNMin1 BulanPeriode = "BULAN_N_MIN_1"
)

type PjbbKurs string

const (
	PjbbKursTengah PjbbKurs = "Tengah"
	PjbbKursJisdor PjbbKurs = "Jisdor"
	PjbbKursNo     PjbbKurs = "No"
)

func (t PjbbKurs) IsValid() bool {
	return t == PjbbKursTengah || t == PjbbKursJisdor || t == PjbbKursNo
}

type PjbbTipeKurs string

const (
	PjbbTipeKursHarian  PjbbTipeKurs = "Harian"
	PjbbTipeKursBulanan PjbbTipeKurs = "Bulanan"
)

type Pjbb struct {
	baseapp.BaseModel
	TPemasokId         *uuid.UUID           `gorm:"index;type:uuid" json:"pemasokId"`
	Pemasok            *Organization        `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TKantorIndukId     *uuid.UUID           `gorm:"index;type:uuid" json:"kantorIndukId"`
	KantorInduk        *Organization        `gorm:"foreignKey:TKantorIndukId" json:"kantorInduk"`
	TJudulKontrakId    *uuid.UUID           `gorm:"index;type:uuid" json:"judulKontrakId"`
	JudulKontrak       *ConfigDataInfo      `gorm:"foreignKey:TJudulKontrakId;" json:"judulKontrak"`
	JenisKontrakCode   *string              `gorm:"type:varchar(255)" json:"jenisKontrakCode"`
	JenisKontrak       *ConfigDataInfo      `gorm:"foreignKey:JenisKontrakCode;references:Code;" json:"jenisKontrak"`
	Kurs               PjbbKurs             `gorm:"type:varchar(255)" json:"kurs"`
	TipeKurs           *PjbbTipeKurs        `gorm:"type:varchar(255)" json:"tipeKurs"`
	MenggunakanFc      bool                 `gorm:"type:boolean" json:"menggunakanFc"`
	Fc                 *float64             `gorm:"type:float" json:"fc"`
	Trunc              uint                 `gorm:"type:int" json:"trunc"`
	SkemaKontrakSkBbo1 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo1"`
	SkemaKontrakSkBbo2 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo2"`
	SkemaKontrakSkBbo3 bool                 `gorm:"type:boolean" json:"skemaKontrakSkBbo3"`
	NoPemasok          string               `gorm:"type:varchar(255)" json:"noPemasok"`
	NoPjbb             string               `gorm:"type:varchar(255)" json:"noPjbb"`
	TanggalPjbb        *DateField     `gorm:"type:date" json:"tanggalPjbb"`
	TglBerlakuAwal     *DateField     `gorm:"type:date" json:"tglBerlakuAwal"`
	TglBerlakuAkhir    *DateField     `gorm:"type:date" json:"tglBerlakuAkhir"`
	TJangkaWaktuId     *uuid.UUID           `gorm:"type:uuid" json:"jangkaWaktuId"`
	JangkaWaktu        *JangkaWaktu         `gorm:"foreignKey:TJangkaWaktuId" json:"jangkaWaktu"`
	VolumePjbb         float64              `gorm:"type:float" json:"volumePjbb"`
	VolumePjbbPerTahun float64              `gorm:"type:float" json:"volumePjbbPerTahun"`
	Items              []PjbbItem           `gorm:"foreignKey:TPjbbId" json:"items"`
	PjbbRumus          []PjbbRumus          `gorm:"foreignKey:TPjbbId" json:"pjbbRumus"`
	SumberTambang      []PjbbSumberTambang  `gorm:"foreignKey:TPjbbId" json:"sumberTambang"`
	PjbbHarga          []PjbbHarga          `gorm:"foreignKey:TPjbbId" json:"pjbbHargas"`
	PjbbHargaTriwulan  []PjbbHargaTriwulan  `gorm:"foreignKey:TPjbbId" json:"PjbbHargaTriwulan"`
	PjbbAmandemen      []PjbbAmandemen      `gorm:"foreignKey:TPjbbId" json:"pjbbAmandemen"`
	JenisBatuBara      []PjbbJenisBatubara  `gorm:"foreignKey:TPjbbId" json:"jenisBatuBara"`
	NoRekening         []PjbbNoRekening     `gorm:"foreignKey:TPjbbId" json:"noRekening"`
	VolumePerTahun     []PjbbVolumePerTahun `gorm:"foreignKey:TPjbbId" json:"volumePerTahun"`
}

func (PjbbJenisBatubara) TableName() string {
	return "t_pjbb_jenis_batubara"
}

type PjbbJenisBatubara struct {
	baseapp.BaseModel
	TPjbbId          *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb          `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TJenisBatuBaraId uuid.UUID      `gorm:"index;type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara    *JenisBatuBara `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
}

func (PjbbVolumePerTahun) TableName() string {
	return "t_pjbb_volume_per_tahun"
}

type PjbbVolumePerTahun struct {
	baseapp.BaseModel
	TPjbbId          *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb          `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPembangkitId    *uuid.UUID     `gorm:"index" json:"pembangkitId"`
	Pembangkit       *Organization  `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	VolumePerTahun   float64        `gorm:"type:float" json:"volumePerTahun"`
	TPjbbAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
}

func (PjbbSumberTambang) TableName() string {
	return "t_pjbb_sumber_tambang"
}

type PjbbSumberTambang struct {
	baseapp.BaseModel
	TPjbbId          *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb          `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TSumberTambangId uuid.UUID      `gorm:"index;type:uuid" json:"sumberTambangId"`
	SumberTambang    *SumberTambang `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	//for amandemen sumber tambang, move to PjbbAmandemenSumberTambang
	TPjbbAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
}

func (PjbbNoRekening) TableName() string {
	return "t_pjbb_no_rekening"
}

type PjbbNoRekening struct {
	baseapp.BaseModel
	TPjbbId             *uuid.UUID        `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb                *Pjbb             `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TMasterNoRekeningId uuid.UUID         `gorm:"index;type:uuid" json:"masterNoRekeningId"`
	MasterNoRekening    *MasterNoRekening `gorm:"foreignKey:TMasterNoRekeningId" json:"masterNoRekening"`
	//for amandemen sumber tambang, move to PjbbAmandemenSumberTambang
	TPjbbAmandemenId *uuid.UUID     `gorm:"index;type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen    *PjbbAmandemen `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
}

func (PjbbItem) TableName() string {
	return "t_pjbb_item"
}

type PjbbItem struct {
	baseapp.BaseModel
	TPjbbId          uuid.UUID        `gorm:"index;type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb            `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	SkemaKontrakCode *JenisPengiriman `gorm:"type:varchar(255)" json:"skemaKontrakCode"`
	SkemaKontrak     *ConfigDataInfo  `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`

	// Common fields
	TDokumenKontrakId  *uuid.UUID    `gorm:"type:uuid" json:"dokumenKontrakId"`
	DokumenKontrak     *Media        `gorm:"foreignKey:TDokumenKontrakId" json:"dokumenKontrak"`
	DokSkabSyaratBayar bool          `gorm:"type:bool" json:"dokSkabSyaratBayar"`
	AdaCsr             bool          `gorm:"type:bool" json:"adaCsr"`
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

	// FOB specific fields
	DendaMuat             bool    `gorm:"type:bool" json:"dendaMuat"`
	Lhv                   bool    `gorm:"type:bool" json:"lhv"`
	MaksHariDendaVessel   float64 `gorm:"type:float" json:"maksHariDendaVessel"`
	MaksHariDendaSpb      float64 `gorm:"type:float" json:"maksHariDendaSpb"`
	MaksHariDendaTongkang float64 `gorm:"type:float" json:"maksHariDendaTongkang"`
	ToleransiDenda        float64 `gorm:"type:float" json:"toleransiDenda"`
	PersentaseDenda       float64 `gorm:"type:float" json:"persentaseDenda"`
	AdaRepro              bool    `gorm:"type:bool" json:"adaRepro"`

	// CIF specific fields
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

	Pembangkits []PjbbItemPembangkit `gorm:"foreignKey:TPjbbItemId" json:"pembangkit"`
	Desdems     []PjbbItemDesdem     `gorm:"foreignKey:TPjbbItemId" json:"desdems"`
}

func (PjbbItemPembangkit) TableName() string {
	return "t_pjbb_item_pembangkit"
}

type PjbbItemPembangkit struct {
	baseapp.BaseModel
	TPjbbId         uuid.UUID     `gorm:"index;type:uuid" json:"pjbbId"`
	TPjbbItemId     uuid.UUID     `gorm:"index;type:uuid" json:"pjbbItemId"`
	PjbbItem        *PjbbItem     `gorm:"foreignKey:TPjbbItemId" json:"pjbbItem"`
	TOrganizationId uuid.UUID     `gorm:"index" json:"organizationId"`
	Organization    *Organization `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (PjbbItemDesdem) TableName() string {
	return "t_pjbb_item_desdem"
}

type PjbbItemDesdem struct {
	baseapp.BaseModel
	TPjbbId               uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbId"`
	TPjbbItemId           uuid.UUID                  `gorm:"index;type:uuid" json:"pjbbItemId"`
	PjbbItem              *PjbbItem                  `gorm:"foreignKey:TPjbbItemId" json:"pjbbItem"`
	HargaDemurageVessel   float64                    `gorm:"type:float" json:"hargaDemurageVessel"`
	HargaDespatchVessel   float64                    `gorm:"type:float" json:"hargaDespatchVessel"`
	BatasHariVessel       int                        `gorm:"type:int" json:"batasHariVessel"`
	HargaDemurageSpb      float64                    `gorm:"type:float" json:"hargaDemurageSpb"`
	HargaDespatchSpb      float64                    `gorm:"type:float" json:"hargaDespatchSpb"`
	BatasHariSpb          int                        `gorm:"type:int" json:"batasHariSpb"`
	HargaDemurageTongkang float64                    `gorm:"type:float" json:"hargaDemurageTongkang"`
	HargaDespatchTongkang float64                    `gorm:"type:float" json:"hargaDespatchTongkang"`
	BatasHariTongkang     int                        `gorm:"type:int" json:"batasHariTongkang"`
	TIsianMulaiId         *uuid.UUID                 `gorm:"type:uuid" json:"isianMulaiId"`
	IsianMulai            *ConfigData                `gorm:"foreignKey:TIsianMulaiId" json:"isianMulai"`
	TIsianSelesaiId       *uuid.UUID                 `gorm:"type:uuid" json:"isianSelesaiId"`
	IsianSelesai          *ConfigData                `gorm:"foreignKey:TIsianSelesaiId" json:"isianSelesai"`
	Pembangkits           []PjbbItemDesdemPembangkit `gorm:"foreignKey:TPjbbItemDesdemId" json:"pembangkits"`
}

func (PjbbItemDesdemPembangkit) TableName() string {
	return "t_pjbb_item_desdem_pembangkit"
}

type PjbbItemDesdemPembangkit struct {
	baseapp.BaseModel
	TPjbbId           uuid.UUID       `gorm:"index;type:uuid" json:"pjbbId"`
	TPjbbItemDesdemId uuid.UUID       `gorm:"index;type:uuid" json:"pjbbItemDesdemId"`
	PjbbItemDesdem    *PjbbItemDesdem `gorm:"foreignKey:TPjbbItemDesdemId" json:"pjbbItemDesdem"`
	TOrganizationId   uuid.UUID       `gorm:"index;type:uuid" json:"organizationId"`
	Organization      *Organization   `gorm:"foreignKey:TOrganizationId" json:"organization"`
}
