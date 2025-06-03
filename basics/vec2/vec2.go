package vec2

import "github.com/hajimehoshi/ebiten/v2"

// gonna keep this as a relatively constant struct rn so you cant do += or similar operations
type Vec2 struct {
	X, Y float32
}

func (v Vec2) F64() (x float64, y float64) {
	return float64(v.X), float64(v.Y)
}

func UnitUp() Vec2 {
	return Vec2{0, 1}
}

func UnitRight() Vec2 {
	return Vec2{1, 0}
}

func UnitLeft() Vec2 {
	return Vec2{-1, 0}
}

func UnitDown() Vec2 {
	return Vec2{0, -1}
}

// {0,0}
func Zero() Vec2 {
	return Vec2{0, 0}
}

// {1,1}
func One() Vec2 {
	return Vec2{1, 1}
}

func FromF(f float32) Vec2 {
	return Vec2{f, f}
}

func FromI(i int) Vec2 {
	f := float32(i)
	return FromF(f)
}

// a + b
func Sum(a, b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

// a * b
func Product(a, b Vec2) Vec2 {
	return Vec2{a.X * b.X, a.Y * b.Y}
}

// a - b
func Difference(a, b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

// a / b; no checking for div by 0
func Quotient(a, b Vec2) Vec2 {
	return Vec2{a.X / b.X, a.Y / b.Y}
}

func MatMult(mat ebiten.GeoM, vec Vec2) Vec2 {
	x, y := mat.Apply(float64(vec.X), float64(vec.Y))
	return Vec2{float32(x), float32(y)}
}
