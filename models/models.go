package models

import (
	"github.com/thavel/goban/pkg/database"
)

var (
	// Tables of SQL entities.
	Tables = []interface{}{
		new(User),
		new(Role),
		new(Team),
		new(Absence), new(Reason),
	}
	// FKeys are foreign keys.
	FKeys = []database.FKey{
		database.FKey{
			Model: &User{},
			Args:  [...]string{"role", "roles(name)", "SET NULL", "NO ACTION"},
		},
		database.FKey{
			Model: &Absence{},
			Args:  [...]string{"user_id", "users(id)", "CASCADE", "NO ACTION"},
		},
		database.FKey{
			Model: &Absence{},
			Args:  [...]string{"reason_id", "reasons(id)", "CASCADE", "NO ACTION"},
		},
	}
)

// Model is the base model for all data models.
type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}
