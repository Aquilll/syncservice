package queue

import (
    "testing"
    "syncservice/models"
)

func TestPartitionedQueue(t *testing.T) {
    pq := NewPartitionedQueue([]string{"provider1"}, 10)
    customer := models.InternalCustomer{ID: "cust123"}
    pq.Enqueue("provider1", "cust123", customer)

    q := pq.GetOrCreateQueue("provider1", "cust123")
    result := <-q

    if result.ID != "cust123" {
        t.Errorf("expected customer ID 'cust123', got %s", result.ID)
    }
}
