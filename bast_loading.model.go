package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (BastLoading) TableName() string {
	return "t_bast_laoding"
}

type BastLoading struct {
	baseapp.BaseModel
	NoBast                     string            `gorm:"type:varchar(255)" json:"noBast"`
	TJadwalPengirimanId        *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman           *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPemasokId                 *uuid.UUID        `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                    *Organization     `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TMasterKapalId             *uuid.UUID        `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal                *MasterKapal      `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TPelabuhanMuatId           *uuid.UUID        `gorm:"type:uuid" json:"pelabuhanMuatId"`
	PelabuhanMuat              *PelabuhanMuat    `gorm:"foreignKey:TPelabuhanMuatId" json:"pelabuhanMuat"`
	VolumeBl                   *float32          `gorm:"type:float" json:"volumeBl"`
	TanggalBl                  *DateField  `gorm:"type:date" json:"tanggalBl"`
	TglBaLoading               *DateField  `gorm:"type:date" json:"tglBaLoading"`
	NamaPihak1                 *string           `gorm:"type:varchar(255)" json:"namaPihak1"`
	JabatanPihak1              *string           `gorm:"type:varchar(255)" json:"jabatanPihak1"`
	NamaPihak2                 *string           `gorm:"type:varchar(255)" json:"namaPihak2"`
	JabatanPihak2              *string           `gorm:"type:varchar(255)" json:"jabatanPihak2"`
	TLoadingInfoFobId          *uuid.UUID        `gorm:"type:uuid" json:"loadingInfoId"`
	LoadingInfo                *LoadingInfo      `gorm:"foreignKey:TLoadingInfoFobId" json:"loadingInfo"`
	NoDokBastLoading           *string           `gorm:"type:varchar(255)" json:"noDokBastLoading"`
	TDokBastLoadingId          *uuid.UUID        `gorm:"type:uuid" json:"dokBastLoadingId"`
	DokBastLoading             *Media            `gorm:"foreignKey:TDokBastLoadingId" json:"dokBastLoading"`
	BastDiperiksa              *bool             `gorm:"type:bool" json:"bastDiperiksa"`
	KeteranganPeriksaBast      *string           `gorm:"type:text" json:"keteranganPeriksaBast"`
	NoDokBaDendaMuat           *string           `gorm:"type:varchar(255)" json:"noDokBaDendaMuat"`
	TDokBaDendaMuatId          *uuid.UUID        `gorm:"type:uuid" json:"dokBaDendaMuatId"`
	DokBaDendaMuat             *Media            `gorm:"foreignKey:TDokBaDendaMuatId" json:"dokBaDendaMuat"`
	TDokRincianBaDendaMuatId   *uuid.UUID        `gorm:"type:uuid" json:"dokRincianBaDendaMuatId"`
	DokRincianBaDendaMuat      *Media            `gorm:"foreignKey:TDokRincianBaDendaMuatId" json:"dokRincianBaDendaMuat"`
	DendaMuatDiperiksa         *bool             `gorm:"type:bool" json:"dendaMuatDiperiksa"`
	KeteranganPeriksaDendaMuat *string           `gorm:"type:text" json:"keteranganPeriksaDendaMuat"`
	DokumenDibuat              bool              `gorm:"type:bool;default:false" json:"dokumenDibuat"`
}
