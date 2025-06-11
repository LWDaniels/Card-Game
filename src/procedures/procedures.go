package procedures

import (
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
)

var drawQuery = donburi.NewQuery(filter.Contains(components.Sprite))

func DrawSprites(w donburi.World, screen *ebiten.Image) {
	drawQuery.Each(w, func(e *donburi.Entry) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM = utils.GetGeoM(e)
		op.Filter = ebiten.FilterLinear // since we're not a pixel art game
		sprite := components.Sprite.Get(e)
		op.ColorScale = sprite.Tint
		screen.DrawImage(sprite.Image, op)
	})
}

var hoverQuery = donburi.NewQuery(filter.Contains(components.Interactable))

func TriggerInteractables(w donburi.World) {
	mouseX, mouseY := ebiten.CursorPosition()
	hoverQuery.Each(w, func(e *donburi.Entry) {
		g := utils.GetGeoM(e)
		g.Invert()
		localMouseX, localMouseY := g.Apply(float64(mouseX), float64(mouseY))
		bounds := components.Bounds(e)
		hovering := 0 <= localMouseX && bounds.X >= localMouseX &&
			0 <= localMouseY && bounds.Y >= localMouseY
		interactable := components.Interactable.Get(e)
		localMousePos := math.NewVec2(localMouseX, localMouseY)
		if hovering {
			if !interactable.Hovered {
				interactable.Hovered = true
				interactable.OnEnter(e, localMousePos)
			}
			interactable.DuringHover(e, localMousePos)
		} else if interactable.Hovered {
			interactable.Hovered = false
			interactable.OnExit(e, localMousePos)
		}
	})
}
