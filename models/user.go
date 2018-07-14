package models

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"

	validator "gopkg.in/validator.v2"
)

type User struct {
	Model
	Email     string   `json:"email" gorm:"not null;unique" validate:"email"`
	Security  Security `json:"security" gorm:"embedded"`
	Role      string   `json:"role"`
	Firstname *string  `json:"firstname" validate:"min=1,max=25"`
	Lastname  *string  `json:"lastname" validate:"min=1,max=25"`
}

type Security struct {
	Password string `json:"-"`
	Reset    bool   `json:"reset"`
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
				return errors.New("invalid email format")
			}
			return nil
		},
	)
}

func (o *User) Validate() error {
	return validator.Validate(o)
}

func NewUser(data []byte) (*User, error) {
	entity := &User{}
	if err := json.Unmarshal(data, entity); err != nil {
		return nil, err
	}
	return entity, entity.Validate()
}
