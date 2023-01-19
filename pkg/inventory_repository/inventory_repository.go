package inventoryrepository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

// InventoryRepository is a struct that contains a database connection.
// It is used to interact with the database and perform operations such as retrieving and storing data related to inventory.
type InventoryRepository struct {
	db *gorm.DB
}

// This function is part of the InventoryRepository struct and is used to create a new inventory item in the products table.
// It takes in an inventory item of type models.Inventory as an argument and creates it in the database using the db object from the InventoryRepository struct.
// The result of the creation is returned as an error.
func (i *InventoryRepository) Create(inventory models.Inventory) error {
	result := i.db.Table("products").Create(&inventory)
	return result.Error
}

//This function is part of the InventoryRepository and is used to read a single inventory item from the products table in a database.
// It takes an integer id as an argument and returns a models.Inventory object and an error.
// The function uses the First() method of the db object to query the products table for the item with the given id, and assigns it to the inventory variable.
// The result of this query is then returned along with any associated error.
func (i *InventoryRepository) Read(id int) (models.Inventory, error) {
	var inventory models.Inventory
	result := i.db.Table("products").First(&inventory, id)
	return inventory, result.Error
}

// This function is part of the InventoryRepository and is used to update the quantity of an inventory item.
// It takes in a models.Inventory object as a parameter and uses the gorm library to update the quantity in the "products" table.
// If there is an error, it returns the error. If no rows are affected, it returns an error saying that the product doesn't exist.
// Otherwise, it returns nil.
func (i *InventoryRepository) Update(inventory models.Inventory) error {
	result := i.db.Table("products").Model(&inventory).
		Update("quantity", gorm.Expr("quantity - ?", inventory.Quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product doesn't exist")
	}

	return nil
}

// This function establishes a connection to a PostgreSQL database using the GORM library.
// It prints out log messages to indicate the progress of the connection and returns an error if one occurs.
// The user, password, host, and database name are hardcoded into the function for simplicity.
func (i *InventoryRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	// Ideally configuration should be in config
	log.Printf("latest host no gorm.model")
	i.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// This function is a method of the InventoryRepository type. It closes the database connection associated with the InventoryRepository instance (i). The function takes a pointer to an InventoryRepository instance as its argument and calls the Close() method on the database object associated with it.
func (i *InventoryRepository) Close() {
	i.db.Close()
}
