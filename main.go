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

func clear() {
	fmt.Print("\033[H\033[2J")
}

func draw(particles []Particle) {
	for _, p := range particles {
		fmt.Printf("\033[%d;%dH", p.row, p.column)
		fmt.Printf(p.value)
	}
}

func main() {
	clear()
	p1 := Particle{"X", 0, 50}
	p2 := Particle{"*", 5, 20}

	scene := []Particle{p1, p2}

	for range 10 {
		for p := range scene {
			scene[p].column++
		}
		draw(scene)
		time.Sleep(500 * time.Millisecond)
		clear()
	}
}