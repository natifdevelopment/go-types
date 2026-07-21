package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (CoaLoading) TableName() string {
	return "t_coa_loading"
}

type CoaLoading struct {
	baseapp.BaseModel
	TJadwalPengirimanId        *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman           *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId                 *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                    *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TPembangkitId              *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                 *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId             *uuid.UUID        `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                *MasterKapal      `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TSurveyorId                *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor                   *Organization     `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	NoSertCoa                  *string           `gorm:"type:text" json:"noSertCoa"`
	Tanggal                    *DateField  `gorm:"type:date" json:"tanggal"`
	TFileCoaId                 *uuid.UUID        `gorm:"type:uuid" json:"fileCoaId"`
	FileCoa                    *Media            `gorm:"foreignKey:TFileCoaId" json:"fileCoa"`
	SubmitUlang                bool              `gorm:"type:bool;default:false" json:"submitUlang"`
	DibutuhkanPerbaikan        bool              `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan           *string           `gorm:"text" json:"catatanPerbaikan"`
	DibutuhkanPerbaikanFileCoa bool              `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileCoa"`
	CatatanPerbaikanFileCoa    *string           `gorm:"text" json:"catatanPerbaikanFileCoa"`
	CoaItems                   []CoaLoadingItem  `gorm:"foreignKey:TCoaLoadingId; constraint:OnDelete:CASCADE" json:"coaItems"`
}

func (CoaLoadingItem) TableName() string {
	return "t_coa_loading_item"
}

type CoaLoadingItem struct {
	baseapp.BaseModel
	TCoaLoadingId  *uuid.UUID  `gorm:"type:uuid" json:"coaLoadingId"`
	CoaLoading     *CoaLoading `gorm:"foreignKey:TCoaLoadingId" json:"coaLoading"`
	TUraianId      *uuid.UUID  `gorm:"type:uuid" json:"uraianId"`
	Uraian         *Uraian     `gorm:"foreignKey:TUraianId" json:"uraian"`
	Hasil          *float32    `gorm:"type:float" json:"hasil"`
	BatasPenolakan *string     `gorm:"type:text" json:"batasPenolakan"`
	IsPassed       bool        `gorm:"type:bool;default:false" json:"isPassed"`
	Description    *string     `gorm:"type:text" json:"description"`
}
