build:
	cd install && go build -o ../setup ./main.go
	chmod +x ./setup

run:
	make build
	./setup

test:
	make build
	DRY_RUN=true ./setup
