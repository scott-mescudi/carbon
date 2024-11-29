package carbon

import (
	"fmt"
	"os"
)

// BackupToFile creates a backup file on disk with the given file name
// and stores all the key-value pairs from the store in the file in the format {key=value}.
// This function does not store the expiration times of the keys, only the values.
func (s *CarbonStore) BackupToFile(BackupFileName string) error {
	f, err := os.Create(BackupFileName)
	if err != nil {
		return err
	}
	defer f.Close()


	s.store.Range(func(key, value any) bool {
		carb := value.(CarbonValue)
		_, err = fmt.Fprintf(f, "{%v=%v}", key, carb.Value)
		if err != nil {
			return false
		}
		return true
	})

	return nil
}
