package main

import (
	"conway/pkg/automaton"
	"fmt"
	"conway/pkg/world"
	"conway/pkg/patterns"
)

func main() {
	c := automaton.NewConway(world.NewTorus(120, 40), patterns.Methuselah{})
	c.Seed()
	for ; ; {
		c.Tick()
		fmt.Println(c)
	}
}
