package main

import (
    "fmt"   
    // "math/rand"
    "time"
    // "bufio"
    // "os"
    // "container/list"
    // "strings"
)

const X_INPUT string = "X"
const O_INPUT string = "O"
const UNDERSCORE string = "_"



type MoveTesting struct {
    notation string
    x_coor int
    y_coor int
    move_letter string
    future_score chan int
    actual_score int
}


func PrintTableState(tempArr [3][3]string) () {
    fmt.Println("   A B C")
    for itr, _ := range tempArr {
        fmt.Printf("%d", itr)
        fmt.Printf(" ")
        fmt.Println(tempArr[itr])
    }
}




func DidAnyoneWin(tempArr [3][3]string) (string) {
    // time.Sleep(1 * time.Second)
    for _, whoWon := range [2]string{"X", "O"} {
        for _, item := range [3]int{0,1,2} {
            if tempArr[item][0] == whoWon && tempArr[item][1] == whoWon && tempArr[item][2] == whoWon {
                return whoWon
            }
            if tempArr[0][item] == whoWon && tempArr[1][item] == whoWon && tempArr[2][item] == whoWon {
                return whoWon
            }
        }

        if tempArr[0][0] == whoWon && tempArr[1][1] == whoWon && tempArr[2][2] == whoWon {
            return whoWon
        }
        if tempArr[0][2 - 0] == whoWon && tempArr[1][2 - 1] == whoWon && tempArr[2][2 - 2] == whoWon {
            return whoWon
        }
    }

    var available = 0

    for x, _ := range [3]int{0,1,2} {
        for y, _ := range [3]int{0,1,2} {
            if tempArr[x][y] == "_" {
                available += 1
            }
        }
    }
    
    if available > 0 {
        return "_"
    } else {
        return "0"        
    }

}

func InitTable() ([3][3]string) {
    var returnMe [3][3]string

    for vk, v := range returnMe {
        for vvk, _ := range v {
            returnMe[vk][vvk] = "_"
        }
    }

    return returnMe
}

func IsHumanTurn(human_count string, human_turn string, next_turn string) (bool) {
    return human_count == "2" || (human_count == "1" && human_turn == next_turn)
}

func transformHumanInputToStruct(stringInput string) (*MoveTesting) {
    var new_move = MoveTesting{
        notation: stringInput,
        x_coor: int(stringInput[1]-48),
        y_coor: int(stringInput[0]-65),
        move_letter: "_",
        future_score: make(chan int),
        actual_score: 0,
    }

    return &new_move
}

func GetNextTurnLetter(next_turn_letter string) (string) {
    if next_turn_letter == X_INPUT {
        return O_INPUT
    } else {
        return X_INPUT
    }
}
