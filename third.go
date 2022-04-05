package main

import (
    "fmt"   
    // "bufio"
    // "os"
    // "strings"
)

type TestStruct struct {
    first string
    second int
}

func main() {

    accept := "what"
    var more = "than"
    accept = "hello"
    more = "less"
    fmt.Printf("%s\n", accept)
    fmt.Printf("%s\n", more)

    var first_struct = &TestStruct{
        first: "world",
        second: 32,
    }

    var another_reference = first_struct
    another_reference.first = "no _way"

    fmt.Printf("%s\n", first_struct.first)
}

