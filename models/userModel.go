package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClerkID   string    `json:"clerkID" gorm:"type:text;unique;not null"`
	Name 	  string    `json:"name" gorm:"type:text;not null"`
	Email 	  string    `json:"email" gorm:"type:text;unique;not null"`
	Status    string    `json:"status" gorm:"type:text;default:'offline';check:status IN ('studying', 'on_break', 'offline')"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
