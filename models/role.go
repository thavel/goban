package models

import (
	"encoding/json"

	validator "gopkg.in/validator.v2"
)

// Role data model.
type Role struct {
	Model
	Name string `gorm:"not null;unique" validate:"min=3,max=25" json:"name"`
}

// Validate a role.
func (r *Role) Validate() error {
	return validator.Validate(r)
}

// NewRole creates a new role data model instance.
func NewRole(data []byte) (*Role, error) {
	var role Role
	if err := json.Unmarshal(data, &role); err != nil {
		return nil, err
	}
	return &role, role.Validate()
}
