package inventoryrepository

import "github.com/satheeshds/sellerapp/pkg/models"

type InventoryRepository struct {
}

func (i *InventoryRepository) Create(inventory models.Inventory) error {
	return nil
}

func (i *InventoryRepository) Read(pid int) (models.Inventory, error) {
	return models.Inventory{}, nil
}

func (i *InventoryRepository) Update(inventory models.Inventory) error {
	return nil
}

func (i *InventoryRepository) Open() error {
	return nil
}

func (i *InventoryRepository) Close() {

}
