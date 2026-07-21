package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

type ApprovalPemakaian struct {
	baseapp.BaseModel
	NoPemakaian     string          `gorm:"type:varchar(255)" json:"NoPemakaian"`
	Unit            string          `gorm:"type:varchar(255)" json:"Unit"`
	Tanggal         DateField `gorm:"type:date" json:"Tanggal"`
	VolumePemakaian float64         `gorm:"type:float" json:"VolumePemakaian"`
}

func (ApprovalPemakaian) TableName() string {
	return "t_approval_pemakaian"
}
