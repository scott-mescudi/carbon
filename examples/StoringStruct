package main

import (
	"fmt"
	"github.com/scott-mescudi/carbon"
)

// Define a struct to be stored in the cache
type myStruct struct {
	id       int
	username string
}

func main() {
	// Create a new Carbon cache store without the background cleaner.
	cdb := carbon.NewCarbonStore(carbon.NoClean)

	// Create an instance of 'myStruct' to store in the cache.
	ms := myStruct{id: 1, username: "scott-mescudi"}

	// Store the struct in the cache under the key "user1" with no expiry time.
	cdb.Set("user1", ms, carbon.NoExpiry)

	// Retrieve the value associated with the "user1" key.
	value, err := cdb.Get("user1")
	if err != nil {
		// If there's an error while retrieving, print it and return.
		fmt.Println("Error retrieving value:", err)
		return
	}

	// Check if the retrieved value is nil to prevent panics.
	if value == nil {
		fmt.Println("Cache value is nil.")
		return
	}

	// Since the Get function returns a value of type 'any',
	// we need to check if it's not nil and then type assert it to our struct.
	castedValue, ok := value.(myStruct)
	if !ok {
		// If the type assertion fails, print an error message.
		fmt.Println("Failed to assert value to 'myStruct'")
		return
	}

	// Now we can safely access the fields of 'myStruct'.
	fmt.Println("ID:", castedValue.id, "Username:", castedValue.username)
}
