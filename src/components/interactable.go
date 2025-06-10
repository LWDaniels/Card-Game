package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type InteractableData struct {
	HoverCallback func(self *donburi.Entry, localMousePos math.Vec2)
}

// requires transform and sprite
var Interactable = donburi.NewComponentType[InteractableData]()

func InitInteractable(e *donburi.Entry, HoverCallback func(*donburi.Entry, math.Vec2)) {
	Interactable.Get(e).HoverCallback = HoverCallback
}
