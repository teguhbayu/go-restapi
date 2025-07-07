package models

import (
	"time"

	"github.com/google/uuid"
)

type Books struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"default:now()"`
	Title        string
	YearReleased int
	Synopsis     string
}
