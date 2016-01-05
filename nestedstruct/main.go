package main

import "fmt"

type Kitchen struct {
	numOfPlates int
}

type House struct {
	Kitchen
	numOfRooms int
}

func main() {
	h := House{Kitchen{10}, 3}
	fmt.Println("Number of rooms :", h.numOfRooms)
	fmt.Println("Number of Plates:", h.Kitchen.numOfPlates)
}
