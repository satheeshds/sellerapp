package inventoryrepository

import "github.com/satheeshds/sellerapp/pkg/models"

type IInventoryRepository interface {
	Create(models.Inventory) error
	Read(int) (models.Inventory, error)
	Update(models.Inventory) error
	Open() error
	Close()
}
