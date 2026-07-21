package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (PjbbHarga) TableName() string {
	return "t_pjbb_harga"
}

type PjbbHarga struct {
	baseapp.BaseModel
	TPemasokId       *uuid.UUID       `gorm:"type:uuid" json:"pemasokId"`
	Pemasok          *Organization    `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId    *uuid.UUID       `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit       *Organization    `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPjbbId          *uuid.UUID       `gorm:"type:uuid" json:"pjbbId"`
	Pjbb             *Pjbb            `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	SkemaKontrakCode *JenisPengiriman `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak     *ConfigDataInfo  `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TPelabuhanMuatId *uuid.UUID       `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat    *PelabuhanMuat   `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekCvId        *uuid.UUID       `gorm:"type:uuid" json:"spekCvId"`
	SpekCv           *Spek            `gorm:"foreignKey:TSpekCvId" json:"spekCv"`
	TModaId          *uuid.UUID       `gorm:"type:uuid" json:"modaId"`
	Moda             *Moda            `gorm:"foreignKey:TModaId" json:"moda"`
	HargaFob         float64          `gorm:"type:float" json:"hargaFob"`
	HsdBase          float64          `gorm:"type:float" json:"hsdBase"`
	HrgTrans         float64          `gorm:"type:float" json:"hrgTrans"`
	HrgPbm           float64          `gorm:"type:float" json:"hrgPbm"`
	FaktorX          float64          `gorm:"type:float" json:"faktorX"`
	FaktorY          float64          `gorm:"type:float" json:"faktorY"`
	MfoBase          float64          `gorm:"type:float" json:"mfoBase"`
	FaktorY2         float64          `gorm:"type:float" json:"faktorY2"`
	HrgTransTrucking *float64         `gorm:"type:float" json:"hrgTransTrucking"`
	Lhv              *float64         `gorm:"type:float" json:"lhv"`
}
