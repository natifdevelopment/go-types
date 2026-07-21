package types
import baseapp "github.com/natifdevelopment/go-baseapp"
func (JenisPembayaran) TableName() string {
	return "t_jenis_pembayaran"
}

type JenisPembayaran struct {
	baseapp.BaseModel
	JenisPembayaran string `gorm:"type:varchar(255)" json:"jenisPembayaran"`
	NamaSingkat     string `gorm:"type:varchar(255)" json:"namaSingkat"`
	Keterangan      string `gorm:"type:text" json:"keterangan"`
}
