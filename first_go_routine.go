package main

import (
	"fmt"
	"time"
	"runtime"
	// "sync"
)

func main() {
    fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

	fmt.Println("Started Main")

	var what [32]chan int

	for itr, _ := range what {
		what[itr] = make(chan int)
		go calcSquare(itr, what[itr])
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Finished Main")

	for itr, item := range what {
		var temp_val = <-item
		fmt.Printf("%d %d result\n", itr, temp_val)
		fmt.Printf("%d %d result\n", itr, temp_val)
		// fmt.Printf("%d %d result\n", itr, <-item)
	}
}

func calcSquare(number int, cubeop chan int) {
	time.Sleep(2 * time.Second)
	var result = number * number
	cubeop <- result
}
