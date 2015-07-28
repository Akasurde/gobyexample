package main

import "fmt"

func main() {
	beyondhello()
}

func beyondhello() {
	var x int
	x = 3 
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("Sum: ", sum, "\nProduct: ", prod)
}

func learnMultiple(x, y int) (sum, prod int){
	return x + y, x * y
}
