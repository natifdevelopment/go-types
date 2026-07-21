package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PsaLoadingRoa) TableName() string {
	return "t_psa_loading_roa"
}

type PsaLoadingRoa struct {
	baseapp.BaseModel
	TJadwalPengirimanId    *uuid.UUID             `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman       *JadwalPengiriman      `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	Td                     *DateField       `gorm:"type:date" json:"td"`
	TSurveyorId            *uuid.UUID             `gorm:"type:uuid" json:"surveyorId"`
	Surveyor               *Organization          `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TMasterKapalId         *uuid.UUID             `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal            *MasterKapal           `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	NoSertCoa              *string                `gorm:"type:text" json:"noSertCoa"`
	BatasPenolakanAshBawah *float32               `gorm:"type:float" json:"batasPenolakanAshBawah"`
	BatasPenolakanCvBawah  *float32               `gorm:"type:float" json:"batasPenolakanCvBawah"`
	BatasPenolakanHgiBawah *float32               `gorm:"type:float" json:"batasPenolakanHgiBawah"`
	BatasPenolakanTmBawah  *float32               `gorm:"type:float" json:"batasPenolakanTmBawah"`
	BatasPenolakanAshAtas  *float32               `gorm:"type:float" json:"batasPenolakanAshAtas"`
	BatasPenolakanCvAtas   *float32               `gorm:"type:float" json:"batasPenolakanCvAtas"`
	BatasPenolakanHgiAtas  *float32               `gorm:"type:float" json:"batasPenolakanHgiAtas"`
	BatasPenolakanTmAtas   *float32               `gorm:"type:float" json:"batasPenolakanTmAtas"`
	Disetujui              bool                   `gorm:"type:bool" json:"disetujui"`
	Keterangan             *string                `gorm:"type:text" json:"keterangan"`
	TSurveyorPembangkitId  *uuid.UUID             `gorm:"type:uuid" json:"surveyorPembangkitId"`
	SurveyorPembangkit     *Organization          `gorm:"foreignKey:TSurveyorPembangkitId" json:"surveyorPembangkit"`
	TCoaLoadingId          *uuid.UUID             `gorm:"type:uuid" json:"coaLoadingId"`
	CoaLoading             *CoaLoading            `gorm:"foreignKey:TCoaLoadingId" json:"coaLoading"`
	NorIzinSandarBongkar   []NorIzinSandarBongkar `gorm:"foreignKey:TPsaLoadingRoaId; constraint:OnDelete:CASCADE" json:"norIzinSandarBongkar"`
	PsaLoadingRoaItems     []PsaLoadingRoaItem    `gorm:"foreignKey:TPsaLoadingRoaId; constraint:OnDelete:CASCADE" json:"coaItems"`
}

func (PsaLoadingRoaItem) TableName() string {
	return "t_psa_loading_roa_item"
}

type PsaLoadingRoaItem struct {
	baseapp.BaseModel
	TPsaLoadingRoaId *uuid.UUID     `gorm:"type:uuid" json:"psaLoadingRoaId"`
	PsaLoadingRoa    *PsaLoadingRoa `gorm:"foreignKey:TPsaLoadingRoaId" json:"psaLoadingRoa"`
	TUraianId        *uuid.UUID     `gorm:"type:uuid" json:"uraianId"`
	Uraian           *Uraian        `gorm:"foreignKey:TUraianId" json:"uraian"`
	Hasil            *float32       `gorm:"type:float" json:"hasil"`
	BatasPenolakan   *string        `gorm:"type:text" json:"batasPenolakan"`
	IsPassed         bool           `gorm:"type:bool;default:false" json:"isPassed"`
	Description      *string        `gorm:"type:text" json:"description"`
}
