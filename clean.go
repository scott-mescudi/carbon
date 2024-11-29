package carbon

import (
	"fmt"
	"time"
)


func (s *CarbonStore) cleanStore(cleanFrequency time.Duration) {
	ticker := time.NewTicker(cleanFrequency)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopChan:
			fmt.Println("Stopping cleanStore goroutine...")
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


func (s *CarbonStore) StopCleaner() {
	close(s.stopChan)
}

