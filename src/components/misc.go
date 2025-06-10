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
