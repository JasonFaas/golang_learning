package main

import (
    "fmt"   
    "math/rand"
    "time"
    "bufio"
    "os"
    "container/list"
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

func ListAvailableMoves(tempArr [3][3]string) (*list.List) {
    fmt.Println("Availabe Moves:")
    l := list.New()
    for _, item_1 := range [3]int{0,1,2} {
        for _, item_2 := range [3]int{0,1,2} {
            if tempArr[item_2][item_1] == "_" {
                var temp_string = string(item_1 + 65) + string(item_2 + 48)
                fmt.Println(temp_string)
                l.PushFront(temp_string)
            }
        }
    }
    fmt.Println("Availabe Moves Over\n\n")
    return l
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

func DecideMoveRandom(available *list.List) (string) {
    var move_to_make_itr = rand.Intn(available.Len())
    fmt.Printf("Move to make random %d\n", move_to_make_itr)
    

    for move_to_make_itr > 0 {
        // fmt.Printf("Removing element %d\n", move_to_make_itr)
        available.Remove(available.Front())
        move_to_make_itr -= 1
    }


    fmt.Printf("Move for computer is: ")
    fmt.Println(available.Front().Value)

    return available.Front().Value.(string)
}

func DecideMoveIfWinningOrRandom(available *list.List, tempArr [3][3]string) (string) {
    blockingMoves := list.New()

    // TODO: Loop through all scenarios and determine if winning move available, then return that move
    var test_move = available.Front()
    for test_move != nil && available.Len() > 1 {
        var next_move = test_move.Value.(string)

        tempArr[next_move[1]-48][next_move[0]-65] = "X"
        if DidAnyoneWin(tempArr, "X") {
            return next_move
        }

        tempArr[next_move[1]-48][next_move[0]-65] = "O"
        if DidAnyoneWin(tempArr, "O") {
            // TODO: Instead of below, add blocking moves
            blockingMoves.PushFront(next_move)
        }


        tempArr[next_move[1]-48][next_move[0]-65] = "_"
        test_move = test_move.Next()

        fmt.Printf("%d blockingMoves\n", blockingMoves.Len())
    }

    if blockingMoves.Len() > 0 {
        fmt.Println("Getting blocking move")
        return blockingMoves.Front().Value.(string)
    }

    return DecideMoveRandom(available)
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

    for !(x_win || o_win || ListAvailableMoves(tttArr).Len() == 0) {
        var next_move = ""
        if next_turn == "X" {


            var available = ListAvailableMoves(tttArr)

            if available.Len() == 0 {
                fmt.Println("No avilable moves")
                break
            }

            // next_move = DecideMoveRandom(available)
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

