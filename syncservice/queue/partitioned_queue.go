package queue

import (
	"syncservice/models"
)

type PartitionedQueue struct {
	Partitions map[string]map[string]chan models.InternalCustomer
	BufferSize int
}

func NewPartitionedQueue(providers []string, bufferSize int) *PartitionedQueue {
	partitions := make(map[string]map[string]chan models.InternalCustomer)
	for _, p := range providers {
		partitions[p] = make(map[string]chan models.InternalCustomer)
	}
	return &PartitionedQueue{Partitions: partitions, BufferSize: bufferSize}
}

func (pq *PartitionedQueue) GetOrCreateQueue(providerName, customerID string) chan models.InternalCustomer {
	customerQueues, exists := pq.Partitions[providerName]
	if !exists {
		return nil
	}
	queue, exists := customerQueues[customerID]
	if !exists {
		queue = make(chan models.InternalCustomer, pq.BufferSize)
		customerQueues[customerID] = queue
	}
	return queue
}

func (pq *PartitionedQueue) Enqueue(providerName, customerID string, customer models.InternalCustomer) {
	queue := pq.GetOrCreateQueue(providerName, customerID)
	if queue != nil {
		queue <- customer
	}
}
