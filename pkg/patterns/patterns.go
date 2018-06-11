package patterns

import (
	"conway/pkg/world"
)

// Some sample automata, for testing, demo, etc.

type Seeder interface {
	Seed(worldptr *world.World)
}

type Methuselah struct{}

func (m Methuselah) Seed(worldptr *world.World) {
	w := *worldptr
	x, y := w.Center()
	w.SetCell(x, y)
	w.SetCell(x-1, y)
	w.SetCell(x, y-1)
	w.SetCell(x, y+1)
	w.SetCell(x+1, y+1)
}

type Triomino struct{}

func (m Triomino) Seed(worldptr *world.World) {
	w := *worldptr
	x, y := w.Center()
	w.SetCell(x, y)
	w.SetCell(x, y+1)
	w.SetCell(x, y-1)
}

type Glider struct{}

func (g Glider) Seed(worldptr *world.World) {
	w := *worldptr
	x, y := w.Center()
	w.SetCell(x, y)
	w.SetCell(x, y+1)
	w.SetCell(x, y-1)
	w.SetCell(x-1, y-1)
	w.SetCell(x-2, y)
}

type TranslatedGlider struct{}

func (g TranslatedGlider) Seed(worldPtr *world.World) {
	w := *worldPtr
	x, y := w.Center()
	x++
	y--
	w.SetCell(x, y)
	w.SetCell(x, y+1)
	w.SetCell(x, y-1)
	w.SetCell(x-1, y-1)
	w.SetCell(x-2, y)
}
