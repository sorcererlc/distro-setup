build:
	cd install && go build -o ../setup ./main.go
	chmod +x ./setup

install:
	make build
	./setup

test:
	make build
	TEST=true ./setup
