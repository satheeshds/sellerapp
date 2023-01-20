package orderrepository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

// OrderRepository is a struct that contains a pointer to a gorm.DB object.
// This struct is used to store and retrieve data from the database.
type OrderRepository struct {
	db *gorm.DB
}

// This function is part of the OrderRepository and is used to open a connection to a PostgreSQL database.
// It uses the gorm library to open the connection with the given credentials (user, password, host, dbname, sslmode).
// If an error occurs while establishing the connection, it will be logged and returned.
func (o *OrderRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	o.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// This function closes the database connection associated with the OrderRepository object.
// It takes a pointer to an OrderRepository object as an argument and calls the Close() method on the database object associated with it.
func (o *OrderRepository) Close() {
	o.db.Close()
}

// This function creates an order in the database using the OrderRepository.
// It takes a pointer to a models.Order as an argument and returns an error if one occurs.
// The result of the database operation is stored in the variable "result" and its associated error is returned.
func (o *OrderRepository) Create(order *models.Order) error {
	result := o.db.Omit("status").Create(order)

	return result.Error
}

// This function is part of the OrderRepository struct and is used to read an order from the database.
// It takes an integer id as a parameter and returns a models.Order object and an error.
// The function first creates an empty models.Order object, then uses the db field of the OrderRepository struct to query for the order with the given id.
// The result of this query is stored in result, and then both the order object and result.Error are returned from the function.
func (o *OrderRepository) Read(id int) (models.Order, error) {
	var order models.Order
	result := o.db.First(&order, id)

	return order, result.Error
}

// This function updates an existing order in the database. It takes in an integer id and a models.Order object as parameters.
// It uses the db object to access the orders table and then uses the Model() and Update() functions to update the order with the given parameters.
// Finally, it returns any errors that may have occurred during the update process.
func (o *OrderRepository) Update(id int, order models.Order) error {
	result := o.db.Table("orders").Model(&order).Update(order)

	return result.Error
}

// This function deletes an order from the database based on its ID.
// It first reads the order from the database using the Read() function, then uses the Delete() function from the database to delete it.
// It returns an error if there is one.
func (o *OrderRepository) Delete(id int) error {
	order, err := o.Read(id)
	if err != nil {
		return err
	}
	result := o.db.Table("orders").Model(&order).Delete(&order)
	return result.Error
}
