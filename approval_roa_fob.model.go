package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (ApprovalRoaFob) TableName() string {
	return "t_approval_roa_fob"
}

type ApprovalRoaFob struct {
	baseapp.BaseModel
	TJadwalPengirimanId   *uuid.UUID          `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman      *JadwalPengiriman   `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TSurveyorPembangkitId *uuid.UUID          `gorm:"type:uuid" json:"surveyorPembangkitId"`
	SurveyorPembangkit    *Organization       `gorm:"foreignKey:TSurveyorPembangkitId" json:"surveyorPembangkit"`
	LoadingInfo           []LoadingInfo       `gorm:"foreignKey:TApprovalRoaFobId; constraint:OnDelete:CASCADE" json:"loadingInfo"`
	RoaItems              []RoaPembangkitItem `gorm:"foreignKey:TApprovalRoaFobId; constraint:OnDelete:CASCADE" json:"roaItems"`
}

func (RoaPembangkitItem) TableName() string {
	return "t_roa_pembangkit_item"
}

type RoaPembangkitItem struct {
	baseapp.BaseModel
	TApprovalRoaFobId *uuid.UUID      `gorm:"type:uuid" json:"approvalRoaFobId"`
	ApprovalRoaFob    *ApprovalRoaFob `gorm:"foreignKey:TApprovalRoaFobId" json:"approvalRoaFob"`
	TUraianId         *uuid.UUID      `gorm:"type:uuid" json:"uraianId"`
	Uraian            *Uraian         `gorm:"foreignKey:TUraianId" json:"uraian"`
	Hasil             *float32        `gorm:"type:float" json:"hasil"`
	BatasPenolakan    *string         `gorm:"type:text" json:"batasPenolakan"`
	IsPassed          bool            `gorm:"type:bool;default:false" json:"isPassed"`
	Description       *string         `gorm:"type:text" json:"description"`
}
