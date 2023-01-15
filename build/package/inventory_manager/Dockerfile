FROM golang:1.18 AS build
WORKDIR /go/src
#COPY go ./go
COPY ./cmd/inventory_manager ./cmd/inventory_manager
COPY ./pkg ./pkg

WORKDIR /go/src/cmd/inventory_manager
ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o inventory_manager .
FROM scratch AS runtime
COPY --from=build /go/src/cmd/inventory_manager/inventory_manager ./
ENTRYPOINT ["./inventory_manager"]