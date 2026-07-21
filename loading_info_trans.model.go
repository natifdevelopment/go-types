package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (LoadingInfoTrans) TableName() string {
	return "t_loading_info_trans"
}

type LoadingInfoTrans struct {
	baseapp.BaseModel
	TJadwalRakorId                    *uuid.UUID                      `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor                       *JadwalRakor                    `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TJadwalPengirimanId               *uuid.UUID                      `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman                  *JadwalPengiriman               `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId                     *uuid.UUID                      `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                        *Organization                   `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TMasterKapalId                    *uuid.UUID                      `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                       *MasterKapal                    `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId                 *uuid.UUID                      `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang                    *MasterTongkang                 `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	TSpekId                           *uuid.UUID                      `gorm:"type:uuid" json:"spekId"`
	Spek                              *Spek                           `gorm:"foreignKey:TSpekId" json:"spek"`
	EtaRakor                          *DateField                `gorm:"type:date" json:"etaRakor"`
	EtaRevisi                         *DateField                `gorm:"type:date" json:"etaRevisi"`
	TaPelMuat                         *DateTimeField            `gorm:"type:timestamp" json:"taPelMuat"`
	Tendered                          *DateTimeField            `gorm:"type:timestamp" json:"tendered"`
	Accepted                          *DateTimeField            `gorm:"type:timestamp" json:"accepted"`
	Sandar                            *DateTimeField            `gorm:"type:timestamp" json:"sandar"`
	TglMuat                           *DateTimeField            `gorm:"type:timestamp" json:"tglMuat"`
	SelesaiMuat                       *DateTimeField            `gorm:"type:timestamp" json:"selesaiMuat"`
	VolumeBl                          *float64                        `gorm:"type:float" json:"volumeBl"`
	TglBl                             *DateField                `gorm:"type:date" json:"tglBl"`
	TPjabId                           *uuid.UUID                      `gorm:"type:uuid" json:"pjabId"`
	Pjab                              *Pjab                           `gorm:"foreignKey:TPjabId" json:"pjab"`
	TPjabAmandemenId                  *uuid.UUID                      `gorm:"type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen                     *PjabAmandemen                  `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	TPelabuhanMuatId                  *uuid.UUID                      `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat                     *PelabuhanMuat                  `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TModaId                           *uuid.UUID                      `gorm:"type:uuid" json:"modaId"`
	Moda                              *Moda                           `gorm:"foreignKey:TModaId" json:"moda"`
	AgenPelapor                       *string                         `gorm:"type:varchar(255)" json:"agenPelapor"`
	Td                                *DateTimeField            `gorm:"type:timestamp" json:"td"`
	TFileNorLoadingSofId              *uuid.UUID                      `gorm:"type:uuid" json:"fileNorLoadingSofId"`
	FileNorLoadingSof                 *Media                          `gorm:"foreignKey:TFileNorLoadingSofId" json:"fileNorLoadingSof"`
	DibutuhkanPerbaikanFileNorSof     bool                            `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileNorSof"`
	CatatanPerbaikanFileNorSof        *string                         `gorm:"text" json:"catatanPerbaikanFileNorSof"`
	NoManifest                        *string                         `gorm:"type:varchar(255)" json:"noManifest"`
	NoBl                              *string                         `gorm:"type:varchar(255)" json:"noBl"`
	TFileManifestBlId                 *uuid.UUID                      `gorm:"type:uuid" json:"fileManifestBlId"`
	FileManifestBl                    *Media                          `gorm:"foreignKey:TFileManifestBlId" json:"fileManifestBl"`
	DibutuhkanPerbaikanFileManifestBl bool                            `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileManifestBl"`
	CatatanPerbaikanFileManifestBl    *string                         `gorm:"text" json:"catatanPerbaikanFileManifestBl"`
	Tpb                               *string                         `gorm:"type:varchar(255)" json:"tpb"`
	TFileTpbId                        *uuid.UUID                      `gorm:"type:uuid" json:"fileTpbId"`
	FileTpb                           *Media                          `gorm:"foreignKey:TFileTpbId" json:"fileTpb"`
	DibutuhkanPerbaikanFileTpb        bool                            `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikanFileTpb"`
	CatatanPerbaikanFileTpb           *string                         `gorm:"text" json:"catatanPerbaikanFileTpb"`
	Skab                              *string                         `gorm:"type:varchar(255)" json:"skab"`
	TFileManifestBlSkabId             *uuid.UUID                      `gorm:"type:uuid" json:"fileManifestBlSkabId"` // @deprecated
	FileManifestBlSkab                *Media                          `gorm:"foreignKey:TFileManifestBlSkabId" json:"fileManifestBlSkab"`
	TSurveyorId                       *uuid.UUID                      `gorm:"type:uuid" json:"surveyorId"`
	Surveyor                          *Organization                   `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	TApprovalRoaFobId                 *uuid.UUID                      `gorm:"type:uuid" json:"approvalRoaFobId"`
	ApprovalRoaFob                    *ApprovalRoaFob                 `gorm:"foreignKey:TApprovalRoaFobId" json:"approvalRoaFob"`
	NorLoading                        []NorLoading                    `gorm:"foreignKey:TLoadingInfoTransId" json:"norLoading"`
	SumberTambang                     []LoadingInfoTransSumberTambang `gorm:"foreignKey:TLoadingInfoTransId" json:"sumberTambang"`

	// Approval
	SubmitUlang         bool    `gorm:"type:bool;default:false" json:"submitUlang"`
	DibutuhkanPerbaikan bool    `gorm:"type:bool;default:false" json:"dibutuhkanPerbaikan"`
	CatatanPerbaikan    *string `gorm:"text" json:"catatanPerbaikan"`
}

func (LoadingInfoTransSumberTambang) TableName() string {
	return "t_loading_info_trans_sumber_tambang"
}

type LoadingInfoTransSumberTambang struct {
	baseapp.BaseModel
	TLoadingInfoTransId *uuid.UUID        `gorm:"type:uuid" json:"loadingInfoTransId"`
	LoadingInfoTrans    *LoadingInfoTrans `gorm:"foreignKey:TLoadingInfoTransId" json:"loadingInfoTrans"`
	TSumberTambangId    *uuid.UUID        `gorm:"type:uuid" json:"sumberTambangId"`
	SumberTambang       *SumberTambang    `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
}
