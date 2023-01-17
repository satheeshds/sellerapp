module github.com/satheeshds/sellerapp/cmd/order_processor

go 1.18

require (
	github.com/satheeshds/sellerapp/pkg/common v0.0.0
	github.com/satheeshds/sellerapp/pkg/models v0.0.0
	github.com/satheeshds/sellerapp/pkg/order_repository v0.0.0
	github.com/satheeshds/sellerapp/pkg/order_service v0.0.0 
	github.com/satheeshds/sellerapp/pkg/api_clients v0.0.0
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
)

replace github.com/satheeshds/sellerapp/pkg/common v0.0.0 => ../../pkg/common

replace github.com/satheeshds/sellerapp/pkg/models v0.0.0 => ../../pkg/models

replace github.com/satheeshds/sellerapp/pkg/order_repository v0.0.0 => ../../pkg/order_repository

replace github.com/satheeshds/sellerapp/pkg/order_service v0.0.0 => ../../pkg/order_service
replace github.com/satheeshds/sellerapp/pkg/api_clients v0.0.0 => ../../pkg/api_clients
