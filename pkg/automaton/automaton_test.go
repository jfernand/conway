package automaton_test

import (
	"testing"
	"conway/pkg/automaton"
	"conway/pkg/world"
	"conway/pkg/patterns"
)

// The glider pattern repeats itself, translated, after four iterations.
// These iterations go through all neighbourhood cases
func TestGenerations(t *testing.T) {
	glider := automaton.NewConway(world.NewTorus(20, 10), patterns.Glider{})
	glider1 := automaton.NewConway(world.NewTorus(20, 10), patterns.TranslatedGlider{})

	for i := 0; i < 4; i++ {
		glider.Tick()
	}

	if glider.Signature() != glider1.Signature() {
		t.Error("Glider glides as expected")
	}
}
