package constants

import "math"

const (
	// aspect = screen height / width
	DefaultAspect float32 = 19.5 / 9.0 // can maybe make this not constant if we want a modifiable window size; aspect ratio is for iphone reference rn
	worldHeight   int     = 720
)

func WorldWidth() int {
	return int(math.Floor(float64(worldHeight) / float64(DefaultAspect)))
}

func WorldHeight() int {
	return worldHeight
}

// tick this up with global IDs
var idCounter uint64 = 0

func NextID() uint64 {
	idCounter++
	return idCounter
}
