package types
import (
	baseapp "github.com/natifdevelopment/go-baseapp"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Data Connector Type
type ConnectorType string

const (
	ConnectorTypeApi         ConnectorType = "api"
	ConnectorTypePullRequest ConnectorType = "pull_request"
	ConnectorTypeKafka       ConnectorType = "kafka"
	ConnectorTypeMqtt        ConnectorType = "mqtt"
	ConnectorTypeGoogleSheet ConnectorType = "google_sheet"
	ConnectorTypeExcel       ConnectorType = "excel"
)

func (ct ConnectorType) IsValid() bool {
	return ct == ConnectorTypeApi ||
		ct == ConnectorTypePullRequest ||
		ct == ConnectorTypeKafka ||
		ct == ConnectorTypeGoogleSheet ||
		ct == ConnectorTypeExcel ||
		ct == ConnectorTypeMqtt
}

func (ct ConnectorType) String() string {
	return string(ct)
}

// Data Connector Method
type ConnectorMethod string

const (
	ConnectorMethodGET   ConnectorMethod = "GET"
	ConnectorMethodPOST  ConnectorMethod = "POST"
	ConnectorMethodPUT   ConnectorMethod = "PUT"
	ConnectorMethodPATCH ConnectorMethod = "PATCH"
)

func (cm ConnectorMethod) String() string {
	return string(cm)
}

func (cm ConnectorMethod) IsValid() bool {
	return cm == ConnectorMethodGET ||
		cm == ConnectorMethodPOST ||
		cm == ConnectorMethodPUT ||
		cm == ConnectorMethodPATCH
}

// Data Connector Status
type ConnectorStatus string

const (
	ConnectorStatusActive   ConnectorStatus = "active"
	ConnectorStatusInactive ConnectorStatus = "inactive"
)

func (cs ConnectorStatus) IsValid() bool {
	return cs == ConnectorStatusActive || cs == ConnectorStatusInactive
}

func (o ConnectorStatus) String() string {
	return string(o)
}

// Data Connector AuthenticationType
type ConnectorAuthenticationType string

const (
	NoneAuthentication            ConnectorAuthenticationType = "NONE"
	BasicAuthentication           ConnectorAuthenticationType = "BASIC_AUTH"
	AccessTokenAuthentication     ConnectorAuthenticationType = "ACCESS_TOKEN"
	KafkaSaslPlainAuthentication  ConnectorAuthenticationType = "SASL_PLAIN"
	KafkaSaslSha256Authentication ConnectorAuthenticationType = "SASL_SCRAM256"
	KafkaSaslSha512Authentication ConnectorAuthenticationType = "SASL_SCRAM512"
)

func (at ConnectorAuthenticationType) IsValid() bool {
	switch at {
	case BasicAuthentication:
		return true
	default:
		return false
	}
}

// Data Connector Scheduled Type
type ConnectorScheduledType string

const (
	ConnectorScheduledMonthly   ConnectorScheduledType = "monthly"
	ConnectorScheduledWeekly    ConnectorScheduledType = "weekly"
	ConnectorScheduledDaily     ConnectorScheduledType = "daily"
	ConnectorScheduledRetention ConnectorScheduledType = "retention"
	ConnectorScheduledEvent     ConnectorScheduledType = "event"
)

func (at ConnectorScheduledType) IsValid() bool {
	switch at {
	case ConnectorScheduledMonthly, ConnectorScheduledWeekly,
		ConnectorScheduledDaily, ConnectorScheduledRetention, ConnectorScheduledEvent:
		return true
	default:
		return false
	}
}

// Data Connector DB Target
type KafkaOffset string

const (
	KafkaOffsetLatest KafkaOffset = "latest"
	KafkaOffsetOldest KafkaOffset = "oldest"
)

func (at KafkaOffset) IsValid() bool {
	return at == KafkaOffsetLatest || at == KafkaOffsetOldest
}

type KafkaType string

const (
	KafkaTypeProducer   KafkaType = "PRODUCER"
	KafkaTypeSubscriber KafkaType = "SUBSCRIBER"
)

func (at KafkaType) IsValid() bool {
	return at == KafkaTypeProducer || at == KafkaTypeSubscriber
}

func (DataConnector) TableName() string {
	return "t_data_connector"
}

type DataConnector struct {
	baseapp.BaseModel
	TipeKonektor         ConnectorType               `gorm:"type:varchar(50)" json:"type"`
	Name                 string                      `gorm:"type:varchar(255)" json:"name"`
	Code                 string                      `gorm:"type:varchar(255)" json:"code"`
	Description          *string                     `gorm:"type:text" json:"description"`
	AuthenticationType   ConnectorAuthenticationType `gorm:"type:varchar(100); default:'basic_authentication'" json:"authenticationType"`
	AuthUser             *string                     `gorm:"type:text" json:"authUser"`
	AuthPassword         *string                     `gorm:"type:text" json:"-"`
	AccessTokenMethod    *string                     `gorm:"type:varchar(12)" json:"accessTokenMethod"`
	AccessTokenUrl       *string                     `gorm:"type:varchar(455)" json:"accessTokenUrl"`
	AccessTokenPath      *string                     `gorm:"type:varchar(45)" json:"accessTokenPath"`
	AccessTokenHeader    *string                     `gorm:"type:varchar(45)" json:"accessTokenHeader"`
	AccessTokenUserKey   *string                     `gorm:"type:varchar(45)" json:"accessTokenUserKey"`
	AccessTokenPwdKey    *string                     `gorm:"type:varchar(45)" json:"accessTokenPwdKey"`
	ScheduledType        ConnectorScheduledType      `gorm:"type:varchar(50);" json:"scheduledType"` // Store type of scheduled (monthly, weekly, daily, retention)
	ScheduledDate        *int                        `gorm:"type:int" json:"scheduledDate"`          // Store day of month (1-31) -> for (monthly)
	ScheduledDay         *int                        `gorm:"type:int" json:"scheduledDay"`           // Store day of week (0-6 for Sunday-Saturday) -> for (weekly)
	ScheduledHour        *string                     `gorm:"type:varchar(45)" json:"scheduledHour"`  // Store hour and minute (HH:mm) -> for (monthly, weekly, daily)
	RetentionAt          *time.Time                  `gorm:"type:time" json:"retentionAt"`           // Checkpoint when the last time this retentian run
	RetentionEveryMinute *int                        `gorm:"type:int" json:"retentionEveryMinute"`   // Store hour and minute (HH:mm) -> for (retention), every selected time

	// Used for kafka and mqtt to detect wheter this data connector already subscribed to topic
	// this is not guarantee the, to check if success subscribe please refer to IsSubscribeSuccess
	IsSubscribed bool `gorm:"type:boolean;default:false" json:"isSubscribed"`

	// Set to true when confiration changes and need to restart
	// after successfully restared, it will set to false again
	IsRestartRequired bool `gorm:"type:boolean;default:false" json:"isRestartRequired"`

	// Detect if connection are subscribed successfully and ready to accept connection
	IsSubscribeSuccess bool `gorm:"type:boolean;default:false" json:"isSubscribeSuccess"`

	// Flag the service is available and can be run
	// default is false and if user stopped the service, this is set to false
	// until user start again the service
	IsAvailable bool `gorm:"type:boolean;default:false" json:"isAvailable"`

	// Flag that this is service is currently starting / restarting
	// so event stopped will be ignored until it come back online
	IsStarting bool `gorm:"type:boolean;default:false" json:"isStarting"`

	// TLS Configuration
	UseTls              bool       `gorm:"type:boolean;default:false" json:"useTls"`
	UseTlsPublicKey     bool       `gorm:"type:boolean;default:false" json:"useTlsPublicKey"`
	TlsServerName       *string    `gorm:"type:varchar(255)" json:"tlsServerName"`
	TTlsPublicKeyId     *uuid.UUID `gorm:"type:uuid" json:"tlsPublicKeyId"`
	TlsPublicKey        *Media     `gorm:"foreignKey:TTlsPublicKeyId" json:"tlsPublicKey"`
	TlsSkipVerification bool       `gorm:"type:boolean;default:false" json:"tlsSkipVerification"`

	IsError              bool                     `gorm:"type:boolean;default:false" json:"isError"`
	IsActive             bool                     `gorm:"type:boolean;default:false" json:"isActive"`
	KafkaType            *KafkaType               `gorm:"type:varchar(45)" json:"kafkaType"`
	ErrorMessage         *string                  `gorm:"type:text" json:"errorMessage"`
	ErrorCount           int                      `gorm:"type:int;default:0" json:"errorCount"`
	KafkaBrokerAddresses *string                  `gorm:"type:varchar(355)" json:"kafkaBrokerAddresses"`
	KafkaTopics          *string                  `gorm:"type:varchar(255)" json:"kafkaTopics"`
	KafkaConsumerGroup   *string                  `gorm:"type:varchar(75)" json:"kafkaConsumerGroup"`
	KafkaOffset          *KafkaOffset             `gorm:"type:varchar(45)" json:"kafkaOffset"`
	KafkaClient          *string                  `gorm:"type:varchar(75)" json:"kafkaClient"`
	MqttBrokerAddress    *string                  `gorm:"type:varchar(355)" json:"mqttBrokerAddress"`
	MqttTopics           *string                  `gorm:"type:varchar(255)" json:"mqttTopics"`
	MqttClient           *string                  `gorm:"type:varchar(75)" json:"mqttClient"`
	Priority             *int                     `gorm:"type:bigint" json:"priority"`
	HttpReqMethod        ConnectorMethod          `gorm:"type:varchar(50)" json:"httpReqMethod"`
	EndpointUrl          string                   `gorm:"type:varchar(455)" json:"endpointUrl"`
	RequestBody          json.RawMessage          `gorm:"type:json" json:"requestBody"`
	RequestInternval     int64                    `gorm:"type:bigint; default:1800" json:"requestInterval"`
	ConnectorStatus      ConnectorStatus          `gorm:"type:varchar(50)" json:"connectorStatus"`
	ApiKey               string                   `gorm:"type:varchar(255)" json:"-"`
	ApiKeyEnabled        bool                     `gorm:"boolean" json:"apiKeyEnabled"`
	ApiLoginEnabled      bool                     `gorm:"boolean" json:"apiLoginEnabled"`
	RequestHeaders       []DataConnectorReqHeader `gorm:"foreignKey:TDataConnectorId;constraint:OnDelete:CASCADE" json:"requestHeaders"`
}

func (DataConnectorReqHeader) TableName() string {
	return "t_data_connector_req_header"
}

type DataConnectorReqHeader struct {
	baseapp.BaseModel
	TDataConnectorId *uuid.UUID     `gorm:"type:uuid" json:"dataConnectorId"`
	DataConnector    *DataConnector `gorm:"foreignKey:TDataConnectorId" json:"dataConnector"`
	HeaderKey        *string        `gorm:"type:varchar(255)" json:"headerKey"`
	HeaderValue      *string        `gorm:"type:varchar(255)" json:"headerValue"`
}
