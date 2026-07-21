package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"

	"github.com/google/uuid"
)

func (AccountSession) TableName() string {
	return "t_account_session"
}

type AccountSession struct {
	baseapp.BaseModel
	TAccountId *uuid.UUID   `gorm:"type:uuid" json:"accountId"`
	Account    *AccountInfo `gorm:"foreignKey:TAccountId" json:"account"`
	PeerId     string       `gorm:"type:varchar(255)" json:"peerId"`
	SessionId  string       `gorm:"type:varchar(255)" json:"sessionId"`
	TDeviceId  *uuid.UUID   `gorm:"type:uuid" json:"deviceId"`
	LastSeen   *time.Time   `gorm:"type:timestamp" json:"lastSeen"`
}
