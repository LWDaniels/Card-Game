package game

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/constants"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// note that this is a fixed update
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	op := &eb.DrawImageOptions{}

	screen.DrawImage(assets.GetTexture(textures.Gopher), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WorldWidth(), constants.WorldHeight()
}
