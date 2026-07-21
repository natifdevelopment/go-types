package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"github.com/google/uuid"
)

func (Page) TableName() string {
	return "t_page"
}

type Page struct {
	baseapp.BaseModel
	Name         string       `gorm:"type:varchar(255)" json:"name"`
	Code         string       `gorm:"type:varchar(255)" json:"code"`
	Url          string       `gorm:"type:varchar(255)" json:"url"`
	Description  string       `gorm:"type:varchar(255)" json:"description"`
	TDirectoryId uuid.UUID    `gorm:"type:uuid" json:"directoryId"`
	Directory    *Directory   `gorm:"foreignKey:TDirectoryId" json:"page"`
	PageAccess   []PageAccess `gorm:"foreignKey:TPageId;constraint:OnDelete:CASCADE" json:"pageAccess"`
	Activity     []Activity   `gorm:"foreignKey:TPageId;constraint:OnDelete:CASCADE" json:"activities"`
}

func (PageInfo) TableName() string {
	return "t_page"
}

type PageInfo struct {
	baseapp.BaseModel
	Uuid        string `gorm:"type:varchar(85)" json:"uuid"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Code        string `gorm:"type:varchar(255);unique" json:"code"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

type PageAccessType string

const (
	pageAccessType1 PageAccessType = "organization"
	pageAccessType2 PageAccessType = "account"
)

func (o PageAccessType) IsValid() bool {
	return o == pageAccessType1 || o == pageAccessType2
}

func (o PageAccessType) IsOrganization() bool {
	return o == pageAccessType1
}

func (o PageAccessType) IsAccount() bool {
	return o == pageAccessType2
}

func (PageAccess) TableName() string {
	return "t_page_access"
}

type PageAccess struct {
	baseapp.BaseModel
	Type            PageAccessType `gorm:"type:varchar(50)" json:"type"`
	TPageId         *uuid.UUID     `gorm:"type:uuid" json:"pageId"`
	TOrganizationId *uuid.UUID     `gorm:"type:uuid" json:"organizationId"`
	Organization    Organization   `gorm:"foreignKey:TOrganizationId;" json:"organization"`
	TAccountId      *uuid.UUID     `gorm:"type:uuid" json:"accountId"`
	Account         Account        `gorm:"foreignKey:TAccountId;" json:"account"`
}
