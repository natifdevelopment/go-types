package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (NorIzinSandarBongkar) TableName() string {
	return "t_nor_izin_sandar_bongkar"
}

type NorIzinSandarBongkar struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID           `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman    `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId       *uuid.UUID           `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization        `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId      *uuid.UUID           `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal         *MasterKapal         `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	Td                  *DateField     `gorm:"type:date" json:"td"`
	TglEta              *DateField     `gorm:"type:date" json:"tglEta"`
	TPelabuhanMuatId    *uuid.UUID           `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat       *PelabuhanMuat       `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	AgenPelapor         *string              `gorm:"type:varchar(255)" json:"agenPelapor"`
	TaTglJam            *DateTimeField `gorm:"type:timestamp" json:"taTglJam"`
	TFileNorId          *uuid.UUID           `gorm:"type:uuid" json:"fileNorId"`
	FileNor             *Media               `gorm:"foreignKey:TFileNorId" json:"fileNor"`
	NoBl                *string              `gorm:"type:varchar(255)" json:"noBl"`
	TglBl               *DateField     `gorm:"type:date" json:"tglBl"`
	TDokBlId            *uuid.UUID           `gorm:"type:uuid" json:"dokBlId"`
	DokBl               *Media               `gorm:"foreignKey:TDokBlId" json:"dokBl"`

	// @deprecated
	// now use seprate dok bl and skab (skab based on sumber tambang loading info)
	// no skab moved to .sumberTambang
	NoSkab               *string    `gorm:"type:varchar(255)" json:"noSkab"`
	TDokBlManifestSkabId *uuid.UUID `gorm:"type:uuid" json:"dokBlManifestSkabId"`
	DokBlManifestSkab    *Media     `gorm:"foreignKey:TDokBlManifestSkabId" json:"dokBlManifestSkab"`

	NoSib                                *string          `gorm:"type:varchar(255)" json:"noSib"`
	TglSib                               *DateField `gorm:"type:date" json:"tglSib"`
	TFileIzinId                          *uuid.UUID       `gorm:"type:uuid" json:"fileIzinId"`
	FileIzin                             *Media           `gorm:"foreignKey:TFileIzinId" json:"fileIzin"`
	TCoaLoadingId                        *uuid.UUID       `gorm:"type:uuid" json:"coaLoadingId"`
	CoaLoading                           *CoaLoading      `gorm:"foreignKey:TCoaLoadingId" json:"coaLoading"`
	TPsaLoadingRoaId                     *uuid.UUID       `gorm:"type:uuid" json:"psaLoadingRoaId"`
	PsaLoadingRoa                        *PsaLoadingRoa   `gorm:"foreignKey:TPsaLoadingRoaId" json:"psaLoadingRoa"`
	SubmitUlang                          bool             `gorm:"type:bool;default:false" json:"submitUlang"`
	DibutuhkanPerbaikan                  bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan                     *string          `gorm:"text" json:"catatanPerbaikan"`
	DibutuhkanPerbaikanFileNor           bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileNor"`
	CatatanPerbaikanFileNor              *string          `gorm:"text" json:"catatanPerbaikanFileNor"`
	DibutuhkanPerbaikanDokBlManifestSkab bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanDokBlManifestSkab"`
	CatatanPerbaikanDokBlManifestSkab    *string          `gorm:"text" json:"catatanPerbaikanDokBlManifestSkab"`
	DibutuhkanPerbaikanFileIzin          bool             `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileIzin"`
	CatatanPerbaikanFileIzin             *string          `gorm:"text" json:"catatanPerbaikanFileIzin"`
	TNorLoadingId                        *uuid.UUID       `gorm:"type:uuid" json:"norLoadingId"`
	NorLoading                           *NorLoading      `gorm:"foreignKey:TNorLoadingId" json:"norLoading"`
	NoBongkar                            *string          `gorm:"type:varchar(255)" json:"noBongkar"`

	// BAST Pejabat
	TBastPejabatId             *uuid.UUID           `gorm:"type:uuid" json:"bastPejabatId"`
	BastPejabat                *Pejabat_surat_kuasa `gorm:"foreignKey:TBastPejabatId" json:"bastPejabat"`
	TPenyesuaianHargaPejabatId *uuid.UUID           `gorm:"type:uuid" json:"penyesuaianHargaPejabatId"`
	PenyesuaianHargaPejabat    *Pejabat_surat_kuasa `gorm:"foreignKey:TPenyesuaianHargaPejabatId" json:"penyesuaianHargaPejabat"`

	NorIzinSandarBongkarPembangkit []Nor_izin_sandar_bongkar_pembangkit `gorm:"foreignKey:TNorIzinSandarBongkarId; constraint:OnDelete:CASCADE" json:"norIzinSandarBongkarPembangkit"`
	SumberTambang                  []NorIzinSandarBongkarSumberTambang  `gorm:"foreignKey:TNorIzinSandarBongkarId; constraint:OnDelete:CASCADE" json:"sumberTambang"`
}

func (NorIzinSandarBongkarSumberTambang) TableName() string {
	return "t_nor_izin_sandar_bongkar_sumber_tambang"
}

type NorIzinSandarBongkarSumberTambang struct {
	baseapp.BaseModel
	TNorIzinSandarBongkarId *uuid.UUID            `gorm:"type:uuid" json:"norIzinSandarBongkarId"`
	NorIzinSandarBongkar    *NorIzinSandarBongkar `gorm:"foreignKey:TNorIzinSandarBongkarId" json:"norIzinSandarBongkar"`
	AdaDokSkab              bool                  `gorm:"type:bool" json:"adaDokSkab"`
	TSumberTambangId        *uuid.UUID            `gorm:"type:uuid" json:"sumberTambangId"`
	SumberTambang           *SumberTambang        `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	TDokSkabId              *uuid.UUID            `gorm:"type:uuid" json:"dokSkabId"`
	DokSkab                 *Media                `gorm:"foreignKey:TDokSkabId" json:"dokSkab"`
	NoSkab                  *string               `gorm:"type:varchar(255)" json:"noSkab"`
}
