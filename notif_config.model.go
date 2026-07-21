package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (NotifConfig) TableName() string {
	return "t_notif_config"
}

type NotifConfig struct {
	baseapp.BaseModel
	Name            string           `gorm:"type:varchar(255)" json:"name"`
	Code            string           `gorm:"type:varchar(255)" json:"code"`
	Priority        int64            `gorm:"type:int; default:0; unique;" json:"priority"`
	TOrganizationId *uuid.UUID           `gorm:"type:uuid" json:"organizationId"`
	Organization    OrganizationInfo `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	NotifChannel    []NotifChannel   `gorm:"foreignKey:TNotifConfigId;constraint:OnDelete:CASCADE" json:"notifChannels"`
}
