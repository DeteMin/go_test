package main

import (
	"time"
)

func main() {
	defer println("in main")
	go func() {
		defer println("in goroutin")
		panic("")
	}()

	time.Sleep(1 * time.Second)
}
