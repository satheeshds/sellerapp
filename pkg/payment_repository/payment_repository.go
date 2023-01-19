package paymentrepository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

// PaymentRepository is a struct that stores a pointer to a gorm.DB object.
// This struct is used to store and access data related to payments from the database.
type PaymentRepository struct {
	db *gorm.DB
}

// This function opens a connection to a Postgres database using the GORM library.
// The connection is stored in the PaymentRepository struct. If an error occurs, it is logged and returned.
func (p *PaymentRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	p.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// This function closes the database connection associated with the PaymentRepository object.
// It takes a pointer to a PaymentRepository object as its argument and calls the Close() method on the database object stored in the PaymentRepository's db field.
func (p *PaymentRepository) Close() {
	p.db.Close()
}

// This function is a method of the PaymentRepository struct. It takes in a pointer to a models.Payment struct and creates a payment record in the database.
// It returns an error if there is an issue with creating the record.
func (p *PaymentRepository) Create(payment *models.Payment) error {
	result := p.db.Create(payment)

	return result.Error
}

// This function is part of the PaymentRepository type. It takes an integer id as an argument and returns a models.Payment struct and an error.
// It uses the db variable from the PaymentRepository to query the database for a payment with the given id and assign it to the payment struct.
// The function then returns the payment struct and any errors that may have occurred during the query.
func (p *PaymentRepository) Read(id int) (models.Payment, error) {
	var payment models.Payment

	result := p.db.First(&payment, id)
	return payment, result.Error
}

// This function updates a Payment in the PaymentRepository. It takes a pointer to a Payment model as an argument and uses the Model method of the database to update the payment.
// The result of the update is then returned as an error.
func (p *PaymentRepository) Update(payment *models.Payment) error {
	result := p.db.Model(payment).Update(payment)

	return result.Error
}
