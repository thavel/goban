package models

import (
	"github.com/thavel/goban/pkg/database"
)

var (
	// Tables of SQL entities.
	Tables = []interface{}{
		new(User),
		new(Role),
	}
	// FKeys are foreign keys.
	FKeys = []database.FKey{
		database.FKey{
			Model: &User{},
			Args:  [...]string{"role", "roles(name)", "SET NULL", "NO ACTION"},
		},
	}
)

// Model is the base model for all data models.
type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}
