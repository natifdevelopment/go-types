package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (Transportir) TableName() string {
	return "t_transportir"
}

type Transportir struct {
	baseapp.BaseModel
	TMasterEpiTransportirId uuid.UUID             `gorm:"type:uuid" json:"masterEpiTransportirId"`
	MasterEpiTransportir    *MasterEpiTransportir `gorm:"foreignKey:TMasterEpiTransportirId;" json:"masterEpiTransportir"`
	NamaTransportir         string                `gorm:"type:varchar(255)" json:"namaTransportir"`
	Email                   string                `gorm:"type:varchar(255)" json:"email"`
	Alamat                  string                `gorm:"type:text" json:"alamat"`
	Telp                    string                `gorm:"type:varchar(20)" json:"telp"`
	StatusTransportir       bool                  `gorm:"type:bool;default:false" json:"statusTransportir"`
	TOrganizationId         *uuid.UUID            `gorm:"type:uuid" json:"organizationId"`
	Organization            *Organization         `gorm:"foreignKey:TOrganizationId;" json:"organization"`
}
