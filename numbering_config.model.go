package types
import baseapp "github.com/natifdevelopment/go-baseapp"
type NumberingConfigType string

const (
	NumberingConfigTypeDynamic NumberingConfigType = "dynamic"
	NumberingConfigTypeDate    NumberingConfigType = "date"
	NumberingConfigTypeNumber  NumberingConfigType = "number"
	NumberingConfigTypeCustom  NumberingConfigType = "custom"
)

func (ct NumberingConfigType) IsValid() bool {
	return ct == NumberingConfigTypeDynamic ||
		ct == NumberingConfigTypeDate ||
		ct == NumberingConfigTypeNumber ||
		ct == NumberingConfigTypeCustom
}

func (ct NumberingConfigType) String() string {
	return string(ct)
}

func (NumberingConfig) TableName() string {
	return "t_numbering_config"
}

type NumberingConfig struct {
	baseapp.BaseModel
	Type        NumberingConfigType `gorm:"type:varchar(50)" json:"type"`
	Label       string              `gorm:"type:varchar(255)" json:"label"`
	DateFormat  string              `gorm:"type:varchar(255)" json:"dateFormat"`
	DbTableName string              `gorm:"type:varchar(255)" json:"dbTableName"`
	DbFieldName string              `gorm:"type:varchar(255)" json:"dbFieldName"`
	UpLevel     int16               `gorm:"type:int" json:"upLevel"`
	DigitString int16               `gorm:"type:int;default:0" json:"digitString"`
	StartNumber int16               `gorm:"type:int" json:"startNumber"`
	MaxNumber   int16               `gorm:"type:int" json:"maxNumber"`
	DigitNumber int16               `gorm:"type:int" json:"digitNumber"`
	CustomValue string              `gorm:"type:varchar(255)" json:"customValue"`
}
