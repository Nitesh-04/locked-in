package models

import (
	"time"

	"github.com/google/uuid"
)

type StudySession struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID  `json:"userId" gorm:"type:uuid;not null"`
	GroupID   uuid.UUID  `json:"groupId" gorm:"type:uuid;not null"`
	StartedAt time.Time  `json:"startedAt" gorm:"not null"`
	EndedAt   *time.Time `json:"endedAt" gorm:"default:null"`
	Duration  *int       `json:"duration" gorm:"default:null"`

	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Group Group `json:"group" gorm:"foreignKey:GroupID"`
}
