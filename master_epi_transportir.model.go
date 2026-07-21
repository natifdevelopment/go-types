package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (MasterEpiTransportir) TableName() string {
	return "t_master_epi_transportir"
}

type MasterEpiTransportir struct {
	baseapp.BaseModel
	IdTransportir     uuid.UUID        `gorm:"type:uuid" json:"idTransportir"`
	NamaTransportir   *string          `gorm:"type:varchar(255)" json:"namaTransportir"`
	NamaSingkat       *string          `gorm:"type:varchar(255)" json:"namaSingkat"`
	Email             *string          `gorm:"type:varchar(255)" json:"email"`
	Alamat            *string          `gorm:"type:text" json:"alamat"`
	Telp              *string          `gorm:"type:varchar(20)" json:"telp"`
	StatusTransportir bool             `gorm:"type:bool;default:false" json:"statusTransportir"`
	MasaBerlakuAwal   *DateField `gorm:"type:date" json:"masaBerlakuAwal"`
	MasaBerlakuAkhir  *DateField `gorm:"type:date" json:"masaBerlakuAkhir"`
	TPpnId            *uuid.UUID       `gorm:"type:uuid" json:"ppnId"`
	Ppn               *Ppn             `gorm:"foreignKey:TPpnId" json:"ppn"`
	TaxCodeWapu       bool             `gorm:"type:boolean" json:"taxCodeWapu"`
	VendorCode        *string          `gorm:"type:varchar(255)" json:"vendorCode"`
	SuffixFakturPajak *string          `gorm:"type:varchar(10)" json:"suffixFakturPajak"`
	PrefixFakturPajak *string          `gorm:"type:varchar(10)" json:"prefixFakturPajak"`
}
