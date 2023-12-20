# Build
buildForLinux:
	GOOS=linux GOARCH=amd64 go build -o first-service cmd/app/main.go

buildForMacARM:
	GOOS=darwin GOARCH=arm64 go build -o first-service cmd/app/main.go

buildForWindows:
	set GOOS=windows \
	set	GOARCH=amd64 \
	go build -o first-service cmd/app/main.go


protocGen:
	protoc protoMessages/*.proto \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--proto_path=.
