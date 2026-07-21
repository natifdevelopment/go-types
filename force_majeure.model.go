package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (ForceMajeure) TableName() string {
	return "t_force_majeure"
}

type ForceMajeure struct {
	baseapp.BaseModel
	PemasokId uuid.UUID       `gorm:"type:uuid" json:"PemasokId"`
	Pemasok   Pemasok         `gorm:"foreignKey:PemasokId" json:"Pemasok"`
	TglBl     DateField `gorm:"type:date" json:"tglBl"`
	Hari      int             `gorm:"type:int" json:"hari"`
}
