package logic

type Trigger uint

const (
	TriggerResolve  Trigger = iota
	TriggerNextPlay         // triggers BEFORE the TriggerResolve effect of the next card
	TriggerDraw
	// add others as needed
)
