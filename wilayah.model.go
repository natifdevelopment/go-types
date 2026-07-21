package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Wilayah) TableName() string {
	return "t_wilayah"
}

type Wilayah struct {
	baseapp.BaseModel
	Nama      string `gorm:"type:varchar(255)" json:"nama"`
	Deskripsi string `gorm:"type:varchar(255)" json:"deskripsi"`
}
