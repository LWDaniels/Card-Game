package scenes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/factory"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
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
	// really need to make some other function for this to avoid duplication
	mouseX, mouseY := ebiten.CursorPosition()
	hoverQuery := donburi.NewQuery(filter.Contains(components.Sprite, components.Interactable, transform.Transform))
	hoverQuery.Each(g.World, func(e *donburi.Entry) {
		g := utils.GetGeoM(e)
		g.Invert()
		localMouseX, localMouseY := g.Apply(float64(mouseX), float64(mouseY))
		bounds := components.Sprite.Get(e).Image.Bounds()
		if float64(bounds.Min.X) <= localMouseX && float64(bounds.Max.X) >= localMouseX &&
			float64(bounds.Min.Y) <= localMouseY && float64(bounds.Max.Y) >= localMouseY {
			components.Interactable.Get(e).HoverCallback(e, math.NewVec2(localMouseX, localMouseY))
		}
	})
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	// will want to wrap stuff like this up in systems or something similar
	drawQuery := donburi.NewQuery(filter.Contains(components.Sprite))
	drawQuery.Each(g.World, func(e *donburi.Entry) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM = utils.GetGeoM(e)
		screen.DrawImage(components.Sprite.Get(e).Image, op)
	})
}
