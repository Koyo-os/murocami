start:
	./scripts/start.sh
deps:
	go mod download
proto:
	protoc --go_out=. --go-grpc_out=. internal/network_agent_client/proto/nagent.proto