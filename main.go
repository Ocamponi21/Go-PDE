package main

import "fmt"

//this simple application demonstrates different components in Go
// To start it is vital that we import the fmt library that hold the formatting for the print statements in Go
// This function is iterating throught a series of number displaying them as they go.
// When it reaches lucky number 4 it will display a special output.
//This uses a combination of for loop and if statement to demonstarte simple decision structures
func main() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Println("Hello Lucky Number 4")
		}
		if i != 4 {
			fmt.Print("Hello: ")
			fmt.Println(i)
		}
	}
}
