run-app:
	go run ./internal/cmd/app/*.go
build-app:
	go build -o bin/app ./internal/cmd/app/*.go
run-script:
	go run ./internal/cmd/script/*.go