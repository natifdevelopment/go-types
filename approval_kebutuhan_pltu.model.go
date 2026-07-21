package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type ApprovalKebutuhanPembangkitPersetujuan string

const (
	ApprovalKebutuhanPembangkitPersetujuanSetuju      ApprovalKebutuhanPembangkitPersetujuan = "setuju"
	ApprovalKebutuhanPembangkitPersetujuanTidakSetuju ApprovalKebutuhanPembangkitPersetujuan = "tidak setuju"
)

func (o ApprovalKebutuhanPembangkitPersetujuan) IsValid() bool {
	switch o {
	case ApprovalKebutuhanPembangkitPersetujuanSetuju, ApprovalKebutuhanPembangkitPersetujuanTidakSetuju:
		return true
	default:
		return false
	}
}

func (ApprovalKebutuhanPembangkit) TableName() string {
	return "t_approval_kebutuhan_pembangkit"
}

type ApprovalKebutuhanPembangkit struct {
	baseapp.BaseModel
	NoJadwal      string                                 `gorm:"type:varchar(255)" json:"noJadwal"`
	TglPenerimaan DateField                        `gorm:"type:date" json:"tglPenerimaan"`
	Pembangkit    string                                 `gorm:"type:varchar(255)" json:"pembangkit"`
	PemasokId     uuid.UUID                              `gorm:"type:uuid" json:"PemasokId"`
	Pemasok       Pemasok                                `gorm:"foreignKey:PemasokId" json:"Pemasok"`
	VolumeKirim   float64                                `gorm:"type:float" json:"volumeKirim"`
	Persetujuan   ApprovalKebutuhanPembangkitPersetujuan `gorm:"type:varchar(255)" json:"persetujuan"`
}
