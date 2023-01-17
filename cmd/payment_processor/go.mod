module github.com/satheeshds/sellerapp/cmd/payment_processor

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	github.com/satheeshds/sellerapp/pkg/common v0.0.0
	github.com/satheeshds/sellerapp/pkg/models v0.0.0
	github.com/satheeshds/sellerapp/pkg/payment_repository v0.0.0
	github.com/satheeshds/sellerapp/pkg/payment_service v0.0.0
)

require (
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
)

replace github.com/satheeshds/sellerapp/pkg/common v0.0.0 => ../../pkg/common

replace github.com/satheeshds/sellerapp/pkg/models v0.0.0 => ../../pkg/models

replace github.com/satheeshds/sellerapp/pkg/payment_repository v0.0.0 => ../../pkg/payment_repository

replace github.com/satheeshds/sellerapp/pkg/payment_service v0.0.0 => ../../pkg/payment_service
