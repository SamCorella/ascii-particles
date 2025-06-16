package main

import (
	"fmt"
	"time"
)

type Particle struct {
	value string
	column int
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	clear()
	p := Particle{"X", 0}

	for col := 0; col < 10; col++ {
		fmt.Printf("\033[%dC", col)
		fmt.Print(p.value)
		time.Sleep(500 * time.Millisecond)
		clear()
	}
}