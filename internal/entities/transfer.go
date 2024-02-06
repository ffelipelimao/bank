package entities

import (
	"errors"
	"time"
)

var (
	ErrInvalidValue       = errors.New("invalid value")
	ErrInvalidType        = errors.New("invalid type")
	ErrInvalidDescription = errors.New("invalid description")
)

type Transfer struct {
	Value       int64     `json:"valor"`
	UserID      int64     `json:"cliente_id,omitempty"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func (t *Transfer) Validate() error {
	if t.Value < 0 {
		return ErrInvalidValue
	}

	if t.Type != "c" && t.Type != "d" {
		return ErrInvalidType
	}

	if len(t.Description) > 10 {
		return ErrInvalidDescription
	}

	return nil
}

func (t *Transfer) IsDebit() bool {
	return t.Type == "d"
}
