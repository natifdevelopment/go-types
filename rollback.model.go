package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type RollbackStatus string

const (
	RollbackStatusAktif   RollbackStatus = "AKTIF"   // bank garansi aktif
	RollbackStatusSelesai RollbackStatus = "SELESAI" // masa belaku bank garansi sudah lewat
)

func (Rollback) TableName() string {
	return "t_rollback"
}

type Rollback struct {
	baseapp.BaseModel
	RollbackStatus  RollbackStatus `gorm:"type:varchar(255);default:'AKTIF';" json:"rollbackStatus"`
	TRefId          *uuid.UUID     `gorm:"type:uuid" json:"refId"`
	TPageActivityId *uuid.UUID     `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity  `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
	Items           []RollbackItem `gorm:"foreignKey:TRollbackId;constraint:OnDelete:CASCADE" json:"items"`
}

func (RollbackItem) TableName() string {
	return "t_rollback_item"
}

type RollbackItem struct {
	baseapp.BaseModel
	TRollbackId     *uuid.UUID    `gorm:"type:uuid" json:"rollbackId"`
	Rollback        *Rollback     `gorm:"foreignKey:TRollbackId" json:"rollback"`
	TPageActivityId *uuid.UUID    `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
}
