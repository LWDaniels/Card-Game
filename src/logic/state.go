package logic

import (
	"context"
	"fmt"

	"github.com/LWDaniels/Card-Game/src/constants"
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

func (bs *BoardState) Transition(event string) {
	err := bs.Phase.Event(context.Background(), event)
	if err != nil {
		fmt.Println(err)
	}
}

// a callback; not to be used really (could make private but w/e)
func (bs *BoardState) EnterState(e *fsm.Event) {
	fmt.Println("enter ", e.Dst)
	switch e.Dst {
	case PhasePass:
		PassPhaseBegin(bs)
	case PhasePlay:
	case PhaseEnd:
		//TODO
	}
}

// a callback; not to be used really (could make private but w/e)
func (bs *BoardState) LeaveState(e *fsm.Event) {
	fmt.Println("leave ", e.Src)
	switch e.Src {
	case PhaseStart:
		StartGame(bs) // handles initialization... eventually
	case PhasePass:
	case PhasePlay:
		PlayPhaseEnd(bs)
	}
}

func NewBoardState(startingDeck structures.Stack[*CardInstance]) *BoardState {
	bs := BoardState{Deck: startingDeck}

	bs.Phase = fsm.NewFSM(PhaseStart, BoardEvents,
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { bs.EnterState(e) },
			"leave_state": func(_ context.Context, e *fsm.Event) { bs.LeaveState(e) },
		}, // can also do special stuff on specific transitions if needed; for now I am just handling enter/exit as they are tho
	)

	for range 4 { // just doing 4 players as a magic number for now
		bs.Players = append(bs.Players, Player{Health: constants.DefaultHealth}) // could use a NewPlayer() func prob
	}
	return &bs
}

type Player struct {
	Hand            []*CardInstance // will no longer include actively resolving card
	PassPile        []*CardInstance
	Triggers        map[Trigger][]Ability // the triggers that the player has queued up; fifo meaning they will be put on the stack in order from oldest to newest (resolving newest to oldest)
	Health, Victory int
	// will almost certainly need more things
}
