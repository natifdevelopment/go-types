package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (NorLoading) TableName() string {
	return "t_nor_loading"
}

type NorLoading struct {
	baseapp.BaseModel
	TJadwalPengirimanId       *uuid.UUID                  `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman          *JadwalPengiriman           `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	Disetujui                 bool                        `gorm:"type:bool" json:"disetujui"`
	Keterangan                *string                     `gorm:"type:text" json:"keterangan"`
	TSurveyorPembangkitId     *uuid.UUID                  `gorm:"type:uuid" json:"surveyorPembangkitId"`
	SurveyorPembangkit        *Organization               `gorm:"foreignKey:TSurveyorPembangkitId" json:"surveyorPembangkit"`
	TLoadingInfoTransId       *uuid.UUID                  `gorm:"type:uuid" json:"loadingInfoTransId"`
	LoadingInfoTrans          *LoadingInfoTrans           `gorm:"foreignKey:TLoadingInfoTransId" json:"loadingInfoTrans"`
	NorIzinSandarBongkarTrans []NorIzinSandarNongkarTrans `gorm:"foreignKey:TNorLoadingId; constraint:OnDelete:CASCADE" json:"norIzinSandarBongkarTrans"`
	NorIzinSandarBongkar      []NorIzinSandarBongkar      `gorm:"foreignKey:TNorLoadingId; constraint:OnDelete:CASCADE" json:"norIzinSandarBongkar"`
}
