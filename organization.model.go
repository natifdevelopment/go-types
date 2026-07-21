package types

import (
	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (Organization) TableName() string {
	return "t_organization"
}

type Organization struct {
	baseapp.BaseModel
	TParentId             *uuid.UUID                `gorm:"type:uuid" json:"parentId"`
	Parent                *Organization             `gorm:"foreignKey:TParentId;" json:"parent"`
	TAccessLevelId        *uuid.UUID                `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel           *AccessLevel              `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
	UniversalId           *string                   `gorm:"type:varchar(255)" json:"universalId"`
	Name                  string                    `gorm:"type:varchar(255)" json:"name"`
	Description           *string                   `gorm:"type:varchar(455)" json:"description"`
	TPembangkit2Id        *uuid.UUID                `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit2           *Pembangkit               `gorm:"foreignKey:TPembangkit2Id;" json:"pembangkit"`
	TPemasok2Id           *uuid.UUID                `gorm:"type:uuid" json:"pemasokId"`
	Pemasok2              *Pemasok                  `gorm:"foreignKey:TPemasok2Id;" json:"pemasok"`
	TSurveyor2Id          *uuid.UUID                `gorm:"type:uuid" json:"surveyorId"`
	Surveyor2             *MasterEpiSurveyor        `gorm:"foreignKey:TSurveyor2Id;" json:"surveyor"`
	TKantorInduk2Id       *uuid.UUID                `gorm:"type:uuid" json:"kantorIndukId"`
	KantorInduk2          *KantorInduk              `gorm:"foreignKey:TKantorInduk2Id;" json:"kantorInduk"`
	TWilayah2Id           *uuid.UUID                `gorm:"type:uuid" json:"wilayahId"`
	Wilayah2              *Wilayah                  `gorm:"foreignKey:TWilayah2Id;" json:"wilayah"`
	TTransportir2Id       *uuid.UUID                `gorm:"type:uuid" json:"transportirId"`
	Transportir2          *MasterEpiTransportir     `gorm:"foreignKey:TTransportir2Id;" json:"transportir"`
	TRegional2Id          *uuid.UUID                `gorm:"type:uuid" json:"regionalId"`
	Regional2             *Regional                 `gorm:"foreignKey:TRegional2Id;" json:"regional"`
	Children              []OrganizationInfo        `gorm:"foreignKey:TParentId;" json:"children"`
	AccessGroups          []AccessGroupOrganization `gorm:"foreignKey:TOrganizationId; constraint:OnDelete:CASCADE" json:"accessGroups"`
	PageActivity          []OrganizationAccess      `gorm:"foreignKey:TOrganizationId; constraint:OnDelete:CASCADE" json:"accessPages"`
	AccessLevels          []AccessLevel             `gorm:"foreignKey:TOrganizationId; constraint:OnDelete:CASCADE" json:"accessLevels"`
	RencanaPasokan        []RencanaPasokan          `gorm:"foreignKey:TOrganizationId; constraint:OnDelete:CASCADE" json:"rencanaPasokan"`
	EntryJadwalPemasok    []EntryJadwal             `gorm:"foreignKey:TPemasokId; constraint:OnDelete:CASCADE" json:"entryJadwalPemasok"`
	EntryJadwalPembangkit []EntryJadwal             `gorm:"foreignKey:TPembangkitId; constraint:OnDelete:CASCADE" json:"entryJadwalPembangkit"`
	// PelimpahanPembangkit  []PelimpahanPembangkit    `gorm:"foreignKey:TPembangkitId; constraint:OnDelete:CASCADE" json:"pelimpahanPembangkit"`
	Account []Account `gorm:"foreignKey:TOrganizationId; constraint:OnDelete:CASCADE" json:"-"`

	// Iteration for doc generator
	NoBonkar          int                                 `json:"-"`
	NoBastCif         int                                 `json:"-"`
	NoBaDk            int                                 `json:"-"`
	NoBaMuat          int                                 `json:"-"`
	NoBadk            int                                 `json:"-"`
	HopRencanaPasokan *RencanaPasokanProyeksiHopStokAkhir `gorm:"-" json:"hopRencanaPasokan"`
}

func (OrganizationInfo) TableName() string {
	return "t_organization"
}

type OrganizationInfo struct {
	baseapp.BaseModel
	Name           string            `gorm:"type:varchar(255)" json:"name"`
	Description    *string           `gorm:"type:varchar(255)" json:"description"`
	TParentId      *uuid.UUID        `gorm:"type:uuid" json:"parentId"`
	Parent         *OrganizationInfo `gorm:"foreignKey:TParentId;" json:"parent"`
	TAccessLevelId *uuid.UUID        `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel    *AccessLevel      `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}

func (OrganizationNameAndLevel) TableName() string {
	return "t_organization"
}

type OrganizationNameAndLevel struct {
	ID             uuid.UUID            `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	Name           string               `gorm:"type:varchar(255)" json:"name"`
	TAccessLevelId *uuid.UUID           `gorm:"type:uuid" json:"-"`
	AccessLevel    *AccessLevelNameOnly `gorm:"foreignKey:TAccessLevelId;" json:"accessLevel"`
}

func (OrganizationInfoAccessLevel) TableName() string {
	return "t_organization"
}

type OrganizationInfoAccessLevel struct {
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Code        string `gorm:"type:varchar(255)" json:"code"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

type AccessType string

const (
	AccessTypeUnset    AccessType = "unset"
	AccessTypeAllow    AccessType = "allow"
	AccessTypeDisallow AccessType = "disallow"
	AccessTypeInherit  AccessType = "inherit"
)

func (a AccessType) IsValid() bool {
	return a == AccessTypeUnset || a == AccessTypeAllow || a == AccessTypeDisallow || a == AccessTypeInherit
}

func (OrganizationAccess) TableName() string {
	return "t_organization_access"
}

type OrganizationAccess struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID           `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization        `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	TPageActivityId *uuid.UUID           `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity        `gorm:"foreignKey:TPageActivityId;" json:"pageActivity"`
	TParentId       *uuid.UUID           `gorm:"type:uuid" json:"parentId"`
	Parent          *OrganizationAccess  `gorm:"foreignKey:TParentId;" json:"parent"`
	AccessType      AccessType           `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
	Children        []OrganizationAccess `gorm:"foreignKey:TParentId;constraint:OnDelete:CASCADE" json:"children"`
}

type OrganizationAccessType = AccessType

const (
	OrganizationAccessTypeUnset    = AccessTypeUnset
	OrganizationAccessTypeAllow    = AccessTypeAllow
	OrganizationAccessTypeDisallow = AccessTypeDisallow
	OrganizationAccessTypeInherit  = AccessTypeInherit
)

func (OrganizationAccessGroup) TableName() string {
	return "t_organization_access_group"
}

type OrganizationAccessGroup struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID `gorm:"type:uuid" json:"organizationId"`
	TAccessGroupId  *uuid.UUID `gorm:"type:uuid" json:"accessGroupId"`
}

func (OrganizationAccessPage) TableName() string {
	return "t_organization_access_page"
}

type OrganizationAccessPage struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	TPageId         *uuid.UUID             `gorm:"type:uuid" json:"pageId"`
	AccessType      OrganizationAccessType `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
}

func (OrganizationAccessPageInfo) TableName() string {
	return "t_organization_access_page"
}

type OrganizationAccessPageInfo struct {
	baseapp.BaseModel
	TOrganizationId *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	AccessType      OrganizationAccessType `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
}

func (OrganizationAccessActivity) TableName() string {
	return "t_organization_access_activity"
}

type OrganizationAccessActivity struct {
	baseapp.BaseModel
	TOrganizationAccessPageId *uuid.UUID             `gorm:"type:uuid" json:"organizationAccessPageId"`
	TOrganizationId           *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	TPageId                   *uuid.UUID             `gorm:"type:uuid" json:"pageId"`
	TActivityId               *uuid.UUID             `gorm:"type:uuid" json:"activityId"`
	AccessType                OrganizationAccessType `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
}

func (OrganizationAccessActivityItem) TableName() string {
	return "t_organization_access_activity_item"
}

type OrganizationAccessActivityItem struct {
	baseapp.BaseModel
	TOrganizationAccessActivityId *uuid.UUID             `gorm:"type:uuid" json:"organizationAccessActivityId"`
	TOrganizationId               *uuid.UUID             `gorm:"type:uuid" json:"organizationId"`
	TPageId                       *uuid.UUID             `gorm:"type:uuid" json:"pageId"`
	TActivityId                   *uuid.UUID             `gorm:"type:uuid" json:"activityId"`
	TActivityItemId               *uuid.UUID             `gorm:"type:uuid" json:"activityItemId"`
	AccessType                    OrganizationAccessType `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
}
