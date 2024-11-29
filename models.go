package carbon

import (
	"sync"
	"time"
)

type CarbonStore struct {
	store sync.Map
	stopChan chan struct{}
}

type CarbonValue struct {
	Value any
	Expiry *time.Time
}