package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Opel) TableName() string {
	return "t_opel"
}

type Opel struct {
	baseapp.BaseModel
	NamaPerusahaan string       `gorm:"type:varchar(255)" json:"namaPerusahaan"`
	NamaPic        string       `gorm:"type:varchar(255)" json:"namaPic"`
	TPembangkitId  uuid.UUID    `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit     Organization `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
}
