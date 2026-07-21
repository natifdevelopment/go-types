package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type ApprovalStrategyType string

const (
	ApprovalStrategyTypeOneByOne ApprovalStrategyType = "BERURUTAN"
	ApprovalStrategyTypeSendAll  ApprovalStrategyType = "BERSAMAAN"
)

func (o ApprovalStrategyType) IsValid() bool {
	return o == ApprovalStrategyTypeOneByOne || o == ApprovalStrategyTypeSendAll
}

func (o ApprovalStrategyType) Value() (driver.Value, error) {
	return string(o), nil
}

func (o *ApprovalStrategyType) Scan(value interface{}) error {
	*o = ApprovalStrategyType(value.(string))
	return nil
}

func (ApprovalStrategy) TableName() string {
	return "t_approval_strategy"
}

type ApprovalStrategy struct {
	baseapp.BaseModel
	Name          string                         `gorm:"type:varchar(255)" json:"name"`
	Code          string                         `gorm:"type:varchar(255)" json:"code"`
	TDirectoryId  *uuid.UUID                     `gorm:"type:uuid" json:"directoryId"`
	Directory     *PageActivity                  `gorm:"foreignKey:TDirectoryId" json:"directory"`
	TMenuId       *uuid.UUID                     `gorm:"type:uuid" json:"menuId"`
	Menu          *PageActivity                  `gorm:"foreignKey:TMenuId" json:"menu"`
	ApprovalType  ApprovalStrategyType           `gorm:"type:varchar(100)" json:"approvalType"`
	AllowPass     bool                           `gorm:"type:boolean; default: false" json:"allowPass"`
	ApprovalLevel []ApprovalStrategyLevel        `gorm:"foreignKey:TApprovalStrategyId; constraint:OnDelete:CASCADE" json:"approvalLevels"`
	Organizations []ApprovalStrategyOrganization `gorm:"foreignKey:TApprovalStrategyId; constraint:OnDelete:CASCADE" json:"organization"`
}

func (ApprovalStrategyInfo) TableName() string {
	return "t_approval_strategy"
}

type ApprovalStrategyInfo struct {
	Name        string     `gorm:"type:varchar(255)" json:"name"`
	Code        string     `gorm:"type:varchar(255)" json:"code"`
	Description string     `gorm:"type:varchar(255)" json:"description"`
	TActivityId *uuid.UUID `gorm:"type:uuid" json:"activityId"`
}

func (ApprovalStrategyOrganization) TableName() string {
	return "t_approval_strategy_organization"
}

type ApprovalStrategyOrganization struct {
	baseapp.BaseModel
	TApprovalStrategyId *uuid.UUID        `gorm:"type:uuid" json:"approvalStrategyId"`
	TOrganizationId     *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization        *OrganizationInfo `gorm:"foreignKey:TOrganizationId" json:"organization"`
}

func (ApprovalStrategyLevel) TableName() string {
	return "t_approval_strategy_level"
}

type ApprovalStrategyLevel struct {
	baseapp.BaseModel
	TApprovalStrategyId *uuid.UUID             `gorm:"type:uuid" json:"approvalStrategyId"`
	ApprovalCode        []ApprovalStrategyCode `gorm:"foreignKey:TApprovalStrategyLevelId; constraint:OnDelete:CASCADE" json:"approvalCodes"`
}

func (ApprovalStrategyCode) TableName() string {
	return "t_approval_strategy_code"
}

type ApprovalStrategyCode struct {
	baseapp.BaseModel
	TApprovalStrategyId      *uuid.UUID    `gorm:"type:uuid" json:"approvalStrategyId"`
	TApprovalStrategyLevelId *uuid.UUID    `gorm:"type:uuid" json:"approvalStrategyLevelId"`
	TApprovalCodeId          *uuid.UUID    `gorm:"type:uuid" json:"approvalCodeId"`
	ApprovalCode             *ApprovalCode `gorm:"foreignKey:TApprovalCodeId;" json:"approvalCode"`

	ApprovalStatus ApprovalStatus `gorm:"-" json:"approvalStatus"`
	SubmitUlang    bool           `gorm:"-" json:"submitUlang"`
	Comment        *string        `gorm:"-" json:"comment"`
	ApprovalDate   *time.Time     `gorm:"-" json:"approvalDate"`
}
