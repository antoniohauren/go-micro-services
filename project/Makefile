FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp

up:
	docker compose up -d
	@echo "Done!"

down:
	docker compose down
	@echo "Done!"

up_build: build_broker build_auth build_logger
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

build_broker:
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_auth:
	cd ../auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

start: build_front
	cd ../front-end && ./${FRONT_END_BINARY} &
	@echo -e "\nDone!"

stop:
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Done!"

clean:
	rm ../broker-service/${brokerApp}
	rm ../auth-service/${authApp}
	rm ../logger-service/${loggerApp}
	rm ../front-end/${FRONT_END_BINARY}