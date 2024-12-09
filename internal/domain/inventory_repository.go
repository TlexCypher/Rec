package domain

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type InventoryRepository interface {
	Save(entity.Inventory) error
	FindById(valueobject.InventoryId) entity.Inventory
	FindAll() []entity.Inventory
	Update(entity.Inventory) error
	DeleteById(valueobject.InventoryId) error
}
