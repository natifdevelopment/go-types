package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (MasterKapal) TableName() string {
	return "t_master_kapal"
}

type MasterKapal struct {
	baseapp.BaseModel
	TMasterEpiKapalId *uuid.UUID      `gorm:"type:uuid" json:"masterEpiKapalId"`
	MasterEpiKapal    *MasterEpiKapal `gorm:"foreignKey:TMasterEpiKapalId" json:"masterEpiKapal"`
	TPemilikId        *uuid.UUID      `gorm:"type:uuid" json:"pemilikId"`
	Pemilik           *Organization   `gorm:"foreignKey:TPemilikId" json:"pemilik"`
	TTransportirId    *uuid.UUID      `gorm:"type:uuid" json:"transportirId"`
	Transportir       *Organization   `gorm:"foreignKey:TTransportirId" json:"transportir"`
	StatusKapal       bool            `gorm:"type:bool" json:"statusKapal"`
	TOrganizationId   *uuid.UUID      `gorm:"type:uuid" json:"organizationId"`
	Organization      *Organization   `gorm:"foreignKey:TOrganizationId" json:"organization"`
}
