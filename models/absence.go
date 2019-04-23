package models

import (
	"encoding/json"
	"fmt"
	"time"

	validator "gopkg.in/validator.v2"
)

type Absence struct {
	Model
	UserID   uint      `json:"user"`
	ReasonID uint      `json:"reason"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

func (a *Absence) Validate() error {
	if a.From.After(a.To) {
		return fmt.Errorf("'from' before 'to'")
	}
	return validator.Validate(a)
}

func NewAbsence(data []byte) (*Absence, error) {
	var entity Absence
	if err := json.Unmarshal(data, &entity); err != nil {
		return nil, err
	}
	return &entity, entity.Validate()
}

type Reason struct {
	Model
	Label      string `json:"label" gorm:"not null" validate:"min=1,max=25"`
	Deductable bool   `json:"deductable"`
}

func (r *Reason) Validate() error {
	return validator.Validate(r)
}

func NewReason(data []byte) (*Reason, error) {
	var entity Reason
	if err := json.Unmarshal(data, &entity); err != nil {
		return nil, err
	}
	return &entity, entity.Validate()
}
