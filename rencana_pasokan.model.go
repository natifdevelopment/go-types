package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type StatusRencanaPasokan string

const (
	StatusRencanaPasokanUnconfirmed StatusRencanaPasokan = "unconfirmed"
	StatusRencanaPasokanConfirmed   StatusRencanaPasokan = "confirmed"
	StatusRencanaPasokanApproved    StatusRencanaPasokan = "approved"
	StatusRencanaPasokanRejected    StatusRencanaPasokan = "rejected"
)

func (o StatusRencanaPasokan) IsValid() bool {
	return o == StatusRencanaPasokanUnconfirmed || o == StatusRencanaPasokanConfirmed ||
		o == StatusRencanaPasokanApproved || o == StatusRencanaPasokanRejected
}

func (RencanaPasokan) TableName() string {
	return "t_rencana_pasokan"
}

type RencanaPasokan struct {
	baseapp.BaseModel
	TJadwalPraRakorId     *uuid.UUID                      `gorm:"type:uuid" json:"jadwalPraRakorId"`
	JadwalPraRakor        *JadwalPraRakor                 `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalPraRakor"`
	TOrganizationId       *uuid.UUID                      `gorm:"type:uuid" json:"organizationId"` // PLTU
	Organization          *Organization                   `gorm:"foreignKey:TOrganizationId" json:"organization"`
	Periode               *YearMonthField           `gorm:"type:date" json:"periode"`
	KapasitasPembangkitan float64                         `gorm:"type:float" json:"kapasitasPembangkitan"`
	RencanaOperasiKit     float64                         `gorm:"type:float" json:"rencanaOperasiKit"`
	Sfc                   float64                         `gorm:"type:float" json:"sfc"`
	Cf                    float64                         `gorm:"type:float" json:"cf"`
	RencanaKonsumsi       float64                         `gorm:"type:float" json:"rencanaKonsumsi"`
	StokAwalBulan         float64                         `gorm:"type:float" json:"stokAwalBulan"`
	HopAwalBulan          float64                         `gorm:"type:float" json:"hopAwalBulan"`
	StokAkhirBulan        float64                         `gorm:"type:float" json:"stokAkhirBulan"`
	HopAkhirBulan         float64                         `gorm:"type:float" json:"hopAkhirBulan"`
	ListPemasok           []RencanaPasokanPemasok         `gorm:"foreignKey:TRencanaPasokanId" json:"listPemasok"`
	ListJadwal            []RencanaPasokanJadwal          `gorm:"foreignKey:TRencanaPasokanId" json:"listJadwal"`
	HariOperasiUnit       []RencanaPasokanHariOperasiUnit `gorm:"foreignKey:TRencanaPasokanId" json:"hariOperasiUnit"`
	StatusRencanaPasokan  StatusRencanaPasokan            `gorm:"type:varchar(45);default:'unconfirmed'" json:"statusRencanaPasokan"`
	Hop                   float64                         `gorm:"-" json:"hop"`
}

func (RencanaPasokanHariOperasiUnit) TableName() string {
	return "t_rencana_pasokan_unit"
}

type RencanaPasokanHariOperasiUnit struct {
	baseapp.BaseModel
	TRencanaPasokanId *uuid.UUID      `gorm:"index" json:"rencanaPasokanId"`
	RencanaPasokan    *RencanaPasokan `gorm:"foreignKey:TRencanaPasokanId" json:"rencanaPasokan"`
	TMasterUnitId     *uuid.UUID      `gorm:"index" json:"masterUnitId"`
	MasterUnit        *MasterUnit     `gorm:"foreignKey:TMasterUnitId" json:"masterUnit"`
	KodeMesin         *string         `gorm:"type:varchar(255)" json:"kodeMesin"`
	HariOperasi       int             `gorm:"type:int" json:"hariOperasi"`
}

func (RencanaPasokanPemasok) TableName() string {
	return "t_rencana_pasokan_pemasok"
}

type RencanaPasokanPemasok struct {
	baseapp.BaseModel
	TRencanaPasokanId      *uuid.UUID             `gorm:"index" json:"rencanaPasokanId"`
	RencanaPasokan         *RencanaPasokan        `gorm:"foreignKey:TRencanaPasokanId" json:"rencanaPasokan"`
	TPjbbId                *uuid.UUID             `gorm:"index" json:"pjbbId"`
	Pjbb                   *Pjbb                  `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPemasokId             *uuid.UUID             `gorm:"index" json:"pemasokId"` // Dari PJBB
	Pemasok                *Organization          `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	SkemaKontrakCode       *JenisPengiriman       `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak           *ConfigDataInfo        `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TSpekId                *uuid.UUID             `gorm:"type:uuid" json:"spekId"` // Dari gcv
	Spek                   *Spek                  `gorm:"foreignKey:TSpekId" json:"spek"`
	TJenisBatuBaraId       *uuid.UUID             `gorm:"type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara          *JenisBatuBara         `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
	SaldoPasokan           *float64               `gorm:"type:float" json:"saldoPasokan"`
	TglTargetLoading       *DateField       `gorm:"type:date" json:"tglTargetLoading"`
	TglKonfirmasi          DateField        `gorm:"type:date" json:"tglKonfirmasi"`
	EstimasiSelesaiBongkar DateField        `gorm:"type:date" json:"estimasiSelesaiBongkar"`
	MaksBongkarHarian      float64                `gorm:"type:float" json:"maksBongkarHarian"`
	VolumeTotal            float64                `gorm:"type:float" json:"volumeTotal"` // Total kebutuhan yang di input pembangkit
	IsCarryOver            bool                   `json:"isCarryOver"`
	TotalHariBongkar       float64                `gorm:"type:float" json:"totalHariBongkar"`
	AlokasiPasokan         float64                `gorm:"type:float" json:"alokasiPasokan"` // Yang ditetapkan oleh epi
	RealisasiPasokan       float64                `gorm:"type:float" json:"realisasiPasokan"`
	JumlahPengiriman       int                    `gorm:"type:int;default:1" json:"jumlahPengiriman"`
	Gcv                    float64                `gorm:"type:float" json:"gcv"`
	Jadwal                 []RencanaPasokanJadwal `gorm:"foreignKey:TRencanaPasokanPemasokId" json:"jadwal"`
}

func (RencanaPasokanJadwal) TableName() string {
	return "t_rencana_pasokan_jadwal"
}

type RencanaPasokanJadwal struct {
	baseapp.BaseModel
	TRencanaPasokanPemasokId *uuid.UUID             `gorm:"index" json:"rencanaPasokanPemasokId"`
	RencanaPasokanPemasok    *RencanaPasokanPemasok `gorm:"foreignKey:TRencanaPasokanPemasokId" json:"rencanaPasokanPemasok"`
	TRencanaPasokanId        *uuid.UUID             `gorm:"index" json:"rencanaPasokanId"`
	RencanaPasokan           *RencanaPasokan        `gorm:"foreignKey:TRencanaPasokanId" json:"rencanaPasokan"`
	TPemasokId               *uuid.UUID             `gorm:"index" json:"pemasokId"`
	Pemasok                  *Organization          `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	BongkarHariKe            int                    `gorm:"type:int" json:"bongkarHariKe"`
	TotalHariBongkar         int                    `gorm:"type:int" json:"totalHariBongkar"`
	BongkarPengirimanKe      int                    `gorm:"type:int" json:"bongkarPengirimanKe"`
	TotalPengiriman          int                    `gorm:"type:int" json:"totalPengiriman"`
	PemasokKe                int                    `gorm:"type:int" json:"pemasokKe"`
	VolumeHarian             float64                `gorm:"type:float" json:"volumeHarian"`
	TotalVolume              float64                `gorm:"type:float" json:"totalVolume"`
	KapasitasBongkar         float64                `gorm:"type:float" json:"kapasitasBongkar"` // Kapasitas bongkar pembangkit ketika data ini dibuat
	Gcv                      float64                `gorm:"type:float" json:"gcv"`
	TSpekId                  *uuid.UUID             `gorm:"type:uuid" json:"spekId"` // Get spek from GCV
	Spek                     *Spek                  `gorm:"foreignKey:TSpekId" json:"spek"`
	TJenisBatuBaraId         *uuid.UUID             `gorm:"type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara            *JenisBatuBara         `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
	TglBongkar               DateField        `gorm:"type:date" json:"tglBongkar"`
	StatusJadwal             StatusRencanaPasokan   `gorm:"type:varchar(45);default:'unconfirmed'" json:"statusJadwal"`
	Hop                      float64                `gorm:"-" json:"hop"`
}

// @deprecated
// func (RencanaPasokanJadwalKalender) TableName() string {
// 	return "v_rencana_pasokan_jadwal_kalender"
// }

// type RencanaPasokanJadwalKalender struct {
// 	TRencanaPasokanPemasokId *uuid.UUID             `gorm:"index" json:"rencanaPasokanPemasokId"`
// 	RencanaPasokanPemasok    *RencanaPasokanPemasok `gorm:"foreignKey:TRencanaPasokanPemasokId" json:"rencanaPasokanPemasok"`
// 	TRencanaPasokanId        *uuid.UUID             `gorm:"index" json:"rencanaPasokanId"`
// 	RencanaPasokan           *RencanaPasokan        `gorm:"foreignKey:TRencanaPasokanId" json:"rencanaPasokan"`
// 	TPemasokId               *uuid.UUID             `gorm:"index" json:"pemasokId"`
// 	Pemasok                  *Organization          `gorm:"foreignKey:TPemasokId" json:"pemasok"`
// 	TotalPengiriman          int                    `gorm:"type:int" json:"totalPengiriman"`
// 	BongkarPengirimanKe      int                    `gorm:"type:int" json:"bongkarPengirimanKe"`
// 	TotalHariBongkar         int                    `gorm:"type:int" json:"totalHariBongkar"`
// 	TotalVolume              float64                `gorm:"type:float" json:"totalVolume"`
// 	KapasitasBongkar         float64                `gorm:"type:float" json:"kapasitasBongkar"`
// 	SpekBatubara             string                 `gorm:"type:varchar(255)" json:"spekBatubara"`
// 	TglMulai                 DateField        `gorm:"type:date" json:"tglMulai"`
// 	TglSelesai               DateField        `gorm:"type:date" json:"tglSelesai"`
// 	StatusJadwal             string                 `gorm:"type:varchar(45);default:'unconfirmed'" json:"statusJadwal"`
// }

type ListRencanaPasokan struct {
	Spek        float64 `json:"spek"`
	VolumeTotal float64 `json:"volumeTotal"`
}

type RencanaPasokanProyeksiHopKaloriTertimbang struct {
	TotalKaloriTertimbang float64              `json:"totalKaloriTertimbang"`
	RencanaPasokan        []ListRencanaPasokan `json:"rencanaPasokan"`
}

type RencanaPasokanProyeksiHopStokAkhir struct {
	Hop                 float64 `json:"hop"`
	StokAkhir           float64 `json:"stokAkhir"`
	MaksPemakaianHarian float64 `json:"maksPemakaianHarian"`
	TotalHari           int     `json:"totalHari"`
}
