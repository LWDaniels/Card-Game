package main

import (
	"log"

	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/constants"
	"github.com/LWDaniels/Card-Game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.WorldWidth(), constants.WorldHeight())
	ebiten.SetWindowTitle("Card Game Prototype")

	// just loading everything on startup rn but not smart obv
	assets.LoadAll()

	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}

	// not sure if this is necessary but I don't want to memory leak on GPU :)
	assets.UnloadTextures()
}
