module github.com/satheeshds/sellerapp/cmd/inventory_manager

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	github.com/satheeshds/sellerapp/pkg/common v0.0.0
	github.com/satheeshds/sellerapp/pkg/inventory_repository v0.0.0
	github.com/satheeshds/sellerapp/pkg/inventory_service v0.0.0
	github.com/satheeshds/sellerapp/pkg/models v0.0.0
)

replace github.com/satheeshds/sellerapp/pkg/common v0.0.0 => ../../pkg/common

replace github.com/satheeshds/sellerapp/pkg/models v0.0.0 => ../../pkg/models

replace github.com/satheeshds/sellerapp/pkg/inventory_repository v0.0.0 => ../../pkg/inventory_repository

replace github.com/satheeshds/sellerapp/pkg/inventory_service v0.0.0 => ../../pkg/inventory_service
