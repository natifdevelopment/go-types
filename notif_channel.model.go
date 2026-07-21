package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (NotifChannel) TableName() string {
	return "t_notif_channel"
}

type NotifChannel struct {
	baseapp.BaseModel
	TNotifConfigId     *uuid.UUID    `gorm:"type:uuid" json:"notifConfigId"`
	NotifConfig        NotifConfig   `gorm:"foreignKey:TNotifConfigId;" json:"notifConfig"`
	TPageId            *uuid.UUID    `gorm:"type:uuid;" json:"pageId"`
	TActivityId        *uuid.UUID    `gorm:"type:uuid;" json:"activityId"`
	Page               *PageInfo     `gorm:"foreignKey:TPageId;" json:"page"`
	Message            string        `gorm:"type:text" json:"message"`
	Activity           *ActivityInfo `gorm:"foreignKey:TActivityId;" json:"activity"`
	Name               string        `gorm:"type:varchar(255)" json:"name"`
	Code               string        `gorm:"type:varchar(255)" json:"code"`
	IconName           *string       `gorm:"type:varchar(255)" json:"iconName"`
	DisplayCloseButton bool          `gorm:"type:boolean; default: true;" json:"displayCloseButton"`
	UseCustomIcon      bool          `gorm:"type:boolean; default: false;" json:"useCustomIcon"`
	TCustomIconMediaId *uuid.UUID    `gorm:"type:uuid" json:"customIconMediaId"`
	CustomIcon         *Media        `gorm:"foreignKey:TCustomIconMediaId;" json:"customIcon"`
	// TOrganizationId         *int64                    `gorm:"type:bigint; default: 1" json:"organizationId"`
	TOrganizationId         *uuid.UUID                `gorm:"type:uuid" json:"organizationId"`
	Organization            *OrganizationInfo         `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	BgColor                 *string                   `gorm:"type:varchar(255)" json:"bgColor"`
	TextColor               *string                   `gorm:"type:varchar(255)" json:"textColor"`
	IconColor               *string                   `gorm:"type:varchar(255)" json:"iconColor"`
	IconBgColor             *string                   `gorm:"type:varchar(255)" json:"iconBgColor"`
	NotifChannelAccessLevel []NotifChannelAccessLevel `gorm:"foreignKey:TNotifChannelId; constraint:OnDelete:CASCADE" json:"notifChannelAccessLevels"`
	NotifChannelUser        []NotifChannelUser        `gorm:"foreignKey:TNotifChannelId; constraint:OnDelete:CASCADE" json:"notifChannelUsers"`
}

func (NotifChannelInfo) TableName() string {
	return "t_notif_channel"
}

type NotifChannelInfo struct {
	baseapp.BaseModel
	Name               string        `gorm:"type:varchar(255)" json:"name"`
	Code               string        `gorm:"type:varchar(255)" json:"code"`
	BgColor            *string       `gorm:"type:varchar(255)" json:"bgColor"`
	TextColor          *string       `gorm:"type:varchar(255)" json:"textColor"`
	IconColor          *string       `gorm:"type:varchar(255)" json:"iconColor"`
	IconBgColor        *string       `gorm:"type:varchar(255)" json:"iconBgColor"`
	IconName           *string       `gorm:"type:varchar(255)" json:"iconName"`
	DisplayCloseButton bool          `gorm:"type:boolean; default: true;" json:"displayCloseButton"`
	UseCustomIcon      bool          `gorm:"type:boolean; default: false;" json:"useCustomIcon"`
	TCustomIconMediaId *uuid.UUID    `gorm:"type:uuid" json:"customIconMediaId"`
	CustomIcon         *Media        `gorm:"foreignKey:TCustomIconMediaId;" json:"customIcon"`
	TPageId            *uuid.UUID    `gorm:"type:uuid;" json:"pageId"`
	Page               *PageInfo     `gorm:"foreignKey:TPageId;" json:"page"`
	TActivityId        *uuid.UUID    `gorm:"type:uuid;" json:"activityId"`
	Activity           *ActivityInfo `gorm:"foreignKey:TActivityId;" json:"activity"`
}

func (NotifChannelAccessLevel) TableName() string {
	return "t_notif_channel_access_level"
}

type NotifChannelAccessLevel struct {
	baseapp.BaseModel
	TNotifChannelId uuid.UUID        `gorm:"type:uuid;" json:"notifChannelId"`
	TAccessLevelId  uuid.UUID        `gorm:"type:uuid;" json:"accessLevelId"`
	AccessLevel     *AccessLevelInfo `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}

func (NotifChannelUser) TableName() string {
	return "t_notif_channel_user"
}

type NotifChannelUser struct {
	baseapp.BaseModel
	TNotifChannelId uuid.UUID `gorm:"type:uuid;" json:"notifChannelId"`
	TAccountId      uuid.UUID `gorm:"type:uuid;" json:"accountId"`
	Account         *Account  `gorm:"foreignKey:TAccountId;" json:"account"`
}
