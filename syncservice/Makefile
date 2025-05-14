.PHONY: docker-build docker-run docker-test docker-test-integration

# Build the Docker image
docker-build:
	docker build -t syncservice .

# Run the Docker container with environment file
docker-run:
	docker run --env-file .env -p $$(grep ^PORT .env | cut -d '=' -f2):$$(grep ^PORT .env | cut -d '=' -f2) syncservice

# Run all tests
docker-test:
	docker run --rm syncservice go test ./...

# Run only integration test
docker-test-integration:
	docker run --rm syncservice go test ./integration/...

