package world_test

import (
	"testing"
	"conway/pkg/world"
)

func TestWrap(t *testing.T) {
	tiny := world.NewTorus(1, 1)
	tiny.SetCell(0, 0)
	tiny.Tick()
	if tiny.GetNeighbours(0, 0) != 8 {
		t.Error("The hallway of mirrors, I am my own neighbors", tiny.GetNeighbours(0, 0))
	}

	small := world.NewTorus(3, 3)
	small.SetCell(0, 0)
	small.SetCell(0, 1)
	small.SetCell(1, 0)

	small.Tick()
	if small.GetNeighbours(2, 2) != 3 {
		t.Error("Far corner neighborhood count is correct", small.GetNeighbours(2, 2))
	}
}

func TestNeighborCount(t *testing.T) {
	AssertCount(t, 0, 0, 1)
	AssertCount(t, 1, 0, 1)
	AssertCount(t, 2, 0, 1)
	AssertCount(t, 0, 1, 1)
	AssertCount(t, 1, 1, 0)
	AssertCount(t, 2, 1, 1)
	AssertCount(t, 0, 2, 1)
	AssertCount(t, 1, 2, 1)
	AssertCount(t, 2, 2, 1)
}

func AssertCount(t *testing.T, x, y world.Coordinate, n uint8) {
	small := world.NewTorus(3, 3)
	small.SetCell(x, y)
	small.Tick()

	if small.GetNeighbours(1, 1) != n {
		t.Error("Correct neighbors returned for ", x, y, n)
	}

}
