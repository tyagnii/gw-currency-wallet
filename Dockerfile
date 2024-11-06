FROM golang:1.23.2

WORKDIR /usr/gw-currency-wallet

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

# CMD ["./bin/server", "serve"]

