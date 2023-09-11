package expired_lock

import "testing"

func TestExpiredLock(t *testing.T) {
	var mu ExpiredLock
	mu.Lock(1)
	mu.Lock(0)
	mu.Unlock()
}
