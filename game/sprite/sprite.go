package sprite

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/basics/transform"
	"github.com/LWDaniels/Card-Game/basics/vec2"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*transform.TransformContainerDefault
	Texture *eb.Image
}

func NewSprite(texKey textures.TextureKey) Sprite {
	return NewSpriteX(texKey, vec2.One(), 0, vec2.Zero())
}

// new sprite with extra parameters; parameters applied in SRT order
// anchor should ([0,1],[0,1]) pref, top left is (0,0) and coordinates are after scaling
func NewSpriteX(texKey textures.TextureKey, scale vec2.Vec2, rotation float32, anchor vec2.Vec2) Sprite {
	im := assets.GetTexture(texKey)
	imSize := im.Bounds().Size()
	anchor = vec2.Product(anchor, scale)
	return Sprite{transform.NewTCD(scale, rotation,
		vec2.Product(vec2.NewVec2(float32(-imSize.X), float32(-imSize.Y)), anchor), 0),
		im}
}

// can add other stuff but for now I'll just leave it public and have it be
