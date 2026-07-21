package types
import baseapp "github.com/natifdevelopment/go-baseapp"
type JangkaWaktu struct {
	baseapp.BaseModel
	JenisKontrak string `gorm:"type:varchar(255)" json:"jenisKontrak"`
	NamaSingkat  string `gorm:"type:varchar(255)" json:"namaSingkat"`
	Keterangan   string `gorm:"type:text" json:"keterangan"`
}

func (JangkaWaktu) TableName() string {
	return "t_jangka_waktu"
}
