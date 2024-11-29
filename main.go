package main

import (
	"fmt"
	"time"

	z "github.com/scott-mescudi/carbon/internal"
)




func main() {
	key := "ehllo"
	value := 12

	cdb := z.NewCarbonStore(1 * time.Second)
	cdb.Set(key, value, 40*time.Second)
	cdb.Set(value, key, 40*time.Second)

	val, err := cdb.Get(key)
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}

	cdb.BackupToFile("hello.txt")

	fmt.Println(val)
}

// TODO: add encryption and compression to backup
// TODO: make func to get a newcarbonstore from file