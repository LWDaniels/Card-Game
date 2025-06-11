package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func InitTransform(e *donburi.Entry, localScale math.Vec2, localRot float64, localPos math.Vec2) {
	t := transform.GetTransform(e)
	t.LocalScale = localScale
	t.LocalRotation = localRot
	t.LocalPosition = localPos
}

// returns a vec2 of (width, height) without any transform scaling
// will panic if shit goes wrong lol
func Bounds(e *donburi.Entry) math.Vec2 {
	if e.HasComponent(Sprite) {
		im := Sprite.Get(e).Image
		return math.NewVec2(float64(im.Bounds().Dx()), float64(im.Bounds().Dy()))
	} else if e.HasComponent(NinePatch) {
		bottomRight := NinePatch.Get(e).children[8]
		return transform.GetTransform(bottomRight).LocalPosition.Add(Bounds(bottomRight))
	} else {
		// hopefully don't reach here lol
		return math.NewVec2(0, 0)
	}
}
