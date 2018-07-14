package validation

import (
	"fmt"
	"reflect"

	validator "gopkg.in/validator.v2"
)

func NewValidator(name string, registry interface{}) {
	validator.SetValidationFunc(
		name,
		func(v interface{}, param string) error {
			st := reflect.ValueOf(v)
			if st.Kind() == reflect.Ptr {
				if st.Pointer() == 0 {
					return nil
				}
				st = st.Elem()
			}
			if st.Kind() != reflect.String {
				return validator.ErrUnsupported
			}
			value := st.String()
			switch reflect.TypeOf(registry).Kind() {
			case reflect.Slice:
				enum := reflect.ValueOf(registry)
				for i := 0; i < enum.Len(); i++ {
					if value == enum.Index(i).String() {
						return nil
					}
				}
			}
			return fmt.Errorf("unknown value '%s'", value)
		},
	)
}
