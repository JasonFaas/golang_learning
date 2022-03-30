package main

import (
    "fmt"   
    "bufio"
    "os"
    "strings"
)

func main() {

    // Accept input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What is your move? ")
    next_move, _ := reader.ReadString('\n')

    // Only stop accepting input on string "exit"
    for next_move != "exit\n" {
        fmt.Print("You moved to " + next_move)

        fmt.Println()
        fmt.Print("What is your move? ")
        next_move, _ = reader.ReadString('\n')
    }


}

