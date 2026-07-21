package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (StokOpname) TableName() string {
	return "t_stok_opname"
}

type StokOpname struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization    *OrganizationInfo `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	NoSertifikat    string            `gorm:"type:varchar(255)" json:"noSertifikat"`
	TSurveyorId     *uuid.UUID        `gorm:"type:uuid" json:"surveyorId"`
	Surveyor        *Organization     `gorm:"foreignKey:TSurveyorId;" json:"surveyor"`
	TanggalStok     DateField   `gorm:"type:date" json:"tanggalStok"`
	StokAwal        float64           `gorm:"type:float" json:"stokAwal"`
	StokOpname      float64           `gorm:"type:float" json:"stokOpname"`
	StockAdjustment float64           `gorm:"type:float" json:"stockAdjustment"`
	KirimKeSapErp   bool              `gorm:"type:bool" json:"kirimKeSapErp"`
	FileOpnameId    *uuid.UUID        `gorm:"type:uuid" json:"fileOpnameId"`
	FileOpname      *Media            `gorm:"foreignKey:FileOpnameId;" json:"fileOpname"`
	Keterangan      *string           `gorm:"type:text" json:"keterangan"`
	MatDocNo        *string           `gorm:"type:varchar(255)" json:"matDocNo"`

	// SAP Integration
	TIntegrasiReqId  *uuid.UUID       `gorm:"type:uuid" json:"integrasiReqId"`
	IntegrasiReq     *IntegrasiStatus `gorm:"foreignKey:TIntegrasiReqId" json:"integrasiReq"`
	TIntegrasiRespId *uuid.UUID       `gorm:"type:uuid" json:"integrasiRespId"`
	IntegrasiResp    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiRespId" json:"integrasiResp"`
	SapStatus        ResponseStatus   `gorm:"type:varchar(255);default:'CREATED'" json:"sapStatus"`
	SapMatDocNo      *string          `gorm:"type:varchar(255);" json:"sapMatDocNo"`
	SapDocYear       *int             `gorm:"type:int;" json:"sapDocYear"`
	SapLineNo        *string          `gorm:"type:varchar(255);" json:"sapLineNo"`
}
