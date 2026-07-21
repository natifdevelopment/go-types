package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Umpire) TableName() string {
	return "t_umpire"
}

type Umpire struct {
	baseapp.BaseModel
	TPengajuanUmpireId  *uuid.UUID        `gorm:"type:uuid" json:"pengajuanUmpireId"`
	PengajuanUmpire     *PengajuanUmpire  `gorm:"foreignKey:TPengajuanUmpireId" json:"pengajuanUmpire"`
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId          *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok             *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TMasterKapalId      *uuid.UUID        `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal         *MasterKapal      `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TSurveyorId         *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor            *Organization     `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	NoSertifikatCoa     *string           `gorm:"type:varchar(255)" json:"noSertifikatCoa"`
	// TCoaCowCfrId         *uuid.UUID        `gorm:"type:uuid" json:"coaCowCfrId"`
	// CoaCowCfr            *Coa_cow_CFR      `gorm:"foreignKey:TCoaCowCfrId" json:"coaCowCfr"`
	TCoaCowCifId         *uuid.UUID       `gorm:"type:uuid" json:"coaCowCifId"`
	CoaCowCif            *CoaCow          `gorm:"foreignKey:TCoaCowCifId" json:"coaCowCif"`
	TCoaCowFobId         *uuid.UUID       `gorm:"type:uuid" json:"coaCowFobId"`
	CoaCowFob            *CoaCowFob       `gorm:"foreignKey:TCoaCowFobId" json:"coaCowFob"`
	TCoaCowTransId       *uuid.UUID       `gorm:"type:uuid" json:"coaCowTransId"`
	CoaCowTrans          *CoaCowTrans     `gorm:"foreignKey:TCoaCowTransId" json:"coaCowTrans"`
	VolumeDs             string           `gorm:"type:varchar(255)" json:"volumeDs"`
	TSurveyorUmpireId    *uuid.UUID       `gorm:"type:uuid" json:"surveyorUmpireId"`
	SurveyorUmpire       *Organization    `gorm:"foreignKey:TSurveyorUmpireId" json:"surveyorUmpire"`
	TglPermohonanUmpire  *DateField `gorm:"type:date" json:"tglPermohonanUmpire"`
	TempatUmpire         *string          `gorm:"type:varchar(255)" json:"tempatUmpire"`
	NoCoaUmpire          string           `gorm:"type:varchar(255)" json:"noCoaUmpire"`
	TanggalSelesaiUmpire DateField  `gorm:"type:date" json:"tanggalSelesaiUmpire"`
	NoBaUmpire           string           `gorm:"type:varchar(255)" json:"noBaUmpire"`
	TanggalBaUmpire      DateField  `gorm:"type:date" json:"tanggalBaUmpire"`
	TFileCoaUmpireId     *uuid.UUID       `gorm:"type:uuid" json:"fileCoaUmpireId"`
	FileCoaUmpire        *Media           `gorm:"foreignKey:TFileCoaUmpireId" json:"fileCoaUmpire"`
	TFileBaUmpireId      *uuid.UUID       `gorm:"type:uuid" json:"fileBaUmpireId"`
	FileBaUmpire         *Media           `gorm:"foreignKey:TFileBaUmpireId" json:"fileBaUmpire"`
	UmpireUraian         []UmpireUraian   `gorm:"foreignKey:TUmpireId; constraint:OnDelete:CASCADE" json:"uraianItems"`
}

func (UmpireUraian) TableName() string {
	return "t_umpire_uraian"
}

type UmpireUraian struct {
	baseapp.BaseModel
	TUmpireId *uuid.UUID `gorm:"type:uuid" json:"umpireId"`
	Umpire    *Umpire    `gorm:"foreignKey:TUmpireId" json:"umpire"`
	// TUraianId *uuid.UUID `gorm:"type:uuid" json:"uraianId"`
	// Uraian    *Uraian    `gorm:"foreignKey:TUraianId" json:"uraian"`
	TCoaCowItemId uuid.UUID     `gorm:"type:uuid" json:"coaCowItemId"`
	CoaCowItem    CoaCowCifItem `gorm:"foreignKey:TCoaCowItemId" json:"coaCowItem"`
}
