package valueobject

import (
	"fmt"
	"log/slog"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Role struct {
	LiteralBase[string]
}

func NewRoleFromString(src string) (Role, error) {
	e := Role{LiteralBase: newLiteralBase(src)}
	if err := e.Validate(); err != nil {
		return Role{}, err
	}
	return e, nil
}

func (e *Role) Value() string {
	return e.v
}

func (e *Role) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.v,
			validation.Required,
			validation.By(func(value interface{}) error {
				r, ok := value.(string)
				if !ok {
					slog.Error("Role", "invalid tyep, expected string but got %T, value:%v\n", value, value)
					return fmt.Errorf("invalid type, expected string but got %T", value)
				}
				if r != "admin" && r != "operator" {
					return fmt.Errorf("No such role: %v", r)
				}
				return nil
			}),
		),
	)
}
