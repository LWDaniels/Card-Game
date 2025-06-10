package scenes

import (
	"image/color"

	"github.com/LWDaniels/Card-Game/src/archetypes"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type MainMenuScene struct {
	World donburi.World
}

func NewMainMenuScene() *MainMenuScene {
	m := &MainMenuScene{donburi.NewWorld()}

	e := archetypes.Button.Spawn(m.World)
	im := components.Sprite.Get(e).Image
	im.Fill(color.White)
	transform.SetWorldPosition(e,
		math.NewVec2(float64(constants.WorldWidth()/2-im.Bounds().Dx()/2),
			float64(constants.WorldHeight()/2-im.Bounds().Dy()/2)))
	components.Interactable.Get(e).HoverCallback = func(entry *donburi.Entry) {
		components.Sprite.Get(entry).Image.Fill(color.RGBA{255, 0, 0, 255})
	}
	return m
}

func (m *MainMenuScene) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	hoverQuery := donburi.NewQuery(filter.Contains(components.Sprite, components.Interactable, transform.Transform))
	hoverQuery.Each(m.World, func(e *donburi.Entry) {
		g := utils.GetGeoM(e)
		g.Invert()
		localMouseX, localMouseY := g.Apply(float64(mouseX), float64(mouseY))
		bounds := components.Sprite.Get(e).Image.Bounds()
		if float64(bounds.Min.X) <= localMouseX && float64(bounds.Max.X) >= localMouseX &&
			float64(bounds.Min.Y) <= localMouseY && float64(bounds.Max.Y) >= localMouseY {
			components.Interactable.Get(e).HoverCallback(e)
		}
	})
	return nil
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {
	// will want to wrap stuff like this up in systems or something similar
	drawQuery := donburi.NewQuery(filter.Contains(components.Sprite))
	drawQuery.Each(m.World, func(e *donburi.Entry) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM = utils.GetGeoM(e)
		screen.DrawImage(components.Sprite.Get(e).Image, op)
	})
}
