package apiclients

type InventoryClient struct{}

func (i *InventoryClient) BlockProduct(product_id, product_quantity int) bool {
	return false
}

func (i *InventoryClient) UnblockProduct(product_id, product_quantity int) bool {
	return false
}
