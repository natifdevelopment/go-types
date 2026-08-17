package types

import (
	"time"

	"github.com/google/uuid"
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (Account) TableName() string {
	return "t_account"
}

type SSOType string

const (
	TypeSSO        SSOType = "SSO"
	TypeClassicSSO SSOType = "CLASIC"
)

func (o SSOType) IsValid() bool {
	return o == TypeSSO || o == TypeClassicSSO
}

func (o SSOType) IsClassicSSO() bool {
	return o == TypeClassicSSO
}

type Account struct {
	baseapp.BaseModel
	Name                   EncryptedField    `gorm:"type:text" json:"name"`
	NameHash               string            `gorm:"type:varchar(455)"`
	Description            *string           `gorm:"type:text" json:"description"`
	Address                *EncryptedField   `gorm:"type:text" json:"address"`
	AddressHash            string            `gorm:"type:varchar(455)"`
	TOrganizationId        *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization           *OrganizationInfo `gorm:"foreignKey:TOrganizationId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"organization"`
	TAccessLevelId         *uuid.UUID        `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel            *AccessLevelInfo  `gorm:"foreignKey:TAccessLevelId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"accessLevel"`
	Email                  EncryptedField    `gorm:"type:text" json:"email"`
	EmailHash              string            `gorm:"type:varchar(455);unique"`
	PhoneNumber            *EncryptedField   `gorm:"type:text" json:"phoneNumber"`
	PhoneNumberHash        string            `gorm:"type:varchar(455)"`
	Password               string            `gorm:"type:varchar(255)"`
	PasswordExpiryDate     *time.Time        `gorm:"type:timestamp" json:"passwordExpiryDate"`
	Pin                    *string           `gorm:"type:varchar(255)"`
	BiometricEnabled       bool              `gorm:"type:boolean;default:false"`
	BiometricSession       *string           `gorm:"type:varchar(645)"`
	BiometricCredential    *string           `gorm:"type:varchar(645)"`
	IsPwdDefault           bool              `gorm:"type:boolean;default:false" json:"IsPwdDefault"`
	LastEmailVerified      *time.Time        `gorm:"type:timestamp" json:"lastEmailVerified"`
	IsUserActive           bool              `gorm:"type:boolean;default:true" json:"isUserActive"`
	Nip                    *EncryptedField   `gorm:"type:text" json:"nip"`
	NipHash                *string           `gorm:"type:varchar(455)"`
	SSOType                SSOType           `gorm:"type:varchar(50)" json:"ssoType"`
	SSOId                  *string           `gorm:"type:text" json:"ssoId"`
	IsBlockedPermanent     bool              `gorm:"type:boolean" json:"isBlockedPermanent"`
	BlockedPermanentAt     *time.Time        `gorm:"type:timestamp" json:"blockedPermanentAt"`
	BlockedPermanentReason *string           `gorm:"type:text" json:"blockedPermanentReason"`
	BlockedUntil           *time.Time        `gorm:"type:timestamp" json:"blockedUntil"`
	IsServiceAccount       bool              `gorm:"type:boolean;default:false" json:"isServiceAccount"`

	// TOTP / Authenticator app (RFC 6238) fields.
	// TotpSecret is AES-256-GCM encrypted at rest (configs.Encrypt) and only
	// decrypted in memory when validating a code. TotpEnabled is false until
	// the user completes setup by verifying a code from their authenticator.
	TotpSecret         *string    `gorm:"type:text" json:"-"` // encrypted base32 secret
	TotpEnabled        bool       `gorm:"type:boolean;default:false" json:"totpEnabled"`
	TotpVerifiedAt     *time.Time `gorm:"type:timestamp" json:"-"`     // when setup was confirmed
	TotpBackupCodes    *string    `gorm:"type:text" json:"-"`          // JSON array of argon2id-hashed codes
	TotpFailedAttempts int        `gorm:"type:int;default:0" json:"-"` // failed verify attempts (lockout)
	TotpLockedUntil    *time.Time `gorm:"type:timestamp" json:"-"`     // lockout expiry
	// OtpPreference controls which channel(s) are used at login.
	// "auto" (default): authenticator primary, email fallback.
	// "email_auto": email primary, authenticator fallback.
	// "authenticator": authenticator only (requires TotpEnabled).
	// "email": email only.
	OtpPreference string `gorm:"type:varchar(20);default:'auto'" json:"otpPreference"`

	// TotpPolicyAppliedAt records when the TOTP enrollment policy started
	// applying to this user. The 30-day grace period is counted from this
	// timestamp, NOT from CreatedAt. This allows existing users (created
	// before the policy was introduced) to get a fresh grace period.
	// If nil, falls back to CreatedAt for backward compatibility.
	TotpPolicyAppliedAt *time.Time `gorm:"type:timestamp" json:"-"`

	AccessGroup        []AccessGroupAccount `gorm:"foreignKey:TAccountId; constraint:OnDelete:CASCADE" json:"accessGroups"`
	PageActivity       []AccountAccess      `gorm:"foreignKey:TAccountId; constraint:OnDelete:CASCADE" json:"accessPages"`
	ApprovalCodeMember []ApprovalCodeMember `gorm:"foreignKey:TAccountId; constraint:OnDelete:CASCADE" json:"approvalCodes"`
	AccountSession     []AccountSession     `gorm:"foreignKey:TAccountId; constraint:OnDelete:CASCADE" json:"accountSession"`
	AccountTotpLogs    []AccountTotpLog     `gorm:"foreignKey:TAccountId; constraint:OnDelete:CASCADE" json:"-"`
}

func (AccountInfo) TableName() string {
	return "t_account"
}

type AccountInfo struct {
	baseapp.BaseModel
	Name            EncryptedField    `gorm:"type:varchar(255)" json:"name"`
	Code            string            `gorm:"type:varchar(255)" json:"code"`
	Description     string            `gorm:"type:varchar(255)" json:"description"`
	Email           EncryptedField    `gorm:"type:varchar(255)" json:"email"`
	TOrganizationId *uuid.UUID        `gorm:"type:uuid" json:"organizationId"`
	Organization    *OrganizationInfo `gorm:"foreignKey:TOrganizationId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"organization"`
	TAccessLevelId  *uuid.UUID        `gorm:"type:uuid" json:"accessLevelId"`
	AccessLevel     *AccessLevelInfo  `gorm:"foreignKey:TAccessLevelId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"accessLevel"`
}

func (AccountAccess) TableName() string {
	return "t_account_access"
}

type AccountAccess struct {
	baseapp.BaseModel
	TAccountId      uuid.UUID       `gorm:"type:uuid" json:"accountId"`
	Account         *Account        `gorm:"foreignKey:TAccountId;" json:"account"`
	TPageActivityId *uuid.UUID      `gorm:"type:uuid" json:"pageActivityId"`
	PageActivity    *PageActivity   `gorm:"foreignKey:TPageActivityId;" json:"pageActivity"`
	TParentId       *uuid.UUID      `gorm:"type:uuid" json:"parentId"`
	Parent          *AccountAccess  `gorm:"foreignKey:TParentId;" json:"parent"`
	AccessType      AccessType      `gorm:"type:varchar(255);default:'unset'" json:"accessType"`
	Children        []AccountAccess `gorm:"foreignKey:TParentId;constraint:OnDelete:CASCADE" json:"children"`
}

type OtpType string

const (
	OtpLogin      OtpType = "login"
	OtpForgotPass OtpType = "forgotpass"
	OtpGeneral    OtpType = "general"
)

func (o OtpType) IsValid() bool {
	return o == OtpLogin || o == OtpForgotPass
}

func (AccountOTP) TableName() string {
	return "t_account_otp"
}

type AccountOTP struct {
	baseapp.BaseModel
	Type         OtpType   `gorm:"type:varchar(50)" json:"type"`
	TAccountId   uuid.UUID `gorm:"type:uuid" json:"accountId"`
	Passcode     string    `gorm:"type:varchar(455)" json:"passCode"`
	PasscodeHash string    `gorm:"type:varchar(455)"`
	ExpiryDate   time.Time `gorm:"type:timestamp" json:"expiryDate"`
	IsUsed       bool      `gorm:"type:boolean;default:false" json:"isUsed"`
}

func (AccountPasswordHistory) TableName() string {
	return "t_account_password_history"
}

type AccountPasswordHistory struct {
	baseapp.BaseModel
	TAccountId uuid.UUID `gorm:"type:uuid" json:"accountId"`
	Password   string    `gorm:"type:varchar(255)"`
	ExpiryDate time.Time `gorm:"type:timestamp" json:"expiryDate"`
}

func (AccountLoginAttempt) TableName() string {
	return "t_account_login_attempt"
}

type AccountLoginAttempt struct {
	baseapp.BaseModel
	TAccountId   uuid.UUID  `gorm:"type:uuid" json:"accountId"`
	Account      *Account   `gorm:"foreignKey:TAccountId;" json:"account"`
	Attempt      int        `gorm:"type:int" json:"attempt"`
	LastAttempt  *time.Time `gorm:"type:timestamp" json:"lastAttempt"`
	BlockedUntil *time.Time `gorm:"type:timestamp" json:"blockedUntil"`
}

func (AccountActivationToken) TableName() string {
	return "t_account_activation_token"
}

type AccountActivationToken struct {
	baseapp.BaseModel
	TAccountId uuid.UUID `gorm:"type:uuid;not null" json:"accountId"`
	TokenHash  string    `gorm:"type:varchar(64);unique;not null"`
	ExpiryDate time.Time `gorm:"type:timestamp;not null" json:"expiryDate"`
	IsUsed     bool      `gorm:"type:boolean;default:false" json:"isUsed"`
}

func (AccountDevice) TableName() string {
	return "t_account_device"
}

type AccountDevice struct {
	baseapp.BaseModel
	TAccountId    uuid.UUID    `gorm:"type:uuid" json:"accountId"`
	Account       *AccountInfo `gorm:"foreignKey:TAccountId" json:"account"`
	UserAgent     string       `gorm:"type:text" json:"userAgent"`
	DeviceId      string       `gorm:"type:text" json:"deviceId"`
	AccessToken   *string      `gorm:"type:text" json:"accessToken"`
	RefreshToken  *string      `gorm:"type:text" json:"refreshToken"`
	LastSeen      time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastSeen"`
	LastLogin     time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastLogin"`
	LoggedOutDate *time.Time   `gorm:"type:timestamp;default:NULL" json:"loggedOutDate"`
}

// OtpPreference values control which verification channel(s) are used at login.
const (
	OtpPrefEmail         = "email"         // email OTP only
	OtpPrefAuthenticator = "authenticator" // authenticator app only (requires TotpEnabled)
	OtpPrefAuto          = "auto"          // authenticator primary, email fallback (DEFAULT)
	OtpPrefEmailAuto     = "email_auto"    // email primary, authenticator fallback
)

// IsValidOtpPreference returns true if the value is a recognised OTP preference.
func IsValidOtpPreference(v string) bool {
	return v == OtpPrefEmail || v == OtpPrefAuthenticator || v == OtpPrefAuto || v == OtpPrefEmailAuto
}

// TotpLogAction identifies the kind of TOTP operation recorded in the audit log.
type TotpLogAction string

const (
	TotpActionSetupInitiated TotpLogAction = "setup_initiated"
	TotpActionSetupCompleted TotpLogAction = "setup_completed"
	TotpActionSetupAbandoned TotpLogAction = "setup_abandoned"
	TotpActionDisabled       TotpLogAction = "disabled"
	TotpActionUsedLogin      TotpLogAction = "used_login"
	TotpActionUsedRecovery   TotpLogAction = "used_recovery"
	TotpActionVerifyFailed   TotpLogAction = "verify_failed"
	TotpActionBackupRegen    TotpLogAction = "backup_regen"
	TotpActionLockedOut      TotpLogAction = "locked_out"
)

func (AccountTotpLog) TableName() string {
	return "t_account_totp_log"
}

// AccountTotpLog records every TOTP-related event for audit/security purposes
// (OWASP A09:2021 — Security Logging and Monitoring). The code value itself is
// never logged; only the action and whether it succeeded.
type AccountTotpLog struct {
	baseapp.BaseModel
	TAccountId   uuid.UUID     `gorm:"type:uuid;index" json:"accountId"`
	Action       TotpLogAction `gorm:"type:varchar(50)" json:"action"`
	IpAddress    string        `gorm:"type:varchar(45)" json:"ipAddress"`
	UserAgent    string        `gorm:"type:varchar(512)" json:"userAgent"`
	Success      bool          `gorm:"type:boolean;default:false" json:"success"`
	ErrorMessage *string       `gorm:"type:text" json:"errorMessage,omitempty"`
}
