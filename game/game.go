package game

import (
	"sort"

	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/constants"
	"github.com/LWDaniels/Card-Game/game/card"
	"github.com/LWDaniels/Card-Game/game/item"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	items []item.Item
}

type Drawable struct {
	GeoM    eb.GeoM
	Texture *eb.Image
	Z       float32
}

func NewGame() *Game {
	g := Game{}
	g.items = make([]item.Item, 0)
	g.items = append(g.items, card.NewCard())

	return &g
}

// note that this is a fixed update
func (g *Game) Update() error {
	g.HandleInput()

	for n := range g.items {
		g.items[n].Update()
	}
	return nil
}

func (g *Game) HandleInput() {
	if !inpututil.IsMouseButtonJustPressed(eb.MouseButton0) {
		return
	}

	// will need to change for mobile
	mouseXInt, mouseYInt := eb.CursorPosition()
	mousePos := vec2.Vec2{X: float32(mouseXInt), Y: float32(mouseYInt)}
	for n := range g.items {
		switch g.items[n].(type) {
		// add more types if desired
		case *card.Card:
			c := g.items[n].(*card.Card)
			if c.InBounds(mousePos) {
				// shift right by 10
				c.SetPos(vec2.Sum(c.Pos(), vec2.FromF(10)))
			}
		}
	}
}

func (g *Game) Draw(screen *eb.Image) {
	op := &eb.DrawImageOptions{}

	drawables := make([]Drawable, len(g.items))
	for n := range g.items {
		texScaleX, texScaleY := g.items[n].TexScale()
		texScaleM := eb.GeoM{}
		texScaleM.Scale(texScaleX, texScaleY)
		drawables[n].GeoM = texScaleM
		drawables[n].GeoM.Concat(g.items[n].GeoM())
		drawables[n].Texture = g.items[n].Texture()
		drawables[n].Z = g.items[n].Z()
	}

	sort.Slice(drawables, func(a, b int) bool { return drawables[a].Z < drawables[b].Z })

	for n := range drawables {
		op.GeoM = drawables[n].GeoM
		screen.DrawImage(drawables[n].Texture, op)
	}
	screen.DrawImage(assets.GetTexture(textures.Gopher), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WorldWidth(), constants.WorldHeight()
}
