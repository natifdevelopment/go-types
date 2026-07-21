package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (SurveyorIndependenPembangkit) TableName() string {
	return "t_surveyor_independen_pembangkit"
}

type SurveyorIndependenPembangkit struct {
	baseapp.BaseModel
	TMasterEpiSurveyorId *uuid.UUID         `gorm:"type:uuid" json:"masterEpiSurveyorId"`
	MasterEpiSurveyor    *MasterEpiSurveyor `gorm:"foreignKey:TMasterEpiSurveyorId;" json:"masterEpiSurveyor"`
	TOrganizationId      *uuid.UUID         `gorm:"type:uuid" json:"organizationId"`
	Organization         *Organization      `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	StatusSurveyor       bool               `gorm:"type:boolean" json:"statusSurveyor"`
}
