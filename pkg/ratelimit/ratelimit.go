package ratelimit

import (
	"sync"
	"time"
)

// RequestPerMinute is the user-defined limit of 15 requests per minute.
const RequestPerMinute = 15

// MinTimeBetweenRequests is the required minimum duration between consecutive throttled calls (60s / 15 = 4s).
var MinTimeBetweenRequests = time.Minute / RequestPerMinute

var (
	// lastRequestTime tracks when the last call *started*.
	lastRequestTime = time.Time{}
	mu              sync.Mutex
)

// Throttle waits for the minimum required time to pass since the last call
// started, ensuring the rate limit of RequestPerMinute is not exceeded.
// The function f will be executed immediately after the throttle check passes.
func Throttle(f func()) {
	mu.Lock()
	defer mu.Unlock()

	// Calculate time elapsed since the last request started.
	elapsed := time.Since(lastRequestTime)

	// If the elapsed time is less than the minimum required time, sleep the difference.
	if elapsed < MinTimeBetweenRequests {
		sleepDuration := MinTimeBetweenRequests - elapsed
		time.Sleep(sleepDuration)
	}

	// Update the last request time to the current time (after any potential sleep).
	lastRequestTime = time.Now()

	// Unlock to allow other goroutines to check the rate limit while 'f' runs.
	// This is crucial for handling concurrent requests efficiently.
	mu.Unlock()

	// Execute the function that contains the actual API call.
	f()
}