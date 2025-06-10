package components

import (
	"github.com/yohamta/donburi"
)

type InteractableData struct {
	HoverCallback func(*donburi.Entry)
}

// requires transform and sprite
var Interactable = donburi.NewComponentType[InteractableData]()
