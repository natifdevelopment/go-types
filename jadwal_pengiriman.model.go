package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"
"github.com/google/uuid"
)

type ShipmentStatus string

const (
	ShipmentStatusKapalKaram   ShipmentStatus = "KAPAL_KARAM"
	ShipmentStatusBelumLoading ShipmentStatus = "BELUM_LOADING"
	ShipmentStatusLoading      ShipmentStatus = "LOADING"
	ShipmentStatusCoa          ShipmentStatus = "COA"
	ShipmentStatusPsa          ShipmentStatus = "PSA"
	ShipmentStatusNor          ShipmentStatus = "NOR"
	ShipmentStatusNorLoading   ShipmentStatus = "NOR_LOADING"
	ShipmentStatusBongkar      ShipmentStatus = "BONGKAR"
	ShipmentStatusPenerimaan   ShipmentStatus = "PENERIMAAN"
	ShipmentStatusBast         ShipmentStatus = "BAST"
	ShipmentStatusRoa          ShipmentStatus = "ROA"
	ShipmentStatusRoaCfr       ShipmentStatus = "ROA_CFR"
	ShipmentStatusPropose      ShipmentStatus = "PROPOSE"
	ShipmentStatusInvoice      ShipmentStatus = "INVOICE"
	ShipmentStatusPembayaran   ShipmentStatus = "PEMBAYARAN"
	ShipmentStatusSpt          ShipmentStatus = "SPT"
	ShipmentStatusNodin        ShipmentStatus = "NODIN"
	ShipmentStatusVerifikasi   ShipmentStatus = "VERIFIKASI"
	ShipmentStatusPelunasan    ShipmentStatus = "PELUNASAN"
	ShipmentStatusSelesai      ShipmentStatus = "SELESAI"
	ShipmentStatusRejected     ShipmentStatus = "REJECTED"
)

type StatusBankGransi string

const (
	StatusBankGransiBelumBerlaku  StatusBankGransi = "BELUM_BERLAKU"  // bank garansi belum berlaku untuk tanggal yg ditetapkan
	StatusBankGransiAktif         StatusBankGransi = "AKTIF"          // bank garansi aktif
	StatusBankGransiExpired       StatusBankGransi = "KADALUARSA"     // masa belaku bank garansi sudah lewat
	StatusBankGransiTidakTersedia StatusBankGransi = "TIDAK_TERSEDIA" // pada kontrak bank garansi tidak ditetapkan
)

type JenisPengiriman string

const (
	JenisPengirimanCif   JenisPengiriman = "bbo//cd/skema-kontrak/sk-bbo-1"
	JenisPengirimanFob   JenisPengiriman = "bbo//cd/skema-kontrak/sk-bbo-2"
	JenisPengirimanCfr   JenisPengiriman = "bbo//cd/skema-kontrak/sk-bbo-3"
	JenisPengirimanTrans JenisPengiriman = "bbo//cd/skema-kontrak/sk-bbo-4"
)

func (o JenisPengiriman) IsValid() bool {
	return o == JenisPengirimanCif || o == JenisPengirimanFob || o == JenisPengirimanTrans || o == JenisPengirimanCfr
}

type StatusKelengkapan string

const (
	StatusKelengkapanLengkap      StatusKelengkapan = "LENGKAP"
	StatusKelengkapanBelumLengkap StatusKelengkapan = "BELUM_LENGKAP"
)

func (o StatusKelengkapan) IsValid() bool {
	return o == StatusKelengkapanLengkap || o == StatusKelengkapanBelumLengkap
}

func (JadwalPengiriman) TableName() string {
	return "t_jadwal_pengiriman"
}

type JadwalPengiriman struct {
	baseapp.BaseModel
	TJadwalRakorId         *uuid.UUID            `gorm:"type:uuid" json:"jadwalRakorId"`
	JadwalRakor            *JadwalRakor          `gorm:"foreignKey:TJadwalRakorId" json:"jadwalRakor"`
	TEntryJadwalId         *uuid.UUID            `gorm:"type:uuid" json:"entryJadwalId"`
	EntryJadwal            *EntryJadwal          `gorm:"foreignKey:TEntryJadwalId" json:"entryJadwal"`
	TShipmentPenugasanId   *uuid.UUID            `gorm:"type:uuid" json:"shipmentPenugasanId"`
	ShipmentPenugasan      *ShipmentPenugasan    `gorm:"foreignKey:TShipmentPenugasanId" json:"shipmentPenugasan"`
	TPjbbId                *uuid.UUID            `gorm:"type:uuid" json:"pjbbId"`
	Pjbb                   *Pjbb                 `gorm:"foreignKey:TPjbbId" json:"pjbb"`
	TPjbbAmandemenId       *uuid.UUID            `gorm:"type:uuid" json:"pjbbAmandemenId"`
	PjbbAmandemen          *PjbbAmandemen        `gorm:"foreignKey:TPjbbAmandemenId" json:"pjbbAmandemen"`
	TPjabId                *uuid.UUID            `gorm:"type:uuid" json:"pjabId"`
	Pjab                   *Pjab                 `gorm:"foreignKey:TPjabId" json:"pjab"`
	TPjabAmandemenId       *uuid.UUID            `gorm:"type:uuid" json:"pjabAmandemenId"`
	PjabAmandemen          *PjabAmandemen        `gorm:"foreignKey:TPjabAmandemenId" json:"pjabAmandemen"`
	IsBaseJadwal           bool                  `gorm:"type:bool;default:true" json:"isBaseJadwal"`
	Periode                *YearMonthField `gorm:"type:date" json:"periode"`
	NoJadwal               string                `gorm:"type:varchar(255)" json:"noJadwal"`
	NoJadwalCount          int                   `gorm:"type:int" json:"-"`
	NoPengiriman           *string               `gorm:"type:varchar(255)" json:"noPengiriman"`
	NoPengirimanCount      int                   `gorm:"type:int" json:"-"`
	TOrganizationId        *uuid.UUID            `gorm:"type:uuid" json:"organizationId"` // EPI
	Organization           *Organization         `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TPembangkitId          *uuid.UUID            `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit             *Organization         `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	TPemasokId             *uuid.UUID            `gorm:"type:uuid" json:"pemasokId"`
	Pemasok                *Organization         `gorm:"foreignKey:TPemasokId" json:"pemasok"`
	TTransportirId         *uuid.UUID            `gorm:"type:uuid" json:"transportirId"`
	Transportir            *Organization         `gorm:"foreignKey:TTransportirId" json:"transportir"`
	TCurrentPicOrgId       *uuid.UUID            `gorm:"type:uuid" json:"currentPicOrgId"`
	CurrentPicOrg          *Organization         `gorm:"foreignKey:TCurrentPicOrgId" json:"currentPicOrg"`
	TJadwalPengirimanFobId *uuid.UUID            `gorm:"type:uuid" json:"jadwalPengirimanFobId"`
	JadwalPengirimanFob    *JadwalPengiriman     `gorm:"foreignKey:TJadwalPengirimanFobId" json:"jadwalPengirimanFob"`
	TJenisBatuBaraId       *uuid.UUID            `gorm:"index;type:uuid" json:"jenisBatuBaraId"`
	JenisBatuBara          *JenisBatuBara        `gorm:"foreignKey:TJenisBatuBaraId" json:"jenisBatuBara"`
	SkemaKontrakCode       *JenisPengiriman      `gorm:"type:varchar(255)" json:"skemaKontrakCode"` // CIF/FOB/CFR
	SkemaKontrak           *ConfigDataInfo       `gorm:"foreignKey:SkemaKontrakCode;references:Code;" json:"skemaKontrak"`
	TSpekId                *uuid.UUID            `gorm:"type:uuid" json:"spekId"`
	Spek                   *Spek                 `gorm:"foreignKey:TSpekId" json:"spek"`
	DueDate                time.Time             `gorm:"type:timestamp" json:"dueDate"`
	TargetLoading          *DateField      `gorm:"type:date" json:"targetLoading"`
	TanggalEta             *DateField      `gorm:"type:date" json:"tanggalEta"`
	VolumeConfirm          float64               `gorm:"type:float" json:"volumeConfirm"`
	VolumePerLotPjbb       float64               `gorm:"type:float" json:"volumePerLotPjbb"`
	VolumePerLot           float64               `gorm:"type:float" json:"volumePerLot"`
	TMasterKapalId         *uuid.UUID            `gorm:"type:uuid" json:"masterKapalId"`
	MasterKapal            *MasterKapal          `gorm:"foreignKey:TMasterKapalId" json:"masterKapal"`
	TMasterTongkangId      *uuid.UUID            `gorm:"type:uuid" json:"masterTongkangId"`
	MasterTongkang         *MasterTongkang       `gorm:"foreignKey:TMasterTongkangId" json:"masterTongkang"`
	ClosedAt               *time.Time            `gorm:"type:timestamp" json:"closedAt"`
	TClosedById            *uuid.UUID            `gorm:"type:uuid" json:"closedById"`
	ClosedBy               *Account              `gorm:"foreignKey:TClosedById" json:"closedBy"`
	LastReminder           *time.Time            `gorm:"type:timestamp" json:"lastReminder"`
	StatusJadwalPengiriman ShipmentStatus        `gorm:"type:varchar(255);default:'BELUM_LOADING';" json:"statusJadwalPengiriman"`
	TPageActivityId        *uuid.UUID            `gorm:"type:uuid" json:"pageActivityId"` // Current page
	PageActivity           *PageActivity         `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
	SelesaiPenyerahanCfr   bool                  `gorm:"bool;default:false" json:"selesaiPenyerahanCfr"`

	// Surveyor
	// 1. SurveyorPembangkit digunakan untuk upload COW & COW oleh surveyor
	// 2. SurveyorBongkar akan ter isi jika pada kontrak ada repro
	TSurveyorPembangkitId *uuid.UUID    `gorm:"type:uuid" json:"surveyorPembangkitId"`
	SurveyorPembangkit    *Organization `gorm:"foreignKey:TSurveyorPembangkitId" json:"surveyorPembangkit"`
	TSurveyorBongkarId    *uuid.UUID    `gorm:"type:uuid" json:"surveyorBongkarId"`
	SurveyorBongkar       *Organization `gorm:"foreignKey:TSurveyorBongkarId" json:"surveyorBongkar"`

	LoadingInfoCif                 []LoadingInfoCif                     `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"loadingInfoCif"`
	LoadingInfoTrans               []LoadingInfoTrans                   `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"loadingInfoTrans"`
	LoadingInfoFob                 []LoadingInfo                        `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"loadingInfoFob"`
	CoaLoading                     []CoaLoading                         `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"coaLoading"`
	CatatBongkarDs                 []CatatBongkarDs                     `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"catatBongkarDs"`
	NorIzinSandarBongkar           []NorIzinSandarBongkar               `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"norIzinSandarBongkar"`
	NorIzinSandarBongkarTrans      []NorIzinSandarNongkarTrans          `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"norIzinSandarBongkarTrans"`
	NorIzinSandarBongkarPembangkit []Nor_izin_sandar_bongkar_pembangkit `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"norIzinSandarBongkarPembangkit"`
	Roa                            []Roa                                `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"roa"`
	ApprovalRoaFob                 []ApprovalRoaFob                     `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"approvalRoaFob"`
	ApprovalRoaCfr                 []ApprovalRoaCfr                     `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"approvalRoaCfr"` // @deprecated
	NorLoading                     []NorLoading                         `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"norLoading"`
	RoaCfr                         []UploadRoaCfr                       `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"roaCfr"`
	BastCif                        []BastCif                            `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"bastCif"`
	BastLoading                    []BastLoading                        `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"bastLoading"`
	PengajuanUmpire                []PengajuanUmpire                    `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"pengajuanUmpire"`
	PsaLoadingRoa                  []PsaLoadingRoa                      `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"psaLoadingRoa"`
	CoaCowCif                      []CoaCow                             `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"coaCowCif"`
	CoaCowFob                      []CoaCowFob                          `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"coaCowFob"`
	CoaCowTrans                    []CoaCowTrans                        `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"coaCowTrans"`
	// CoaCowCfr                      []Coa_cow_CFR                        `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"coaCowCfr"`
	ProposePerhitungan      []ProposePerhitungan      `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"proposePerhitungan"`
	PencatatanPenerimaanCfr []PencatatanPenerimaanCfr `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"pencatatanPenerimaanCfr"`
	BastCifApproval         []BastCifApproval         `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"bastCifApproval"`
	LoadingInfoCfr          []LoadingInfoCfr          `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"loadingInfoCfr"`
	Stats                   *JadwalPengirimanStats    `gorm:"-" json:"stats,omitempty"`
	StatusKelengkapan       *StatusKelengkapan        `gorm:"-" json:"statusKelengkapan,omitempty"`
	PipPenerimaan           []PipPenerimaan           `gorm:"foreignKey:TJadwalPengirimanId;constraint:OnDelete:CASCADE" json:"pipPenerimaan"`

	// Manual Propose
	StatusManualPropose StatusPerhitungan `gorm:"type:varchar(255);default:'BELUM_DIHITUNG'" json:"statusManualPropose"` // sudah_dihitung/belum_dihitung
	StatusProposeTermin StatusPerhitungan `gorm:"type:varchar(255);default:'BELUM_DIHITUNG'" json:"statusProposeTermin"` // sudah_dihitung/belum_dihitung

	// @deprecated
	// skarang split di pra rakor
	// TLoadingInfoCfrId              *uuid.UUID                           `gorm:"type:uuid" json:"loadingInfoCfrId"`
	// LoadingInfoCfr                 *LoadingInfoCfr                      `gorm:"foreignKey:TLoadingInfoCfrId" json:"loadingInfoCfr"`
}

type JadwalPengirimanStats struct {
	JumlahPengiriman int     `json:"jumlahPengiriman"`
	TotalVolume      float64 `json:"totalVolume"`
}

func (JadwalPengirimanSimple) TableName() string {
	return "t_jadwal_pengiriman"
}

type JadwalPengirimanSimple struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	NoJadwal     string    `gorm:"type:varchar(255)" json:"noJadwal"`
	NoPengiriman *string   `gorm:"type:varchar(255)" json:"noPengiriman"`
}
