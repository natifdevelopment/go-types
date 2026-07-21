package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Pembangkit) TableName() string {
	return "t_pembangkit"
}

type Pembangkit struct {
	baseapp.BaseModel
	Nama                   string           `gorm:"type:varchar(255)" json:"nama"`
	NamaPembangkit         string           `gorm:"type:varchar(255)" json:"namaPembangkit"`
	NamaPendek             string           `gorm:"type:varchar(255)" json:"namaPendek"`
	NamaSingkat            string           `gorm:"type:varchar(255)" json:"namaSingkat"`
	Alamat                 string           `gorm:"type:text" json:"alamat"`
	KodeSentral            *string          `gorm:"type:varchar(255)" json:"kodeSentral"`
	NoTelepon              *string          `gorm:"type:varchar(255)" json:"noTelepon"`
	KepemilikanCode        *string          `gorm:"type:varchar(255)" json:"kepemilikanCode"`
	Kepemilikan            *ConfigData      `gorm:"foreignKey:KepemilikanCode;references:Code;" json:"kepemilikan"`
	PengelolaCode          *string          `gorm:"type:varchar(255)" json:"pengelolaCode"`
	Pengelola              *ConfigData      `gorm:"foreignKey:PengelolaCode;references:Code;" json:"pengelola"`
	TWilayahId             *uuid.UUID       `gorm:"type:uuid" json:"wilayahId"`
	Wilayah                *Wilayah         `gorm:"foreignKey:TWilayahId" json:"wilayah"`
	JenisPembangkitCode    *string          `gorm:"type:varchar(255)" json:"jenisPembangkitCode"`
	JenisPembangkit        *ConfigData      `gorm:"foreignKey:JenisPembangkitCode;references:Code;" json:"jenisPembangkit"`
	DayaTerpasang          float32          `gorm:"type:float" json:"dayaTerpasang"`
	DayaMampuNetto         float32          `gorm:"type:float" json:"dayaMampuNetto"`
	Kapasitas              string           `gorm:"type:varchar(255)" json:"kapasitas"`
	MinimumStok            float64          `gorm:"type:float" json:"minimumStok"`
	HariOperasi            float64          `gorm:"type:float" json:"hariOperasi"`
	TKantorIndukId         *uuid.UUID       `gorm:"type:uuid" json:"kantorIndukId"`
	KantorInduk            *Organization    `gorm:"foreignKey:TKantorIndukId" json:"kantorInduk"`
	Latitude               float64          `gorm:"type:float" json:"latitude"`
	Longitude              float64          `gorm:"type:float" json:"longitude"`
	Kota                   string           `gorm:"type:varchar(255)" json:"kota"`
	KdPembangkit           string           `gorm:"type:varchar(255)" json:"kdPembangkit"`
	Tipikal                float64          `gorm:"type:float" json:"tipikal"`
	TotalDaya              float64          `gorm:"type:float" json:"totalDaya"`
	CapacityFactor         float64          `gorm:"type:float" json:"capacityFactor"`
	Fsc                    float64          `gorm:"type:float" json:"fsc"`
	ZonaWaktuCode          *string          `gorm:"type:varchar(255)" json:"zonaWaktuCode"`
	ZonaWaktu              *ConfigData      `gorm:"foreignKey:ZonaWaktuCode;references:Code;" json:"zonaWaktu"`
	MasaBerlakuAwal        *DateField `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir       *DateField `gorm:"type:date" json:"masaBerlakuAkhir"`
	KapasitasBongkarHarian float64          `gorm:"type:float" json:"kapasitasBongkarHarian"`
	MaksPemakaianHarian    float64          `gorm:"type:float" json:"maksPemakaianHarian"`
	KapasitasStockpile     *float64         `gorm:"type:float" json:"kapasitasStockpile"`
	DigunakanBbo           bool             `gorm:"type:boolean" json:"digunakanBbo"`
	StatusPembangkit       bool             `gorm:"type:boolean" json:"statusPembangkit"`
	StatusDashboard        bool             `gorm:"type:boolean" json:"statusDashboard"`
	MulutTambang           bool             `gorm:"type:boolean" json:"mulutTambang"`
	TRegionalId            *uuid.UUID       `gorm:"type:uuid" json:"regionalId"`
	Regional               *Regional        `gorm:"foreignKey:TRegionalId" json:"regional"`

	// Organization that was created for this Pembangkit
	Organization *Organization `gorm:"foreignKey:TPembangkit2Id" json:"organization"`

	// Kode BAST
	KodeBastCif        *string `gorm:"type:varchar(255)" json:"kodeBastCif"`
	KodeBaBongkarTrans *string `gorm:"type:varchar(255)" json:"kodeBaBongkarTrans"`

	// SAP
	Plant       *string `gorm:"type:varchar(255)" json:"plant"`
	Sloc        *string `gorm:"type:varchar(255)" json:"sloc"`
	CostCenter  *string `gorm:"type:varchar(255)" json:"costCenter"`
	PlantGudang *string `gorm:"type:varchar(255)" json:"plantGudang"`
	SlocGudang  *string `gorm:"type:varchar(255)" json:"slocGudang"`

	// PLTU Proyek
	StatusPltuProyek     bool       `gorm:"type:boolean" json:"statusPltuProyek"`
	ProyekCode           *string    `gorm:"type:varchar(255)" json:"proyekCode"`
	TSapInboundProjectId *uuid.UUID `gorm:"type:uuid" json:"sapInboundProjectId"`

	PipPembangkitMap []PipPembangkitMap `gorm:"foreignKey:TPembangkitId;constraint:OnDelete:CASCADE" json:"-"`
}
