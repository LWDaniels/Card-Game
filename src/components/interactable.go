package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type InteractableData struct {
	Hovered       bool // whether the cursor is currently over this
	HoverCallback func(self *donburi.Entry, localMousePos math.Vec2)
}

// requires transform and sprite (should allow ninepatch as sprite alternative)
var Interactable = donburi.NewComponentType[InteractableData]()

func InitInteractable(e *donburi.Entry, HoverCallback func(*donburi.Entry, math.Vec2)) {
	i := Interactable.Get(e)
	i.Hovered = false
	i.HoverCallback = HoverCallback
}
