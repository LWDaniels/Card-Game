package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type GameScene struct {
	World donburi.World
}

func NewGameScene() *GameScene {
	g := &GameScene{donburi.NewWorld()}
	return g
}

func (g *GameScene) Update() error {
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {

}
