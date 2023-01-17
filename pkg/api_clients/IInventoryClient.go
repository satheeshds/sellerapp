package apiclients

type IInventoryClient interface {
	BlockProduct(int, int) bool
	UnblockProduct(int, int) bool
}
