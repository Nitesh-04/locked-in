package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMembership struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `json:"userId" gorm:"type:uuid;not null"`
	GroupID   uuid.UUID `json:"groupId" gorm:"type:uuid;not null"`
	Status    string    `json:"status" gorm:"type:text;default:'offline';check:status IN ('studying', 'on_break', 'offline')"`
	JoinedAt  time.Time `json:"joinedAt" gorm:"autoCreateTime"`

	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Group Group `json:"group" gorm:"foreignKey:GroupID"`
}
