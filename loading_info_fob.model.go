package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (LoadingInfo) TableName() string {
	return "t_loading_info_fob"
}

type LoadingInfo struct {
	baseapp.BaseModel
	TJadwalRakorId                    *uuid.UUID                 `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor                       *JadwalRakor               `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TJadwalPengirimanId               *uuid.UUID                 `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman                  *JadwalPengiriman          `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId                     *uuid.UUID                 `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                        *Organization              `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	EtaRakor                          *DateField           `gorm:"type:date" json:"etaRakor"`
	EtaRevisi                         *DateField           `gorm:"type:date" json:"etaRevisi"`
	TaPelMuat                         *DateTimeField       `gorm:"type:timestamp" json:"taPelMuat"`
	Tendered                          *DateTimeField       `gorm:"type:timestamp" json:"tendered"`
	Accepted                          *DateTimeField       `gorm:"type:timestamp" json:"accepted"`
	Sandar                            *DateTimeField       `gorm:"type:timestamp" json:"sandar"`
	TglMuat                           *DateTimeField       `gorm:"type:timestamp" json:"tglMuat"`
	SelesaiMuat                       *DateTimeField       `gorm:"type:timestamp" json:"selesaiMuat"`
	VolumeBl                          float64                    `gorm:"type:float" json:"volumeBl"`
	TglBl                             *DateField           `gorm:"type:date" json:"tglBl"`
	TPjbbId                           *uuid.UUID                 `gorm:"type:uuid" json:"pjbbId"`
	Pjbb                              *Pjbb                      `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPelabuhanMuatId                  *uuid.UUID                 `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat                     *PelabuhanMuat             `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TModaId                           *uuid.UUID                 `gorm:"type:uuid" json:"modaId"`
	Moda                              *Moda                      `gorm:"foreignKey:TModaId" json:"moda"`
	AgenPelapor                       *string                    `gorm:"type:varchar(255)" json:"agenPelapor"`
	Td                                *DateTimeField       `gorm:"type:timestamp" json:"td"`
	TFileNorLoadingSofId              *uuid.UUID                 `gorm:"type:uuid" json:"fileNorLoadingSofId"`
	FileNorLoadingSof                 *Media                     `gorm:"foreignKey:TFileNorLoadingSofId" json:"fileNorLoadingSof"`
	DibutuhkanPerbaikanFileNorSof     bool                       `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileNorSof"`
	CatatanPerbaikanFileNorSof        *string                    `gorm:"text" json:"catatanPerbaikanFileNorSof"`
	NoManifest                        *string                    `gorm:"type:varchar(255)" json:"noManifest"`
	NoBl                              *string                    `gorm:"type:varchar(255)" json:"noBl"`
	Tpb                               *string                    `gorm:"type:varchar(255)" json:"tpb"`
	TFileManifestBlId                 *uuid.UUID                 `gorm:"type:uuid" json:"fileManifestBlId"`
	FileManifestBl                    *Media                     `gorm:"foreignKey:TFileManifestBlId" json:"fileManifestBl"`
	DibutuhkanPerbaikanFileManifestBl bool                       `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileManifestBl"`
	CatatanPerbaikanFileManifestBl    *string                    `gorm:"text" json:"catatanPerbaikanFileManifestBl"`
	Skab                              *string                    `gorm:"type:varchar(255)" json:"skab"`
	TFileManifestBlSkabId             *uuid.UUID                 `gorm:"type:uuid" json:"fileManifestBlSkabId"`                      // @deprecated
	FileManifestBlSkab                *Media                     `gorm:"foreignKey:TFileManifestBlSkabId" json:"fileManifestBlSkab"` // @deprecated
	TSurveyorId                       *uuid.UUID                 `gorm:"type:uuid" json:"surveyorId"`
	Surveyor                          *Organization              `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TApprovalRoaFobId                 *uuid.UUID                 `gorm:"type:uuid" json:"approvalRoaFobId"`
	ApprovalRoaFob                    *ApprovalRoaFob            `gorm:"foreignKey:TApprovalRoaFobId" json:"approvalRoaFob"`
	BastLoading                       []BastLoading              `gorm:"foreignKey:TLoadingInfoFobId" json:"bastLoading"`
	SumberTambang                     []LoadingInfoSumberTambang `gorm:"foreignKey:TLoadingInfoFobId" json:"sumberTambang"`
}

func (LoadingInfoSumberTambang) TableName() string {
	return "t_loading_info_sumber_tambang"
}

type LoadingInfoSumberTambang struct {
	baseapp.BaseModel
	TLoadingInfoFobId *uuid.UUID     `gorm:"type:uuid" json:"loadingInfoId"`
	LoadingInfo       *LoadingInfo   `gorm:"foreignKey:TLoadingInfoFobId" json:"loadingInfo"`
	TSumberTambangId  *uuid.UUID     `gorm:"type:uuid" json:"sumberTambangId"`
	SumberTambang     *SumberTambang `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	VolumeBl          float64        `gorm:"type:float" json:"volumeBl"`
}
