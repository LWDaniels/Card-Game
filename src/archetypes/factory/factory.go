package factory

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/src/archetypes"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

// shouldn't create a card by itself for sprite positioning reasons
func createCard(w donburi.World) *donburi.Entry {
	card := archetypes.Card.Spawn(w)
	im := assets.GetTexture(textures.BlackLotus)
	components.InitSprite(card, im)
	scale := math.NewVec2(.2, .2)
	components.InitTransform(card, scale,
		0, scale.Mul(math.NewVec2(float64(-im.Bounds().Dx()/2), float64(-im.Bounds().Dy()/2))))
	components.InitInteractable(card, func(e *donburi.Entry, localMousePos math.Vec2) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			parent, _ := transform.GetParent(e) // assumes a parent for reasons :)
			transform.GetTransform(parent).LocalRotation += 1
		}
	})

	return card
}

func CreateCardContainer(w donburi.World, pos math.Vec2) *donburi.Entry {
	parent := w.Entry(w.Create(transform.Transform))
	components.InitTransform(parent, math.NewVec2(1, 1), 0, pos)
	transform.AppendChild(parent, createCard(w), false)
	return parent
}
