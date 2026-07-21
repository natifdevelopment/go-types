package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

type ApprovalHargaCfr struct {
	baseapp.BaseModel
	NoPengiriman          string          `gorm:"type:varchar(255)" json:"noPengiriman"`
	NoBast                string          `gorm:"type:varchar(255)" json:"noBast"`
	TglBast               DateField `gorm:"type:date" json:"tglBast"`
	DendaKeterlambatan    uint64          `gorm:"type:bigint" json:"dendaKeterlambatan"`
	TotalTagihan          uint64          `gorm:"type:bigint" json:"totalTagihan"`
	MenyetujuiHitungHarga bool            `gorm:"type:boolean" json:"menyetujuiHitungHarga"`
	PejabatBaph           string          `gorm:"type:varchar(255)" json:"pejabatBaph"`
}

func (ApprovalHargaCfr) TableName() string {
	return "t_approval_harga_cfr"
}
