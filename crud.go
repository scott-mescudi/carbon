package carbon

import (
	"fmt"
	"time"
	
)

const (
	NoExpiry = -1
	NoClean = -1
)

func (s *CarbonStore)Set(key, value any, expiry time.Duration) {
	if expiry == NoExpiry {
		s.store.Store(key, CarbonValue{Value: value, Expiry: nil})
		return
	}

	exp := time.Now().Add(expiry)
	s.store.Store(key, CarbonValue{Value: value, Expiry: &exp })

}

func (s *CarbonStore)Get(key any) (value any, err error) {
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

	return v.(CarbonValue).Value , nil
}

func (s *CarbonStore) Delete(key any) {
    s.store.Delete(key)
}

func (s *CarbonStore) ClearStore() {
    s.store.Clear()
}

func (s *CarbonStore) CloseStore() {
	s.StopCleaner()
    s.store.Clear()
}

//debug
func (s *CarbonStore) Printall() {
	s.store.Range(func(key, value any) bool{
		fmt.Println(key, value)
			return true
	})
}