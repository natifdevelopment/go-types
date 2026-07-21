package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Numbering) TableName() string {
	return "t_numbering"
}

type Numbering struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	TActivityId     *uuid.UUID        `gorm:"type:uuid" json:"activityId"`
	TPageId         *uuid.UUID        `gorm:"type:uuid" json:"pageId"`
	Page            *PageInfo         `gorm:"foreignKey:TPageId;" json:"page"`
	Activity        *ActivityInfo     `gorm:"foreignKey:TActivityId;" json:"activity"`
	Organization    *OrganizationInfo `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	DbFieldName     string            `gorm:"type:varchar(255)" json:"dbFieldName"`
	Name            string            `gorm:"type:varchar(255)" json:"name"`
	Code            *string           `gorm:"type:varchar(255)" json:"code"`
	NumberingFormat string            `gorm:"type:varchar(255)" json:"numberingFormat"`
}
