package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (PipPembangkit) TableName() string {
	return "t_pip_pembangkit"
}

type PipPembangkit struct {
	baseapp.BaseModel
	UnitId           *string            `gorm:"type:varchar(255)" json:"unitId"`
	UnitKd           *string            `gorm:"type:varchar(255)" json:"unitKd"`
	UnitNama         *string            `gorm:"type:varchar(255)" json:"unitNama"`
	Jenis            *string            `gorm:"type:varchar(255)" json:"jenis"`
	ResponseHash     *string            `gorm:"type:varchar(255)" json:"-"`
	PipPembangkitMap []PipPembangkitMap `gorm:"foreignKey:TPipPembangkitId;constraint:OnDelete:CASCADE" json:"-"`
}
