# Run application in development mode
dev:
	air
# Generate or update swagger docs
swag:
	swag init --dir ./cmd/server,./internal/handlers,./internal/services,./internal/routes,./internal/payments/proto

proto:
	protoc --go_out=.  proto/payments.proto
	# protoc --go-grpc_out=internal/payments --go-grpc_opt=paths=source_relative proto/payments.proto
	# need both