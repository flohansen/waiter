FROM golang:1.22-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd/waiter/main.go cmd/waiter/main.go
RUN CGO_ENABLED=0 go build -o waiter cmd/waiter/main.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /usr/src/app/waiter /waiter

ENTRYPOINT ["/waiter"]
