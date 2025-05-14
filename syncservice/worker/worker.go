package worker

import (
	"fmt"
	"sync"
	"time"

	"syncservice/models"
	"syncservice/provider"
	"syncservice/queue"
	"syncservice/transformer"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rps int) *RateLimiter {
	tokens := make(chan struct{}, rps)
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rps))
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
			}
		}
	}()
	return &RateLimiter{tokens: tokens}
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

func StartWorkerForProvider(providerName string, providerImpl provider.CRMProvider, pq *queue.PartitionedQueue, rps int) {
	rateLimiter := NewRateLimiter(rps)

	processedQueues := make(map[string]bool)
	var queueMutex sync.Mutex

	go func() {
		for {
			time.Sleep(time.Second)

			customerQueues := pq.Partitions[providerName]
			queueMutex.Lock()

			for customerID, ch := range customerQueues {
				if !processedQueues[customerID] {
					processedQueues[customerID] = true

					go func(cid string, q chan models.InternalCustomer) {
						fmt.Printf("[Worker] Started worker for %s:%s\n", providerName, cid)
						for customer := range q {
							rateLimiter.Wait()
							ext := transformer.ToExternal(customer)
							// calling only the update method of providers for the demo.
							if err := providerImpl.Update(ext); err != nil {
								fmt.Printf("[Worker %s:%s] Error: %v\n", providerName, cid, err)
							}
						}
					}(customerID, ch)
				}
			}

			queueMutex.Unlock()
		}
	}()
}
