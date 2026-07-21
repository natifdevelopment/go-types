package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type OutlineAgreement struct {
	baseapp.BaseModel
	TPemasokId *uuid.UUID      `gorm:"type:uuid;uniqueIndex:pemasok_no_kontrak_no_oa_sap" json:"pemasokId"`
	Pemasok    *Organization   `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	NoKontrak  *string         `gorm:"type:varchar(255);uniqueIndex:pemasok_no_kontrak_no_oa_sap" json:"noKontrak"`
	Tipe       JenisPengiriman `gorm:"type:varchar(25)" json:"tipe"`
	NoOaSAP    *string         `gorm:"type:varchar(255);uniqueIndex:pemasok_no_kontrak_no_oa_sap" json:"noOaSap"`
}

func (OutlineAgreement) TableName() string {
	return "t_outline_agreement"
}
