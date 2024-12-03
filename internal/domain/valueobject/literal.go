package valueobject

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log/slog"
)

type LiteralOnly interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | []byte
}

type LiteralBase[T LiteralOnly] struct {
	v T
}

func newLiteralBase[T LiteralOnly](value T) LiteralBase[T] {
	return LiteralBase[T]{
		v: value,
	}
}

func (e *LiteralBase[T]) Scan(src interface{}) error {
	switch v := src.(type) {
	case T:
		e.v = v
	default:
		slog.Error("LiteralBase[T]", "Failed to scan value.", fmt.Sprintf("Value:%v\n", v))
		return fmt.Errorf("Failed to scan value:%v\n", v)
	}
	return nil
}

func (e LiteralBase[T]) Value() (driver.Value, error) {
	return json.Marshal(e.v)
}

func (e *LiteralBase[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.v)
}

func (e LiteralBase[T]) UnmarhsalJSON(data []byte) error {
	var tmp T
	if err := json.Unmarshal(data, &tmp); err != nil {
		slog.Error("LiteralBase[T]", "Failed to unmarshal value.", fmt.Sprintf("Value:%v\n", data))
		return fmt.Errorf("Failed to unmarshal LiteralBase[T]: %w", err)
	}
	e.v = tmp
	return nil
}
