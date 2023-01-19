package inventoryservice

import "github.com/satheeshds/sellerapp/pkg/models"

// IInventoryService is an interface for managing inventory items.
// CreateInventory creates a new inventory item with the given models.Inventory object. It returns an error if the creation fails.
// BlockInventory blocks an existing inventory item with the given models.Inventory object. It returns an error if the blocking fails.
// GetInventory retrieves an existing inventory item with the given product ID (pid). It returns a models.Inventory object and any error encountered during retrieval.
type IInventoryService interface {
	CreateInventory(*models.Inventory) error
	BlockInventory(*models.Inventory) error
	GetInventory(pid int) (models.Inventory, error)
}
