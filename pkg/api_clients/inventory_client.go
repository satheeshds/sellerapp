package apiclients

import (
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type InventoryClient struct{}

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

func (i *InventoryClient) UnblockProduct(product_id, product_quantity int) error {
	return i.BlockProduct(product_id, -product_quantity)
}
