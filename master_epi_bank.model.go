package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (MasterEpiBank) TableName() string {
	return "t_master_epi_bank"
}

type MasterEpiBank struct {
	baseapp.BaseModel
	Nama string `gorm:"type:varchar(255)" json:"nama"`
}
