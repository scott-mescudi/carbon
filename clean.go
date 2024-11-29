package carbon

import (
	"time"
)

// cleanStore periodically cleans up expired keys from the store based on the specified frequency.
// The function runs in a separate goroutine and checks for expired keys at regular intervals defined by cleanFrequency.
// It stops when the StopCleaner function is called (via the s.stopChan channel).
func (s *CarbonStore) cleanStore(cleanFrequency time.Duration) {
	ticker := time.NewTicker(cleanFrequency)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopChan:
			return 
		case <-ticker.C:
			s.store.Range(func(key, value any) bool {
				if value.(CarbonValue).Expiry == nil {
					return true
				}

				if time.Since(*value.(CarbonValue).Expiry) > 0 {
					s.store.Delete(key)
				}

				return true
			})
		}
	}
}

// StopCleaner stops the periodic cleaning of expired keys by closing the stop channel.
// It effectively halts the cleanStore function's loop.
func (s *CarbonStore) StopCleaner() {
	close(s.stopChan)
}
