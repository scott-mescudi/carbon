package carbon

import (
	"fmt"
	"sync"
	"time"
)


func NewCarbonStore(cleanFrequency time.Duration) *CarbonStore {
	z := make(chan struct{})
	s := CarbonStore{Store: sync.Map{}, stopChan: z}
	go s.cleanStore(cleanFrequency)
	return &s
}

//debug
func (s *CarbonStore) Printall() {
	s.Store.Range(func(key, value any) bool{
		fmt.Println(key, value)
			return true
	})
}

func (s *CarbonStore)Set(key, value any, expiry time.Duration) {
	s.Store.Store(key, CarbonValue{Value: value, Expiry: time.Now().Add(expiry)})
}

func (s *CarbonStore)Get(key any) (value any, err error) {
	v, ok := s.Store.Load(key)
	if v == nil || !ok {
		return nil, fmt.Errorf("'%v' does not exist", key)
	}

	carb := v.(CarbonValue)
	if carb.Expiry.Before(time.Now()) {
        s.Store.Delete(key)
        return nil, fmt.Errorf("'%v' is expired", key)
    }

	return v.(CarbonValue).Value , nil
}

func (s *CarbonStore) Delete(key any) {
    s.Store.Delete(key)
}

func (s *CarbonStore) CloseStore() {
	s.StopCleaner()
    s.Store.Clear()
}

