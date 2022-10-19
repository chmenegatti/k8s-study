package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()

	return scope.SetColumn("ID", id)
}
