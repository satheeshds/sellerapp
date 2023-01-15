package inventoryservice

import (
	inventoryrepository "github.com/satheeshds/sellerapp/pkg/inventory_repository"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type InventoryService struct {
	Repository inventoryrepository.IInventoryRepository
}

func (i *InventoryService) CreateInventory(inventory models.Inventory) error {
	return nil
}

func (i *InventoryService) BlockInventory(inventory models.Inventory) error {
	return nil
}

func (i *InventoryService) GetInventory(pid int) (models.Inventory, error) {
	return models.Inventory{}, nil
}
