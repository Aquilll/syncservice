package worker

import (
    "testing"
    "time"
)

func TestRateLimiter(t *testing.T) {
    limiter := NewRateLimiter(2)
    start := time.Now()
    limiter.Wait()
    limiter.Wait()
    limiter.Wait()
    elapsed := time.Since(start)

    if elapsed < 500*time.Millisecond {
        t.Errorf("rate limiter did not apply expected delay")
    }
}
