package world

import (
	"fmt"
)

type Coordinate uint32

// TODO implement a different world model, to validate this api

type World interface {
	GetNeighbours(x, y Coordinate) uint8
	GetCell(x, y Coordinate) uint8
	SetCell(x, y Coordinate)
	KillCell(x, y Coordinate)
	SpareCell(x, y Coordinate)
	Center() (Coordinate, Coordinate)
	Size() (Coordinate, Coordinate)
	Tick()
	Signature() uint64
}

// grid == current state as seen by e.g. GetNeighbours
// next grid == where the next state is computed, e.g. SetCell, etc.
type Torus struct {
	grid                    [][] uint8
	nextGrid                [][] uint8
	width, height           Coordinate
	signature, tmpSignature uint64
}

func (world Torus) Center() (Coordinate, Coordinate) {
	return world.width / 2, world.height / 2
}

func (world Torus) Size() (Coordinate, Coordinate) {
	return world.width, world.height
}

func (world Torus) GetNeighbours(x, y Coordinate) uint8 {
	xLower, yLower := world.normalize(x-1, y-1)
	xUpper, yUpper := world.normalize(x+1, y+1)
	return world.grid[xLower][yLower] +
		world.grid[x][yLower] +
		world.grid[xUpper][yLower] +
		world.grid[xLower][y] +
		world.grid[xUpper][y] +
		world.grid[xLower][yUpper] +
		world.grid[x][yUpper] +
		world.grid[xUpper][yUpper]
}

func (world Torus) normalize(x, y Coordinate) (Coordinate, Coordinate) {
	return (x + world.width) % world.width, (y + world.height) % world.height
}

func (world Torus) GetCell(x, y Coordinate) uint8 {
	localX, localY := world.normalize(x, y)
	return world.grid[localX][localY]
}

func (world *Torus) SetCell(x, y Coordinate) {
	localX, localY := world.normalize(x, y)
	world.nextGrid[localX][localY] = 1
	world.tmpSignature += uint64(localX)*uint64(world.height-1) + uint64(localY)
}

func (world *Torus) SpareCell(x, y Coordinate) {
	localX, localY := world.normalize(x, y)
	world.nextGrid[localX][localY] = world.grid[localX][localY]
	if cell := world.grid[localX][localY] == 1; cell {
		world.tmpSignature += uint64(localX)*uint64(world.height-1) + uint64(localY)
	}
}

func (world *Torus) KillCell(x, y Coordinate) {
	localX, localY := world.normalize(x, y)
	world.nextGrid[localX][localY] = 0
}

func (world *Torus) Tick() {
	world.nextGrid, world.grid = world.grid, world.nextGrid
	world.signature = world.tmpSignature
	world.tmpSignature = 0
}

func (world Torus) String() string {
	var i, j Coordinate
	var buf string
	for j = world.height - 1; j > 0; j-- {
		for i = 0; i < world.width; i++ {
			buf += fmt.Sprintf("%v", renderCell(world.grid[i][j]))
		}
		buf += "\n"
	}
	return buf
}

func (world Torus) Signature() uint64 {
	return world.signature
}

func renderCell(cell uint8) string {
	if cell == 1 {
		return "*"
	}
	return "_"
}

func NewTorus(w, h Coordinate) *Torus {

	f := &Torus{width: w, height: h, signature: 0}
	f.grid = make([][]uint8, f.width)
	for i := range f.grid {
		f.grid[i] = make([]uint8, f.height)
	}
	f.nextGrid = make([][]uint8, f.width)
	for i := range f.nextGrid {
		f.nextGrid[i] = make([]uint8, f.height)
	}
	return f
}
