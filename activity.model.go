package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (Activity) TableName() string {
	return "t_activity"
}

type Activity struct {
	baseapp.BaseModel
	TPageId              uuid.UUID              `gorm:"type:uuid" json:"pageId"`
	Page                 *PageInfo              `gorm:"foreignKey:TPageId" json:"page"`
	Name                 string                 `gorm:"type:varchar(255)" json:"name"`
	Code                 string                 `gorm:"type:varchar(255)" json:"code"`
	Description          string                 `gorm:"type:varchar(255)" json:"description"`
	ConnectQuery         bool                   `gorm:"type:boolean" json:"connectQuery"`
	ActivityItem         []ActivityItem         `gorm:"foreignKey:TActivityId;constraint:OnDelete:CASCADE" json:"activityItems"`
	ActivityItemSelected []ActivityItemSelected `gorm:"foreignKey:TActivityId;constraint:OnDelete:CASCADE"` // Deprecated
	ActivityConfigItem   []ActivityConfigItem   `gorm:"foreignKey:TActivityId;constraint:OnDelete:CASCADE"`
	ActivityAccess       []ActivityAccess       `gorm:"foreignKey:TActivityId;constraint:OnDelete:CASCADE"`
	// ApprovalStrategies   []ApprovalStrategy     `gorm:"foreignKey:TActivityId;constraint:OnDelete:CASCADE" json:"approvalStrategies"`
}

func (ActivityInfo) TableName() string {
	return "t_activity"
}

type ActivityInfo struct {
	baseapp.BaseModel
	Uuid        string `gorm:"type:varchar(85)" json:"uuid"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Code        string `gorm:"type:varchar(255);unique" json:"code"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

func (ActivityItemSelected) TableName() string {
	return "t_activity_item_selected"
}

type ActivityItemSelected struct {
	baseapp.BaseModel
	TActivityId     *uuid.UUID   `gorm:"type:uuid" json:"activityId"`
	TActivityItemId *uuid.UUID   `gorm:"type:uuid" json:"activityItemId"`
	ActivityItem    ActivityItem `gorm:"foreignKey:TActivityItemId" json:"activityItem"`
}

func (ActivityConfigItem) TableName() string {
	return "t_activity_config_item"
}

type ActivityConfigItem struct {
	baseapp.BaseModel
	TActivityId   *uuid.UUID     `gorm:"type:uuid" json:"activityId"`
	TConfigDataId *uuid.UUID     `gorm:"type:uuid" json:"configDataId"`
	ConfigData    ConfigDataInfo `gorm:"foreignKey:TConfigDataId" json:"configData"`
}

type ActivityAccessType string

const (
	activityAccessType1 ActivityAccessType = "organization"
	activityAccessType2 ActivityAccessType = "account"
)

func (o ActivityAccessType) IsValid() bool {
	return o == activityAccessType1 || o == activityAccessType2
}

func (o ActivityAccessType) IsOrganization() bool {
	return o == activityAccessType1
}

func (o ActivityAccessType) IsAccount() bool {
	return o == activityAccessType2
}

func (ActivityAccess) TableName() string {
	return "t_activity_access"
}

type ActivityAccess struct {
	baseapp.BaseModel
	TActivityId     *uuid.UUID         `gorm:"type:uuid" json:"activityId"`
	Type            ActivityAccessType `gorm:"type:varchar(50)" json:"type"`
	TOrganizationId *uuid.UUID         `gorm:"type:uuid" json:"organizationId"`
	Organization    Organization       `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	TAccountId      *uuid.UUID         `gorm:"type:uuid" json:"accountId"`
	Account         Account            `gorm:"foreignKey:TAccountId;" json:"account"`
}

type ActivityShipmentCif struct {
	JadwalPengirimanTanggal               *string `json:"jadwalPengirimanTanggal"`
	JadwalPengirimanOleh                  *string `json:"jadwalPengirimanOleh"`
	LoadingInfoCifTanggal                 *string `json:"loadingInfoCifTanggal"`
	LoadingInfoCifOleh                    *string `json:"loadingInfoCifOleh"`
	CoaLoadingTanggal                     *string `json:"coaLoadingTanggal"`
	CoaLoadingOleh                        *string `json:"coaLoadingOleh"`
	PsaLoadingRoaTanggal                  *string `json:"psaLoadingRoaTanggal"`
	PsaLoadingRoaOleh                     *string `json:"psaLoadingRoaOleh"`
	NorIzinSandarBongkarTanggal           *string `json:"norIzinSandarBongkarTanggal"`
	NorIzinSandarBongkarOleh              *string `json:"norIzinSandarBongkarOleh"`
	NorIzinSandarBongkarPembangkitTanggal *string `json:"norIzinSandarBongkarPembangkitTanggal"`
	NorIzinSandarBongkarPembangkitOleh    *string `json:"norIzinSandarBongkarPembangkitOleh"`
	CatatBongkarDsTanggal                 *string `json:"catatBongkarDsTanggal"`
	CatatBongkarDsOleh                    *string `json:"catatBongkarDsOleh"`
	BastCifTanggal                        *string `json:"bastCifTanggal"`
	BastCifOleh                           *string `json:"bastCifOleh"`
}

type ActivityShipmentFob struct {
	JadwalPengirimanTanggal               *string `json:"jadwalPengirimanTanggal"`
	JadwalPengirimanOleh                  *string `json:"jadwalPengirimanOleh"`
	RoaTanggal                            *string `json:"roaTanggal"`
	RoaOleh                               *string `json:"roaOleh"`
	ApprovalRoaFobTanggal                 *string `json:"approvalRoaFobTanggal"`
	ApprovalRoaFobOleh                    *string `json:"approvalRoaFobOleh"`
	LoadingInfoFobTanggal                 *string `json:"loadingInfoFobTanggal"`
	LoadingInfoFobOleh                    *string `json:"loadingInfoFobOleh"`
	BastLoadingFobTanggal                 *string `json:"bastLoadingFobTanggal"`
	BastLoadingFobOleh                    *string `json:"bastLoadingFobOleh"`
	LoadingInfoTransTanggal               *string `json:"loadingInfoTransTanggal"`
	LoadingInfoTransOleh                  *string `json:"loadingInfoTransOleh"`
	NorLoadingTanggal                     *string `json:"norLoadingTanggal"`
	NorLoadingOleh                        *string `json:"norLoadingOleh"`
	NorIzinSandarBongkarTransTanggal      *string `json:"norIzinSandarBongkarTransTanggal"`
	NorIzinSandarBongkarTransOleh         *string `json:"norIzinSandarBongkarTransOleh"`
	NorIzinSandarBongkarPembangkitTanggal *string `json:"norIzinSandarBongkarPembangkitTanggal"`
	NorIzinSandarBongkarPembangkitOleh    *string `json:"norIzinSandarBongkarPembangkitOleh"`
	CatatBongkarDsTanggal                 *string `json:"catatBongkarDsTanggal"`
	CatatBongkarDsOleh                    *string `json:"catatBongkarDsOleh"`
}

type ActivityShipmentCfr struct {
	JadwalPengirimanTanggal        *string `json:"jadwalPengirimanTanggal"`
	JadwalPengirimanOleh           *string `json:"jadwalPengirimanOleh"`
	UploadRoaCfrTanggal            *string `json:"uploadRoaCfrTanggal"`
	UploadRoaCfrOleh               *string `json:"uploadRoaCfrOleh"`
	ApprovalRoaCfrTanggal          *string `json:"approvalRoaCfrTanggal"`
	ApprovalRoaCfrOleh             *string `json:"approvalRoaCfrOleh"`
	PencatatanPenerimaanCfrTanggal *string `json:"pencatatanPenerimaanCfrTanggal"`
	PencatatanPenerimaanCfrOleh    *string `json:"pencatatanPenerimaanCfrOleh"`
	BaSerahTerimaCfrTanggal        *string `json:"baSerahTerimaCfrTanggal"`
	BaSerahTerimaCfrOleh           *string `json:"baSerahTerimaCfrOleh"`
}
