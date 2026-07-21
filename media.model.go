package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"time"
"github.com/google/uuid"
)

func (Media) TableName() string {
	return "t_media"
}

type Media struct {
	baseapp.BaseModel
	Name         string `gorm:"type:varchar(245)" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	Filetype     string `gorm:"type:varchar(255)" json:"filetype"`
	Mimetype     string `gorm:"type:varchar(145)" json:"mimetype"`
	Storage      string `gorm:"type:varchar(45)" json:"storage"`
	Path         string `gorm:"type:varchar(255)" json:"path"`
	Extentsion   string `gorm:"type:varchar(10)" json:"extentsion"`
	Url          string `gorm:"type:varchar(455)" json:"url"`
	DownloadUrl  string `gorm:"type:varchar(455)" json:"downloadUrl"`
	Filename     string `gorm:"type:varchar(245)" json:"filename"`
	Filesize     int64  `gorm:"type:bigint;default:0" json:"filesize"`
	Tmp          string `gorm:"type:varchar(245)" json:"tmp"`
	OriginalName string `gorm:"type:varchar(245)" json:"originalName"`

	// PDF Metadata
	PdfCreator      *string              `gorm:"type:varchar(655)" json:"pdfCreator"`
	PdfProducer     *string              `gorm:"type:varchar(655)" json:"pdfProducer"`
	PdfCreatedDate  *DateTimeField `gorm:"type:timestamp" json:"pdfCreatedDate"`
	PdfModifiedDate *DateTimeField `gorm:"type:timestamp" json:"pdfModifiedDate"`

	AppSwitcher        []AppSwitcher        `gorm:"foreignKey:TMediaId;constraint:OnDelete:CASCADE"`
	DokumenKolaborator []DokumenKolaborator `gorm:"foreignKey:TMediaId;constraint:OnDelete:CASCADE" json:"kokumenKolaborator"`
}

func (MediaSimple) TableName() string {
	return "t_media"
}

type MediaSimple struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(245)" json:"name"`
	Url         string    `gorm:"type:varchar(455)" json:"url"`
	DownloadUrl string    `gorm:"type:varchar(455)" json:"downloadUrl"`
}

type TipeKolaborator string

const (
	TipeKolaboratorMaker   TipeKolaborator = "MAKER"
	TipeKolaboratorChecker TipeKolaborator = "CHECKER"
)

func (o TipeKolaborator) IsValid() bool {
	return o == TipeKolaboratorMaker || o == TipeKolaboratorChecker
}

func (DokumenKolaborator) TableName() string {
	return "t_dokumen_kolaborator"
}

type DokumenKolaborator struct {
	baseapp.BaseModel
	TMediaId        *uuid.UUID      `gorm:"type:uuid" json:"mediaId"`
	Media           *Media          `gorm:"foreignKey:TMediaId;" json:"media"`
	TAccountId      *uuid.UUID      `gorm:"type:uuid" json:"accountId"`
	Account         *Account        `gorm:"foreignKey:TAccountId" json:"account"`
	TOrganizationId *uuid.UUID      `gorm:"type:uuid" json:"organizationId"`
	Organization    *Organization   `gorm:"foreignKey:TOrganizationId" json:"organization"`
	TAccessLevelId  *uuid.UUID      `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevel    `gorm:"foreignKey:TAccessLevelId" json:"accessLevel"`
	TipeKolaborasi  TipeKolaborator `gorm:"type:varchar(255)" json:"tipeKolaborasi"`
}

func (MediaAccessToken) TableName() string {
	return "t_media_access_token"
}

type MediaAccessToken struct {
	baseapp.BaseModel
	TMediaId    *uuid.UUID `gorm:"type:uuid" json:"mediaId"`
	Media       *Media     `gorm:"foreignKey:TMediaId;" json:"media"`
	AccessToken string     `gorm:"type:varchar(255)" json:"accessToken"`
	ExpiredAt   time.Time  `gorm:"type:timestamp" json:"expiredAt"`
}
