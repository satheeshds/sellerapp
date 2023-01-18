package paymentrepository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type PaymentRepository struct {
	db *gorm.DB
}

func (p *PaymentRepository) Open() error {
	log.Printf("establishing db connection")
	var err error
	// Ideally configuration should be in config
	log.Printf("latest host no gorm.model")
	p.db, err = gorm.Open("postgres", "user=user password=password host=db dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	p.db.LogMode(true)
	return err
}

func (p *PaymentRepository) Close() {
	p.db.Close()
}

func (p *PaymentRepository) Create(payment *models.Payment) error {
	result := p.db.Create(payment)

	return result.Error
}

func (p *PaymentRepository) Read(id int) (models.Payment, error) {
	var payment models.Payment

	result := p.db.First(&payment, id)
	return payment, result.Error
}

func (p *PaymentRepository) Update(payment *models.Payment) error {
	result := p.db.Model(payment).Update(payment)

	return result.Error
}
