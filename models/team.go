package models

import (
	"encoding/json"

	validator "gopkg.in/validator.v2"
)

type Team struct {
	Model
	Name  string `json:"name" gorm:"not null;unique" validate:"min=3,max=25"`
	Size  uint   `json:"size" gorm:"-"`
	Users []User `json:"-"`
}

func (t *Team) Validate() error {
	return validator.Validate(t)
}

func NewTeam(data []byte) (*Team, error) {
	var entity Team
	if err := json.Unmarshal(data, &entity); err != nil {
		return nil, err
	}
	return &entity, entity.Validate()
}
