package stressgen

import (
	"math/rand"
)

// GenerateConfigObject generates and returns a single configuration object,
// using the given random number generator to choose what kind of object
// to return and how to populate it.
func GenerateConfigObject(rnd *rand.Rand, namer *Namer) ConfigObject {
	const (
		chooseVariable int = 0
		chooseOutput   int = 1
	)
	which := decideIndex(rnd, []int{
		chooseVariable: 1,
		chooseOutput:   1,
	})
	switch which {
	case chooseVariable:
		return GenerateConfigVariable(rnd, namer)
	case chooseOutput:
		return GenerateConfigOutput(rnd, namer)
	default:
		// This suggests either a bug in decideIndex or in our call
		// to decideIndex.
		panic("invalid decision")
	}
}
