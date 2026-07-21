package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PelunasanPenagihan) TableName() string {
	return "t_pelunasan_penagihan"
}

type PelunasanPenagihan struct {
	baseapp.BaseModel
	NoClearing       *string                  `gorm:"type:varchar(255)" json:"noClearing"`
	TglPembayaran    *DateField         `gorm:"type:date" json:"tglPembayaran"`
	TglPosting       *DateField         `gorm:"type:date" json:"tglPosting"`
	SapInvCc         *string                  `gorm:"type:varchar(255)" json:"sapInvCc"`
	TDokPembayaranId *uuid.UUID               `gorm:"type:uuid" json:"dokPembayaranId"`
	DokPembayaran    *Media                   `gorm:"foreignKey:TDokPembayaranId" json:"dokPembayaran"`
	Items            []PelunasanPenagihanItem `gorm:"foreignKey:TPelunasanPenagihanId" json:"items"`
	// DokPembayaran []PelunasanPenagihan_dok_pembayaran `gorm:"foreignKey:TPelunasanPenagihanId" json:"dokPembayaran"`
}

func (PelunasanPenagihanItem) TableName() string {
	return "t_pelunasan_penagihan_item"
}

type PelunasanPenagihanItem struct {
	baseapp.BaseModel
	TPelunasanPenagihanId *uuid.UUID          `gorm:"type:uuid" json:"pelunasanPenagihanId"`
	PelunasanPenagihan    *PelunasanPenagihan `gorm:"foreignKey:TPelunasanPenagihanId;constraint:OnDelete:CASCADE" json:"pelunasanPenagihan"`
	TProposePerhitunganId *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`

	// @deprecated
	// Bisa set lunas per invoice harusnya
	TPermohonanPembayaranId *uuid.UUID            `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
}

func (PelunasanPenagihan_dok_pembayaran) TableName() string {
	return "t_pelunasan_penagihan_dok_pembayaran"
}

type PelunasanPenagihan_dok_pembayaran struct {
	baseapp.BaseModel
	TPelunasanPenagihanId *uuid.UUID          `gorm:"type:uuid" json:"pelunasanPenagihanId"`
	PelunasanPenagihan    *PelunasanPenagihan `gorm:"foreignKey:TPelunasanPenagihanId;constraint:OnDelete:CASCADE" json:"pelunasanPenagihan"`
	TDokPembayaranId      *uuid.UUID          `gorm:"type:uuid" json:"dokPembayaranId"`
	DokPembayaran         *Media              `gorm:"foreignKey:TDokPembayaranId" json:"dokPembayaran"`
	TProposePerhitunganId *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`

	// @deprecated
	// Bisa set lunas per invoice harusnya
	TPermohonanPembayaranId *uuid.UUID            `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
}
