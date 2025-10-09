package logic

import (
	"math/rand/v2"

	"github.com/LWDaniels/Card-Game/src/constants"
)

// returns a list which is [cardList] with [card] removed, if it is there
func removeCard(card *CardInstance, cardList []*CardInstance) []*CardInstance {
	// I think this can be done without having to return a new slice, but w/e
	output := make([]*CardInstance, len(cardList)-1)
	for _, c := range cardList {
		if c.Id == card.Id {
			continue
		}
		output = append(output, c)
	}
	return output
}

func PlayCard(card *CardInstance, state *BoardState) {
	// requires the stack to be empty before this is played
	// clearing stack just in case
	state.Stack.Clear()

	// remove card from hand or pass pile, whichever it happens to be in
	state.Players[state.ActivePlayerIndex].Hand = removeCard(card, state.Players[state.ActivePlayerIndex].Hand)
	state.Players[state.ActivePlayerIndex].PassPile = removeCard(card, state.Players[state.ActivePlayerIndex].PassPile)

	// could factor into a function since it's also used in copying abilities, but that would require scope stuff
	state.Stack.PushBack(
		Ability{Trigger: TriggerResolve,
			BoundEffect: func() {
				effect, ok := card.Preset.Effects[TriggerResolve]
				if ok {
					effect(state, state.ActivePlayerIndex, card)
				}
				PostResolve(card, state)
			},
		})
}

func PopStack(state *BoardState) {
	ability := state.Stack.Pop()
	if ability != nil {
		ability.BoundEffect()
	}
}

// only done when the stack is empty (to prevent this from being used on copies; this should only be used on real cards)
func PostResolve(card *CardInstance, state *BoardState) {
	state.StackCard = nil
	state.Waiting = append(state.Waiting, card)
	for t, e := range card.Preset.Effects {
		state.Players[state.ActivePlayerIndex].Triggers[t] = append(state.Players[state.ActivePlayerIndex].Triggers[t],
			Ability{Trigger: t, BoundEffect: func() {
				e(state, state.ActivePlayerIndex, card)
			}})
	}
}

// maybe should rename since this is drawing from the deck, not rendering to screen
// active player draws 1 card
func Draw(state *BoardState) {
	card := state.Deck.Pop()
	if card == nil {
		return
	}
	state.Players[state.ActivePlayerIndex].Hand = append(state.Players[state.ActivePlayerIndex].Hand, *card)
	// TODO: trigger draw abilities
}

// shuffles all hands/pass-piles into Waiting and puts it on the bottom of the deck
func PlayPhaseEnd(state *BoardState) {
	for i := range state.Players { // I don't assume the order matters
		state.Waiting = append(state.Waiting, state.Players[i].Hand...)
		state.Players[i].Hand = make([]*CardInstance, 0)
		state.Waiting = append(state.Waiting, state.Players[i].PassPile...)
		state.Players[i].PassPile = make([]*CardInstance, 0)
		// maybe can trigger some stuff here?
	}

	rand.Shuffle(len(state.Waiting), func(i, j int) {
		state.Waiting[i], state.Waiting[j] = state.Waiting[j], state.Waiting[i]
	})

	state.Deck.PushListBack(state.Waiting)
	state.Waiting = make([]*CardInstance, 0)
}

// upgrades cards and passes out hands
func PassPhaseBegin(state *BoardState) {
	for range constants.UpgradesPerTurn {
		state.Deck.Contents[rand.N(state.Deck.Size())].Upgrade() // not handling the case of upgrading something that's fully upgraded; if that happens it happens rn
	}
	// assumes all hands/pass-piles are empty, waiting is empty, stack is empty, and deck is full
	for range constants.CardsInHand {
		for i := range state.Players { // idk if first draw should rotate or what
			state.ActivePlayerIndex = i
			Draw(state)
		}
	}
	// leaves the active player index on the last player; change if desired
}

/*
TODO:
- EVERYTHING about the passing phase
- Add larger game loop that progresses from phase to phase, triggering these systems at the right times (and changing active player and such)
- Add deck generation (make the deck at the start of the game based on the pool of cards)
- Actually integrate stuff nto game logic
*/
