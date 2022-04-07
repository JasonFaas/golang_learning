package main

import (
    // "fmt"   
    "math/rand"
    // "time"
    // "bufio"
    // "os"
    "container/list"
    // "strings"
)

func GetBotInput(tempArr [3][3]string, next_turn_letter string, first_move bool) (*MoveTesting) {
    var availableMoves = ListAvailableMoves(tempArr, first_move, next_turn_letter)
    // var move_to_make = DecideMoveRandom(availableMoves)

    var move_to_make = DecideMoveIfWinningOrRandom(availableMoves, tempArr)

    return move_to_make
}

func ListAvailableMoves(tempArr [3][3]string, board_open bool, next_turn_letter string) (*list.List) {
    // fmt.Println("Availabe Moves:")
    l := list.New()

    for _, item_1 := range [3]int{0,1,2} {
        for _, item_2 := range [3]int{0,1,2} {
            if tempArr[item_2][item_1] == "_" {
                var temp_string = string(item_1 + 65) + string(item_2 + 48)
                if board_open && !(temp_string == "A0" || temp_string == "A1" || temp_string == "B1") {
                    continue
                }
                // fmt.Println(temp_string)

                var new_move = MoveTesting{
                    notation: temp_string,
                    x_coor: item_2,
                    y_coor: item_1,
                    move_letter: next_turn_letter,
                    future_score: make(chan int),
                    actual_score: 0,
                }

                l.PushFront(&new_move)
            }
        }
    }
    // fmt.Println("Availabe Moves Over\n\n")
    return l
}


func WhoWinGoRoutine(tempArr [3][3]string, move_consider *MoveTesting) () {
    var result = 0

    tempArr[move_consider.x_coor][move_consider.y_coor] = move_consider.move_letter

    var thisMoveResult = DidAnyoneWin(tempArr)

    if thisMoveResult == move_consider.move_letter {
        result = 1
    } else if thisMoveResult == "0" {
        result = 0
    } else {
        var whatwhat = GetBotInput(tempArr, GetNextTurnLetter(move_consider.move_letter), false)
        result = whatwhat.actual_score * -1
    }
    move_consider.future_score <- result
    return
}



// func DecideMoveRandom(available *list.List) (*MoveTesting) {
//     var move_to_make_itr = rand.Intn(available.Len())
//     fmt.Printf("Move to make random %d\n", move_to_make_itr)
    

//     for move_to_make_itr > 0 {
//         // fmt.Printf("Removing element %d\n", move_to_make_itr)
//         available.Remove(available.Front())
//         move_to_make_itr -= 1
//     }


//     fmt.Printf("Move for computer is: ")
//     fmt.Println(available.Front().Value)

//     return available.Front().Value.(*MoveTesting)
// }



func DecideMoveIfWinningOrRandom(available *list.List, tempArr [3][3]string) (*MoveTesting) {
    
    var test_move = available.Front()
    for test_move != nil {
        var next_move = test_move.Value.(*MoveTesting)
        go WhoWinGoRoutine(tempArr, next_move)

        test_move = test_move.Next()
    }

    var best_move = available.Front().Value.(*MoveTesting)
    best_move.actual_score = <-best_move.future_score

    test_move = available.Front().Next()
    for test_move != nil {
        var next_move = test_move.Value.(*MoveTesting)
        next_move.actual_score = <-next_move.future_score

        if next_move.actual_score > best_move.actual_score {
            best_move = next_move
        } else if next_move.actual_score == best_move.actual_score && rand.Intn(2) == 1 {
            best_move = next_move
        }

        test_move = test_move.Next()
    }

    // All moves of equal value, return first move in list
    return best_move
}

