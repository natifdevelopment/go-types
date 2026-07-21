package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (ConfigData) TableName() string {
	return "t_config_data"
}

type ConfigData struct {
	baseapp.BaseModel
	TConfigDataId *uuid.UUID      `gorm:"type:uuid" json:"configDataId"`
	ConfigData    *ConfigDataInfo `gorm:"foreignKey:TConfigDataId;" json:"configData"`
	Name          string          `gorm:"type:varchar(255);" json:"name"`
	Code          string          `gorm:"type:varchar(255);unique;" json:"code"`
	ParentCode    *string         `gorm:"type:varchar(255);" json:"parentCode"`
	Description   *string         `gorm:"type:text" json:"description"`
	Color         *string         `gorm:"type:varchar(45)" json:"color"`
	Unit          *string         `gorm:"type:varchar(45)" json:"unit"`
	DataType      string          `gorm:"type:varchar(100); default:'String'" json:"dataType"` // Default text
}

func (ConfigDataInfo) TableName() string {
	return "t_config_data"
}

type ConfigDataInfo struct {
	baseapp.BaseModel
	Name        string `gorm:"type:varchar(255);" json:"name"`
	Code        string `gorm:"type:varchar(255);unique;" json:"code"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Color       string `gorm:"type:varchar(255)" json:"color"`
	Unit        string `gorm:"type:varchar(255)" json:"unit"`
	DataType    string `gorm:"default:'Text'" json:"dataType"`
}
