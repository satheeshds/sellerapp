package inventoryrepository

import "github.com/satheeshds/sellerapp/pkg/models"

// This is an interface for an Inventory Repository. It contains methods for creating, reading, updating, opening and closing an inventory.
// The Create() method takes a models.Inventory as a parameter and returns an error.
// The Read() method takes an int as a parameter and returns a models.Inventory and an error.
// The Update() method takes a models.Inventory as a parameter and returns an error.
// The Open() method does not take any parameters and returns an error.
// The Close() method does not take any parameters and does not return anything.
type IInventoryRepository interface {
	Create(models.Inventory) error
	Read(int) (models.Inventory, error)
	Update(models.Inventory) error
	Open() error
	Close()
}
