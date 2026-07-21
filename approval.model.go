package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (Approval) TableName() string {
	return "t_approval"
}

type Approval struct {
	baseapp.BaseModel
	Uuid            string            `gorm:"type:varchar(85)" json:"uuid"`
	Code            string            `gorm:"type:varchar(255)" json:"code"`
	TOrganizationId *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization    *OrganizationInfo `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPageId         *uuid.UUID        `gorm:"type:uuid" json:"pageId"`
	Page            *PageInfo         `gorm:"foreignKey:TPageId" json:"page"`
	TActivityId     *uuid.UUID        `gorm:"type:uuid" json:"activityId"`
	Activity        *ActivityInfo     `gorm:"foreignKey:TActivityId" json:"activity"`
	ApprovalStatus  ApprovalStatus    `gorm:"type:varchar(100)" json:"approvalStatus"`
	Comment         *string           `gorm:"type:text" json:"comment"`
	DataUuid        *uuid.UUID        `gorm:"type:uuid" json:"dataUuid"`
	DataId          *uuid.UUID        `gorm:"type:uuid" json:"dataId"`
	TAccountId      *uuid.UUID        `gorm:"type:uuid" json:"accountId"`
	Account         *AccountInfo      `gorm:"foreignKey:TAccountId" json:"account"`
}
