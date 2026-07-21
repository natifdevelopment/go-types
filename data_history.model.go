package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func (DataHistory) TableName() string {
	return "t_data_history"
}

type DataHistory struct {
	baseapp.BaseModel
	Table           string          `gorm:"type:varchar(255)" json:"-"`
	Method          string          `gorm:"type:varchar(255)" json:"-"`
	Type            string          `gorm:"type:varchar(255)" json:"type"`
	ReferenceId     uuid.UUID       `gorm:"type:varchar(255)" json:"-"`
	Payload         datatypes.JSON  `gorm:"type:jsonb" json:"-"`
	PrevData        *datatypes.JSON `gorm:"type:jsonb" json:"-"`
	TOrganizationId *uuid.UUID      `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization   `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPageActivityId *uuid.UUID      `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity   `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
	Version         string          `gorm:"type:varchar(255)" json:"version"`
	AppVersion      string          `gorm:"type:varchar(255)" json:"appVersion"`
}
