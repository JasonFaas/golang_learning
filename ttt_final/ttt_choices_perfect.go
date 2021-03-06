package main

import (
    "fmt"   
    "math/rand"
    "time"
    // "bufio"
    "os"
    // "container/list"
    // "strings"
)


func main() {
    rand.Seed(time.Now().UnixNano())

    var tttArr [3][3]string = InitTable()

    PrintTableState(tttArr)

    var human_count = HowManyHumans()
    var human_turn = UNDERSCORE
    if human_count == "1" {
        human_turn = WhichTurnIsHuman()
    }

    fmt.Printf("%s %s\n", human_count, human_turn)

    // Loop for turns
    var next_turn_letter = X_INPUT
    // var available = list.New() //ListAvailableMoves(tttArr, true)
    var winner = "_"
    var next_move_struct *MoveTesting
    var first_move = true
    for winner == "_" { //&& available.Len() > 0 {
        if IsHumanTurn(human_count, human_turn, next_turn_letter) {
            next_move_struct = GetHumanInput(next_turn_letter)
        } else {
            next_move_struct = GetBotInput(tttArr, next_turn_letter, first_move)
            first_move = false
        }

        if tttArr[next_move_struct.x_coor][next_move_struct.y_coor] != "_" {
            fmt.Printf("Not a valid position: %s %s\n", next_move_struct.notation)
            os.Exit(63)
        }
        
        tttArr[next_move_struct.x_coor][next_move_struct.y_coor] = next_turn_letter

        next_turn_letter = GetNextTurnLetter(next_turn_letter)

        PrintTableState(tttArr)

        winner = DidAnyoneWin(tttArr)
        if winner != "_" {
            break
        }
    }

    if winner == X_INPUT || winner == O_INPUT {
        fmt.Printf("%s won!!!\n", winner)
    } else {
        fmt.Printf("Tie :(\n")
    }
}
