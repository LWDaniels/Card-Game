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
	g := ebiten.GeoM{}
	scale := transform.WorldScale(e)
	rot := transform.WorldRotation(e)
	pos := transform.WorldPosition(e)
	g.Scale(scale.X, scale.Y)
	g.Rotate(rot)
	g.Translate(pos.X, pos.Y)

	return g
}
