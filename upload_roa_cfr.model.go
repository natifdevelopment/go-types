package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (UploadRoaCfr) TableName() string {
	return "t_upload_roa_cfr"
}

type UploadRoaCfr struct {
	baseapp.BaseModel
	TJadwalPengirimanId *uuid.UUID        `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman    *JadwalPengiriman `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TPembangkitId       *uuid.UUID        `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit          *Organization     `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	NoLot               string            `gorm:"type:varchar(255)" json:"noLot"`
	TSurveyorId         *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor            *Organization     `gorm:"foreignKey:TSurveyorId" json:"surveyor"`
	NoRoa               string            `gorm:"type:varchar(255)" json:"noRoa"`
	Gcv                 string            `gorm:"type:varchar(255)" json:"gcv"`
	Tm                  string            `gorm:"type:varchar(255)" json:"tm"`
	Ash                 string            `gorm:"type:varchar(255)" json:"ash"`
	Ts                  string            `gorm:"type:varchar(255)" json:"ts"`
	GcvPjbb             string            `gorm:"type:varchar(255)" json:"gcvPjbb"`
	TmPjbb              string            `gorm:"type:varchar(255)" json:"tmPjbb"`
	AshPjbb             string            `gorm:"type:varchar(255)" json:"ashPjbb"`
	TsPjbb              string            `gorm:"type:varchar(255)" json:"tsPjbb"`
	TglSampling         DateField   `gorm:"type:date" json:"tglSampling"`
	TFileRoaId          *uuid.UUID        `gorm:"type:uuid" json:"fileRoaId"`
	FileRoa             *Media            `gorm:"foreignKey:TFileRoaId" json:"fileRoa"`
	ApprovalRoaCfr      []ApprovalRoaCfr  `gorm:"foreignKey:TRoaCfrId" json:"approvalRoaCfr"`
}
