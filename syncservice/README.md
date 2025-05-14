# SyncService

# Pre-requisite

Docker should be installed and running in the system.

## Usage

1. Run using make file

```bash
make docker-build # To build the docker image.
make docker-run # To run the server
make docker-test-integration # To run the integration test.
make docker-test # To run the entire test.
```

2. Build & Run with Docker

```bash
docker build -t syncservice .
docker run --env-file .env -p 8080:8080 syncservice
```

3. Example API Requests

```bash
curl -X POST http://localhost:8081/internal/crud -H "Content-Type: application/json" -d '{"id": "cust123", "first_name": "John", "last_name": "Doe", "email": "john@example.com", "phone_number": "1234567890"}'
curl -X POST http://localhost:8081/webhook -H "Content-Type: application/json" -d '{"customer_id": "cust456", "full_name": "Jane Smith", "email_address": "jane@example.com", "phone": "9876543210", "last_modified": "2024-05-10T12:00:00Z"}'
```

## General troubleshooting
- If you encounter issues with docker-run, ensure that the port 8081 is available or change the port inside the .env file to an available one.
    - When you change the port make sure you use the same port in the curl command too.
