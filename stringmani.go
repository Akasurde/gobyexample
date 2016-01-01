package main

import "fmt"

func main() {
    var str string
    str = "Hello World"

    fmt.Printf("The str : %s\n", str)
    fmt.Printf("Hex Bytes :")
    for i:= 0; i < len(str); i++ {
        fmt.Printf("%x ", str[i])
    }
}
