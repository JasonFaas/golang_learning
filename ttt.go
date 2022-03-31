package main

import (
    "fmt"   
    "math/rand"
    "time"
    "bufio"
    "os"
    // "strings"
)

func PrintTableState(tempArr [3][3]string) () {
    fmt.Println("   A B C")
    for itr, _ := range tempArr {
        fmt.Printf("%d", itr)
        fmt.Printf(" ")
        fmt.Println(tempArr[itr])
    }
}

func DidAnyoneWin(tempArr [3][3]string, whoWon string) (bool) {
    for _, item := range [3]int{0,1,2} {
        if tempArr[item][0] == whoWon && tempArr[item][1] == whoWon && tempArr[item][2] == whoWon {
            return true
        }
        if tempArr[0][item] == whoWon && tempArr[1][item] == whoWon && tempArr[2][item] == whoWon {
            return true
        }
    }

    if tempArr[0][0] == whoWon && tempArr[1][1] == whoWon && tempArr[2][2] == whoWon {
        return true
    }
    if tempArr[0][2 - 0] == whoWon && tempArr[1][2 - 1] == whoWon && tempArr[2][2 - 2] == whoWon {
        return true
    }
    return false
}


func main() {

    reader := bufio.NewReader(os.Stdin)

    // TODO: V1

    // double 3x3 array
    var tttArr [3][3]string

    // random population of array with blank/X/O
    rand.Seed(time.Now().UnixNano())
    for vk, v := range tttArr {
        for vvk, _ := range v {
            // var value = rand.Intn(3)
            tttArr[vk][vvk] = "_"
            // if value == 0 {                                
            //     tttArr[vk][vvk] = "X"
            // } else if value == 1 {                
            //     tttArr[vk][vvk] = "O"
            // }
        }
    }

    // print random array
    PrintTableState(tttArr)

    var next_turn = "X"

    // did anyone win?
    var x_win = DidAnyoneWin(tttArr, "X")
    var o_win = DidAnyoneWin(tttArr, "O")

    for !(x_win || o_win) {
        if next_turn == "X" {
            fmt.Println("What position for X? ")
        } else {
            fmt.Println("What position for O? ")
        }

        fmt.Println("Put next move in format \"D3\"")

        next_move, _ := reader.ReadString('\n')
        fmt.Printf("Next move is %s\n", next_move)

        // var example_string = "D3\n"
        // fmt.Printf("ascii value of letter: %d\n", example_string[0])
        // fmt.Printf("ascii value of number: %d\n", example_string[1])

        fmt.Printf("")

        if next_move == "Exit\n" {
            break
        }

        if tttArr[next_move[1]-48][next_move[0]-65] != "_" {
            fmt.Printf("Not a valid position: %s %s\n", next_move, tttArr[next_move[1]-48][next_move[0]-65])
        } else if next_turn == "X" {
            tttArr[next_move[1]-48][next_move[0]-65] = next_turn
            next_turn = "O"
        } else {
            tttArr[next_move[1]-48][next_move[0]-65] = next_turn
            next_turn = "X"
        }



        PrintTableState(tttArr)

        x_win = DidAnyoneWin(tttArr, "X")
        o_win = DidAnyoneWin(tttArr, "O")
    }

    if x_win {
        fmt.Println("X won")
    }
    if o_win {
        fmt.Println("O won")
    }

    // exit program


    // TODO: V2

    // keep running program until a winner occurs, max <input> times


    // TODO: V3

    // able to run this program more quickly with threads or processes?


    // TODO: v4

    // play against self


    // TODO: v5 

    // play against computer that never loses







    // // Accept input
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("What is your move? ")
    // next_move, _ := reader.ReadString('\n')

    // // Only stop accepting input on string "exit"
    // for next_move != "exit\n" {
    //     fmt.Print("You moved to " + next_move)

    //     fmt.Println()
    //     fmt.Print("What is your move? ")
    //     next_move, _ = reader.ReadString('\n')
    // }


}

