package main

import "flag"
import "fmt"

func main() {
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string svar")

    flag.Parse()
    fmt.Println("Word: ", *wordPtr)
    fmt.Println("Numb: ", *numbPtr)
    fmt.Println("Bool: ", *boolPtr)
    fmt.Println("svar: ", svar)
    fmt.Println("Tail: ", flag.Args())
}
