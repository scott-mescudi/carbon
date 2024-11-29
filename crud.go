package carbon

import (
	"fmt"
	"time"
)

const (
	// NoExpiry is a constant used to indicate that the value should not expire.
	NoExpiry = -1
	// NoClean is a constant used to indicate that there is no cleaning operation set.
	NoClean = -1
)

// Set adds a key-value pair to the store with an optional expiration time.
// If the expiration is set to carbon.NoExpiry, the key-value pair will remain in the store indefinitely.
func (s *CarbonStore) Set(key, value any, expiry time.Duration) {
	if expiry == NoExpiry {
		s.store.Store(key, CarbonValue{Value: value, Expiry: nil})
		return
	}

	exp := time.Now().Add(expiry)
	s.store.Store(key, CarbonValue{Value: value, Expiry: &exp })
}

// Get retrieves a value by its key from the store.
// If the key does not exist, it returns nil and an error.
// If the key is expired, it deletes the key and returns nil and an error.
func (s *CarbonStore) Get(key any) (value any, err error) {
	v, ok := s.store.Load(key)
	if v == nil || !ok {
		return nil, fmt.Errorf("'%v' does not exist", key)
	}

	carb := v.(CarbonValue)
	if carb.Expiry != nil {
		if carb.Expiry.Before(time.Now()) {
			s.store.Delete(key)
			return nil, fmt.Errorf("'%v' is expired", key)
		}	
	}

	return carb.Value, nil
}

// Delete removes a key-value pair from the store by its key.
func (s *CarbonStore) Delete(key any) {
    s.store.Delete(key)
}

// ClearStore removes all keys and values from the store, freeing up memory.
func (s *CarbonStore) ClearStore() {
    s.store.Clear()
}

// CloseStore stops any active cleaner goroutines (if present) and clears the store to free memory.
func (s *CarbonStore) CloseStore() {
	s.StopCleaner()
    s.store.Clear()
}

// Printall prints all key-value pairs in the store to stdout.
func (s *CarbonStore) Printall() {
	s.store.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
