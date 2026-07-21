package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

type TipePejabat string

const (
	TipePejabatPLN      TipePejabat = "INTERNAL"
	TipePejabatExternal TipePejabat = "EXTERNAL"
)

func (t TipePejabat) IsValid() bool {
	return t == TipePejabatPLN || t == TipePejabatExternal
}

type PejabatAms struct {
	baseapp.BaseModel
	TOrganizationId   *uuid.UUID       `gorm:"type:uuid" json:"organizationId"`
	Organization      *Organization    `gorm:"foreignKey:TOrganizationId" json:"organization"`
	Nama              string           `json:"nama" gorm:"type:varchar(255)"`
	Nip               *string          `json:"nip" gorm:"type:varchar(255);null;unique"`
	Email             *string          `json:"email" gorm:"type:varchar(255);null;unique"`
	Jabatan           string           `json:"jabatan" gorm:"type:varchar(255)"`
	PltJabatan        *string          `json:"pltJabatan" gorm:"type:varchar(255)"`
	RolePejabat       *string          `json:"rolePejabat" gorm:"type:varchar(255)"`
	NoSuratTugas      string           `json:"noSuratTugas" gorm:"type:varchar(255)"`
	TglSuratTugas     DateField  `json:"tglSuratTugas" gorm:"type:date;"`
	TFileSuratTugasId *string          `json:"fileSuratTugasId" gorm:"type:uuid"`
	FileSuratTugas    *Media           `gorm:"foreignKey:TFileSuratTugasId" json:"fileSuratTugas"`
	StatusDefault     bool             `json:"statusDefault" gorm:"type:boolean;default:true;"`
	StatusPejabat     bool             `json:"statusPejabat" gorm:"type:boolean;default:true;"`
	TglBerlakuMulai   *DateField `json:"tglBerlakuMulai" gorm:"type:date;"`
	TglBerlakuAkhir   *DateField `json:"tglBerlakuAkhir" gorm:"type:date;"`
	TipeDokumen       *string          `json:"tipeDokumen" gorm:"type:varchar(255);"`
	StatusPeruri      *string          `json:"statusPeruri" gorm:"type:varchar(255)"`
	StatusAms         string           `json:"statusAms" gorm:"type:varchar(255)"`
	TipePejabat       *TipePejabat     `json:"tipePejabat" gorm:"type:varchar(255)"`
}

func (PejabatAms) TableName() string {
	return "t_pejabat_ams"
}
