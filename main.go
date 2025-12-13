package main

import (
	"fmt"
	"time"
)

type Particle struct {
	value string
	row int
	column int
}

type Scene struct {
	particles []Particle
	width int
	height int
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func draw(sc Scene) {
	for _, p := range sc.particles {
		fmt.Printf("\033[%d;%dH", p.row, p.column)
		fmt.Printf(p.value)
	}
	time.Sleep(500 * time.Millisecond)
}

func snow() {
	scene = Scene{[]Particle{}, 100, 100}
	big_flake := Particle{"X", 0, 0}
	small_flake := Particle{"*", 0, 0}
}

func main() {
	clear()
}