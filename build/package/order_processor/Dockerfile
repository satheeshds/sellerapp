FROM golang:1.18 AS build
WORKDIR /go/src
#COPY go ./go
COPY ./cmd/order_processor ./cmd/order_processor
COPY ./pkg ./pkg

WORKDIR /go/src/cmd/order_processor
ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o order_processor .
FROM scratch AS runtime
COPY --from=build /go/src/cmd/order_processor/order_processor ./
ENTRYPOINT ["./order_processor"]