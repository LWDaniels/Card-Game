package item

import "github.com/hajimehoshi/ebiten/v2"

type Item interface {
	// order of what gets updated first is undefined
	Update()

	// will add scene graph system later
	// Parent() *Item
	// Children() []*Item

	// the texture to draw with; if nil then don't draw
	Texture() *ebiten.Image

	// the transform to draw with (relative to parent); if nil then don't draw
	GeoM() ebiten.GeoM

	// returns the Z position of the item (relative to parent)
	// more negative is closer to the screen
	// (rendered last) while more positive is further from the screen (rendered first)
	// tied Z will draw in undefined order
	Z() float32
}
