package logic

import (
	"context"

	"github.com/looplab/fsm"
)

type Phase string

// type BoardEvent string // maybe do this too

const (
	PhaseStart = "Start"
	PhasePass  = "Pass"
	PhasePlay  = "Play"
	PhaseEnd   = "End"
)

var BoardEvents = fsm.Events{ // will prob need lots of changes
	{Name: "start to pass", Src: []string{PhaseStart}, Dst: PhasePass},
	{Name: "pass to play", Src: []string{PhasePass}, Dst: PhasePlay},
	{Name: "play to pass", Src: []string{PhasePlay}, Dst: PhasePass},
	{Name: "end game", Src: []string{PhasePlay, PhasePass}, Dst: PhaseEnd},
}

type BoardState struct {
	Players           []Player        // Players[0] is always the local player
	Deck              []*CardInstance // can add a new type if needed; for now, 0 is the top of the deck
	Stack             []Ability       // first in, last out, like magic stack (active ability to resolve is the last index). There is only ever one Resolve effect on the stack
	StackCard         *CardInstance   // the card that is resolving on the stack, if it exists
	Waiting           []*CardInstance // where cards go until the end of the turn, where they are shuffled together and added to the bottom of the deck
	ActivePlayerIndex int
	Phase             *fsm.FSM
}

func (bs *BoardState) EnterState(e *fsm.Event) {
	// TODO
	// will prob put all major game logic here
}

func NewBoardState() BoardState {
	// TODO: do non-fsm initialization
	bs := BoardState{}

	bs.Phase = fsm.NewFSM(PhaseStart, BoardEvents,
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { bs.EnterState(e) },
		},
	)
	return bs
}

type Player struct {
	Hand            []*CardInstance // will no longer include actively resolving card
	PassPile        []*CardInstance
	Health, Victory int
	// will almost certainly need more things
}
