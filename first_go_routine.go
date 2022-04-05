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
	// // what[0] = 3
	// fmt.Printf("%d\n", what[0])
	// fmt.Printf("%d\n", what[1])

	for itr, _ := range what {
		what[itr] = make(chan int)
		go calcSquare(itr, what[itr])
	}


	squareOfFour := make(chan int)
	go calcSquare(4, squareOfFour)
	fmt.Printf("%T\n", squareOfFour)

	// time.Sleep(2 * time.Second)
	fmt.Println("Finished Main")

	realSquareOfFour := <- squareOfFour
	fmt.Printf("%d result\n", realSquareOfFour)
	fmt.Printf("%T\n", realSquareOfFour)
	fmt.Printf("%T\n", squareOfFour)


	for itr, item := range what {
		fmt.Printf("%d %d result\n", itr, <-item)
	}
}

func first_go_routine(id int) {
	fmt.Printf("Started Go Routine %d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished Go Routine %d\n", id)
}

func calcSquare(number int, cubeop chan int) {
	time.Sleep(2 * time.Second)
	var result = number * number
	cubeop <- result
}

// TODO: Test waiting for all goroutines completing

// TODO: Get return values back from goroutines
