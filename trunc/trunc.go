package main

import "fmt"
import "math"

func main() {
	var userInput float64
	fmt.Println("enter a number: ")
	fmt.Scan(&userInput)
	fmt.Println(math.Floor(userInput))
}