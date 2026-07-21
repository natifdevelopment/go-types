package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (NorIzinSandarNongkarTrans) TableName() string {
	return "t_nor_izin_sandar_bongkar_trans"
}

type NorIzinSandarNongkarTrans struct {
	baseapp.BaseModel
	TJadwalPengirimanId            *uuid.UUID                           `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman               *JadwalPengiriman                    `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId                  *uuid.UUID                           `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                     *Organization                        `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId                 *uuid.UUID                           `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                    *MasterKapal                         `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	Td                             *DateField                     `gorm:"type:date" json:"td"`
	TglEta                         *DateField                     `gorm:"type:date" json:"tglEta"`
	TPelabuhanMuatId               *uuid.UUID                           `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat                  *PelabuhanMuat                       `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	AgenPelapor                    *string                              `gorm:"type:varchar(255)" json:"agenPelapor"`
	TaTglJam                       *DateTimeField                 `gorm:"type:timestamp" json:"taTglJam"`
	TFileNorId                     *uuid.UUID                           `gorm:"type:uuid" json:"fileNorId"`
	FileNor                        *Media                               `gorm:"foreignKey:TFileNorId" json:"fileNor"`
	NoBl                           *string                              `gorm:"type:varchar(255)" json:"noBl"`
	TglBl                          *DateField                     `gorm:"type:date" json:"tglBl"`
	NoSkab                         *string                              `gorm:"type:varchar(255)" json:"noSkab"`
	TDokBlManifestSkabId           *uuid.UUID                           `gorm:"type:uuid" json:"dokBlManifestSkabId"`
	DokBlManifestSkab              *Media                               `gorm:"foreignKey:TDokBlManifestSkabId" json:"dokBlManifestSkab"`
	NoSib                          *string                              `gorm:"type:varchar(255)" json:"noSib"`
	TglSib                         *DateField                     `gorm:"type:date" json:"tglSib"`
	TFileIzinId                    *uuid.UUID                           `gorm:"type:uuid" json:"fileIzinId"`
	FileIzin                       *Media                               `gorm:"foreignKey:TFileIzinId" json:"fileIzin"`
	TBaSerahTerimaId               *uuid.UUID                           `gorm:"type:uuid" json:"baSerahTerimaId"`
	BaSerahTerima                  *Account                             `gorm:"foreignKey:TBaSerahTerimaId" json:"baSerahTerima"`
	TPenyesuaianHargaId            *uuid.UUID                           `gorm:"type:uuid" json:"penyesuaianHargaId"`
	PenyesuaianHarga               *Account                             `gorm:"foreignKey:TPenyesuaianHargaId" json:"penyesuaianHarga"`
	TNorLoadingId                  *uuid.UUID                           `gorm:"type:uuid" json:"norLoadingId"`
	NorLoading                     *NorLoading                          `gorm:"foreignKey:TNorLoadingId" json:"norLoading"`
	NorIzinSandarBongkarPembangkit []Nor_izin_sandar_bongkar_pembangkit `gorm:"foreignKey:TNorIzinSandarBongkarTransId; constraint:OnDelete:CASCADE" json:"norIzinSandarBongkarPembangkit"`

	// Approval
	SubmitUlang         bool    `gorm:"type:bool;default:false" json:"submitUlang"`
	DibutuhkanPerbaikan bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan    *string `gorm:"text" json:"catatanPerbaikan"`

	// Pejabat
	TBastPejabatId             *uuid.UUID           `gorm:"type:uuid" json:"bastPejabatId"`
	BastPejabat                *Pejabat_surat_kuasa `gorm:"foreignKey:TBastPejabatId" json:"bastPejabat"`
	TPenyesuaianHargaPejabatId *uuid.UUID           `gorm:"type:uuid" json:"penyesuaianHargaPejabatId"`
	PenyesuaianHargaPejabat    *Pejabat_surat_kuasa `gorm:"foreignKey:TPenyesuaianHargaPejabatId" json:"penyesuaianHargaPejabat"`
}
