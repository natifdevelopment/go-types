package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (ShipmentPenugasan) TableName() string {
	return "t_shipment_penugasan"
}

type ShipmentPenugasan struct {
	baseapp.BaseModel
	TPembangkitId     *uuid.UUID         `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit        *Organization      `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId        *uuid.UUID         `gorm:"type:uuid" json:"pemasokId"`
	Pemasok           *Organization      `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TTransportirId    *uuid.UUID         `gorm:"type:uuid" json:"transportirId"`
	Transportir       *Organization      `gorm:"foreignKey:TTransportirId" json:"transportir"`
	SkemaKontrakCode  *JenisPengiriman   `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak      *ConfigDataInfo    `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	VolumeConfirm     float64            `gorm:"type:float" json:"volumeConfirm"`
	TMasterKapalId    *uuid.UUID         `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal       *MasterKapal       `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId *uuid.UUID         `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang    *MasterTongkang    `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	TanggalEta        *DateField   `gorm:"type:date" json:"tanggalEta"`
	TargetLoading     *DateField   `gorm:"type:date" json:"targetLoading"`
	JadwalPengiriman  []JadwalPengiriman `gorm:"foreignKey:TShipmentPenugasanId;constraint:OnDelete:CASCADE" json:"jadwalPengiriman"`
}
