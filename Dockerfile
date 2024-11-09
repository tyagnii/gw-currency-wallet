FROM golang:1.23.2 as build
WORKDIR /usr/gw-currency-wallet

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

FROM scratch
WORKDIR /
# binary
COPY --from=build /usr/gw-currency-wallet/bin/server  /
# configuration file
COPY --from=build /usr/gw-currency-wallet/config.env  /
# keep migrations in the same dir hirarchy
ADD ./internal/db/migrations/*.sql  /internal/db/migrations/

CMD ["/server", "serve"]

