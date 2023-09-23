package main

import (
	"fmt"
	"log"
)

func main() {
	event, fn, err := InitEvent("")
	if err != nil {
		log.Fatal(err)
	}
	fn()
	fmt.Println(event)
}
