package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (ActivityItem) TableName() string {
	return "t_activity_item"
}

type ActivityItem struct {
	baseapp.BaseModel
	TActivityId  uuid.UUID    `gorm:"type:uuid" json:"activityId"`
	Activity     ActivityInfo `gorm:"foreignKey:TActivityId" json:"activity"`
	Name         string       `gorm:"type:varchar(255)" json:"name"`
	Code         string       `gorm:"type:varchar(255)" json:"code"`
	Description  string       `gorm:"type:varchar(255)" json:"description"`
	ConnectQuery bool         `gorm:"type:boolean" json:"connectQuery"`
}
