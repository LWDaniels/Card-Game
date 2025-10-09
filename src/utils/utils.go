package utils

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

// source: https://www.youtube.com/watch?v=LSNQuFEDOyQ
func ExpDecayF(a, b, decay float64) float64 {
	return b + (a-b)*math.Exp(-decay/60)
}

func ExpDecayVec2(a, b dmath.Vec2, decay float64) dmath.Vec2 {
	return b.Add(a.Sub(b).MulScalar(math.Exp(-decay / 60)))
}

// TODO: various easing functions

func LerpF(a, b, interp float64) float64 {
	return a + (b-a)*interp
}

func LerpVec2(a, b dmath.Vec2, interp float64) dmath.Vec2 {
	return a.Add(b.Sub(a).MulScalar(interp))
}

// requires transform component
func GetGeoM(e *donburi.Entry) ebiten.GeoM {
	// not particularly efficient... idk why transform doesn't store this info already
	// (I could calculate it by hand from what it stores but I'm too lazy)
	// transform.World... stuff seems unreliable tbh
	g := ebiten.GeoM{}
	t := transform.GetTransform(e)
	g.Scale(t.LocalScale.X, t.LocalScale.Y)
	g.Rotate(t.LocalRotation)
	g.Translate(t.LocalPosition.X, t.LocalPosition.Y)

	if parent, ok := transform.GetParent(e); ok {
		g.Concat(GetGeoM(parent))
	}
	return g
}
