package logic

import (
	"github.com/LWDaniels/Card-Game/assets/textures"
	"github.com/LWDaniels/Card-Game/src/constants"
)

type CardTarget uint // could redo to be a *bool if needed

const (
	TargetLeft CardTarget = iota
	TargetRight
	TargetNone
)

// For an intance of a card, not for a type of card/card preset
type CardInstance struct {
	Id     uint64 // unique id for this instance; to be used for comparisons
	Level  int    // in [1,2,3]; could have a more restricted int type if desired
	Target CardTarget
	Preset *CardPreset
	// may need a modifications list
}

func NewInstance(preset *CardPreset, target CardTarget) *CardInstance {
	return &CardInstance{
		Id:     constants.NextID(),
		Level:  1,
		Target: target,
		Preset: preset,
	}
}

func (ci *CardInstance) Upgrade() (couldUpgrade bool) {
	prevLevel := ci.Level
	ci.Level = max(ci.Level+1, 3)
	return prevLevel != ci.Level
}

// originalCard is always the card that caused this effect, not the card this may be effecting
type Effect func(state *BoardState, casterIndex int, originalCard *CardInstance) // no idea what this should return if anything

type CardPreset struct {
	Name           string
	Text           string
	RequiresTarget bool
	Effects        map[Trigger]Effect // when a card resolves, it triggers its TriggerResolve effect, then populates the appropriate event listeners with the other effects
	// Effects could also instead be an Ability[] list, but its fine as-is where it transforms ig
	Texture textures.TextureKey
}

type Ability struct {
	// note: the difference between effects and abilities
	// are that abilities are on the stack (weird naming, ik)
	// thus, abilities have the effect params bound to them
	Trigger     Trigger // not actually sure if we need the trigger inside of here
	BoundEffect func()
}
