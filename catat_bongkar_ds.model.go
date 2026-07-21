package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

type TipeBongkarDs string

const (
	TipeBongkarDsBypass TipeBongkarDs = "BYPASS_BONGKAR"
	TipeBongkarDsCatat  TipeBongkarDs = "CATAT_BONGKAR"
)

func (o TipeBongkarDs) IsValid() bool {
	return o == TipeBongkarDsBypass || o == TipeBongkarDsCatat
}

func (CatatBongkarDs) TableName() string {
	return "t_catat_bongkar_ds"
}

type CatatBongkarDs struct {
	baseapp.BaseModel
	TJadwalPengirimanId      *uuid.UUID           `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman         *JadwalPengiriman    `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	NoBongkar                *string              `gorm:"type:varchar(255)" json:"noBongkar"`
	TipeBongkar              TipeBongkarDs        `gorm:"type:varchar(255)" json:"tipeBongkar"`
	TMasterKapalId           *uuid.UUID           `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal              *MasterKapal         `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId        *uuid.UUID           `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang           *MasterTongkang      `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	NoticeOfTendered         *string              `gorm:"type:varchar(255)" json:"noticeOfTendered"`
	NoRid                    *string              `gorm:"type:varchar(255)" json:"noRid"`
	TJettyDermagaId          *uuid.UUID           `gorm:"type:uuid" json:"jettyDermagaId"`
	JettyDermaga             *MasterUnitJetty     `gorm:"foreignKey:TJettyDermagaId" json:"jettyDermaga"`
	RealisasiSandar          *DateTimeField `gorm:"type:timestamp" json:"realisasiSandar"`
	MulaiBongkar             *DateTimeField `gorm:"type:timestamp" json:"mulaiBongkar"`
	SelesaiBongkar           *DateTimeField `gorm:"type:timestamp" json:"selesaiBongkar"`
	TglCastOff               *DateTimeField `gorm:"type:timestamp" json:"tglCastOff"`
	VolumeBongkar            float64              `gorm:"type:float" json:"volumeBongkar"`
	GangguanPembongkaran     bool                 `gorm:"type:bool;default:false" json:"gangguanPembongkaran"`
	JumlahHariGangguan       *float64             `gorm:"type:float" json:"jumlahHariGangguan"`
	KeteranganGangguan       *string              `gorm:"type:text" json:"keteranganGangguan"`
	KesepakatanRapat         bool                 `gorm:"type:bool;default:false" json:"kesepakatanRapat"`
	TargetPembongkaran       *float64             `gorm:"type:float" json:"targetPembongkaran"`
	TDokNotulenId            *uuid.UUID           `gorm:"type:uuid" json:"dokNotulenId"`
	DokNotulen               *Media               `gorm:"foreignKey:TDokNotulenId" json:"dokNotulen"`
	NoDraughtSurvey          *string              `gorm:"type:varchar(255)" json:"noDraughtSurvey"`
	TDokDraughtSurveyId      *uuid.UUID           `gorm:"type:uuid" json:"dokDraughtSurveyId"`
	DokDraughtSurvey         *Media               `gorm:"foreignKey:TDokDraughtSurveyId" json:"dokDraughtSurvey"`
	TDokTimesheetSofId       *uuid.UUID           `gorm:"type:uuid" json:"dokTimesheetSofId"`
	DokTimesheetSof          *Media               `gorm:"foreignKey:TDokTimesheetSofId" json:"dokTimesheetSof"`
	DibutuhkanPerbaikanDokDs bool                 `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanDokDs"`
	CatatanPerbaikanDokDs    *string              `gorm:"text" json:"catatanPerbaikanDokDs"`
}
