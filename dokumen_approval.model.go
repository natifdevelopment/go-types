package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type ApprovalStatus string

const (
	ApprovalStatusPending         ApprovalStatus = "PENDING"
	ApprovalStatusWaitingApproval ApprovalStatus = "WAITING_APPROVAL"
	ApprovalStatusApproved        ApprovalStatus = "APPROVED"
	ApprovalStatusPassed          ApprovalStatus = "PASSED"
	ApprovalStatusRejected        ApprovalStatus = "REJECTED"
)

func (o ApprovalStatus) IsValid() bool {
	return o == ApprovalStatusPending || o == ApprovalStatusWaitingApproval ||
		o == ApprovalStatusApproved || o == ApprovalStatusRejected
}

func (o ApprovalStatus) Value() (driver.Value, error) {
	return string(o), nil
}

func (o *ApprovalStatus) Scan(value interface{}) error {
	*o = ApprovalStatus(value.(string))
	return nil
}

func (Dokumen_approval) TableName() string {
	return "t_dokumen_approval"
}

type Dokumen_approval struct {
	baseapp.BaseModel
}

// Document viewer
type DocumentViewerItem struct {
	ID                 string               `json:"id"`
	DokParentId        uuid.UUID            `json:"-"`
	Name               string               `json:"name"`
	Url                string               `json:"url"`
	DownloadUrl        string               `json:"downloadUrl"`
	CreatedAt          time.Time            `json:"createdAt"`
	DokumenKolaborator []DokumenKolaborator `json:"dokumenKolaborator"`
	Type               string               `json:"type"`
	Diperiksa          bool                 `json:"diperiksa"`
	SudahDiperbaiki    bool                 `json:"sudahDiperbaiki"`
	Notes              *string              `json:"notes"`
	TMediaId           *uuid.UUID           `json:"mediaId"`
	TRefId             *uuid.UUID           `json:"refId"`

	// Will be deprecated
	// Now use TRefId
	TProposePerhitunganId *uuid.UUID `json:"proposePerhitunganId"`
	TJadwalPengirimanId   *uuid.UUID `json:"jadwalPengirimanId"`

	// Will be deprecated
	NamaDokumen          string `json:"namaDokumen"`
	IsSupportingDocument bool   `json:"isSupportingDocument"`
}

type DokApprovalHistoryType string

const (
	DokApprovalHistoryTypeBuat         DokApprovalHistoryType = "BUAT"
	DokApprovalHistoryTypeSetujui      DokApprovalHistoryType = "SETUJUI"
	DokApprovalHistoryTypeTolak        DokApprovalHistoryType = "TOLAK"
	DokApprovalHistoryTypeEdit         DokApprovalHistoryType = "EDIT"
	DokApprovalHistoryTypePeriksa      DokApprovalHistoryType = "PERIKSA"
	DokApprovalHistoryTypeCatatan      DokApprovalHistoryType = "CATATAN"
	DokApprovalHistoryTypeCatatanRumus DokApprovalHistoryType = "CATATAN_RUMUS"
	DokApprovalHistoryTypePerbaikan    DokApprovalHistoryType = "PERBAIKAN"
)

func (o DokApprovalHistoryType) IsValid() bool {
	return o == DokApprovalHistoryTypeBuat || o == DokApprovalHistoryTypeSetujui || o == DokApprovalHistoryTypeTolak || o == DokApprovalHistoryTypeEdit
}

func (DokApproval) TableName() string {
	return "t_dok_approval"
}

type DokApproval struct {
	baseapp.BaseModel
	TPageActivityId *uuid.UUID    `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity `gorm:"foreignKey:TPageActivityId" json:"pageActivity"`
	FlowCode        string        `gorm:"type:varchar(255)" json:"-"`
	TRefId          *uuid.UUID    `gorm:"type:uuid" json:"-"`
	TTargetId       *uuid.UUID    `gorm:"type:uuid" json:"-"`

	// Filled after approved/rejected
	TOrganizationId *uuid.UUID     `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization  `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId  *uuid.UUID     `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevel   `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId      *uuid.UUID     `gorm:"type:uuid" json:"accountId"`
	Account         *AccountInfo   `gorm:"foreignKey:TAccountId" json:"account"`
	ApprovalStatus  ApprovalStatus `gorm:"type:text" json:"approvalStatus"`
	ApprovalDate    *time.Time     `gorm:"type:timestamp" json:"approvalDate"`
	Comment         *string        `gorm:"type:text" json:"comment"`

	// @deprecated, now if resubmit, just change status to waiting_approval again
	SubmitUlang bool `gorm:"type:boolean;default:false" json:"submitUlang"`

	// @deprecated
	Disetujui *bool `gorm:"type:boolean" json:"disetujui"` // @deprecated, change to approvalStatus

	// Metadata and relations
	ApprovalType        ApprovalStrategyType `gorm:"type:varchar(100)" json:"approvalType"`
	TApprovalCodeId     *uuid.UUID           `gorm:"type:uuid" json:"approvalCodeId"`
	ApprovalCode        *ApprovalCode        `gorm:"foreignKey:TApprovalCodeId" json:"approvalCode"`
	TApprovalStrategyId *uuid.UUID           `gorm:"type:uuid" json:"approvalStrategyId"`
	ApprovalStrategy    *ApprovalStrategy    `gorm:"foreignKey:TApprovalStrategyId" json:"approvalStrategy"`
	TRequesterId        *uuid.UUID           `gorm:"type:uuid" json:"requesterId"`
	Requester           *AccountInfo         `gorm:"foreignKey:TRequesterId" json:"requester"`
	ApprovalLevelIdx    uint                 `gorm:"type:int" json:"approvalLevelIdx"`
	ApprovalCodeIdx     uint                 `gorm:"type:int" json:"approvalCodeIdx"`
	Dokumen             []DokApprovalDokumen `gorm:"foreignKey:TDokApprovalId; constraint:OnDelete:CASCADE" json:"dokumen"`
}

func (DokApprovalDokumen) TableName() string {
	return "t_dok_approval_dokumen"
}

type DokApprovalDokumen struct {
	baseapp.BaseModel
	TDokApprovalId  *uuid.UUID   `gorm:"type:uuid" json:"dokApprovalId"`
	DokApproval     *DokApproval `gorm:"foreignKey:TDokApprovalId" json:"dokApproval"`
	TRefId          *uuid.UUID   `gorm:"type:uuid" json:"-"`
	TTargetId       *uuid.UUID   `gorm:"type:uuid" json:"-"`
	DokParentId     *uuid.UUID   `gorm:"type:uuid" json:"-"`
	TMediaId        *uuid.UUID   `gorm:"type:uuid" json:"mediaId"`
	Media           *Media       `gorm:"foreignKey:TMediaId" json:"media"`
	TipeDokumen     string       `gorm:"type:varchar(255)" json:"tipeDokumen"`
	Keterangan      *string      `gorm:"type:varchar(255)" json:"keterangan"`
	SudahDiperbaiki bool         `gorm:"type:boolean;default:false" json:"sudahDiperbaiki"`
	Diperiksa       bool         `gorm:"type:boolean;default:false" json:"diperiksa"`
}

func (DokApprovalHistory) TableName() string {
	return "t_dok_approval_history"
}

type DokApprovalHistory struct {
	baseapp.BaseModel
	TDokApprovalId   *uuid.UUID             `gorm:"type:uuid" json:"dokApprovalId"`
	DokApproval      *DokApproval           `gorm:"foreignKey:TDokApprovalId" json:"dokApproval"`
	TRefId           *uuid.UUID             `gorm:"type:uuid" json:"-"`
	TTargetId        *uuid.UUID             `gorm:"type:uuid" json:"-"`
	TOrganizationId  *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	OrganizationInfo *OrganizationInfo      `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId   *uuid.UUID             `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel      *AccessLevel           `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TAccountId       *uuid.UUID             `gorm:"type:uuid" json:"accountId"`
	AccountInfo      *AccountInfo           `gorm:"foreignKey:TAccountId" json:"account"`
	Tipe             DokApprovalHistoryType `gorm:"type:varchar(255)" json:"tipe"`
	Keterangan       *string                `gorm:"type:varchar(255)" json:"keterangan"`
	TipeDokumen      *string                `gorm:"type:varchar(255)" json:"tipeDokumen"`
	NamaDokumen      *string                `gorm:"type:varchar(255)" json:"namaDokumen"`
}
