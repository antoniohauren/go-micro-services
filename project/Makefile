FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

up:
	docker compose up -d
	@echo "Done!"

down:
	docker compose down
	@echo "Done!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	cd ../front-end && ./${FRONT_END_BINARY} &
	@echo -e "\nDone!"

## stop: stop the front end
stop:
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Done!"