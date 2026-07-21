package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Triwulan) TableName() string {
	return "t_triwulan"
}

type Triwulan struct {
	baseapp.BaseModel
	Triwulan      string                `gorm:"type:varchar(255)" json:"triwulan"`
	NonTriwulan   bool                  `gorm:"type:bool;default:false" json:"nonTriwulan"`
	TriwulanRange *uint                 `gorm:"type:int" json:"triwulanRange"`
	Sd            *uint                 `gorm:"type:int" json:"sd"`
	Keterangan    *string               `gorm:"type:text" json:"keterangan"`
}
