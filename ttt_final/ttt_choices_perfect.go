package main

import (
    "fmt"   
    "math/rand"
    "time"
    "bufio"
    "os"
    // "container/list"
    // "strings"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    rand.Seed(time.Now().UnixNano())

    var tttArr [3][3]string = InitTable()

    PrintTableState(tttArr)

    var human_count = HowManyHumans()
    var human_turn = UNDERSCORE
    if human_count == "1" {
        human_turn = WhichTurnIsHuman()
    }

    fmt.Printf("%s %s\n", human_count, human_turn)

    os.Exit(0)

    // Loop for turns
    var next_turn = "X"
    var available = ListAvailableMoves(tttArr, true)
    var winner = "_"
    for winner == "_" && available.Len() > 0 {
        var next_move = ""
        if next_turn == "X" {            

            if available.Len() == 0 {
                fmt.Println("No avilable moves")
                break
            }
            next_move = DecideMoveIfWinningOrRandom(available, tttArr)
        } else {
            fmt.Println("What position for O? ")

            fmt.Println("Put next move in format \"D3\"")

            read_string, _ := reader.ReadString('\n')
            next_move = read_string
            fmt.Printf("Next human move is %s\n", next_move)

            fmt.Printf("")

            if next_move == "Exit\n" {
                break
            }
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

        winner = DidAnyoneWin(tttArr)
        if winner != "_" {
            break
        }

        available = ListAvailableMoves(tttArr, false)
    }

    if winner != "_" {
        fmt.Printf("%s won!!!\n", winner)
    } else {
        fmt.Printf("Tie :(\n")
    }
}
