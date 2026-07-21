package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Pejabat_surat_kuasa) TableName() string {
	return "t_pejabat_surat_kuasa"
}

type Pejabat_surat_kuasa struct {
	baseapp.BaseModel
	TSuratKuasaId             *uuid.UUID  `gorm:"type:uuid" json:"suratKuasaId"`
	SuratKuasa                *SuratKuasa `gorm:"foreignKey:TSuratKuasaId" json:"suratKuasa"`
	Nama                      string      `json:"nama" gorm:"type:varchar(255)"`
	Email                     string      `json:"email" gorm:"type:varchar(255)"`
	Jabatan                   string      `json:"jabatan" gorm:"type:varchar(255)"`
	TipeDokumen               *string     `json:"tipeDokumen" gorm:"type:varchar(255);"`
	StatusPeruri              *string     `json:"statusPeruri" gorm:"type:varchar(255)"`
	StatusAms                 string      `json:"statusAms" gorm:"type:varchar(255)"`
	TipePejabat               TipePejabat `json:"tipePejabat" gorm:"type:varchar(255);not null"`
	BaSerahTerima             bool        `json:"baSerahTerima"`
	BaPenyesuaianHarga        bool        `json:"baPenyesuaianHarga"`
	BaKeterlambatan           bool        `json:"baKeterlambatan"`
	RincianDendaKeterlambatan bool        `json:"rincianDendaKeterlambatan"`
}
