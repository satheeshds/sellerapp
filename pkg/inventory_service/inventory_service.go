package inventoryservice

import (
	inventoryrepository "github.com/satheeshds/sellerapp/pkg/inventory_repository"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type InventoryService struct {
	Repository inventoryrepository.IInventoryRepository
}

func (i *InventoryService) CreateInventory(inventory models.Inventory) error {
	return i.Repository.Create(inventory)
}

func (i *InventoryService) BlockInventory(inventory models.Inventory) error {
	return i.Repository.Update(inventory)
}

func (i *InventoryService) GetInventory(pid int) (models.Inventory, error) {
	return i.Repository.Read(pid)
}
