package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (MasterUnit) TableName() string {
	return "t_master_unit"
}

type MasterUnit struct {
	baseapp.BaseModel
	TPembangkitId  *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit     *OrganizationInfo `gorm:"foreignKey:TPembangkitId;" json:"pembangkit"`
	KodeUnit       *string           `gorm:"type:varchar(255)" json:"kodeUnit"`
	Unit           *string           `gorm:"type:varchar(255)" json:"unit"`
	JenisBoiler    *string           `gorm:"type:varchar(255)" json:"jenisBoiler"`
	RentangBoiler  *string           `gorm:"type:varchar(255)" json:"rentangBoiler"`
	DayaTerpasang  *float64          `gorm:"type:float" json:"dayaTerpasang"`
	UnitCode       *string           `gorm:"type:varchar(255)" json:"unitCode"`
	KodeMesin      *string           `gorm:"type:varchar(255)" json:"kodeMesin"`
	NamaMesin      *string           `gorm:"type:varchar(255)" json:"namaMesin"`
	NphrCommission *string           `gorm:"type:varchar(255)" json:"nphrCommission"`
	CostCenter     *string           `gorm:"type:varchar(255)" json:"costCenter"`
}
