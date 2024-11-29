package internal

import (
	"sync"
	"time"
)

type CarbonStore struct {
	Store sync.Map
	stopChan chan struct{}
}

type CarbonValue struct {
	Value any
	Expiry time.Time
}