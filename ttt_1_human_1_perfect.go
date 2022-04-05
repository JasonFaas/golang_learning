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

func ListAvailableMoves(tempArr [3][3]string) (*list.List) {
    fmt.Println("Availabe Moves:")
    l := list.New()
    for _, item_1 := range [3]int{0,1,2} {
        for _, item_2 := range [3]int{0,1,2} {
            if tempArr[item_2][item_1] == "_" {
                var temp_string = string(item_1 + 65) + string(item_2 + 48)
                fmt.Println(temp_string)

                var new_move = MoveTesting{
                    notation: temp_string,
                    x_coor: item_2,
                    y_coor: item_1,
                    move_letter: "_",
                    future_score: make(chan int),
                    actual_score: 0,
                }

                l.PushFront(&new_move)
            }
        }
    }
    fmt.Println("Availabe Moves Over\n\n")
    return l
}

func WouldAnyoneWinStruct(tempArr [3][3]string, move_consider *MoveTesting) () {
    var whathwat = 3
    move_consider.future_score <- whathwat

    //     if WouldAnyoneWin(tempArr, next_move.x_coor, next_move.y_coor, "X") == "X" {
    //         return next_move.notation
    //     }

    //     if WouldAnyoneWin(tempArr, next_move.x_coor, next_move.y_coor, "O") == "O" {
    //         blockingMoves.PushFront(next_move)
    //     }
}

func WouldAnyoneWin(tempArr [3][3]string, next_move_position_x int, next_move_position_y int, next_move_value string) (string) {
    tempArr[next_move_position_x][next_move_position_y] = next_move_value
    return DidAnyoneWin(tempArr)
}

func DidAnyoneWin(tempArr [3][3]string) (string) {
    time.Sleep(2 * time.Second)
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
    
    return "_"
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

    return available.Front().Value.(*MoveTesting)
}

func DecideMoveIfWinningOrRandom(available *list.List, tempArr [3][3]string) (string) {

    // TODO: Loop through all scenarios and determine if winning move available, then return that move

    var test_move = available.Front()
    

    for test_move != nil && available.Len() > 1 {
        var next_move = test_move.Value.(*MoveTesting)

        next_move.move_letter = "X"
        
        if test_move.Value.(*MoveTesting).move_letter != next_move.move_letter {
            os.Exit(1)            
        }

        go WouldAnyoneWinStruct(tempArr, next_move)

        test_move = test_move.Next()
    }

    var best_move = available.Front().Value.(*MoveTesting)
    best_move.actual_score = <-best_move.future_score

    test_move = available.Front().Next()
    for test_move != nil && available.Len() > 1 {
        var next_move = test_move.Value.(*MoveTesting)
        next_move.actual_score = <-next_move.future_score

        if next_move.actual_score > best_move.actual_score {
            best_move = next_move
        }

        test_move = test_move.Next()
    }

    // All moves of equal value, return first move in list
    return best_move.notation
}


func main() {

    reader := bufio.NewReader(os.Stdin)

    // double 3x3 array
    var tttArr [3][3]string

    // random population of array with blank/X/O
    rand.Seed(time.Now().UnixNano())
    for vk, v := range tttArr {
        for vvk, _ := range v {
            tttArr[vk][vvk] = "_"
        }
    }

    PrintTableState(tttArr)

    // Loop for turns
    var next_turn = "X"
    var available = ListAvailableMoves(tttArr)
    var winner = "_"
    for winner == "_" && available.Len() > 0 {
        var next_move = ""
        if next_turn == "X" {            

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

        winner = DidAnyoneWin(tttArr)

        available = ListAvailableMoves(tttArr)
    }

    if winner != "_" {
        fmt.Printf("%s won!!!\n", winner)
    } else {
        fmt.Printf("Tie :(\n")
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

