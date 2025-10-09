package logic

type Trigger uint

const (
	TriggerResolve  Trigger = iota
	TriggerNextPlay         // triggers BEFORE the TriggerResolve effect of the next card
	TriggerDraw             // this is meant to be a trigger on the draw phase; maybe there should be a TriggerPhase (for any phase)?
	// add others as needed
)
