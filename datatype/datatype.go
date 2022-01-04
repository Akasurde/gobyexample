package main

import "fmt"

func main() {
    var x float64 = 20.0
    y := 42
    a, b, c := 10, 12.1, "foo"

    fmt.Println(x)
    fmt.Println(y)
    fmt.Printf("x is of type %T\n", x)
    fmt.Printf("y is of type %T\n", y)
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)

    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    area = LENGTH * WIDTH
    fmt.Printf("Area : %d\n", area)
}
