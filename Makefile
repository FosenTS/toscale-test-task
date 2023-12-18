# Build
buildFirstForLinux:
	GOOS=linux GOARCH=amd64 go build -o first-service first-service/cmd/app/main.go

buildFirstForMacARM:
	GOOS=darwin GOARCH=arm64 go build -o first-service first-service/cmd/app/main.go

buildFirstForWindows:
	set GOOS=windows \
	set	GOARCH=amd64 \
	go build -o first-service first-service/cmd/app/main.go

