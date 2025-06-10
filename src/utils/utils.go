package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

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
		return g
	} else {
		return g
	}
}
