package factory

import (
	"image"

	"github.com/LWDaniels/Card-Game/assets"
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/src/archetypes"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateZone(w donburi.World, topLeft math.Vec2, size image.Point) *donburi.Entry {
	zone := archetypes.Zone.Spawn(w)
	scale := float64(.2)
	scaledWidth, scaledHeight := int(float64(size.X)/scale), int(float64(size.Y)/scale)
	components.InitNinePatch(zone, assets.GetTexture(textures.Border), image.Pt(scaledWidth, scaledHeight))
	components.InitTransform(zone, math.NewVec2(scale, scale), 0, topLeft)
	components.InitInteractable(zone, func(e *donburi.Entry, v math.Vec2) {
		np := components.NinePatch.Get(e)
		tint := ebiten.ColorScale{}
		tint.Scale(1, 0, 0, 1)
		np.SetTint(tint)
	})
	return zone
}

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
