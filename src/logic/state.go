package logic

type BoardState struct {
	Players           []Player        // Players[0] is always the local player
	Deck              []*CardInstance // can add a new type if needed; for now, 0 is the top of the deck
	Stack             []Ability       // first in, last out, like magic stack (active ability to resolve is the last index). There is only ever one Resolve effect on the stack
	StackCard         *CardInstance   // the card that is resolving on the stack, if it exists
	Waiting           []*CardInstance // where cards go until the end of the turn, where they are shuffled together and added to the bottom of the deck
	ActivePlayerIndex int
}

type Player struct {
	Hand            []*CardInstance // will no longer include actively resolving card
	Health, Victory int
	// will almost certainly need more things
}
