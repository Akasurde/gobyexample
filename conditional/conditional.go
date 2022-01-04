package main

import "fmt"

func main() {
    var a int = 100
    var b int = 200

    var ret int
    ret = max(a, b)
    fmt.Printf("Max of a (%d) and b (%d) : %d\n", a, b, ret)
    x, y := swap(a, b)
    fmt.Printf("Swap a (%d) and b (%d) : a (%d) and b (%d)\n", a, b, x, y)
}

func max(num1, num2 int) int {
    var result int
    if (num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result
}

func swap(x, y int) (int, int) {
    return y, x
}
