package inventoryrepository

import "github.com/satheeshds/sellerapp/pkg/models"

type InventoryService struct {
}

func (i *InventoryService) Create(inventory models.Inventory) error {
	return nil
}

func (i *InventoryService) Read(pid int) (models.Inventory, error) {
	return models.Inventory{}, nil
}

func (i *InventoryService) Update(inventory models.Inventory) error {
	return nil
}

func (i *InventoryService) Open() error {
	return nil
}

func (i *InventoryService) Close() {

}
