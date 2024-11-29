package main

import (
	"fmt"
	"github.com/scott-mescudi/carbon"
)

func main() {
	// Create a new Carbon cache store without the background cleaner.
	cdb := carbon.NewCarbonStore(carbon.NoClean)

	// Store basic data types in the cache.

	// Caching a string value with no expiry.
	cdb.Set("greeting", "Hello, Carbon!", carbon.NoExpiry)

	// Caching an integer value with no expiry.
	cdb.Set("user_age", 30, carbon.NoExpiry)

	// Caching a float value with no expiry.
	cdb.Set("user_score", 98.6, carbon.NoExpiry)

	// Retrieve and print the cached values.

	// Retrieve the string value.
	stringValue, err := cdb.Get("greeting")
	if err != nil {
		fmt.Println("Error retrieving string:", err)
		return
	}
	if stringValue == nil {
		fmt.Println("String value is nil.")
		return
	}

	// Note: For basic data types like strings, integers, and floats, 
	// there is no need for type assertion. However, it is recommended 
	// for safety to explicitly check the type to avoid potential issues 
	// if the cache stores unexpected types.

	// Type assert the string value (recommended for safety).
	greeting, ok := stringValue.(string)
	if !ok {
		fmt.Println("Failed to assert string value.")
		return
	}
	fmt.Println("Greeting:", greeting)

	// Retrieve the integer value.
	intValue, err := cdb.Get("user_age")
	if err != nil {
		fmt.Println("Error retrieving integer:", err)
		return
	}
	if intValue == nil {
		fmt.Println("Integer value is nil.")
		return
	}

	// Type assert the integer value (recommended for safety).
	age, ok := intValue.(int)
	if !ok {
		fmt.Println("Failed to assert integer value.")
		return
	}
	fmt.Println("User Age:", age)

	// Retrieve the float value.
	floatValue, err := cdb.Get("user_score")
	if err != nil {
		fmt.Println("Error retrieving float:", err)
		return
	}
	if floatValue == nil {
		fmt.Println("Float value is nil.")
		return
	}

	// Type assert the float value (recommended for safety).
	score, ok := floatValue.(float64)
	if !ok {
		fmt.Println("Failed to assert float value.")
		return
	}
	fmt.Println("User Score:", score)
}
