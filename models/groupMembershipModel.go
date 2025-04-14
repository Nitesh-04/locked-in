package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMembership struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string    `json:"userId" gorm:"type:text;not null"`
	GroupID   uuid.UUID `json:"groupId" gorm:"type:uuid;not null"`
	JoinedAt  time.Time `json:"joinedAt" gorm:"autoCreateTime"`

	User  User  `json:"user" gorm:"foreignKey:UserID;references:ClerkID"`
	Group Group `json:"group" gorm:"foreignKey:GroupID"`
}