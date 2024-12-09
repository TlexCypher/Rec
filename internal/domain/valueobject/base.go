package valueobject

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/gofrs/uuid/v5"
	"github.com/samber/lo"
)

/* This PrimaryIdBase struct is the base for primary-key.*/
type PrimaryIdBase struct {
	uuid.UUID
	value string
}

func newPrimaryIdBase() PrimaryIdBase {
	u := lo.Must(uuid.NewV7())
	return PrimaryIdBase{
		UUID:  u,
		value: u.String(),
	}
}

func newPrimaryIdBaseFromString(value string) (PrimaryIdBase, error) {
	u, err := uuid.FromString(value)
	if err != nil {
		return PrimaryIdBase{}, err
	}
	return PrimaryIdBase{
		UUID:  u,
		value: value,
	}, nil
}

func (e *PrimaryIdBase) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		u, err := uuid.FromString(v)
		if err != nil {
			return err
		}
		e.UUID = u
		e.value = v
		return nil
	case []byte:
		u, err := uuid.FromBytes(v)
		if err != nil {
			return err
		}
		e.UUID = u
		e.value = string(v)
		return nil
	default:
		slog.Error("PrimaryIdBase", "Failed to scan value", fmt.Sprintf("Value:%v\n", value))
		return fmt.Errorf("Failed to scan value\n")
	}
}

func (e PrimaryIdBase) Value() (driver.Value, error) {
	return json.Marshal(e.value)
}

func (e PrimaryIdBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func (e *PrimaryIdBase) UnmarshalJSON(data []byte) error {
	u, err := uuid.FromBytes(data)
	if err != nil {
		slog.Error("PrimaryIdBase", "Failed to unmarshal json", fmt.Sprintf("Value:%v\n", string(data)))
		return err
	}
	e.UUID = u
	e.value = string(data)
	return nil
}

type UserId struct {
	PrimaryIdBase
}

func NewUserId() (UserId, error) {
	u, err := uuid.NewV7()
	if err != nil {
		slog.Error("UserId", "Failed to generate new user id.", err)
		return UserId{}, err
	}
	return UserId{
		PrimaryIdBase{
			UUID:  u,
			value: u.String(),
		},
	}, nil
}

func NewUserIdFromString(src string) (UserId, error) {
	p, err := newPrimaryIdBaseFromString(src)
	if err != nil {
		slog.Error("UserId", "Failed to generate new userId.", err)
		return UserId{}, err
	}
	e := UserId{
		PrimaryIdBase{
			UUID:  p.UUID,
			value: p.value,
		},
	}
	if err := e.Validate(); err != nil {
		slog.Error("UserId", "Failed to validate new generated userId.", err)
		return UserId{}, err
	}
	return e, nil
}

/* UserId is just simple uuid. So, no validation rule is necessary.*/
func (u *UserId) Validate() error {
	return nil
}

type InventoryId struct {
	PrimaryIdBase
}

func NewInventoryId() (InventoryId, error) {
	i, err := uuid.NewV7()
	if err != nil {
		slog.Error("InventoryId", "Failed to generate new inventory id.", err)
		return InventoryId{}, err
	}
	return InventoryId{
		PrimaryIdBase{
			UUID:  i,
			value: i.String(),
		},
	}, nil
}

func NewInventoryIdFromString(src string) (InventoryId, error) {
	p, err := newPrimaryIdBaseFromString(src)
	if err != nil {
		slog.Error("InventoryId", "Failed to generate new inventoryId.", err)
		return InventoryId{}, err
	}
	e := InventoryId{
		PrimaryIdBase{
			UUID:  p.UUID,
			value: p.value,
		},
	}
	if err := e.Validate(); err != nil {
		slog.Error("InventoryId", "Failed to validate new generated inventoryId.", err)
		return InventoryId{}, err
	}
	return e, nil

}

type CategoryId struct {
	PrimaryIdBase
}

func (c *CategoryId) Validate() error {
	return nil
}

func NewCategoryId() (CategoryId, error) {
	c, err := uuid.NewV7()
	if err != nil {
		slog.Error("CategoryId", "Failed to generate new category id.", err)
		return CategoryId{}, err
	}
	return CategoryId{
		PrimaryIdBase{
			UUID:  c,
			value: c.String(),
		},
	}, nil
}

func NewCategoryIdFromString(src string) (CategoryId, error) {
	p, err := newPrimaryIdBaseFromString(src)
	if err != nil {
		slog.Error("CategoryId", "Failed to generate new category id.")
		return CategoryId{}, err
	}
	e := CategoryId{
		PrimaryIdBase{
			UUID:  p.UUID,
			value: p.value,
		},
	}
	if err := e.Validate(); err != nil {
		slog.Error("CategoryId", "Failed to validate new generated category id.")
		return CategoryId{}, err
	}
	return e, nil
}

/*InventoryId is just simple uuid. So, no validation rule is necesary.*/
func (i *InventoryId) Validate() error {
	return nil
}
