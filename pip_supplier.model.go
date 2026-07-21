package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (PipSupplier) TableName() string {
	return "t_pip_supplier"
}

type PipSupplier struct {
	baseapp.BaseModel
	VendorId             *string          `gorm:"type:varchar(255)" json:"vendorId"`
	VendorName           *string          `gorm:"type:varchar(255)" json:"vendorName"`
	VendorTypeLookupCode *string          `gorm:"type:varchar(255)" json:"vendorTypeLookupCode"`
	ResponseHash         *string          `gorm:"type:varchar(255)" json:"-"`
	PipSupplierMap       []PipSupplierMap `gorm:"foreignKey:TPipSupplierId;constraint:OnDelete:CASCADE" json:"-"`
}
