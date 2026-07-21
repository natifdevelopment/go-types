package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (MasterEpiTongkang) TableName() string {
	return "t_master_epi_tongkang"
}

type MasterEpiTongkang struct {
	baseapp.BaseModel
	Nama           string     `gorm:"type:varchar(255)" json:"nama"`
	Kapasitas      *float64   `gorm:"type:decimal(20,3)" json:"kapasitas"`
	StatusTongkang bool       `gorm:"type:bool" json:"statusTongkang"`
	TModaId        *uuid.UUID `gorm:"type:uuid" json:"modaId"`
	Moda           *Moda      `gorm:"foreignKey:TModaId" json:"moda"`
}
