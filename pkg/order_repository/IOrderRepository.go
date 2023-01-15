package orderrepository

import "github.com/satheeshds/sellerapp/pkg/models"

type IOrderRepository interface {
	Open() error
	Create(models.Order) error
	Read(id int) (models.Order, error)
	Update(id int, order models.Order) error
	Delete(id int) error
	Close()
}
