package apiclients

type IInventoryClient interface {
	BlockProduct(int, int) error
	UnblockProduct(int, int) error
}
