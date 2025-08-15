package logic

type CardTarget uint // could redo to be a *bool if needed

const (
	TargetLeft CardTarget = iota
	TargetRight
	TargetNone
)

// For an intance of a card, not for a type of card/card preset
type CardInstance struct {
	Level  int // in [1,2,3]; could have a more restricted int type if desired
	Target CardTarget
	Preset *CardPreset
	// may need a modifications list
}

func (ci *CardInstance) Upgrade() (couldUpgrade bool) {
	prevLevel := ci.Level
	ci.Level = max(ci.Level+1, 3)
	return prevLevel != ci.Level
}

// originalCard is always the card that caused this effect, not the card this may be effecting
type Effect func(state *BoardState, activePlayerIndex int, originalCard *CardInstance, triggerParameters ...any) // no idea what this should return

type CardPreset struct {
	Name           string
	Text           string
	RequiresTarget bool
	Effects        map[Trigger]Effect // when a card is played, it registers all its Effects with the appropriate trigger event listeners
}
