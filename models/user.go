package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"

	"github.com/thavel/goban/pkg/crypto"
	validator "gopkg.in/validator.v2"
)

type User struct {
	Model
	Email     string  `json:"email" gorm:"not null;unique" validate:"email"`
	Password  string  `json:"-" gorm:"not null"`
	Role      *string `json:"role"`
	Firstname *string `json:"firstname" validate:"min=1,max=25"`
	Lastname  *string `json:"lastname" validate:"min=1,max=25"`
}

func init() {
	regex := regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
	validator.SetValidationFunc(
		"email",
		func(v interface{}, param string) error {
			st := reflect.ValueOf(v)
			if st.Kind() != reflect.String {
				return validator.ErrUnsupported
			}
			if !regex.MatchString(st.String()) {
				return fmt.Errorf("invalid email format")
			}
			return nil
		},
	)
}

// BeforeSave runs before insert or update.
func (o *User) BeforeSave() error {
	if !crypto.IsHash(o.Password) {
		o.Password = crypto.Hash(o.Password)
	}
	return nil
}

func (o *User) Validate() error {
	if s := len(o.Password); !crypto.IsHash(o.Password) && (s < 6 || s > 50) {
		return fmt.Errorf("invalid password size")
	}
	return validator.Validate(o)
}

func NewUser(data []byte) (*User, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	var entity User
	if err := json.Unmarshal(data, &entity); err != nil {
		return nil, err
	}
	password, ok := raw["password"]
	if !ok {
		return nil, fmt.Errorf("missing password")
	}
	entity.Password = password.(string)

	return &entity, entity.Validate()
}
