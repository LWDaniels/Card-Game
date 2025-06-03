package card

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	v "github.com/LWDaniels/Card-Game/basics/vec2"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	texture *eb.Image
	// center point
	pos v.Vec2
	vel v.Vec2

	// width and height
	size v.Vec2
}

func NewCard() *Card {
	c := Card{}
	c.texture = assets.GetTexture(textures.Gopher)
	c.pos = v.Zero()
	c.vel = v.Zero()
	c.size = v.NewVec2(50, 75)
	return &c
}

func (c *Card) Pos() v.Vec2 {
	return c.pos
}
func (c *Card) Vel() v.Vec2 {
	return c.vel
}
func (c *Card) SetPos(pos v.Vec2) {
	c.pos = pos
}
func (c *Card) SetVel(vel v.Vec2) {
	c.vel = vel
}

func (c *Card) Update() error {
	dt := 1.0 / float32(eb.TPS())
	dtV := v.FromF(dt)
	c.pos = v.Sum(c.pos, v.Product(c.vel, dtV))
	return nil
}

func (c *Card) InBounds(worldPos v.Vec2) bool {
	m := c.GeoM()
	m.Invert()
	localPos := v.MatMult(m, worldPos)
	return localPos.X >= 0 && localPos.Y >= 0 && localPos.X <= c.size.X && localPos.Y <= c.size.Y
}

func (c *Card) Texture() *eb.Image {
	return c.texture
}

func (c *Card) TexScale() (float64, float64) {
	return float64(c.size.X) / float64(c.texture.Bounds().Dx()),
		float64(c.size.Y) / float64(c.texture.Bounds().Dy())
}

// maybe put a function for global GeoM if we add parents

func (c *Card) GeoM() eb.GeoM {
	m := eb.GeoM{}
	x, y := v.Difference(c.pos, v.Product(v.FromF(.5), c.size)).F64()
	m.Translate(x, y)
	return m
}

func (c *Card) Z() float32 {
	return 0
}
