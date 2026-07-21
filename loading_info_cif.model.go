package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (LoadingInfoCif) TableName() string {
	return "t_loading_info_cif"
}

type LoadingInfoCif struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID           `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman    `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TJadwalRakorId      *uuid.UUID           `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor         *JadwalRakor         `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	Td                  *DateTimeField `gorm:"type:timestamp" json:"td"`
	EtaRakor            *DateField     `gorm:"type:date" json:"etaRakor"`
	EtaRevisi           *DateField     `gorm:"type:date" json:"etaRevisi"`
	MulaiLoading        *DateTimeField `gorm:"type:timestamp" json:"mulaiLoading"`
	SelesaiLoading      *DateTimeField `gorm:"type:timestamp" json:"selesaiLoading"`
	TTransportirId      *uuid.UUID           `gorm:"type:uuid" json:"transportirId"`
	Transportir         *Organization        `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TMasterKapalId      *uuid.UUID           `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal         *MasterKapal         `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId   *uuid.UUID           `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang      *MasterTongkang      `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	TPjbbId             *uuid.UUID           `gorm:"type:uuid" json:"pjbbId"`
	Pjbb                *Pjbb                `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId    *uuid.UUID           `gorm:"type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen       *PjbbAmandemen       `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TPelabuhanMuatId    *uuid.UUID           `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat       *PelabuhanMuat       `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	TSpekId             *uuid.UUID           `gorm:"type:uuid" json:"spekId"`
	Spek                *Spek                `gorm:"foreignKey:TSpekId" json:"spek"`
	TJenisBatuBaraId    *uuid.UUID           `gorm:"type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara       *JenisBatuBara       `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
	TModaId             *uuid.UUID           `gorm:"type:uuid" json:"modaId"`
	Moda                *Moda                `gorm:"foreignKey:TModaId" json:"moda"`
	VolumeBl            float64              `gorm:"type:float" json:"volumeBl"`

	// Pejabat
	TBastPejabatId             *uuid.UUID           `gorm:"type:uuid" json:"bastPejabatId"`
	BastPejabat                *Pejabat_surat_kuasa `gorm:"foreignKey:TBastPejabatId" json:"bastPejabat"`
	TPenyesuaianHargaPejabatId *uuid.UUID           `gorm:"type:uuid" json:"penyesuaianHargaPejabatId"`
	PenyesuaianHargaPejabat    *Pejabat_surat_kuasa `gorm:"foreignKey:TPenyesuaianHargaPejabatId" json:"penyesuaianHargaPejabat"`

	SumberTambang []LoadingInfoCifSumberTambang `gorm:"foreignKey:TLoadingInfoCifId; constraint:OnDelete:CASCADE" json:"sumberTambang"`
}

func (LoadingInfoCifSumberTambang) TableName() string {
	return "t_loading_info_cif_sumber_tambang"
}

type LoadingInfoCifSumberTambang struct {
	baseapp.BaseModel
	TLoadingInfoCifId *uuid.UUID      `gorm:"type:uuid" json:"loadingInfoCifId"`
	LoadingInfoCif    *LoadingInfoCif `gorm:"foreignKey:TLoadingInfoCifId" json:"loadingInfoCif"`
	TSumberTambangId  *uuid.UUID      `gorm:"type:uuid" json:"sumberTambangId"`
	SumberTambang     *SumberTambang  `gorm:"foreignKey:TSumberTambangId" json:"sumberTambang"`
	VolumeBl          float64         `gorm:"type:float" json:"volumeBl"`
}
