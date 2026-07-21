package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (SumberTambang) TableName() string {
	return "t_sumber_tambang"
}

type SumberTambang struct {
	baseapp.BaseModel
	Name         string      `gorm:"type:varchar(255)" json:"name"`
	IsActive     bool        `gorm:"type:boolean;default:true" json:"isActive"`
	Lokasi       *string     `gorm:"type:varchar(255)" json:"lokasi"`
	TipeIzinCode *string     `gorm:"type:varchar(255)" json:"tipeIzinCode"`
	TipeIzin     *ConfigData `gorm:"foreignKey:TipeIzinCode;references:Code;" json:"tipeIzin"`
	Latitude     *float64    `gorm:"type:float" json:"latitude"`
	Longitude    *float64    `gorm:"type:float" json:"longitude"`
}
