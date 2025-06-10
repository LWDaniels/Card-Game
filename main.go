package main

import (
	"image"
	"log"

	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(*ebiten.Image)
}

type Game struct {
	scene  Scene
	bounds image.Rectangle
}

func NewGame() *Game {
	g := &Game{
		scene:  &scenes.MainMenuScene{},
		bounds: image.Rectangle{},
	}
	return g
}

func (g *Game) Update() error {
	return g.scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(constants.WorldWidth(), constants.WorldHeight())
	ebiten.SetWindowTitle("Card Game Prototype")

	// just loading everything on startup rn but not smart obv
	assets.LoadAll()

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}

	// not sure if this is necessary but I don't want to memory leak on GPU :)
	assets.UnloadTextures()
}
