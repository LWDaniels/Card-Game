package card

import (
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/basics/transform"
	v "github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/game/sprite"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	*transform.TransformContainerImplementer
	sprite *sprite.Sprite
	vel    v.Vec2

	// width and height
	size v.Vec2
}

func NewCard() *Card {
	c := &Card{}
	c.TransformContainerImplementer = transform.NewTCIDefault()
	c.Trans.Scale = v.FromF(1)
	s := sprite.NewSpriteX(textures.Gopher, v.FromF(.2), 0, v.NewVec2(.5, .8))
	c.sprite = &s
	c.AddChild(c.sprite)
	c.vel = v.Zero()
	c.size = v.NewVec2(50, 75)
	return c
}

func (c *Card) Update() error {
	dt := 1.0 / float32(eb.TPS())
	dtV := v.FromF(dt)
	c.Trans.Pos = v.Sum(c.Trans.Pos, v.Product(c.vel, dtV))
	return nil
}

func (c *Card) InBounds(worldPos v.Vec2) bool {
	m := c.Trans.GeoM()
	m.Invert()
	localPos := v.MatMult(m, worldPos)
	return localPos.X >= 0 && localPos.Y >= 0 && localPos.X <= c.size.X && localPos.Y <= c.size.Y
}
