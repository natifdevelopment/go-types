package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (MasterUnitJetty) TableName() string {
	return "t_master_unit_jetty"
}

type MasterUnitJetty struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization    *OrganizationInfo `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	KodeJetty       string            `gorm:"type:varchar(255)" json:"kodeJetty"`
	Lokasi          string            `gorm:"type:varchar(255)" json:"lokasi"`
	DraughtJetty    float64           `gorm:"type:float" json:"draughtJetty"`
	NamaPtJm        string            `gorm:"type:varchar(255)" json:"namaPtJm"`
	KepemilikanJm   string            `gorm:"type:varchar(255)" json:"kepemilikanJm"`
}
