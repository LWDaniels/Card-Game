package components

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
	Tint  ebiten.ColorScale
}

// requires transform component
var Sprite = donburi.NewComponentType[SpriteData](SpriteData{assets.GetTexture(textures.Gopher), ebiten.ColorScale{}})

func InitSprite(e *donburi.Entry, Image *ebiten.Image) {
	Sprite.Get(e).Image = Image
}
