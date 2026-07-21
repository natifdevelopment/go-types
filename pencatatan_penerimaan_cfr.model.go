package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
"github.com/google/uuid"
)

func (PencatatanPenerimaanCfr) TableName() string {
	return "t_pencatatan_penerimaan_cfr"
}

type PencatatanPenerimaanCfr struct {
	baseapp.BaseModel
	TJadwalPengirimanId    *uuid.UUID                    `gorm:"type:uuid" json:"jadwalPengirimanId"`
	JadwalPengiriman       *JadwalPengiriman             `gorm:"foreignKey:TJadwalPengirimanId" json:"jadwalPengiriman"`
	TipeBongkar            TipeBongkarDs                 `gorm:"type:varchar(255)" json:"tipeBongkar"`
	NoRid                  *string                       `gorm:"type:varchar(255)" json:"noRid"`
	AlatPengirimanCode     *string                       `gorm:"type:varchar(255)" json:"alatPengirimanCode"`
	AlatPengiriman         *ConfigDataInfo               `gorm:"foreignKey:AlatPengirimanCode;references:Code;" json:"alatPengiriman"`
	TglMulaiPenyerahan     *DateTimeField          `gorm:"type:timestamp" json:"tglMulaiPenyerahan"`
	TglSelesaiPenyerahan   *DateTimeField          `gorm:"type:timestamp" json:"tglSelesaiPenyerahan"`
	LokasiPenyerahan       *string                       `gorm:"type:text" json:"lokasiPenyerahan"`
	VolumePenyerahan       *float32                      `gorm:"type:float" json:"volumePenyerahan"`
	Rit                    *float32                      `gorm:"type:float" json:"rit"`
	SelesaiPenyerahan      bool                          `gorm:"type:boolean" json:"selesaiPenyerahan"`
	GangguanHariPenerimaan *float32                      `gorm:"type:float" json:"gangguanHariPenerimaan"`
	TDocTimesheetId        *uuid.UUID                    `gorm:"type:uuid" json:"docTimesheetId"`
	DocTimesheet           *Media                        `gorm:"foreignKey:TDocTimesheetId" json:"docTimesheet"`
	LokasiPengukuran       *string                       `gorm:"type:text" json:"lokasiPengukuran"`
	NilaiCounterAwal       *float32                      `gorm:"type:float" json:"nilaiCounterAwal"`
	NilaiCounterAkhir      *float32                      `gorm:"type:float" json:"nilaiCounterAkhir"`
	Children               []PencatatanPenerimaanCfrItem `gorm:"foreignKey:TParentId;references:ID;constraint:OnDelete:SET NULL" json:"children"`
}

func (PencatatanPenerimaanCfrItem) TableName() string {
	return "t_pencatatan_penerimaan_cfr_item"
}

type PencatatanPenerimaanCfrItem struct {
	baseapp.BaseModel
	NoRid                  *string                  `gorm:"type:varchar(255)" json:"noRid"`
	AlatPengirimanCode     *string                  `gorm:"type:varchar(255)" json:"alatPengirimanCode"`
	AlatPengiriman         *ConfigDataInfo          `gorm:"foreignKey:AlatPengirimanCode;references:Code;" json:"alatPengiriman"`
	TglMulaiPenyerahan     *DateTimeField     `gorm:"type:timestamp" json:"tglMulaiPenyerahan"`
	TglSelesaiPenyerahan   *DateTimeField     `gorm:"type:timestamp" json:"tglSelesaiPenyerahan"`
	LokasiPenyerahan       *string                  `gorm:"type:text" json:"lokasiPenyerahan"`
	VolumePenyerahan       *float32                 `gorm:"type:float" json:"volumePenyerahan"`
	Rit                    *float32                 `gorm:"type:float" json:"rit"`
	SelesaiPenyerahan      bool                     `gorm:"type:boolean" json:"selesaiPenyerahan"`
	GangguanHariPenerimaan *float32                 `gorm:"type:float" json:"gangguanHariPenerimaan"`
	TDocTimesheetId        *uuid.UUID               `gorm:"type:uuid" json:"docTimesheetId"`
	DocTimesheet           *Media                   `gorm:"foreignKey:TDocTimesheetId" json:"docTimesheet"`
	LokasiPengukuran       *string                  `gorm:"type:text" json:"lokasiPengukuran"`
	NilaiCounterAwal       *float32                 `gorm:"type:float" json:"nilaiCounterAwal"`
	NilaiCounterAkhir      *float32                 `gorm:"type:float" json:"nilaiCounterAkhir"`
	TParentId              *uuid.UUID               `gorm:"type:uuid" json:"parentId"`
	Parent                 *PencatatanPenerimaanCfr `gorm:"foreignKey:TParentId" json:"parent"`
}
