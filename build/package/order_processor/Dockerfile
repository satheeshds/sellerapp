FROM golang:1.18 AS build
WORKDIR /go/src
#COPY go ./go
COPY ./cmd/order_processor .
COPY ./pkg ./pkg

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o order_processor .
FROM scratch AS runtime
COPY --from=build /go/src/order_processor ./
ENTRYPOINT ["./order_processor"]