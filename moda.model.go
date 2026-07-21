package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Moda) TableName() string {
	return "t_moda"
}

type Moda struct {
	baseapp.BaseModel
	Moda                  string          `gorm:"type:varchar(255);unique;" json:"moda"`
	VolumeMin             float64         `gorm:"type:float" json:"volumeMin"`
	VolumeMax             float64         `gorm:"type:float" json:"volumeMax"`
	JenisTransportasiCode *string         `gorm:"type:varchar(255)" json:"jenisTransportasiCode"`
	JenisTransportasi     *ConfigDataInfo `gorm:"foreignKey:JenisTransportasiCode;references:Code;" json:"jenisTransportasi"`
}
