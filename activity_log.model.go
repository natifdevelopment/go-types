package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"

	"github.com/google/uuid"
)

func (ActivityLog) TableName() string {
	return "t_activity_log"
}

type ActivityLogType string

const (
	LogUserActivity          ActivityLogType = "user_activity"
	LogSystemActivityLogType ActivityLogType = "system"
	LogAuthActivityLogType   ActivityLogType = "auth"
	//input other
)

func (o ActivityLogType) IsValid() bool {
	return o == LogUserActivity || o == LogSystemActivityLogType || o == LogAuthActivityLogType
	//add other
}

type ActivityLog struct {
	baseapp.BaseModel
	Type         ActivityLogType  `gorm:"type:varchar(50)" json:"type"`
	IpAddress    string           `gorm:"type:text" json:"ipAddress"`
	UserId       *uuid.UUID       `gorm:"type:uuid" json:"userId"`
	UserName     string           `gorm:"text" json:"userName"`
	UserNameHash string           `gorm:"type:varchar(255)"`
	PageId       *uuid.UUID       `gorm:"type:uuid" json:"pageId"`
	PageName     string           `gorm:"type:varchar(255)" json:"pageName"`
	PageCode     string           `gorm:"type:varchar(255)" json:"pageCode"`
	ActivityId   *uuid.UUID       `gorm:"type:uuid" json:"activityId"`
	ActivityName string           `gorm:"type:varchar(255)" json:"activityName"`
	ActivityCode string           `gorm:"type:varchar(255)" json:"activityCode"`
	StatusCode   int              `gorm:"type:int" json:"statusCode"`
	ErrorMessage *string          `gorm:"text" json:"errorMessage"`
	DataId       *uuid.UUID       `gorm:"type:uuid" json:"dataId"`
	DataName     *string          `gorm:"text" json:"DataName"`
	Data         *json.RawMessage `gorm:"type:json" json:"data"`
	Description  *string          `gorm:"type:text" json:"description"`
	TRefId       *uuid.UUID       `gorm:"type:uuid" json:"-"`
	TTargetId    *uuid.UUID       `gorm:"type:uuid" json:"-"`
	Module       *string          `gorm:"type:varchar(255)" json:"module"`
}
