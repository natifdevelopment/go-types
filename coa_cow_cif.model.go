package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (CoaCow) TableName() string {
	return "t_coa_cow_cif"
}

type CoaCow struct {
	baseapp.BaseModel
	SkemaKontrakCode              *JenisPengiriman  `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak                  *ConfigDataInfo   `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TSurveyorId                   *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor                      *Organization     `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TJadwalPengirimanId           *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman              *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId                    *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                       *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TFileCoaId                    *uuid.UUID        `gorm:"type:uuid" json:"fileCoaId"`
	FileCoa                       *Media            `gorm:"foreignKey:TFileCoaId" json:"fileCoa"`
	NoSertifikatCoa               *string           `gorm:"type:varchar(255)" json:"noSertifikatCoa"`
	NamaKapal                     *string           `gorm:"type:varchar(255)" json:"namaKapal"`
	TanggalCoa                    *DateField  `gorm:"type:timestamp" json:"tanggalCoa"`
	GrossCalorificValueAr         *float32          `gorm:"type:float" json:"grossCalorificValueAr"`
	HardGroveGrindabilityIndex    *float32          `gorm:"type:float" json:"hardGroveGrindabilityIndex"`
	TotalMoistureAr               *float32          `gorm:"type:float" json:"totalMoistureAr"`
	AshContentAr                  *float32          `gorm:"type:float" json:"ashContentAr"`
	SodiumContentInAsh            *float32          `gorm:"type:float" json:"sodiumContentInAsh"`
	SulphurContentAr              *float32          `gorm:"type:float" json:"sulphurContentAr"`
	NitrogenDaf                   *float32          `gorm:"type:float" json:"nitrogenDaf"`
	Slagging                      *string           `gorm:"type:varchar(255)" json:"slagging"`
	UkuranButiranLolosAyakan70mm  *float32          `gorm:"type:float" json:"ukuranButiranLolosAyakan70mm"`
	UkuranButiranLolosAyakan50mm  *float32          `gorm:"type:float" json:"ukuranButiranLolosAyakan50mm"`
	UkuranButiranLolosAyakan32mm  *float32          `gorm:"type:float" json:"ukuranButiranLolosAyakan32mm"`
	UkuranButiranLolosAyakan238mm *float32          `gorm:"type:float" json:"ukuranButiranLolosAyakan238mm"`
	FoulingIndex                  *string           `gorm:"type:varchar(255)" json:"foulingIndex"`
	AshFusionTempratureIdt        *float32          `gorm:"type:float" json:"ashFusionTempratureIdt"`
	TFileCowId                    *uuid.UUID        `gorm:"type:uuid" json:"fileCowId"`
	FileCow                       *Media            `gorm:"foreignKey:TFileCowId" json:"fileCow"`
	NoSertifikatCow               *string           `gorm:"type:varchar(255)" json:"noSertifikatCow"`
	QuantityOfCargo               *float32          `gorm:"type:float" json:"quantityOfCargo"`
	Agreement                     bool              `gorm:"type:boolean" json:"agreement"`
	CoaCowCifItem                 []CoaCowCifItem   `gorm:"foreignKey:TCoaCowCifId; constraint:OnDelete:CASCADE" json:"items"`
}

func (CoaCowCifItem) TableName() string {
	return "t_coa_cow_cif_item"
}

type CoaCowCifItem struct {
	baseapp.BaseModel
	TCoaCowCifId *uuid.UUID `gorm:"type:uuid" json:"coaCowCifId"`
	CoaCowCif    *CoaCow    `gorm:"foreignKey:TCoaCowCifId" json:"coaCowCif"`
	TKualitasId  *uuid.UUID `gorm:"type:uuid" json:"kualitasId"`
	// Kualitas        *PjbbRumusBlokKualitasItem `gorm:"foreignKey:TKualitasId" json:"kualitas"`
	TUraianId       *uuid.UUID              `gorm:"type:uuid" json:"uraianId"`
	Uraian          *Uraian                 `gorm:"foreignKey:TUraianId" json:"uraian"`
	ValueStr        *string                 `gorm:"type:varchar(255)" json:"valueStr"`
	ValueNum        *float64                `gorm:"type:float" json:"valueNum"`
	ValueEnum       *ConfigDataInfo         `gorm:"-" json:"valueEnum"`
	PengajuanUmpire []PengajuanUmpireUraian `gorm:"foreignKey:TCoaCowItemId; constraint:OnDelete:CASCADE" json:"-"`
}
