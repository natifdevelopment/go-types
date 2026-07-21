package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (Directory) TableName() string {
	return "t_directory"
}

type Directory struct {
	baseapp.BaseModel
	Name string `gorm:"type:varchar(255)" json:"name"`
	Code string `gorm:"type:varchar(255)" json:"code"`
	Page []Page `gorm:"foreignKey:TDirectoryId;constraint:OnDelete:CASCADE" json:"pages"`
}
