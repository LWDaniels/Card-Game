package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type MainMenuScene struct {
	World *donburi.World
}

func (m *MainMenuScene) Update() error {
	return nil
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {

}
