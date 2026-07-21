package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

type GroupType string

const (
	GroupTypeGlobal       GroupType = "GLOBAL"
	GroupTypeOrganization GroupType = "ORGANISASI"
	GroupTypeAccessLevel  GroupType = "ACCESSL_LEVEL"
)

func (t GroupType) IsValid() bool {
	return t == GroupTypeOrganization || t == GroupTypeAccessLevel || t == GroupTypeGlobal
}

func (t GroupType) IsOrganization() bool {
	return t == GroupTypeOrganization
}

func (t GroupType) IsAccount() bool {
	return t == GroupTypeAccessLevel
}

func (t GroupType) IsGlobal() bool {
	return t == GroupTypeGlobal
}

func (AccessGroup) TableName() string {
	return "t_access_group"
}

type AccessGroup struct {
	baseapp.BaseModel
	GroupType       GroupType                 `gorm:"type:varchar(45)" json:"groupType"`
	Name            string                    `gorm:"type:varchar(255)" json:"name"`
	Code            string                    `gorm:"type:varchar(255)" json:"code"`
	Description     *string                   `gorm:"type:varchar(100)" json:"description"`
	IsDefault       bool                      `gorm:"type:bool; default:false" json:"isDefault"`
	IsSuperAdmin    bool                      `gorm:"type:bool; default:false" json:"-"`
	TAccessLevelId  *uuid.UUID                `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevel              `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TOrganizationId *uuid.UUID                `gorm:"type:uuid" json:"organizationId"`
	Organization    *OrganizationInfo         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	Accounts        []AccessGroupAccount      `gorm:"foreignKey:TAccessGroupId;constraint:OnDelete:CASCADE" json:"accounts"`
	Organizations   []AccessGroupOrganization `gorm:"foreignKey:TAccessGroupId;constraint:OnDelete:CASCADE" json:"organizations"`
	PageActivity    []AccessGroupAccess       `gorm:"foreignKey:TAccessGroupId; constraint:OnDelete:CASCADE" json:"accessPages"`
}

func (AccessGroupOrganization) TableName() string {
	return "t_access_group_organization"
}

type AccessGroupOrganization struct {
	baseapp.BaseModel
	TAccessGroupId  *uuid.UUID       `gorm:"type:uuid" json:"accessGroupId"`
	AccessGroup     AccessGroup      `gorm:"foreignKey:TAccessGroupId" json:"accessGroup"`
	TOrganizationId *uuid.UUID       `gorm:"type:uuid" json:"organizationId"`
	Organization    OrganizationInfo `gorm:"foreignKey:TOrganizationId;" json:"organization"`
}

func (AccessGroupAccount) TableName() string {
	return "t_access_group_account"
}

type AccessGroupAccount struct {
	baseapp.BaseModel
	TAccessGroupId *uuid.UUID  `gorm:"type:uuid" json:"accessGroupId"`
	AccessGroup    AccessGroup `gorm:"foreignKey:TAccessGroupId" json:"accessGroup"`
	TAccountId     *uuid.UUID  `gorm:"type:uuid" json:"accountId"`
	Account        AccountInfo `gorm:"foreignKey:TAccountId;" json:"account"`
}

func (AccessGroupAccess) TableName() string {
	return "t_access_group_access"
}

type AccessGroupAccess struct {
	baseapp.BaseModel
	TAccessGroupId  *uuid.UUID          `gorm:"type:uuid" json:"accessGroupId"`
	AccessGroup     AccessGroup         `gorm:"foreignKey:TAccessGroupId" json:"accessGroup"`
	TPageActivityId *uuid.UUID          `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity       `gorm:"foreignKey:TPageActivityId;" json:"pageActivity"`
	TParentId       *uuid.UUID          `gorm:"type:uuid" json:"parentId"`
	Parent          *AccessGroupAccess  `gorm:"foreignKey:TParentId;" json:"parent"`
	AccessType      AccessType          `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
	Children        []AccessGroupAccess `gorm:"foreignKey:TParentId;constraint:OnDelete:CASCADE" json:"children"`
}
