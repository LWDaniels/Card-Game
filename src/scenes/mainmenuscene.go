package scenes

import (
	"image/color"

	"github.com/LWDaniels/Card-Game/src/archetypes"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/procedures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type MainMenuScene struct {
	World donburi.World
}

func NewMainMenuScene() *MainMenuScene {
	m := &MainMenuScene{donburi.NewWorld()}

	parent := m.World.Entry(m.World.Create(transform.Transform))
	transform.SetWorldPosition(parent,
		math.NewVec2(float64(constants.WorldWidth()/2),
			float64(constants.WorldHeight()/2)))
	transform.SetWorldRotation(parent, 1)
	e := archetypes.Button.Spawn(m.World)
	transform.AppendChild(parent, e, false)
	im := components.Sprite.Get(e).Image
	t := transform.GetTransform(e)
	t.LocalPosition = math.NewVec2(float64(-im.Bounds().Dx()/2), float64(-im.Bounds().Dy()/2))
	components.Interactable.Get(e).HoverCallback = func(entry *donburi.Entry, localMousePos math.Vec2) {
		components.Sprite.Get(entry).Image.Fill(color.RGBA{255, 0, 0, 255})
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			SetNextScene(GameSceneID)
		}
	}
	return m
}

func (m *MainMenuScene) Update() error {
	procedures.TriggerInteractables(m.World)
	return nil
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {
	// will want to wrap stuff like this up in systems or something similar
	procedures.DrawSprites(m.World, screen)
}
