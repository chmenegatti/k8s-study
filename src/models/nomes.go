package models

import (
	"time"

	"github.com/google/uuid"
)

type Nomes struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid()"`
	Nome      string    `json:"nome,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
