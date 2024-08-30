build:
	cd src && go mod vendor && go mod tidy && go build -o ../setup ./main.go
	chmod +x ./setup

install:
	make build
	./setup

test:
	make build
	TEST=true DEBUG=true ./setup
