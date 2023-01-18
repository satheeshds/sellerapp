package orderrepository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func (o *OrderRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	// Ideally configuration should be in config
	log.Printf("latest host no gorm.model")
	o.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (o *OrderRepository) Close() {
	o.db.Close()
}

func (o *OrderRepository) Create(order *models.Order) error {
	result := o.db.Create(order)

	return result.Error
}

func (o *OrderRepository) Read(id int) (models.Order, error) {
	var order models.Order
	result := o.db.First(&order, id)

	return order, result.Error
}

func (o *OrderRepository) Update(id int, order models.Order) error {
	result := o.db.Table("orders").Model(&order).Update(order)

	return result.Error
}

func (o *OrderRepository) Delete(id int) error {
	order, err := o.Read(id)
	if err != nil {
		return err
	}
	result := o.db.Table("orders").Model(&order).Delete(&order)
	return result.Error
}
