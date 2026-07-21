package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PelabuhanMuat) TableName() string {
	return "t_pelabuhan_muat"
}

type PelabuhanMuat struct {
	baseapp.BaseModel
	TMasterEpiPelabuhanMuatId uuid.UUID               `gorm:"type:uuid" json:"masterEpiPelabuhanMuatId"`
	MasterEpiPelabuhanMuat    *MasterEpiPelabuhanMuat `gorm:"foreignKey:TMasterEpiPelabuhanMuatId;" json:"masterEpiPelabuhanMuat"`
	PbbkbCode                 *string                 `gorm:"type:varchar(255)" json:"pbbkbCode"`
	Pbbkb                     *ConfigDataInfo         `gorm:"foreignKey:PbbkbCode;references:Code;" json:"pbbkb"`
	TOrganizationId           *uuid.UUID              `gorm:"type:uuid" json:"organizationId"`
	Organization              *Organization           `gorm:"foreignKey:TOrganizationId;" json:"organization"`
}
