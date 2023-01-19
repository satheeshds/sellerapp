package inventoryservice

import (
	inventoryrepository "github.com/satheeshds/sellerapp/pkg/inventory_repository"
	"github.com/satheeshds/sellerapp/pkg/models"
)

// InventoryService is a struct that contains a reference to an inventory repository.
// It provides methods for interacting with the repository.
type InventoryService struct {
	Repository inventoryrepository.IInventoryRepository
}

// This function is part of the InventoryService and is used to create an inventory item.
// It takes in a parameter of type models.Inventory and returns an error. The function calls the Create() method of the Repository field of the InventoryService struct.
func (i *InventoryService) CreateInventory(inventory models.Inventory) error {
	return i.Repository.Create(inventory)
}

// This function is part of an InventoryService and is used to block inventory.
// It takes a models.Inventory as an argument and returns an error.
// The function calls the Update method on the Repository associated with the InventoryService to update the inventory.
func (i *InventoryService) BlockInventory(inventory models.Inventory) error {
	return i.Repository.Update(inventory)
}

// This function is part of the InventoryService type and is used to get the inventory for a given product ID (pid).
// It calls the Read() method on the Repository field of the InventoryService type, passing in the pid as an argument.
// The function returns a models.Inventory object and an error, which will be nil if there are no errors.
func (i *InventoryService) GetInventory(pid int) (models.Inventory, error) {
	return i.Repository.Read(pid)
}
