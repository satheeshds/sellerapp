FROM golang:1.18 AS build
WORKDIR /go/src
#COPY go ./go
COPY ./cmd/payment_processor ./cmd/payment_processor
COPY ./pkg ./pkg

WORKDIR /go/src/cmd/payment_processor
ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o payment_processor .
FROM scratch AS runtime
COPY --from=build /go/src/cmd/payment_processor/payment_processor ./
ENTRYPOINT ["./payment_processor"]