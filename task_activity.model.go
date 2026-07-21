package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"

	"github.com/google/uuid"
)

func (Task_activity) TableName() string {
	return "t_task_activity"
}

type Task_activity struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TRefId              *uuid.UUID        `gorm:"type:uuid" json:"-"`
	TOrganizationId     *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization        *Organization     `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId      *uuid.UUID        `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel         *AccessLevel      `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId          *uuid.UUID        `gorm:"type:uuid" json:"accountId"`
	Account             *AccountInfo      `gorm:"foreignKey:TAccountId" json:"account"`
	Notes               *string           `gorm:"type:text" json:"notes"`
	DueDate             time.Time         `gorm:"type:timestamp" json:"dueDate"`
	TPageActivityId     *uuid.UUID        `gorm:"type:uuid" json:"pageActivityId"` // Current page
	PageActivity        *PageActivity     `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
}
