package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type InteractableData struct {
	Hovered     bool // whether the cursor is currently over this
	OnEnter     func(self *donburi.Entry, localMousePos math.Vec2)
	DuringHover func(self *donburi.Entry, localMousePos math.Vec2)
	OnExit      func(self *donburi.Entry, localMousePos math.Vec2)
	// maybe should allow some game parameters like hand (idk an easy way to do this; maybe pass donburi entry?)
}

func none(self *donburi.Entry, localMousePos math.Vec2) {
}

// requires transform and sprite (should allow ninepatch as sprite alternative)
var Interactable = donburi.NewComponentType[InteractableData](InteractableData{false, none, none, none})

func InitInteractable(e *donburi.Entry, OnEnter, DuringHover, OnExit func(self *donburi.Entry, localMousePos math.Vec2)) {
	i := Interactable.Get(e)
	i.Hovered = false
	i.OnEnter, i.DuringHover, i.OnExit = OnEnter, DuringHover, OnExit
}
