# Build binary
.PHONY: build
build:
	@ go build -o ./bin/server ./main.go

# Generate swagger documentation
.PHONY: swagger-doc
swagger-doc:
	@ swag init

# Generate mocks fo db
.PHONY: gen-mock
gen-mock:
	@ mockgen -source internal/db/db.go -destination gen/mock/mock.go

