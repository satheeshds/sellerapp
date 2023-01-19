package apiclients

// InventoryClient is an interface that contains two methods: BlockProduct and UnblockProduct.
// BlockProduct takes two int parameters and returns an error. It is used to block a product from being sold.
// UnblockProduct takes two int parameters and returns an error. It is used to unblock a product so it can be sold again.
type IInventoryClient interface {
	BlockProduct(int, int) error
	UnblockProduct(int, int) error
}
