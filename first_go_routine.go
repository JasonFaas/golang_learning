package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
    fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

	fmt.Println("Started Main")

	var what [1024]int

	for itr, _ := range what {
		go first_go_routine(itr)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Finished Main")
}

func first_go_routine(id int) {
	fmt.Printf("Started Go Routine %d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished Go Routine %d\n", id)
}

// TODO: Test waiting for all goroutines completing

// TODO: Get return values back from goroutines
