package orderrepository

import "github.com/satheeshds/sellerapp/pkg/models"

// IOrderRepository is an interface for interacting with an Order repository.
// Open() opens the repository and returns an error if it fails.
// Create(*models.Order) creates a new Order in the repository and returns an error if it fails.
// Read(id int) reads an existing Order from the repository and returns it along with an error if it fails.
// Update(id int, order models.Order) updates an existing Order in the repository and returns an error if it fails.
// Delete(id int) deletes an existing Order from the repository and returns an error if it fails.
// Close() closes the repository and returns any errors that occur during closing.
type IOrderRepository interface {
	Open() error
	Create(*models.Order) error
	Read(id int) (models.Order, error)
	Update(id int, order models.Order) error
	Delete(id int) error
	Close()
}
