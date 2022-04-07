package main

import (
    "fmt"   
    "math/rand"
    // "time"
    // "bufio"
    "os"
    "container/list"
    // "strings"
)

func GetBotInput(tempArr [3][3]string, next_turn_letter string, first_move bool) (*MoveTesting) {
    var availableMoves = ListAvailableMoves(tempArr, first_move)
    return DecideMoveRandom(availableMoves)
}

func ListAvailableMoves(tempArr [3][3]string, board_open bool) (*list.List) {
    fmt.Println("Availabe Moves:")
    l := list.New()


    for _, item_1 := range [3]int{0,1,2} {
        for _, item_2 := range [3]int{0,1,2} {
            if tempArr[item_2][item_1] == "_" {
                var temp_string = string(item_1 + 65) + string(item_2 + 48)
                if board_open && !(temp_string == "A0" || temp_string == "A1" || temp_string == "B1") {
                    continue
                }
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


func WouldAnyoneWinStruct(tempArr [3][3]string, move_consider *MoveTesting, otherMoves *list.List) () {
    var result = 0


    // TODO: these should be gorountines as well
    if WouldAnyoneWin(tempArr, move_consider.x_coor, move_consider.y_coor, "X") == "X" {
        result = 100
    } else if WouldAnyoneWin(tempArr, move_consider.x_coor, move_consider.y_coor, "O") == "O" {
        result = 50
    } else {
        // TODO: One more step down based on otherMoves
        // Assuming:
        // * otherMoves.Len() > 2
        // * computer goes at indicated location
        // * human goes at list of other locations
        // How many positions remain where computer wins? (this is perfect, go here)
        // How many positions remain where human wins? (block this MOVE!)
        tempArr[move_consider.x_coor][move_consider.y_coor] = "X"
        
        //////////////////////// TODO: otherMoves is still not VALID!!! though maybe can just check if already filled in
        var human_test_move = otherMoves.Front()
        fmt.Println("Starting List")
        for human_test_move != nil && otherMoves.Len() > 1 {
            fmt.Println("1sttt")
            var human_move = human_test_move.Value.(*MoveTesting)

            if human_move.x_coor == move_consider.x_coor && human_move.x_coor == move_consider.y_coor {
                human_test_move = human_test_move.Next()
                continue
            }
            tempArr[human_move.x_coor][human_move.y_coor] = "O"
            
            var second_cmp_move = otherMoves.Front()
            for second_cmp_move != nil {
                fmt.Println("2nd")
                var cmp_move = second_cmp_move.Value.(*MoveTesting)
                if cmp_move.x_coor == move_consider.x_coor && cmp_move.x_coor == move_consider.y_coor || human_move.x_coor == cmp_move.x_coor && human_move.x_coor == cmp_move.y_coor {
                    second_cmp_move = second_cmp_move.Next()
                    continue
                }
                tempArr[cmp_move.x_coor][cmp_move.y_coor] = "X"

                if WouldAnyoneWin(tempArr, move_consider.x_coor, move_consider.y_coor, "X") == "X" {
                    result += 5
                } else if WouldAnyoneWin(tempArr, move_consider.x_coor, move_consider.y_coor, "O") == "O" {
                    result -= 10
                }

                tempArr[cmp_move.x_coor][cmp_move.y_coor] = "_"

                second_cmp_move = second_cmp_move.Next()
            }
                

            tempArr[human_move.x_coor][human_move.y_coor] = "_"

            human_test_move = human_test_move.Next()
        } 
    }

    move_consider.future_score <- result

}

func WouldAnyoneWin(tempArr [3][3]string, next_move_position_x int, next_move_position_y int, next_move_value string) (string) {
    tempArr[next_move_position_x][next_move_position_y] = next_move_value
    return DidAnyoneWin(tempArr)
}



func DecideMoveRandom(available *list.List) (*MoveTesting) {
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

    
    // var otherMoves = 

    // var notTheseMoves := list.New()
    // var temp_move = available.Front()
    // for temp_move != nil && available.Len() > 1 {
    //     otherMoves.PushFront(available.Front().Value)
    // }
    
    var test_move = available.Front()
    for test_move != nil && available.Len() > 1 {
        var next_move = test_move.Value.(*MoveTesting)

        next_move.move_letter = "X"
        
        if test_move.Value.(*MoveTesting).move_letter != next_move.move_letter {
            os.Exit(1)            
        }

        go WouldAnyoneWinStruct(tempArr, next_move, available)

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
        } else if next_move.actual_score == best_move.actual_score && rand.Intn(2) == 1 {
            best_move = next_move
        }

        test_move = test_move.Next()
    }

    // All moves of equal value, return first move in list
    return best_move.notation
}

