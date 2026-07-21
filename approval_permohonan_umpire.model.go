package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type ApprovalPermohonanUmpire struct {
	baseapp.BaseModel
	NoPengiriman    string         `gorm:"type:varchar(255)" json:"NoPengiriman"`
	PemasokId       uuid.UUID      `gorm:"type:uuid" json:"PemasokId"`
	Pemasok         Pemasok        `gorm:"foreignKey:PemasokId" json:"Pemasok"`
	KapalTongkangId uuid.UUID      `gorm:"type:uuid" json:"KapalTongkangId"`
	KapalTongkang   MasterTongkang `gorm:"foreignKey:KapalTongkangId" json:"KapalTongkang"`
	VolumeDS        float64        `gorm:"type:float" json:"VolumeDS"`
	SurveyorId      uuid.UUID      `gorm:"type:uuid" json:"SurveyorId"`
	Surveyor        Surveyor       `gorm:"foreignKey:SurveyorId" json:"Surveyor"`
	NoCOA           string         `gorm:"type:varchar(255)" json:"NoCOA"`
}

func (ApprovalPermohonanUmpire) TableName() string {
	return "t_approval_permohonan_umpire"
}
