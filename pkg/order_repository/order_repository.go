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

	return nil
}

func (o *OrderRepository) Close() {
	o.db.Close()
}

func (o *OrderRepository) Create(order models.Order) error {
	o.db.Create(&order)

	return nil
}

func (o *OrderRepository) Read(id int) (models.Order, error) {
	var order models.Order
	o.db.First(&order, id)

	return order, nil
}

func (o *OrderRepository) Update(id int, order models.Order) error {
	o.db.Table("orders").Model(&order).Update(order)

	return nil
}

func (o *OrderRepository) Delete(id int) error {
	order, err := o.Read(id)
	if err != nil {
		return err
	}
	o.db.Table("orders").Model(&order).Delete(&order)
	return nil
}
