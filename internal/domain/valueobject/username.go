package valueobject

import (
	"log/slog"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Username struct {
	LiteralBase[string]
}

func (e *Username) Value() string {
	return e.v
}

func NewUsernameFromString(src string) (Username, error) {
	e := Username{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := e.Validate(); err != nil {
		slog.Error("Username", "Failed to validate username.", err)
		return Username{}, err
	}
	return e, nil
}

func (e *Username) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(validation.Required, validation.Length(6, 15)),
	)
}
