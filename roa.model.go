package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (Roa) TableName() string {
	return "t_roa"
}

type Roa struct {
	baseapp.BaseModel
	TJadwalPengirimanId        *uuid.UUID           `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman           *JadwalPengiriman    `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TSurveyorId                *uuid.UUID           `gorm:"type:uuid" json:"surveyorId"`
	Surveyor                   *Organization        `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TSurveyorLoadingId         *uuid.UUID           `gorm:"type:uuid" json:"surveyorLoadingId"`
	SurveyorLoading            *Organization        `gorm:"foreignKey:TSurveyorLoadingId" json:"surveyorLoading"`
	NoSertRoa                  *string              `gorm:"type:varchar(255)" json:"noSertRoa"`
	TanggalSertRoa             *DateField     `gorm:"type:date" json:"tanggalSertRoa"`
	TPjbbId                    *uuid.UUID           `gorm:"type:uuid" json:"pjbbId"`
	Pjbb                       *Pjbb                `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId           *uuid.UUID           `gorm:"type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen              *PjbbAmandemen       `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TPelabuhanMuatId           *uuid.UUID           `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat              *PelabuhanMuat       `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId                    *uuid.UUID           `gorm:"type:uuid" json:"spekId"`
	Spek                       *Spek                `gorm:"foreignKey:TSpekId" json:"spek"`
	TModaId                    *uuid.UUID           `gorm:"type:uuid" json:"modaId"`
	Moda                       *Moda                `gorm:"foreignKey:TModaId" json:"moda"`
	TTransportirId             *uuid.UUID           `gorm:"type:uuid" json:"transportirId"`
	Transportir                *Organization        `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TMasterKapalId             *uuid.UUID           `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                *MasterKapal         `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId          *uuid.UUID           `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang             *MasterTongkang      `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	TBastPejabatId             *uuid.UUID           `gorm:"type:uuid" json:"bastPejabatId"`
	BastPejabat                *Pejabat_surat_kuasa `gorm:"foreignKey:TBastPejabatId" json:"bastPejabat"`
	TPenyesuaianHargaPejabatId *uuid.UUID           `gorm:"type:uuid" json:"penyesuaianHargaPejabatId"`
	PenyesuaianHargaPejabat    *Pejabat_surat_kuasa `gorm:"foreignKey:TPenyesuaianHargaPejabatId" json:"penyesuaianHargaPejabat"`
	TFileRoaId                 *uuid.UUID           `gorm:"type:uuid" json:"fileRoaId"`
	FileRoa                    *Media               `gorm:"foreignKey:TFileRoaId" json:"fileRoa"`
	DibutuhkanPerbaikanFileRoa bool                 `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileRoa"`
	CatatanPerbaikanFileRoa    *string              `gorm:"text" json:"catatanPerbaikanFileRoa"`
	SubmitUlang                bool                 `gorm:"type:bool;default:false" json:"submitUlang"`
	DibutuhkanPerbaikan        bool                 `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan           *string              `gorm:"text" json:"catatanPerbaikan"`
	RoaItems                   []RoaItem            `gorm:"foreignKey:TRoaId; constraint:OnDelete:CASCADE" json:"roaItems"`
	SumberTambang              []RoaSumberTambang   `gorm:"foreignKey:TRoaId" json:"sumberTambang"`
}

func (RoaItem) TableName() string {
	return "t_roa_item"
}

type RoaItem struct {
	baseapp.BaseModel
	TRoaId         *uuid.UUID `gorm:"type:uuid" json:"roaId"`
	Roa            *Roa       `gorm:"foreignKey:TRoaId" json:"roa"`
	TUraianId      *uuid.UUID `gorm:"type:uuid" json:"uraianId"`
	Uraian         *Uraian    `gorm:"foreignKey:TUraianId" json:"uraian"`
	Hasil          *float32   `gorm:"type:float" json:"hasil"`
	BatasPenolakan *string    `gorm:"type:text" json:"batasPenolakan"`
	IsPassed       bool       `gorm:"type:bool;default:false" json:"isPassed"`
	Description    *string    `gorm:"type:text" json:"description"`
}

func (RoaSumberTambang) TableName() string {
	return "t_roa_sumber_tambang"
}

type RoaSumberTambang struct {
	baseapp.BaseModel
	TRoaId           *uuid.UUID     `gorm:"index;type:uuid" json:"roaId"`
	Roa              *Roa           `gorm:"foreignKey:TRoaId" json:"roa"`
	AdaDokSkab       bool           `gorm:"type:bool" json:"adaDokSkab"`
	TSumberTambangId *uuid.UUID     `gorm:"index;type:uuid" json:"sumberTambangId"`
	SumberTambang    *SumberTambang `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	TDokSkabId       *uuid.UUID     `gorm:"index;type:uuid" json:"dokSkabId"`
	DokSkab          *Media         `gorm:"foreignKey:TDokSkabId" json:"dokSkab"`
	NoSkab           *string        `gorm:"type:varchar(255)" json:"noSkab"`
}
