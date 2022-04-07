package main

import (
    "fmt"   
    // "math/rand"
    // "time"
    "bufio"
    "os"
    // "container/list"
    // "strings"
)

func HowManyHumans() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many humans? 0,1,2? ")
	read_string, _ := reader.ReadString('\n')
	var return_string = string(read_string[0])
	if return_string != "0" && return_string != "1" && return_string != "2" {
		os.Exit(3)
	}
    return return_string
}

func WhichTurnIsHuman() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which turn is human? %s,%s? ", X_INPUT, O_INPUT)
	read_string, _ := reader.ReadString('\n')
	var return_string = string(read_string[0])
	if return_string != X_INPUT && return_string != O_INPUT {
		os.Exit(2)
	}
    return return_string
}

func GetHumanInput(next_turn string) (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What position for %s?\n", next_turn)

    fmt.Println("Put next move in format \"D3\"")

    read_string, _ := reader.ReadString('\n')
    var next_move_pos = read_string
    fmt.Printf("Next human move is %s\n", next_move_pos)

    fmt.Printf("")

    if next_move_pos == "Exit\n" {
        os.Exit(33)
    }

    return next_move_pos
}