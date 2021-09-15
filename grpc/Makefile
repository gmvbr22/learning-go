build:
	cd service && protoc \
	--go_out=./gen \
	--go_opt=paths=source_relative \
	--go-grpc_out=./gen \
	--go-grpc_opt=paths=source_relative \
	*.proto

run:
	go run server/main.go

run_client:
	go run client/main.go