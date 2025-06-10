package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
}

// requires transform component
var Sprite = donburi.NewComponentType[SpriteData](SpriteData{ebiten.NewImage(50, 50)})
