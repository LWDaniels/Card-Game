package sprite

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/game/transform"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*transform.TransformContainerImplementer
	Texture *eb.Image
}

func NewSprite(texKey textures.TextureKey) Sprite {
	return NewSpriteX(texKey, vec2.One(), 0, vec2.Zero())
}

// new sprite with extra parameters; parameters applied in SRT order
// anchor is in ([0,1],[0,1]) pref, top left is (0,0)
func NewSpriteX(texKey textures.TextureKey, scale vec2.Vec2, rotation float32, anchor vec2.Vec2) Sprite {
	im := assets.GetTexture(texKey)
	imSize := im.Bounds().Size()
	return Sprite{transform.NewTCI(scale, rotation,
		vec2.Product(vec2.NewVec2(float32(-imSize.X), float32(-imSize.Y)), anchor), 0),
		im}
}

// can add other stuff but for now I'll just leave it public and have it be
