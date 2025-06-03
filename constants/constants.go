package constants

const (
	// aspect = screen height / width
	DefaultAspect float32 = 16.0 / 9.0 // can maybe make this not constant if we want a modifiable window size
	worldHeight   int     = 720
)

func WorldWidth() int {
	return int(float32(worldHeight) * DefaultAspect)
}

func WorldHeight() int {
	return worldHeight
}
