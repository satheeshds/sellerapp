package inventoryrepository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type InventoryRepository struct {
	db *gorm.DB
}

func (i *InventoryRepository) Create(inventory models.Inventory) error {
	result := i.db.Table("products").Create(&inventory)
	return result.Error
}

func (i *InventoryRepository) Read(id int) (models.Inventory, error) {
	var inventory models.Inventory
	result := i.db.Table("products").First(&inventory, id)
	return inventory, result.Error
}

func (i *InventoryRepository) Update(inventory models.Inventory) error {
	result := i.db.Table("products").Model(&inventory).Update(inventory)

	return result.Error
}

func (i *InventoryRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	// Ideally configuration should be in config
	log.Printf("latest host no gorm.model")
	i.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (i *InventoryRepository) Close() {
	i.db.Close()
}
