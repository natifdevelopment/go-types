package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Spek) TableName() string {
	return "t_spek"
}

type Spek struct {
	baseapp.BaseModel
	SpekCalori float64 `gorm:"type:float" json:"spekCalori"`
}
