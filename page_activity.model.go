package types
import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)
type PageType string

const (
	PageTypeTransaksi   PageType = "TRANSAKSI"
	PageTypeKonfigurasi PageType = "KONFIGURASI"
	PageTypeMasterData  PageType = "MASTER_DATA"
	PageTypeKontrak     PageType = "KONTRAK"
)

func (o PageType) IsValid() bool {
	return o == PageTypeTransaksi ||
		o == PageTypeKonfigurasi ||
		o == PageTypeMasterData ||
		o == PageTypeKontrak
}

func (PageActivity) TableName() string {
	return "t_page_activity"
}

type PageActivity struct {
	baseapp.BaseModel
	Name                   string            `gorm:"type:varchar(255);uniqueIndex:uq_page_activity_name_parent,where:deleted_at IS NULL" json:"name"`
	Code                   string            `gorm:"type:varchar(255)" json:"code"`
	Type                   string            `gorm:"type:varchar(255);default:'menu'" json:"type"` // header / menu
	PageType               PageType          `gorm:"type:varchar(255)" json:"pageType"`
	OrderNo                int               `gorm:"type:int" json:"orderNo"`
	Icon                   *string           `gorm:"type:varchar(255)" json:"icon"`
	Url                    *string           `gorm:"type:varchar(255)" json:"url"`
	Description            *string           `gorm:"type:varchar(255)" json:"description"`
	IsApprovalRequired     bool              `gorm:"type:boolean;default:false" json:"isApprovalRequired"`
	AllowEditAfterApproved bool              `gorm:"type:boolean;default:false" json:"allowEditAfterApproved"`
	ApprovalAction         *string           `gorm:"type:varchar(255)" json:"approvalAction"`
	TApprovalActivityId    *uuid.UUID        `gorm:"type:uuid" json:"approvalActivityId"`
	TApprovalToId          *uuid.UUID        `gorm:"type:uuid" json:"approvalToId"`
	TApprovalStrategyId    *uuid.UUID        `gorm:"type:uuid" json:"approvalStrategyId"`
	ApprovalStrategy       *ApprovalStrategy `gorm:"foreignKey:TApprovalStrategyId" json:"approvalStrategy"`
	TApprovalRequesterId   *uuid.UUID        `gorm:"type:uuid" json:"approvalRequesterId"`
	TParentId              *uuid.UUID        `gorm:"type:uuid;uniqueIndex:uq_page_activity_name_parent,where:deleted_at IS NULL" json:"parentId"`
	Parent                 *PageActivity     `gorm:"foreignKey:TParentId" json:"parent"`
	TNextFlowId            *uuid.UUID        `gorm:"type:uuid" json:"nextFlowId"`
	NextFlow               *PageActivity     `gorm:"foreignKey:TNextFlowId" json:"nextFlow"`
	TPicAccessLevelId      *uuid.UUID        `gorm:"type:uuid" json:"picAccessLevelId"`
	PicAccessLevel         *AccessLevel      `gorm:"foreignKey:TPicAccessLevelId" json:"picAccessLevel"`
	SlaDays                float64           `gorm:"type:float" json:"slaDays"`
	Children               []PageActivity    `gorm:"foreignKey:TParentId;references:ID;constraint:OnDelete:SET NULL" json:"children"`
	TaskActivity           *Task_activity    `gorm:"-" json:"taskActivity"`
	Rollback               []RollbackItem    `gorm:"foreignKey:TPageActivityId" json:"rollback"`
}
