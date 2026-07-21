package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PenagihanAkuntansi) TableName() string {
	return "t_penagihan_akuntansi"
}

type PenagihanAkuntansi struct {
	baseapp.BaseModel
	TPermohonanPembayaranId *uuid.UUID               `gorm:"type:uuid" json:"permohonanPembayaranId"`
	PermohonanPembayaran    *PermohonanPembayaran    `gorm:"foreignKey:TPermohonanPembayaranId" json:"permohonanPembayaran"`
	Dpp                     *float32                 `gorm:"type:float" json:"dpp"`
	Denda                   *float32                 `gorm:"type:float" json:"denda"`
	PphPasalCode            *string                  `gorm:"type:varchar(255)" json:"pphPasalCode"`
	PphPasal                *ConfigDataInfo          `gorm:"foreignKey:PphPasalCode;references:Code;" json:"pphPasal"`
	NilaiPph                *float32                 `gorm:"type:float" json:"nilaiPph"`
	TotalDibayarkan         *float32                 `gorm:"type:float" json:"totalDibayarkan"`
	TglTerimaTagihan        *DateField         `gorm:"type:date" json:"tglTerimaTagihan"`
	TglJatuhTempo           *DateField         `gorm:"type:date" json:"tglJatuhTempo"`
	NoLembarChecklist       *string                  `gorm:"type:text" json:"noLembarChecklist"`
	TglLembarChecklist      *DateField         `gorm:"type:date" json:"tglLembarChecklist"`
	TPejabatReviewerId      *uuid.UUID               `gorm:"type:uuid" json:"pejabatReviewerId"`
	PejabatReviewer         *PejabatAms              `gorm:"foreignKey:TPejabatReviewerId" json:"pejabatReviewer"`
	TPejabatMengetahuiId    *uuid.UUID               `gorm:"type:uuid" json:"pejabatMengetahuiId"`
	PejabatMengetahui       *PejabatAms              `gorm:"foreignKey:TPejabatMengetahuiId" json:"pejabatMengetahui"`
	Items                   []PenagihanAkuntansiItem `gorm:"foreignKey:TPenagihanAkuntansiId; constraint:OnDelete:CASCADE" json:"items"`
}

func (PenagihanAkuntansiItem) TableName() string {
	return "t_penagihan_akuntansi_item"
}

type PenagihanAkuntansiItem struct {
	baseapp.BaseModel
	TPenagihanAkuntansiId *uuid.UUID          `gorm:"type:uuid" json:"penagihanAkuntansiId"`
	PenagihanAkuntansi    *PenagihanAkuntansi `gorm:"foreignKey:TPenagihanAkuntansiId" json:"penagihanAkuntansi"`
	TProposePerhitunganId *uuid.UUID          `gorm:"type:uuid" json:"proposePerhitunganId"`
	ProposePerhitungan    *ProposePerhitungan `gorm:"foreignKey:TProposePerhitunganId" json:"proposePerhitungan"`
}
