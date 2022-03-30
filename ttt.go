package main

import (
    "fmt"   
    "math/rand"
    "time"
    // "bufio"
    // "os"
    // "strings"
)

func PrintTableState(tempArr [3][3]string) () {
    for itr, _ := range tempArr {
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

    // TODO: V1

    // double 3x3 array
    var tttArr [3][3]string

    // random population of array with blank/X/O
    rand.Seed(time.Now().UnixNano())
    for vk, v := range tttArr {
        for vvk, _ := range v {
            var value = rand.Intn(2)
            if value == 0 {                
                tttArr[vk][vvk] = "X"
            } else {                
                tttArr[vk][vvk] = "O"
            }
        }
    }

    // print random array
    PrintTableState(tttArr)

    // did anyone win?

    if DidAnyoneWin(tttArr, "X") {
        fmt.Println("X won")
    }
    if DidAnyoneWin(tttArr, "O") {
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

