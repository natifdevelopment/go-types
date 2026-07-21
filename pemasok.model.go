package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Pemasok) TableName() string {
	return "t_pemasok"
}

//alamat satu aja

type Pemasok struct {
	baseapp.BaseModel
	NamaPemasok        string           `gorm:"type:varchar(255)" json:"namaPemasok"`
	NamaSingkat        string           `gorm:"type:varchar(100)" json:"namaSingkat"`
	NamaBrand          string           `gorm:"type:varchar(100)" json:"namaBrand"`
	NamaPerusahaan     string           `gorm:"type:varchar(255)" json:"namaPerusahaan"`
	NamaPimpinan       string           `gorm:"type:varchar(255)" json:"namaPimpinan"`
	Alamat             string           `gorm:"type:text" json:"alamat"`
	AlamatPerusahaan   string           `gorm:"type:text" json:"alamatPerusahaan"`
	KodePos            int              `gorm:"type:int" json:"kodePos"`
	Telp               string           `gorm:"type:varchar(25)" json:"telp"`
	Fax                string           `gorm:"type:varchar(25)" json:"fax"`
	Email              string           `gorm:"type:varchar(255)" json:"email"`
	JenisPerizinanCode *string          `gorm:"type:varchar(255)" json:"jenisPerizinanCode"`
	JenisPerizinan     *ConfigData      `gorm:"foreignKey:JenisPerizinanCode;references:Code;" json:"jenisPerizinan"`
	IdModi             *uint32          `gorm:"type:int" json:"idModi"`
	TglCutoff          *DateField `gorm:"type:date" json:"tglCutoff"`
	TaxCodeWapu        bool             `gorm:"type:boolean" json:"taxCodeWapu"`
	TPpnId             *uuid.UUID       `gorm:"type:uuid" json:"ppnId"`
	Ppn                *Ppn             `gorm:"foreignKey:TPpnId" json:"ppn"`
	TPipSupplierId     *uuid.UUID       `gorm:"type:uuid" json:"pipSupplierId"`
	PipSupplier        *PipSupplier     `gorm:"foreignKey:TPipSupplierId" json:"pipSupplier"`
	VendorCode         *string          `gorm:"type:varchar(255)" json:"vendorCode"`
	SuffixFakturPajak  *string          `gorm:"type:varchar(10)" json:"suffixFakturPajak"`
	PrefixFakturPajak  *string          `gorm:"type:varchar(10)" json:"prefixFakturPajak"`
	PipSupplierMap     []PipSupplierMap `gorm:"foreignKey:TPemasokId;constraint:OnDelete:CASCADE" json:"-"`

	TDokumenNPWPId           *uuid.UUID `gorm:"type:uuid" json:"dokumenNPWPId"`
	DokumenNPWP              *Media     `gorm:"foreignKey:TDokumenNPWPId" json:"dokumenNPWP"`
	TDokumenNIBId            *uuid.UUID `gorm:"type:uuid" json:"dokumenNIBId"`
	DokumenNIB               *Media     `gorm:"foreignKey:TDokumenNIBId" json:"dokumenNIB"`
	TDokumenSIUPId           *uuid.UUID `gorm:"type:uuid" json:"dokumenSIUPId"`
	DokumenSIUP              *Media     `gorm:"foreignKey:TDokumenSIUPId" json:"dokumenSIUP"`
	TDokumenTDPId            *uuid.UUID `gorm:"type:uuid" json:"dokumenTDPId"`
	DokumenTDP               *Media     `gorm:"foreignKey:TDokumenTDPId" json:"dokumenTDP"`
	TDokumenAktaPerusahaanId *uuid.UUID `gorm:"type:uuid" json:"dokumenAktaPerusahaanId"`
	DokumenAktaPerusahaan    *Media     `gorm:"foreignKey:TDokumenAktaPerusahaanId" json:"dokumenAktaPerusahaan"`
}
