package factory

import (
	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/src/archetypes"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateCard(w donburi.World, pos math.Vec2) *donburi.Entry {
	card := archetypes.Card.Spawn(w)
	components.InitTransform(card, math.NewVec2(1, 1), 0, pos)
	children, _ := transform.GetChildren(card)
	child := children[0] // only one child upon creation
	components.InitCard(card, child)

	im := assets.GetTexture(textures.BlackLotus)
	components.InitSprite(child, im)
	scale := math.NewVec2(.12, .12)
	components.InitTransform(child, scale,
		0, scale.Mul(math.NewVec2(float64(-im.Bounds().Dx()/2), float64(-im.Bounds().Dy()/2))))
	components.InitInteractable(child, func(e *donburi.Entry, localMousePos math.Vec2) {
		// just setting held value and letting scene handle the rest for now
	})

	return card
}
