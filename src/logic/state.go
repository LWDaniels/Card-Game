package logic

import (
	"context"

	"github.com/LWDaniels/Card-Game/src/logic/structures"
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
	// start -> pass -> play -> end
	{Name: EventStart, Src: []string{PhaseStart}, Dst: PhasePass},
	{Name: EventEnterPlay, Src: []string{PhasePass}, Dst: PhasePlay},
	{Name: EventEnterPass, Src: []string{PhasePlay}, Dst: PhasePass},
	{Name: EventEnd, Src: []string{PhasePlay, PhasePass}, Dst: PhaseEnd},
}

type BoardState struct {
	Players           []Player // Players[0] is always the local player
	Deck              structures.Stack[*CardInstance]
	Stack             structures.Stack[Ability]
	StackCard         *CardInstance   // the card that is resolving on the stack, if it exists
	Waiting           []*CardInstance // where cards go until the end of the turn, where they are shuffled together and added to the bottom of the deck; could also be a stack ig
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
	Triggers        map[Trigger][]Ability // the triggers that the player has queued up; fifo meaning they will be put on the stack in order from oldest to newest (resolving newest to oldest)
	Health, Victory int
	// will almost certainly need more things
}
