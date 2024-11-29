package main

import (
	"time"

	z "github.com/scott-mescudi/carbon/carbon"
	"fmt"
)


func carbon(){
	key := "ehllo"
	value := 12

	cdb := z.NewCarbonStore(1 * time.Second)
	cdb.Set(key, value, 40*time.Second)
	cdb.Set(value, key, 40*time.Second)


	start := time.Now()
	val, err := cdb.Get(key)
	fmt.Println(time.Since(start).Nanoseconds())

	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Println(val)
	cdb.CloseStore()
}


func main() {
	carbon()
}

// TODO: add encryption and compression to backup
// TODO: make func to get a newcarbonstore from file