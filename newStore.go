package carbon

import (
	"os"
	"regexp"
	"sync"
	"time"
)

func NewCarbonStore(cleanFrequency time.Duration) *CarbonStore {
	z := make(chan struct{})
	s := CarbonStore{store: sync.Map{}, stopChan: z}

	if cleanFrequency != NoClean {
		go s.cleanStore(cleanFrequency)
	}
	return &s
}

func ImportStoreFromFile(filepath string, cleanFrequency time.Duration) (*CarbonStore, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}


	z := make(chan struct{})
	s := CarbonStore{store: sync.Map{}, stopChan: z}

	pattern := `\{(\w+)=(\w+)\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(file), -1)
	for _, match := range matches {
		if len(match) == 3 {
			s.Set(match[1], match[2], NoExpiry)
		}
	}


	if cleanFrequency != NoClean {
		go s.cleanStore(cleanFrequency)
	}
	
	return &s, nil
}