package models

import (
	"github.com/thavel/goban/pkg/database"
)

var (
	// Tables of SQL entities.
	Tables = []interface{}{
		new(User),
	}
	// FKeys are foreign keys.
	FKeys = []database.FKey{}
)

// Model is the base model for all data models.
type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}
