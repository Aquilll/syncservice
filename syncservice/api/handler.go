package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"syncservice/models"
	"syncservice/queue"
	"syncservice/transformer"
)

var (
	InternalStore = make(map[string]models.InternalCustomer)
	StoreLock     = sync.RWMutex{}
)

type APIHandler struct {
	PartitionedQueue *queue.PartitionedQueue
}

func (h *APIHandler) CrudHandler(w http.ResponseWriter, r *http.Request) {
	var req models.InternalCustomer
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	req.UpdatedAt = time.Now()
	StoreLock.Lock()
	InternalStore[req.ID] = req
	StoreLock.Unlock()
	for providerName := range h.PartitionedQueue.Partitions {
		h.PartitionedQueue.Enqueue(providerName, req.ID, req)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Record accepted for sync\n")
}

func (h *APIHandler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var external models.ExternalCustomer
	if err := json.NewDecoder(r.Body).Decode(&external); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	internal := transformer.ToInternal(external)
	StoreLock.Lock()
	InternalStore[internal.ID] = internal
	StoreLock.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Webhook received and processed\n")
}
