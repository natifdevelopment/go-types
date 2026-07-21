package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Nor_izin_sandar_bongkar_pembangkit) TableName() string {
	return "t_nor_izin_sandar_bongkar_pembangkit"
}

type Nor_izin_sandar_bongkar_pembangkit struct {
	baseapp.BaseModel
	TJadwalPengirimanId          *uuid.UUID                 `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman             *JadwalPengiriman          `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId                *uuid.UUID                 `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                   *Organization              `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId               *uuid.UUID                 `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                  *MasterKapal               `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TPemasokId                   *uuid.UUID                 `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                      *Organization              `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	Td                           *DateField           `gorm:"type:date" json:"td"`
	TPelabuhanMuatId             *uuid.UUID                 `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat                *PelabuhanMuat             `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TaTglJam                     *DateTimeField       `gorm:"type:timestamp" json:"taTglJam"`
	NoIzin                       *string                    `gorm:"type:varchar(255)" json:"noBl"`
	TglIzin                      *DateField           `gorm:"type:date" json:"tglBl"`
	Disetujui                    bool                       `gorm:"type:bool" json:"disetujui"`
	Keterangan                   *string                    `gorm:"type:text" json:"keterangan"`
	TipePenolakan                *string                    `gorm:"type:varchar(255)" json:"tipePenolakan"`
	TNorIzinSandarBongkarId      *uuid.UUID                 `gorm:"type:uuid" json:"norIzinSandarBongkarId"`
	NorIzinSandarBongkar         *NorIzinSandarBongkar      `gorm:"foreignKey:TNorIzinSandarBongkarId" json:"norIzinSandarBongkar"`
	TNorIzinSandarBongkarTransId *uuid.UUID                 `gorm:"type:uuid" json:"norIzinSandarBongkarTransId"`
	NorIzinSandarBongkarTrans    *NorIzinSandarNongkarTrans `gorm:"foreignKey:TNorIzinSandarBongkarTransId" json:"norIzinSandarBongkarTrans"`
	NoBongkar                    *string                    `gorm:"type:varchar(255)" json:"noBongkar"`
	// CatatBongkarDs               []CatatBongkarDs           `gorm:"foreignKey:TNorIzinSandarBongkarPembangkitId; constraint:OnDelete:CASCADE" json:"catatBongkarDs"`
}
