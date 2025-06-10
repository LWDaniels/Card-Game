package components

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
}

// requires transform component
var Sprite = donburi.NewComponentType[SpriteData](SpriteData{assets.GetTexture(textures.Gopher)})
