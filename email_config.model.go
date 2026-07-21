package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

type EmailConfigType string

const (
	EmailConfigTypeTransactional EmailConfigType = "transactional"
	EmailConfigTypeDefault       EmailConfigType = "default"
)

func (ct EmailConfigType) IsValid() bool {
	return ct == EmailConfigTypeTransactional || ct == EmailConfigTypeDefault
}

func (ct EmailConfigType) String() string {
	return string(ct)
}

func (EmailConfig) TableName() string {
	return "t_email_config"
}

type EmailConfig struct {
	baseapp.BaseModel
	Type                   EmailConfigType          `gorm:"type:varchar(455)" json:"type"`
	EmailTemplateCode      *string                  `gorm:"type:varchar(455)" json:"emailTemplateCode"`
	Subject                string                   `gorm:"type:varchar(455)" json:"subject"`
	BodyHtml               string                   `gorm:"type:text" json:"bodyHtml"`
	BodyText               string                   `gorm:"type:text" json:"bodyText"`
	TOrganizationId        *uuid.UUID               `gorm:"type:uuid;" json:"organizationId"`
	TPageId                *uuid.UUID               `gorm:"type:uuid;" json:"pageId"`
	TActivityId            *uuid.UUID               `gorm:"type:uuid;" json:"activityId"`
	Page                   *PageInfo                `gorm:"foreignKey:TPageId;" json:"page"`
	Activity               *ActivityInfo            `gorm:"foreignKey:TActivityId;" json:"activity"`
	EmailTemplate          *ConfigDataInfo          `gorm:"foreignKey:EmailTemplateCode;references:Code;" json:"emailTemplate"`
	Organization           *OrganizationInfo        `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	EmailConfigEmail       []EmailConfigEmail       `gorm:"foreignKey:TEmailConfigId; constraint:OnDelete:CASCADE" json:"sendToEmails"`
	EmailConfigAccessLevel []EmailConfigAccessLevel `gorm:"foreignKey:TEmailConfigId; constraint:OnDelete:CASCADE" json:"sendToAccessLevels"`
}

func (EmailConfigEmail) TableName() string {
	return "t_email_config_email"
}

type EmailConfigEmail struct {
	baseapp.BaseModel
	TEmailConfigId *uuid.UUID   `gorm:"type:uuid" json:"emailConfigId"`
	EmailConfig    *EmailConfig `gorm:"foreignKey:TEmailConfigId" json:"EmailConfig"`
	Email          string       `gorm:"type:varchar(155)" json:"email"`
}

func (EmailConfigAccessLevel) TableName() string {
	return "t_email_config_access_level"
}

type EmailConfigAccessLevel struct {
	baseapp.BaseModel
	TEmailConfigId *uuid.UUID      `gorm:"type:uuid" json:"emailConfigId"`
	EmailConfig    *EmailConfig    `gorm:"foreignKey:TEmailConfigId" json:"EmailConfig"`
	TAccessLevelId *uuid.UUID      `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel    AccessLevelInfo `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}
