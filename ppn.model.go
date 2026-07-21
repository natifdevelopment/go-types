package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (Ppn) TableName() string {
	return "t_ppn"
}

type Ppn struct {
	baseapp.BaseModel
	Ppn                   float64          `gorm:"type:float" json:"ppn"`
	TaxCodeBatubaraCode   *string          `gorm:"type:varchar(255)" json:"taxCodeBatubaraCode"`
	TaxCodeBatubara       *ConfigData      `gorm:"foreignKey:TaxCodeBatubaraCode;references:Code;" json:"taxCodeBatubara"`
	TaxCodeJasaCode       *string          `gorm:"type:varchar(255)" json:"taxCodeJasaCode"`
	TaxCodeJasa           *ConfigData      `gorm:"foreignKey:TaxCodeJasaCode;references:Code;" json:"taxCodeJasa"`
	TaxCodeWapuBarangCode *string          `gorm:"type:varchar(255)" json:"taxCodeWapuBarangCode"`
	TaxCodeWapuBarang     *ConfigData      `gorm:"foreignKey:TaxCodeWapuBarangCode;references:Code;" json:"taxCodeWapuBarang"`
	TaxCodeWapuJasaCode   *string          `gorm:"type:varchar(255)" json:"taxCodeWapuJasaCode"`
	TaxCodeWapuJasa       *ConfigData      `gorm:"foreignKey:TaxCodeWapuJasaCode;references:Code;" json:"taxCodeWapuJasa"`
	TglBerlakuAwal        *DateField `gorm:"type:date" json:"tglBerlakuAwal"`
	TglBerlakuAkhir       *DateField `gorm:"type:date" json:"tglBerlakuAkhir"`
	IsDefault             bool             `gorm:"type:boolean;default:false" json:"isDefault"`
}
