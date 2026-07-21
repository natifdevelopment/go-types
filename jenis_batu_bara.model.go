package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (JenisBatuBara) TableName() string {
	return "t_jenis_batu_bara"
}

type JenisBatuBara struct {
	baseapp.BaseModel
	JenisBatuBara string  `gorm:"type:varchar(255)" json:"jenisBatuBara"`
	NamaSingkat   string  `gorm:"type:varchar(255)" json:"namaSingkat"`
	Keterangan    string  `gorm:"type:text" json:"keterangan"`
	CaloriStart   float64 `gorm:"type:float" json:"caloriStart"`
	CaloriEnd     float64 `gorm:"type:float" json:"caloriEnd"`
}
