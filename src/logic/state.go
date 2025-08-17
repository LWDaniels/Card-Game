package logic

import (
	"context"

	"github.com/looplab/fsm"
)

// type BoardEvent string // maybe do this too

const (
	PhaseStart = "Start"
	PhasePass  = "Pass"
	PhasePlay  = "Play"
	PhaseEnd   = "Phase End"
	// may need more granular phases for card resolution + triggers and such
	EventStart     = "Event Start"
	EventEnterPlay = "Enter Play"
	EventEnterPass = "Enter Pass"
	EventEnd       = "Event End"
)

var BoardEvents = fsm.Events{ // will prob need lots of changes
	{Name: EventStart, Src: []string{PhaseStart}, Dst: PhasePass},
	{Name: EventEnterPlay, Src: []string{PhasePass}, Dst: PhasePlay},
	{Name: EventEnterPass, Src: []string{PhasePlay}, Dst: PhasePass},
	{Name: EventEnd, Src: []string{PhasePlay, PhasePass}, Dst: PhaseEnd},
}

type BoardState struct {
	Players           []Player        // Players[0] is always the local player
	Deck              []*CardInstance // can add a new type if needed; for now, 0 is the top of the deck
	Stack             []Ability       // first in, last out, like magic stack (active ability to resolve is the last index). There is only ever one Resolve effect on the stack
	StackCard         *CardInstance   // the card that is resolving on the stack, if it exists
	Waiting           []*CardInstance // where cards go until the end of the turn, where they are shuffled together and added to the bottom of the deck
	ActivePlayerIndex int
	Phase             *fsm.FSM // trigger phase changes with Phase.Event(...); things will be triggered appropriately
}

func (bs *BoardState) EnterState(e *fsm.Event) {
	switch e.Dst {
	case PhasePass:
	case PhasePlay:
	case PhaseEnd:
		//TODO
	}
}

func (bs *BoardState) LeaveState(e *fsm.Event) {
	switch e.Src {
	case PhaseStart:
		// StartGame() // TODO
	case PhasePass:
	case PhasePlay:
		PlayPhaseEnd(bs)
	}
}

func NewBoardState() BoardState {
	// TODO: do non-fsm initialization
	bs := BoardState{}

	bs.Phase = fsm.NewFSM(PhaseStart, BoardEvents,
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { bs.EnterState(e) },
			"leave_state": func(_ context.Context, e *fsm.Event) { bs.LeaveState(e) },
		}, // can also do special stuff on specific transitions if needed; for now I am just handling enter/exit as they are tho
	)
	return bs
}

type Player struct {
	Hand            []*CardInstance // will no longer include actively resolving card
	PassPile        []*CardInstance
	Health, Victory int
	// will almost certainly need more things
}
