package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (MasterNoRekening) TableName() string {
	return "t_master_no_rekening"
}

type MasterNoRekening struct {
	baseapp.BaseModel
	TMasterEpiBankId *uuid.UUID     `gorm:"type:uuid" json:"masterEpiBankId"`
	MasterEpiBank    *MasterEpiBank `gorm:"foreignKey:TMasterEpiBankId" json:"masterEpiBank"`
	TOrganizationId  *uuid.UUID     `gorm:"type:uuid" json:"organizationId"`
	Organization     *Organization  `gorm:"foreignKey:TOrganizationId" json:"organization"`
	AtasNama         string         `gorm:"type:varchar(255)" json:"atasNama"`
	NamaBank         string         `gorm:"type:varchar(255)" json:"namaBank"`
	NoRekening       string         `gorm:"type:varchar(255)" json:"noRekening"`
	Alamat           *string        `gorm:"type:varchar(255)" json:"alamat"`
	StatusNoRekening bool           `gorm:"type:boolean" json:"statusNoRekening"`
}
