package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

type ApprovalHargaTrans struct {
	baseapp.BaseModel
	NoPengiriman          string          `gorm:"type:varchar(255)" json:"noPengiriman"`
	NoBABongkar           string          `gorm:"type:varchar(255)" json:"noBaBongkar"`
	TglBABongkar          DateField `gorm:"type:date" json:"tglBaBongkar"`
	TotalTagihan          uint64          `gorm:"type:bigint" json:"totalTagihan"`
	MenyetujuiHitungHarga bool            `gorm:"type:bool" json:"menyetujuiHitungHarga"`
	PejabatBAPH           string          `gorm:"type:varchar(255)" json:"pejabatBaph"`
}

func (ApprovalHargaTrans) TableName() string {
	return "t_approval_harga_TRANS"
}
