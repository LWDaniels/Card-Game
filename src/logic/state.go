package logic

type Deck struct {
	List []*CardInstance // what is left in the deck (not currently in play/hands)
}

type BoardState struct {
	Players []Player // Players[0] is always the local player
	Deck    Deck
}

type Player struct {
	Hand            []*CardInstance // will no longer include actively resolving card
	Health, Victory int
	// will almost certainly need more things
}
