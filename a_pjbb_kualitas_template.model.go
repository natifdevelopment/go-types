package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Pjbb_kualitas_template) TableName() string {
	return "t_pjbb_kualitas_template"
}

type Pjbb_kualitas_template struct {
	baseapp.BaseModel
	Name  *string                    `gorm:"type:varchar(255)" json:"name"`
	Items []PjbbKualitasTemplateItem `gorm:"foreignKey:TPjbbKualitasTemplateId; constraint:OnDelete:CASCADE" json:"items"`
}

func (PjbbKualitasTemplateItem) TableName() string {
	return "t_pjbb_kualitas_template_item"
}

type PjbbKualitasTemplateItem struct {
	baseapp.BaseModel
	Type                    *string                   `gorm:"type:varchar(255)" json:"type"` // penyesuaian/denda_penolakan
	TPjbbKualitasTemplateId uuid.UUID                 `gorm:"type:uuid" json:"pjbbKualitasTemplateId"`
	PjbbKualitasTemplate    *Pjbb_kualitas_template   `gorm:"foreignKey:TPjbbKualitasTemplateId" json:"pjbbKualitasTemplate"`
	TipeKualitas            Pjbb_kualitasItemType     `gorm:"type:varchar(255);default:'none'" json:"tipeKualitas"`
	TUraianId               uuid.UUID                 `gorm:"type:uuid" json:"uraianId"`
	Uraian                  *Uraian                   `gorm:"foreignKey:TUraianId" json:"uraian"`
	TRefId                  uuid.UUID                 `gorm:"type:uuid" json:"refId"`
	BasisCode               *string                   `gorm:"type:varchar(255)" json:"basisCode"`
	Basis                   *ConfigDataInfo           `gorm:"foreignKey:BasisCode;references:Code;" json:"basis"`
	Tipikal                 *string                   `gorm:"type:varchar(255)" json:"tipikal"`
	Penolakan               *string                   `gorm:"type:varchar(255)" json:"penolakan"`
	Simbol                  *string                   `gorm:"type:varchar(255)" json:"simbol"`
	BatasBawah              *string                   `gorm:"type:varchar(255)" json:"batasBawah"`
	BatasAtas               *string                   `gorm:"type:varchar(255)" json:"batasAtas"`
	TipeData                Pjbb_kualitasItemDataType `gorm:"type:varchar(255);default:'Num'" json:"tipeData"`
	TurunHarga              *string                   `gorm:"type:float" json:"turunHarga"`
	Entangled               bool                      `gorm:"type:bool" json:"entangled"`
	OrderNumber             int                       `json:"orderNumber"`
}
