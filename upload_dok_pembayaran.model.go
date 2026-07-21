package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
func (Upload_dok_pembayaran) TableName() string {
	return "t_upload_dok_pembayaran"
}

type Upload_dok_pembayaran struct {
	baseapp.BaseModel
	NoClearing              *string                     `gorm:"type:varchar(255)" json:"noClearing"`
	UploadDokPembayaranItem []Upload_dok_pembayaranItem `gorm:"foreignKey:TUploadDokPembayaranId" json:"items"`
}

func (Upload_dok_pembayaranItem) TableName() string {
	return "t_upload_dok_pembayaran_item"
}

type Upload_dok_pembayaranItem struct {
	baseapp.BaseModel
	TUploadDokPembayaranId  *uuid.UUID             `gorm:"type:uuid" json:"uploadDokPembayaranId"`
	UploadDokPembayaran     *Upload_dok_pembayaran `gorm:"foreignKey:TUploadDokPembayaranId" json:"uploadDokPembayaran"`
	TPermohonanPembayaranId *uuid.UUID             `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran  `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
	TDokPembayaranId        *uuid.UUID             `gorm:"type:uuid" json:"dokPembayaranId"`
	DokPembayaran           *Media                 `gorm:"foreignKey:TDokPembayaranId" json:"dokPembayaran"`
}
