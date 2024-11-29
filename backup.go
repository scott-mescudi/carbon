package carbon

import (
	"fmt"
	"os"
)

func (s *CarbonStore) BackupToFile(BackupFileName string) error {
	f, err := os.Create(BackupFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	s.Store.Range(func(key, value any) bool {
		carb := value.(CarbonValue)
		_, err = fmt.Fprintf(f, "%v=%v\n", key, carb.Value)
		if err != nil {
			return false
		}
		return true
	})

	return nil
}