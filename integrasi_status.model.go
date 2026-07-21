package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type StatusEnum string
type MethodEnum string
type ResponseStatus string

const (
	StatusEnumSuccess StatusEnum = "SUCCESS"
	StatusEnumFailed  StatusEnum = "FAILED"
	StatusEnumError   StatusEnum = "ERROR"
	StatusEnumTimeout StatusEnum = "TIMEOUT"
	StatusEnumSkipped StatusEnum = "SKIPPED"
)

const (
	ResponseStatusCreated         ResponseStatus = "CREATED"
	ResponseStatusWaitingResponse ResponseStatus = "WAITING_REPSONSE"
	ResponseStatusSuccess         ResponseStatus = "SUCCESS"
	ResponseStatusError           ResponseStatus = "ERROR"
	ResponseStatusDeleted         ResponseStatus = "DELETED"
)

const (
	MethodPOST   MethodEnum = "POST"
	MethodGET    MethodEnum = "GET"
	MethodPUT    MethodEnum = "PUT"
	MethodDELETE MethodEnum = "DELETE"
	MethodPATCH  MethodEnum = "PATCH"
)

func (IntegrasiStatus) TableName() string {
	return "t_integrasi_status"
}

type IntegrasiStatus struct {
	baseapp.BaseModel
	TDataConnectorId          *uuid.UUID              `gorm:"type:uuid" json:"dataConnectorId"`
	DataConnector             *DataConnector          `gorm:"foreignKey:TDataConnectorId" json:"dataConnector"`
	BoundingCode              *string                 `gorm:"type:varchar(255)" json:"boundingCode"`
	TPembangkitId             *uuid.UUID              `gorm:"type:uuid" json:"pembangkitId"`
	Pembangkit                *Organization           `gorm:"foreignKey:TPembangkitId" json:"pembangkit"`
	PageCode                  *string                 `gorm:"type:varchar(255)" json:"pageCode"`
	TParentId                 *uuid.UUID              `gorm:"type:uuid" json:"parentId"`
	Parent                    *IntegrasiStatus        `gorm:"foreignKey:TParentId" json:"parent"`
	TargetId                  *uuid.UUID              `gorm:"type:uuid" json:"targetId"`
	TUserId                   *uuid.UUID              `gorm:"type:uuid" json:"userId"`
	User                      *Account                `gorm:"foreignKey:TUserId" json:"user"`
	TOrganizationId           *uuid.UUID              `gorm:"type:uuid" json:"organizationId"`
	Organization              *Organization           `gorm:"foreignKey:TOrganizationId" json:"organization"`
	ServiceId                 *string                 `gorm:"type:varchar(255);default:'0'" json:"serviceId"`
	Method                    MethodEnum              `gorm:"enum('POST','GET','PUT','DELETE','PATCH')" json:"method"`
	Url                       string                  `gorm:"type:text" json:"url"`
	Balance                   float64                 `gorm:"type:float" json:"balance"`
	ResponseStatus            ResponseStatus          `gorm:"type:varchar(255);default:'WAITING_REPSONSE'" json:"responseStatus"`
	ResponseMessage           *string                 `gorm:"type:text" json:"responseMessage"`
	Status                    StatusEnum              `gorm:"enum('SUCCESS','FAILED','TIMEOUT')" json:"status"`
	StatusCode                int                     `gorm:"type:int" json:"statusCode"`
	TrigeredBy                *string                 `gorm:"type:varchar(255)" json:"trigeredBy"`
	ErrorMessage              *string                 `gorm:"type:text" json:"errorMessage"`
	ResponseData              *json.RawMessage        `gorm:"type:json" json:"responseData"`
	Headers                   *json.RawMessage        `gorm:"type:json" json:"headers"`
	RequestData               *json.RawMessage        `gorm:"type:json" json:"requestData"`
	ProcessingTimeMs          int64                   `gorm:"type:bigint" json:"processingTimeMs"`
	BytesSize                 int64                   `gorm:"type:bigint" json:"bytesSize"`
	TIntegrasiStatusSessionId *uuid.UUID              `gorm:"type:uuid" json:"integrasiStatusSessionId"`
	IntegrasiStatusSession    *IntegrasiStatusSession `gorm:"foreignKey:TIntegrasiStatusSessionId" json:"integrasiStatusSession"`
	Childs                    []IntegrasiStatus       `gorm:"foreignKey:TParentId" json:"childs"`
	Items                     []IntegrasiStatusItem   `gorm:"foreignKey:TIntegrasiStatusId" json:"items"`
}

func (IntegrasiStatusItem) TableName() string {
	return "t_integrasi_status_item"
}

type IntegrasiStatusItem struct {
	baseapp.BaseModel
	TIntegrasiStatusId *uuid.UUID       `gorm:"type:uuid" json:"integrasiStatusId"`
	IntegrasiStatus    *IntegrasiStatus `gorm:"foreignKey:TIntegrasiStatusId" json:"integrasiStatus"`
	TResponseId        *uuid.UUID       `gorm:"type:uuid" json:"responseId"`
	Response           *IntegrasiStatus `gorm:"foreignKey:TResponseId" json:"response"`
	RefId              *uuid.UUID       `gorm:"type:uuid" json:"refId"`
	ResponseStatus     ResponseStatus   `gorm:"type:varchar(255);default:'WAITING_REPSONSE'" json:"responseStatus"`
	ResponseMessage    *string          `gorm:"type:text" json:"responseMessage"`
	ResponseData       *json.RawMessage `gorm:"type:json" json:"responseData"`
	RequestData        *json.RawMessage `gorm:"type:json" json:"requestData"`
}

func (IntegrasiStatusSession) TableName() string {
	return "t_integrasi_status_session"
}

type IntegrasiStatusSession struct {
	baseapp.BaseModel
	SessionCode string     `gorm:"type:varchar(255)" json:"sessionCode"`
	Session     string     `gorm:"type:varchar(255)" json:"session"`
	ExpiredAt   time.Time  `gorm:"type:timestamp" json:"expiredAt"`
	IsUsed      bool       `gorm:"type:boolean" json:"isUsed"`
	TargetId    *uuid.UUID `gorm:"type:uuid" json:"targetId"`
}
