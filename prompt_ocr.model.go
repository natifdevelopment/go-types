package types
import (
	"time"

	"github.com/google/uuid"
)

func (PromptOcr) TableName() string {
	return "t_prompt_ocr"
}

type PromptOcr struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null;unique" json:"name"`
	Prompt    string     `gorm:"type:text;not null" json:"prompt"`
	IsActive  bool       `gorm:"type:bool;default:true" json:"isActive"`
	CreatedAt time.Time  `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
