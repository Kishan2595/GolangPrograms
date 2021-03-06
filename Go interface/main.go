package main

import (
	"fmt"
)

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	var player1 BaseGuitarist
	player1.Name = "Anjani"
	player1.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Kishan"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player1)
	guitarists = append(guitarists, player2)
}
