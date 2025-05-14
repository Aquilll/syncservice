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
 
## Walkthrough of the code

SyncService is a bidirectional synchronization service that efficiently handles customer data synchronization between internal systems and external CRM providers.

## Architecture and Component Interaction

### Core Components

1. **API Layer (`api/handler.go`)**
   - Exposes two HTTP endpoints:
     - `/internal/crud`: Processes internal services CRUD operations
     - `/webhook`: Receives external webhooks from CRM systems
   - Transforms and queues data for synchronization
   - Maintains an in-memory store of customer data

2. **Queue System (`queue/partitioned_queue.go`)**
   - Manages partitioned queues based on provider names
   - Each provider partition contains customer-specific channels
       - Customer is a record type that I have considered for this solution.
   - Buffers customer data updates to be processed by workers
   - Provides methods to create/access queues and enqueue items

3. **Worker System (`worker/worker.go`)**
   - Processes queued items for each CRM provider
   - Implements rate limiting to respect API limits of external providers
   - Spawns dedicated goroutines for each customer ID
   - Transforms internal data models to external formats before sending.
       - This is just a basic transformation as the solution revolves around system syncing and not transformation. 

4. **Provider Interfaces (`provider/*.go`)**
   - Defines a common interface (`CRMProvider`) for all CRM systems
   - Implements provider-specific logic (Salesforce, HubSpot)
   - Handles actual API calls to external CRM systems

5. **Data Models (`models/models.go`)**
   - Defines internal and external data structures:
     - `InternalCustomer`: Used within the system
     - `ExternalCustomer`: Format used by external CRM systems

6. **Data Transformer (`transformer/transformer.go`)**
   - Converts between internal and external data formats
   - Handles field mapping and data normalization

### Data Flow

1. **Internal CRUD → External CRMs**:
   - Internal CRUD request hits `/internal/crud` endpoint
   - Data is stored in `InternalStore` and enqueued for all providers
   - Workers process each customer's queue
   - Data is transformed to external format and sent to CRM providers

2. **External Webhook → Internal System**:
   - External webhook hits `/webhook` endpoint
   - External data is transformed to internal format
   - Data is stored in `InternalStore`

### Concurrency and Rate Limiting

- Uses Go channels to manage work queues
- Implements a token-bucket rate limiter for API throttling
- Utilizes goroutines for concurrent processing
- Uses mutex locks to protect shared resources
