package card

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	v "github.com/LWDaniels/Card-Game/basics/vec2"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	texture *eb.Image
	pos     v.Vec2
	vel     v.Vec2
}

func NewCard() *Card {
	c := Card{}
	c.texture = assets.GetTexture(textures.Gopher)
	c.pos = v.Zero()
	c.vel = v.One()
	return &c
}

func (c *Card) Update() {
	dt := 1.0 / float32(eb.TPS())
	dtV := v.FromF(dt)
	c.pos = v.Sum(c.pos, v.Product(c.vel, dtV))
}

func (c *Card) Texture() *eb.Image {
	return c.texture
}

func (c *Card) GeoM() eb.GeoM {
	m := eb.GeoM{}
	x, y := c.pos.F64()
	m.Translate(x, y)
	return m
}

func (c *Card) Z() float32 {
	return 0
}
