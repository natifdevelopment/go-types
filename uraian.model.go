package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type TipeUraian string

const (
	TipeUraianKualitas  TipeUraian = "KUALITAS"
	TipeUraianRumus     TipeUraian = "RUMUS"
	TipeUraianTransaksi TipeUraian = "TRANSAKSI"
)

func (o TipeUraian) IsValid() bool {
	return o == TipeUraianKualitas || o == TipeUraianRumus || o == TipeUraianTransaksi
}

type TipeDataUraian string

const (
	TipeDataUraianNumber TipeDataUraian = "NUMBER"
	TipeDataUraianString TipeDataUraian = "STRING"
	TipeDataUraianEnum   TipeDataUraian = "ENUM"
)

func (o TipeDataUraian) IsValid() bool {
	return o == TipeDataUraianNumber || o == TipeDataUraianString
}

func (Uraian) TableName() string {
	return "t_uraian"
}

type Uraian struct {
	baseapp.BaseModel
	Type           TipeUraian       `gorm:"type:varchar(255);default:'RUMUS'" json:"type"`      // kualitas/rumus
	TipeData       TipeDataUraian   `gorm:"type:varchar(255);default:'NUMBER'" json:"tipeData"` // NUMBER/STRING
	Name           string           `gorm:"type:varchar(255)" json:"name"`
	Description    string           `gorm:"type:varchar(645)" json:"description"`
	Variable       string           `gorm:"type:varchar(145)" json:"variable"`
	SatuanCode     *string          `gorm:"type:varchar(85)" json:"satuanCode"`
	Satuan         *ConfigDataInfo  `gorm:"foreignKey:SatuanCode;references:Code;" json:"satuan"`
	TEnumId        *uuid.UUID       `gorm:"index;type:uuid" json:"enumId"`
	Enum           *ConfigDataInfo  `gorm:"foreignKey:TEnumId" json:"enum"`
	Rumus          string           `gorm:"type:text" json:"rumus"`
	Nilai          string           `gorm:"type:varchar(255)" json:"nilai"`
	BatasanRepro   *float64         `gorm:"type:float" json:"batasanRepro"`
	Hasil          *float64         `gorm:"-" json:"hasil"`
	TAccessLevelId *uuid.UUID       `gorm:"type:uuid;" json:"accessLevelId"`
	AccessLevel    *AccessLevelInfo `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}
