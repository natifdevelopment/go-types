package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (JadwalPraRakor) TableName() string {
	return "t_jadwal_pra_rakor"
}

type JadwalPraRakor struct {
	baseapp.BaseModel
	Judul                  string                   `gorm:"type:varchar(255)" json:"judul"`
	TOrganizationId        *uuid.UUID               `gorm:"type:uuid" json:"organizationId"`
	Organization           *Organization            `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TRegionalId            *uuid.UUID               `gorm:"type:uuid" json:"regionalId"`
	Regional               *Organization            `gorm:"foreignKey:TRegionalId" json:"regional"`
	Tanggal                DateField          `gorm:"type:date" json:"tanggal"`
	Periode                *YearMonthField    `gorm:"type:date" json:"periode"`
	JadwalPraRakorUndangan []JadwalPraRakorUndangan `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalPraRakorUndangan"`
	RencanaPasokan         []RencanaPasokan         `gorm:"foreignKey:TJadwalPraRakorId" json:"rencanaPasokan,omitempty"`
	JadwalRakor            []JadwalRakor            `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalRakor,omitempty"`
}

func (JadwalPraRakorUndangan) TableName() string {
	return "t_jadwal_pra_rakor_undangan"
}

type JadwalPraRakorUndangan struct {
	baseapp.BaseModel
	TJadwalPraRakorId *uuid.UUID      `gorm:"type:uuid" json:"jadwalPraRakorId"`
	JadwalPraRakor    *JadwalPraRakor `gorm:"foreignKey:TJadwalPraRakorId" json:"jadwalPraRakor"`
	TOrganizationId   *uuid.UUID      `gorm:"type:uuid" json:"organizationId"`
	Organization      *Organization   `gorm:"foreignKey:TOrganizationId" json:"organization"`
}
