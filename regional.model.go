package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Regional) TableName() string {
	return "t_regional"
}

type Regional struct {
	baseapp.BaseModel
	Region         string `gorm:"type:varchar(255)" json:"region"`
	Keterangan     string `gorm:"type:text" json:"keterangan"`
	RegionalStatus bool   `gorm:"column:regional_status;type:bool;default:false" json:"regionalStatus"`
}
