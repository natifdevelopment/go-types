package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (Surveyor) TableName() string {
	return "t_surveyor"
}

type Surveyor struct {
	baseapp.BaseModel
	TMasterEpiSurveyorId *uuid.UUID         `gorm:"type:uuid" json:"masterEpiSurveyorId"`
	MasterEpiSurveyor    *MasterEpiSurveyor `gorm:"foreignKey:TMasterEpiSurveyorId;" json:"masterEpiSurveyor"`
	NamaSurveyor         string             `gorm:"type:varchar(255)" json:"namaSurveyor"`
	NamaSingkat          string             `gorm:"type:varchar(255)" json:"namaSingkat"`
	Alamat               string             `gorm:"type:varchar(255)" json:"alamat"`
	ContactPerson        string             `gorm:"type:varchar(255)" json:"contactPerson"`
	Email                string             `gorm:"type:varchar(255)" json:"email"`
	Telp                 string             `gorm:"type:varchar(255)" json:"telp"`
	Keterangan           string             `gorm:"type:varchar(255)" json:"keterangan"`
	TOrganizationId      *uuid.UUID         `gorm:"type:uuid" json:"organizationId"`
	Organization         *Organization      `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	StatusSurveyor       bool               `gorm:"type:bool" json:"statusSurveyor"`
}
