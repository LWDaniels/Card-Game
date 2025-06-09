package card

import (
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/basics/vec2"
	v "github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/game/box"
	"github.com/LWDaniels/Card-Game/game/sprite"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Card struct {
	*box.BoxContainerDefault
	sprite *sprite.Sprite
	vel    v.Vec2
}

func NewCard() *Card {
	c := &Card{}
	c.BoxContainerDefault = box.NewBCDDefault()
	c.BoxContainerDefault.SetSize(v.NewVec2(50, 75))
	c.Trans.Scale = v.FromF(1)
	s := sprite.NewSpriteX(textures.Gopher, v.FromF(.2), 0, v.FromF(0))
	c.sprite = &s
	c.AddChild(c.sprite)
	c.vel = v.Zero()
	return c
}

func (c *Card) Update() error {
	dt := 1.0 / float32(eb.TPS())
	dtV := v.FromF(dt)
	c.Trans.Pos = v.Sum(c.Trans.Pos, v.Product(c.vel, dtV))

	if !inpututil.IsMouseButtonJustPressed(eb.MouseButtonLeft) {
		return nil
	}

	// will need to change for mobile
	mouseXInt, mouseYInt := eb.CursorPosition()
	mousePos := vec2.Vec2{X: float32(mouseXInt), Y: float32(mouseYInt)}

	if c.InBounds(mousePos) {
		c.sprite.Texture.Clear()
	}

	return nil
}
