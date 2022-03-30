package main

import "fmt"

func main() {
    // This is a single line comment
    /* 
        This is a block comment 
    */
    var str1 = "String Variable"
    fmt.Println(str1)
    // fmt.Printf(str1)

    for k, v := range str1 {
        fmt.Printf("k：%d,v：%c == %d\n", k, v, v)
    }
}

