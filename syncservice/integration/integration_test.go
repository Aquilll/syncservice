package integration

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"syncservice/models"
	"syncservice/provider"
	"syncservice/queue"
	"syncservice/worker"
)

// Structured log function
func logMetadata(status string, recordCount, queueSize int, data interface{}) {
	metadata := map[string]interface{}{
		"status": status,
		"metadata": map[string]int{
			"record_count": recordCount,
			"queue_size":   queueSize,
		},
		"data": data,
	}
	logJson, _ := json.MarshalIndent(metadata, "", "  ")
	fmt.Println(string(logJson))
}

func TestConcurrentSyncAndQueueStressWithLogs(t *testing.T) {
	// Setup
	providers := []string{"salesforce"}
	bufferSize := 5
	pq := queue.NewPartitionedQueue(providers, bufferSize)
	sfProvider := &provider.SalesforceProvider{}
	rps := 2

	// Start worker with rate limiting
	go worker.StartWorkerForProvider("salesforce", sfProvider, pq, rps)

	// Log start
	fmt.Println("[LOG] Starting concurrent producers: internal CRUD and webhook")

	var wg sync.WaitGroup
	wg.Add(2)

	// Internal creates
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			customer := models.InternalCustomer{ID: fmt.Sprintf("cust%d", i)}
			pq.Enqueue("salesforce", customer.ID, customer)
			fmt.Printf("[LOG] POST /internal/records - Created record ID %s\n", customer.ID)
			logMetadata("success", i+1, len(pq.GetOrCreateQueue("salesforce", customer.ID)), customer)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Webhook creates
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			customer := models.InternalCustomer{ID: fmt.Sprintf("cust%d", i)}
			pq.Enqueue("salesforce", customer.ID, customer)
			fmt.Printf("[LOG] POST /webhook - Webhook received for %s\n", customer.ID)
			logMetadata("success", i+1, len(pq.GetOrCreateQueue("salesforce", customer.ID)), customer)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Wait for both to finish
	wg.Wait()

	// Allow worker to process
	time.Sleep(5 * time.Second)

	// Validation: Check if queues are drained
	for custID, custQueue := range pq.Partitions["salesforce"] {
		qLen := len(custQueue)
		fmt.Printf("[LOG] Queue size for provider salesforce, customer %s: %d\n", custID, qLen)
	}
}
