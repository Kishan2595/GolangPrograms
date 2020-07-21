package main

import "fmt"

var initCounter int

func init() {
	fmt.Println("Called First in order of declaration")
	initCounter++
}

func init() {
	fmt.Println("Called Second in order of declaration")
	initCounter++
}

func main() {
	fmt.Println("Does nothing of any significance")
	fmt.Printf("Init counter: %d\n", initCounter)
}
