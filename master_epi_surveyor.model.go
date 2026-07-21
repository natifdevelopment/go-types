package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (MasterEpiSurveyor) TableName() string {
	return "t_master_epi_surveyor"
}

type MasterEpiSurveyor struct {
	baseapp.BaseModel
	NamaSurveyor  string `gorm:"type:varchar(255)" json:"namaSurveyor"`
	NamaSingkat   string `gorm:"type:varchar(255)" json:"namaSingkat"`
	Alamat        string `gorm:"type:varchar(255)" json:"alamat"`
	ContactPerson string `gorm:"type:varchar(255)" json:"contactPerson"`
	Email         string `gorm:"type:varchar(255)" json:"email"`
	Telp          string `gorm:"type:varchar(255)" json:"telp"`
	Keterangan    string `gorm:"type:varchar(255)" json:"keterangan"`
}