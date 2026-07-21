package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Notif) TableName() string {
	return "t_notif"
}

type Notif struct {
	baseapp.BaseModel
	TNotifChannelId *uuid.UUID        `gorm:"type:uuid;" json:"notifChannelId"`
	NotifChannel    *NotifChannelInfo `gorm:"foreignKey:TNotifChannelId;" json:"notifChannel"`
	TAccountId      *uuid.UUID        `gorm:"type:uuid;" json:"accountId"`
	Account         *Account          `gorm:"foreignKey:TAccountId;" json:"account"`
	Title           *string           `gorm:"type:varchar(255)" json:"name"`
	Message         *string           `gorm:"type:text" json:"message"`
	Url             *string           `gorm:"type:text" json:"url"`
	IsRead          bool              `gorm:"type:boolean;default:false" json:"isRead"`
	IsNew           bool              `gorm:"type:boolean;default:true" json:"isNew"`
}
