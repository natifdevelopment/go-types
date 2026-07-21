package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (PipSupplierMap) TableName() string {
	return "t_pip_supplier_map"
}

type PipSupplierMap struct {
	baseapp.BaseModel
	TPipSupplierId *uuid.UUID   `gorm:"type:uuid" json:"pipSupplierId"`
	PipSupplier    *PipSupplier `gorm:"foreignKey:TPipSupplierId" json:"pipSupplier"`
	TPemasokId     *uuid.UUID   `gorm:"type:uuid" json:"pemasokId"`
	Pemasok        *Pemasok     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
}
