package scenes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/factory"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/procedures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type GameScene struct {
	World donburi.World
}

func NewGameScene() *GameScene {
	g := &GameScene{donburi.NewWorld()}
	factory.CreateCardContainer(g.World, math.NewVec2(float64(constants.WorldWidth()/2),
		float64(constants.WorldHeight()/2)))
	return g
}

func (g *GameScene) Update() error {
	procedures.TriggerInteractables(g.World)
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	procedures.DrawSprites(g.World, screen)
}
