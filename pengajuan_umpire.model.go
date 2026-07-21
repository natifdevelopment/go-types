package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PengajuanUmpire) TableName() string {
	return "t_pengajuan_umpire"
}

type PengajuanUmpire struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId      *uuid.UUID        `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal         *MasterKapal      `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TSurveyorId         *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor            *Organization     `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	NoSertifikatCoa     *string           `gorm:"type:varchar(255)" json:"noSertifikatCoa"`
	// TCoaCowCfrId          *uuid.UUID              `gorm:"type:uuid" json:"coaCowCfrId"`
	// CoaCowCfr             *Coa_cow_CFR            `gorm:"foreignKey:TCoaCowCfrId" json:"coaCowCfr"`
	TCoaCowCifId          *uuid.UUID              `gorm:"type:uuid" json:"coaCowCifId"`
	CoaCowCif             *CoaCow                 `gorm:"foreignKey:TCoaCowCifId" json:"coaCowCif"`
	TCoaCowFobId          *uuid.UUID              `gorm:"type:uuid" json:"coaCowFobId"`
	CoaCowFob             *CoaCowFob              `gorm:"foreignKey:TCoaCowFobId" json:"coaCowFob"`
	TCoaCowTransId        *uuid.UUID              `gorm:"type:uuid" json:"coaCowTransId"`
	CoaCowTrans           *CoaCowTrans            `gorm:"foreignKey:TCoaCowTransId" json:"coaCowTrans"`
	VolumeDs              *float64                `gorm:"type:float" json:"volumeDs"`
	TSurveyorUmpireId     *uuid.UUID              `gorm:"type:uuid" json:"surveyorUmpireId"`
	SurveyorUmpire        *Organization           `gorm:"foreignKey:TSurveyorUmpireId" json:"surveyorUmpire"`
	TglPermohonanUmpire   *DateField        `gorm:"type:date" json:"tglPermohonanUmpire"`
	TempatUmpire          *string                 `gorm:"type:varchar(255)" json:"tempatUmpire"`
	PengajuanUmpireUraian []PengajuanUmpireUraian `gorm:"foreignKey:TPengajuanUmpireId; constraint:OnDelete:CASCADE" json:"uraianItems"`
	Umpire                *Umpire                 `gorm:"foreignKey:TPengajuanUmpireId" json:"umpire"`
}

func (PengajuanUmpireUraian) TableName() string {
	return "t_pengajuan_umpire_uraian"
}

type PengajuanUmpireUraian struct {
	baseapp.BaseModel
	TPengajuanUmpireId *uuid.UUID       `gorm:"type:uuid" json:"pengajuanUmpireId"`
	PengajuanUmpire    *PengajuanUmpire `gorm:"foreignKey:TPengajuanUmpireId" json:"pengajuanUmpire"`
	TCoaCowItemId      uuid.UUID        `gorm:"type:uuid" json:"coaCowItemId"`
	CoaCowItem         CoaCowCifItem    `gorm:"foreignKey:TCoaCowItemId" json:"coaCowItem"`
}
