package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type ApprovalPo struct {
	baseapp.BaseModel
	NoPengiriman string    `gorm:"type:varchar(255)" json:"NoPengiriman"`
	PemasokId    uuid.UUID `gorm:"type:uuid" json:"PemasokId"`
	Pemasok      Pemasok   `gorm:"foreignKey:PemasokId" json:"Pemasok"`
	GrossPrice   *float64  `gorm:"type:float" json:"GrossPrice"`
	FreightValue *float64  `gorm:"type:float" json:"FreightValue"`
	Quantity     *int64    `gorm:"type:bigint" json:"Quantity"`
}

func (ApprovalPo) TableName() string {
	return "t_approval_po"
}
