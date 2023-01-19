package apiclients

import (
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type InventoryClient struct{}

// BlockProduct is a function of the InventoryClient type that takes an integer product_id and an integer product_quantity as parameters.
// The function uses the resty library to make a PATCH request to an inventory manager server at the given URL with a body containing the product ID and quantity.
// If an error occurs, it will be returned, otherwise if the response status code is 400 (Bad Request), an error is returned with the response string.
// If successful, nil is returned.
func (i *InventoryClient) BlockProduct(product_id, product_quantity int) error {
	client := resty.New()

	resp, err := client.R().
		SetBody(models.Inventory{Id: product_id, Quantity: product_quantity}).
		Patch(fmt.Sprintf("http://inventory_manager:5001/v1/inventory/%d", product_id))

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusBadRequest {
		return fmt.Errorf(resp.String())
	}

	return nil

}

// This function is a method of the InventoryClient type. It takes two parameters, product_id and product_quantity, both of which are integers.
// The function calls the BlockProduct method of the same type, passing in the product_id and a negative value of product_quantity.
// It then returns any error that may have occurred. This function is used to unblock a specified quantity of a product from the inventory.
func (i *InventoryClient) UnblockProduct(product_id, product_quantity int) error {
	return i.BlockProduct(product_id, -product_quantity)
}
