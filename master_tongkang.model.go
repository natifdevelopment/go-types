package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (MasterTongkang) TableName() string {
	return "t_master_tongkang"
}

type MasterTongkang struct {
	baseapp.BaseModel
	TMasterEpiTongkangId uuid.UUID         `gorm:"type:uuid" json:"masterEpiTongkangId"`
	MasterEpiTongkang    MasterEpiTongkang `gorm:"foreignKey:TMasterEpiTongkangId;" json:"masterEpiTongkang"`
	TPemilikId           *uuid.UUID        `gorm:"type:uuid" json:"pemilikId"`
	Pemilik              *Organization     `gorm:"foreignKey:TPemilikId;" json:"pemilik"`
	TTransportirId       *uuid.UUID        `gorm:"type:uuid" json:"transportirId"`
	Transportir          *Organization     `gorm:"foreignKey:TTransportirId;" json:"transportir"`
	StatusTongkang       bool              `gorm:"type:bool" json:"statusTongkang"`
	TOrganizationId      *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization         *Organization     `gorm:"foreignKey:TOrganizationId;" json:"organization"`
}
