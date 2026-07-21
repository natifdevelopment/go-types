package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"

	"github.com/google/uuid"
)

func (PipPenerimaan) TableName() string {
	return "t_pip_penerimaan"
}

type PipPenerimaan struct {
	baseapp.BaseModel
	OperatingUnit          *string           `gorm:"type:varchar(10)" json:"operatingUnit"`
	NoKontrakAdendum       *string           `gorm:"type:varchar(50)" json:"noKontrakAdendum"`
	PoNumber               *string           `gorm:"type:varchar(20)" json:"poNumber"`
	RevisionPoNumber       *string           `gorm:"type:varchar(10)" json:"revisionPoNumber"`
	SupplierID             *string           `gorm:"type:varchar(20)" json:"supplierId"`
	Supplier               *string           `gorm:"type:varchar(255)" json:"supplier"`
	TransactionID          *string           `gorm:"type:varchar(30)" json:"transactionId"`
	TransactionDate        *time.Time        `gorm:"type:date" json:"transactionDate"`
	ReceivedBy             *string           `gorm:"type:varchar(100)" json:"receivedBy"`
	ItemNumber             *string           `gorm:"type:varchar(50)" json:"itemNumber"`
	Satuan                 *string           `gorm:"type:varchar(10)" json:"satuan"`
	Quantity               *float64          `gorm:"type:float" json:"quantity"`
	ItemCost               *float64          `gorm:"type:float" json:"itemCost"`
	AmountLine             *float64          `gorm:"type:float" json:"amountLine"`
	NamaKapal              *string           `gorm:"type:varchar(150)" json:"namaKapal"`
	TanggalPenerimaan      *time.Time        `gorm:"type:timestamp" json:"tanggalPenerimaan"`
	NomorJadwalPengiriman  *string           `gorm:"type:varchar(50)" json:"nomorJadwalPengiriman"`
	NomorPengiriman        *string           `gorm:"type:varchar(50)" json:"nomorPengiriman"`
	Storeroom              *string           `gorm:"type:varchar(50)" json:"storeroom"`
	TrxID                  *string           `gorm:"type:varchar(20)" json:"trxId"`
	Jetty                  *string           `gorm:"type:varchar(255)" json:"jetty"`
	TanggalMulaiBongkar    *time.Time        `gorm:"type:timestamp" json:"tanggalMulaiBongkar"`
	TanggalRealisasiSandar *time.Time        `gorm:"type:timestamp" json:"tanggalRealisasiSandar"`
	ResponseHash           *string           `gorm:"type:varchar(255)" json:"-"`
	TJadwalPengirimanId    *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman       *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
}
