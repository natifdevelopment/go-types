package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

type AppSwitcherBehavior string

const (
	AppSwitcherBehaviorRedirect  AppSwitcherBehavior = "redirect"
	AppSwitcherBehaviorNewTab    AppSwitcherBehavior = "new_tab"
	AppSwitcherBehaviorNewWindow AppSwitcherBehavior = "new_window"
)

func (ct AppSwitcherBehavior) IsValid() bool {
	return ct == AppSwitcherBehaviorRedirect || ct == AppSwitcherBehaviorNewTab || ct == AppSwitcherBehaviorNewWindow
}

func (ct AppSwitcherBehavior) String() string {
	return string(ct)
}

func (AppSwitcher) TableName() string {
	return "t_app_switcher"
}

type AppSwitcher struct {
	baseapp.BaseModel
	TParentId   *uuid.UUID               `gorm:"type:uuid" json:"parentId"`
	Parent      *AppSwitcherInfo         `gorm:"foreignKey:TParentId;" json:"parent"`
	Name        string                   `gorm:"type:varchar(255)" json:"name"`
	Code        string                   `gorm:"type:varchar(255)" json:"code"`
	Description *string                  `gorm:"type:text" json:"description"`
	Url         *string                  `gorm:"type:varchar(455)" json:"url"`
	TMediaId    *uuid.UUID               `gorm:"type:uuid" json:"mediaId"`
	Behavior    AppSwitcherBehavior      `gorm:"type:varchar(255); default:'redirect'" json:"behavior"`
	Logo        *Media                   `gorm:"foreignKey:TMediaId;" json:"logo"`
	AccessGroup []AppSwitcherAccessGroup `gorm:"foreignKey:TAppSwitcherId; constraint:OnDelete:CASCADE" json:"accessGroups"`
	AccessLevel []AppSwitcherAccessLevel `gorm:"foreignKey:TAppSwitcherId; constraint:OnDelete:CASCADE" json:"accessLevels"`
	Children    []AppSwitcher            `gorm:"foreignKey:TParentId;references:ID;constraint:OnDelete:CASCADE" json:"childrens"`
}

func (AppSwitcherInfo) TableName() string {
	return "t_app_switcher"
}

type AppSwitcherInfo struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Code        string    `gorm:"type:varchar(255)" json:"code"`
	Description string    `gorm:"type:text" json:"description"`
}

func (AppSwitcherAccessGroup) TableName() string {
	return "t_app_switcher_access_group"
}

type AppSwitcherAccessGroup struct {
	baseapp.BaseModel
	TAppSwitcherId *uuid.UUID      `gorm:"type:uuid" json:"appSwitcherId"`
	TAccessGroupId *uuid.UUID      `gorm:"type:uuid" json:"accessGroupId"`
	AppSwitcher    AppSwitcherInfo `gorm:"foreignKey:TAppSwitcherId;" json:"appSwitcher"`
	AccessGroup    AccessGroup     `gorm:"foreignKey:TAccessGroupId;" json:"accessGroup"`
}

func (AppSwitcherAccessLevel) TableName() string {
	return "t_app_switcher_access_level"
}

type AppSwitcherAccessLevel struct {
	baseapp.BaseModel
	TAppSwitcherId *uuid.UUID      `gorm:"type:uuid" json:"appSwitcherId"`
	TAccessLevelId *uuid.UUID      `gorm:"type:uuid" json:"accessLevelId"`
	AppSwitcher    AppSwitcherInfo `gorm:"foreignKey:TAppSwitcherId;" json:"appSwitcher"`
	AccessLevel    AccessLevel     `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}
