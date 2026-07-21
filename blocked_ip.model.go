package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (BlockedIp) TableName() string {
	return "t_blocked_ip"
}

type BlockedIp struct {
	baseapp.BaseModel
	IpAddress     string `gorm:"type:varchar(145)" json:"-"`
	BlockedReason string `gorm:"type:text" json:"-"`
}
