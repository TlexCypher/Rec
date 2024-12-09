package infrastructure

import (
	"Vox/db"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type InventoryRepositoryImpl struct {
	db *db.DB
}

func NewInventoryRepositoryImpl(db *db.DB) InventoryRepositoryImpl {
	return InventoryRepositoryImpl{
		db: db,
	}
}

func (ir *InventoryRepositoryImpl) Save(e entity.Inventory) error {
	return nil
}

func (ir *InventoryRepositoryImpl) FindById(inventoryId valueobject.InventoryId) entity.Inventory {
	return entity.Inventory{}
}

func (ir *InventoryRepositoryImpl) FindAll() []entity.Inventory {
	return []entity.Inventory{}
}

func (ir *InventoryRepositoryImpl) Update(e entity.Inventory) error {
	return nil
}

func (ir *InventoryRepositoryImpl) DeleteById(inventoryId valueobject.InventoryId) error {
	return nil
}
