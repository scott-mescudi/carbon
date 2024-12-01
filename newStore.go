package carbon

import (
	"os"
	"regexp"
	"sync"
	"time"
)

// NewCarbonStore creates and returns a new CarbonStore instance.
// If a valid `cleanFrequency` is provided, it starts a goroutine to periodically clean expired keys from the store.
func NewCarbonStore(cleanFrequency time.Duration) *CarbonStore {
	z := make(chan struct{})
	s := CarbonStore{store: sync.Map{}, stopChan: z}


	if cleanFrequency != NoClean {
		go s.cleanStore(cleanFrequency)
	}

	return &s
}

// ImportStoreFromFile reads a file and extracts key-value pairs in the format {key=value} using a regular expression.
// The key-value pairs found in the file are then set in the store with no expiry time (NoExpiry) or a Default expiry time.
// It also starts a cleaner goroutine if `cleanFrequency` is provided.
func ImportStoreFromFile(filepath string, cleanFrequency time.Duration, defaultExpiry time.Duration) (*CarbonStore, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}


	z := make(chan struct{})
	s := CarbonStore{store: sync.Map{}, stopChan: z}

	
	pattern := `\{(\w+)=(\w+)\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(file), -1)
	
	var ss time.Duration
	if defaultExpiry == NoExpiry {
		ss = NoExpiry
	}else{
		ss = defaultExpiry
	}

	
	for _, match := range matches {
		if len(match) == 3 {
			s.Set(match[1], match[2], ss)
		}
	}

	if cleanFrequency != NoClean {
		go s.cleanStore(cleanFrequency)
	}

	return &s, nil
}
