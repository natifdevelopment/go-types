package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (ApprovalCode) TableName() string {
	return "t_approval_code"
}

type ApprovalCode struct {
	baseapp.BaseModel
	TOrganizationId    *uuid.UUID           `gorm:"type:uuid" json:"organizationId"`
	Organization       *OrganizationInfo    `gorm:"foreignKey:TOrganizationId" json:"organization"`
	Name               string               `gorm:"type:varchar(255)" json:"name"`
	Code               string               `gorm:"type:varchar(255);unique" json:"code"`
	Description        string               `gorm:"type:varchar(255)" json:"description"`
	TAccountId         *uuid.UUID           `gorm:"type:uuid" json:"accountId"`
	Account            *AccountInfo         `gorm:"foreignKey:TAccountId;" json:"account"`
	ApprovalCodeMember []ApprovalCodeMember `gorm:"foreignKey:TApprovalCodeId;constraint:OnDelete:CASCADE" json:"approvalCodeMembers"`
}

func (ApprovalCodeMember) TableName() string {
	return "t_approval_code_member"
}

type ApprovalCodeMember struct {
	baseapp.BaseModel
	TAccountId      *uuid.UUID    `gorm:"type:uuid" json:"accountId"`
	TApprovalCodeId uuid.UUID     `gorm:"type:uuid" json:"approvalCodeId"`
	ApprovalCode    *ApprovalCode `gorm:"foreignKey:TApprovalCodeId;" json:"approvalCode"`
	Account         *AccountInfo  `gorm:"foreignKey:TAccountId;" json:"account"`
}
