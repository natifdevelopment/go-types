package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

type LevelType string

const (
	LevelTypeOrganization LevelType = "organization"
	LevelTypeAccount      LevelType = "account"
)

func (t LevelType) IsValid() bool {
	return t == LevelTypeOrganization || t == LevelTypeAccount
}

func (t LevelType) IsOrganization() bool {
	return t == LevelTypeOrganization
}

func (t LevelType) IsAccount() bool {
	return t == LevelTypeAccount
}

func (AccessLevel) TableName() string {
	return "t_access_level"
}

type AccessLevel struct {
	baseapp.BaseModel
	TOrganizationId      *uuid.UUID          `gorm:"type:uuid" json:"organizationId"`
	Organization         *OrganizationInfo   `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	TParentId            *uuid.UUID          `gorm:"type:uuid" json:"parentId"`
	Parent               *AccessLevelInfo    `gorm:"foreignKey:TParentId;" json:"parent"`
	Type                 LevelType           `gorm:"type:varchar(45)" json:"type"`
	Name                 string              `gorm:"type:varchar(255)" json:"name"`
	Code                 string              `gorm:"type:varchar(255)" json:"code"`
	Description          *string             `gorm:"type:varchar(100)" json:"description"`
	ApprovalAction       *string             `gorm:"type:varchar(255)" json:"approvalAction"`
	TApprovalActivityId  *uuid.UUID          `gorm:"type:uuid" json:"approvalActivityId"`
	TApprovalToId        *uuid.UUID          `gorm:"type:uuid" json:"approvalToId"`
	TApprovalStrategyId  *uuid.UUID          `gorm:"type:uuid" json:"approvalStrategyId"`
	TApprovalRequesterId *uuid.UUID          `gorm:"type:uuid" json:"approvalRequesterId"`
	Children             []AccessLevel       `gorm:"foreignKey:TParentId;references:ID;constraint:OnDelete:CASCADE" json:"childrens"`
	PageActivity         []AccessLevelAccess `gorm:"foreignKey:TAccessLevelId; constraint:OnDelete:CASCADE" json:"accessPages"`
}

func (AccessLevelInfo) TableName() string {
	return "t_access_level"
}

type AccessLevelInfo struct {
	baseapp.BaseModel
	Type        string `gorm:"type:varchar(255)" json:"type"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Code        string `gorm:"type:varchar(100)" json:"code"`
	Description string `gorm:"type:varchar(100)" json:"description"`
}

func (AccessLevelNameOnly) TableName() string {
	return "t_access_level"
}

type AccessLevelNameOnly struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	Name string    `gorm:"type:varchar(255)" json:"name"`
}

func (AccessLevelAccess) TableName() string {
	return "t_access_level_page_activity"
}

type AccessLevelAccess struct {
	baseapp.BaseModel
	TAccessLevelId  *uuid.UUID          `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevel        `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
	TPageActivityId *uuid.UUID          `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity       `gorm:"foreignKey:TPageActivityId;" json:"pageActivity"`
	TParentId       *uuid.UUID          `gorm:"type:uuid" json:"parentId"`
	Parent          *AccessLevelAccess  `gorm:"foreignKey:TParentId;" json:"parent"`
	AccessType      AccessType          `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
	Children        []AccessLevelAccess `gorm:"foreignKey:TParentId;constraint:OnDelete:CASCADE" json:"children"`
}
