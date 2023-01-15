package inventoryservice

import "github.com/satheeshds/sellerapp/pkg/models"

type IInventoryService interface {
	CreateInventory(models.Inventory) error
	BlockInventory(models.Inventory) error
	GetInventory(pid int) (models.Inventory, error)
}
