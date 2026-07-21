package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Pjbb_kualitas) TableName() string {
	return "t_pjbb_kualitas"
}

type Pjbb_kualitas struct {
	baseapp.BaseModel
	TPjbbId          *uuid.UUID          `gorm:"type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb               `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TOrganizationId  *uuid.UUID          `gorm:"type:uuid" json:"organizationId"`
	Organization     *Organization       `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPemasokId       *uuid.UUID          `gorm:"type:uuid" json:"pemasokId"`
	Pemasok          *Organization       `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPelabuhanMuatId *uuid.UUID          `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat    *PelabuhanMuat      `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId          *uuid.UUID          `gorm:"type:uuid" json:"spekId"`
	Spek             *Spek               `gorm:"foreignKey:TSpekId" json:"spek"`
	Items            []Pjbb_kualitasItem `gorm:"foreignKey:TPjbb_kualitasId; constraint:OnDelete:CASCADE" json:"items"`
}

type Pjbb_kualitasItemType string

const (
	Pjbb_kualitasItemTypePenyesuaian    Pjbb_kualitasItemType = "PENYESUAIAN"
	Pjbb_kualitasItemTypeDendaPenolakan Pjbb_kualitasItemType = "DENDA_PENOLAKAN"
)

func (d Pjbb_kualitasItemType) IsValid() bool {
	return d == Pjbb_kualitasItemTypePenyesuaian || d == Pjbb_kualitasItemTypeDendaPenolakan
}

type Pjbb_kualitasItemDataType string

const (
	Pjbb_kualitasItemDataTypeNum Pjbb_kualitasItemDataType = "Num"
	Pjbb_kualitasItemDataTypeStr Pjbb_kualitasItemDataType = "Str"
)

func (d Pjbb_kualitasItemDataType) IsValid() bool {
	return d == Pjbb_kualitasItemDataTypeNum || d == Pjbb_kualitasItemDataTypeStr
}

func (Pjbb_kualitasItem) TableName() string {
	return "t_pjbb_kualitas_item"
}

type Pjbb_kualitasItem struct {
	baseapp.BaseModel
	TPjbb_kualitasId uuid.UUID                 `gorm:"type:uuid" json:"pjbbKualitasId"`
	Pjbb_kualitas    *Pjbb_kualitas            `gorm:"foreignKey:TPjbb_kualitasId" json:"pjbbKualitas"`
	TipeKualitas     Pjbb_kualitasItemType     `gorm:"type:varchar(255);default:'none'" json:"tipeKualitas"`
	TUraianId        *uuid.UUID                `gorm:"type:uuid" json:"uraianId"`
	Uraian           *Uraian                   `gorm:"foreignKey:TUraianId" json:"uraian"`
	BasisCode        *string                   `gorm:"type:varchar(255)" json:"basisCode"`
	Basis            *ConfigDataInfo           `gorm:"foreignKey:BasisCode;references:Code;" json:"basis"`
	Tipikal          *string                   `gorm:"type:varchar(255)" json:"tipikal"`
	Penolakan        *string                   `gorm:"type:varchar(255)" json:"penolakan"`
	Simbol           *string                   `gorm:"type:varchar(255)" json:"simbol"`
	BatasBawah       *string                   `gorm:"type:varchar(255)" json:"batasBawah"`
	BatasAtas        *string                   `gorm:"type:varchar(255)" json:"batasAtas"`
	TipeData         Pjbb_kualitasItemDataType `gorm:"type:varchar(255);default:'Num'" json:"tipeData"`
	TurunHarga       *string                   `gorm:"type:float" json:"turunHarga"`
}
