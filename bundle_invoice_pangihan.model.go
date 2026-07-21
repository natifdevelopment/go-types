package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Bundle_invoice_pangihan) TableName() string {
	return "t_bundle_invoice_pangihan"
}

type Bundle_invoice_pangihan struct {
	baseapp.BaseModel
	NoPenagihan               string                      `gorm:"type:varchar(255)" json:"noPenagihan"`
	BundleInvoicePangihanItem []BundleInvoicePangihanItem `gorm:"foreignKey:TBundleInvoicePangihanId; constraint:OnDelete:CASCADE" json:"items"`
	PermohonanPembayaran      []PermohonanPembayaran      `gorm:"foreignKey:TBundleInvoicePenagihanId; constraint:OnDelete:CASCADE" json:"permohonanPembayaran"`
}

func (BundleInvoicePangihanItem) TableName() string {
	return "t_bundle_invoice_pangihan_item"
}

type BundleInvoicePangihanItem struct {
	baseapp.BaseModel
	TBundleInvoicePangihanId *uuid.UUID               `gorm:"index;type:uuid" json:"bundleInvoicePangihanId"`
	BundleInvoicePangihan    *Bundle_invoice_pangihan `gorm:"foreignKey:TBundleInvoicePangihanId" json:"bundleInvoicePangihan"`
	TProposePerhitunganId    *uuid.UUID               `gorm:"index;type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan       *ProposePerhitungan      `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
	// @not_used_yet
	// TUploadInvoicePenagihanId *uuid.UUID                `gorm:"index;type:uuid" json:"uploadInvoicePenagihanId"`
	// UploadInvoicePenagihan    *UploadInvoicePenagihan `gorm:"foreignKey:TUploadInvoicePenagihanId" json:"uploadInvoicePenagihan"`
}
