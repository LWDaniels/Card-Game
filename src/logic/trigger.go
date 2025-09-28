package logic

type Trigger uint

const (
	TriggerResolve  Trigger = iota
	TriggerNextPlay         // triggers BEFORE the TriggerResolve effect of the next card
	TriggerDraw             // not sure if this should mean draw step or draw effect; maybe there should be a TriggerPhase (for any phase)?
	// add others as needed
)
