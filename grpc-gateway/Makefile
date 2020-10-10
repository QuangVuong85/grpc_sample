.ONESHELL:

all: clean compile

compile:
	@echo "Compile binary ..."; \
	./build.sh
	@echo "Build success!";

clean:
	@echo "Cleaning files ..."; \
	find ./build -type f -name '*.yaml' -delete; \
	find ./build -type f -name '*.json' -delete; \
	find ./build -type f -name 'grpc-gateway' -delete; \
	find ./build -type f -name 'services' -delete; \
	sleep 1

run-gateway:
	clear && \
	cd build && \
	./grpc-gateway

run-services:
	clear && \
	cd build && \
	./services