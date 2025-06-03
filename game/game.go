package game

import (
	"sort"

	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/constants"
	"github.com/LWDaniels/Card-Game/game/card"
	"github.com/LWDaniels/Card-Game/game/item"
	eb "github.com/hajimehoshi/ebiten/v2"
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
	for n := range g.items {
		g.items[n].Update()
	}
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	op := &eb.DrawImageOptions{}

	drawables := make([]Drawable, len(g.items))
	for n := range g.items {
		drawables[n].GeoM = g.items[n].GeoM()
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
