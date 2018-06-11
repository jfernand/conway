package automaton

import (
	"conway/pkg/world"
	"conway/pkg/patterns"
	"fmt"
)

// TODO detect loops through signatures, and stop
// TODO implement rules other than Conway's
type Automaton interface {
	Seed()
	Tick()
	Signature()
}

type Conway struct {
	world  world.World
	seeder patterns.Seeder
}

func (automaton Conway) String() string {
	return fmt.Sprint(automaton.world)
}

func (automaton Conway) Seed() {
	automaton.seeder.Seed(&automaton.world)
	automaton.world.Tick() // Shift the seeded pattern into the current grid
}

func (automaton Conway) Tick() {
	w, h := automaton.world.Size()
	var i, j world.Coordinate
	for i = 0; i < w; i++ {
		for j = 0; j < h; j++ {
			automaton.update(i, j)
		}
	}
	automaton.world.Tick()
}

func (automaton Conway) Signature() uint64 {
	return automaton.world.Signature()
}

func (automaton Conway) update(i, j world.Coordinate) {
	switch numNeighbors := automaton.world.GetNeighbours(i, j); numNeighbors {
	case 0, 1:
		automaton.world.KillCell(i, j)
	case 2:
		automaton.world.SpareCell(i, j)
	case 3:
		automaton.world.SetCell(i, j)
	case 4, 5, 6, 7, 8:
		automaton.world.KillCell(i, j)
	}
}

func NewConway(world world.World, seeder patterns.Seeder) *Conway {
	automaton := Conway{world, seeder}
	automaton.Seed()
	return &automaton
}
