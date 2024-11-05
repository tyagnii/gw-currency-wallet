# Build binary
.PHONY: build
build:
	@ go build -o ./bin/server ./main.go

# Generate swagger documentation
.PHONY: swagger-doc
swagger-doc:
	@ swag init

