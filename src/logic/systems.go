package logic

import (
	"math/rand"

	"github.com/LWDaniels/Card-Game/src/constants"
)

func PlayCard(cardIndexInHand int, state *BoardState) {
	// requires the stack to be empty before this is played
	// clearing stack just in case
	state.Stack = make([]Ability, 1)

	card := state.Players[state.ActivePlayerIndex].Hand[cardIndexInHand]
	// remove card from hand
	hand := make([]*CardInstance, len(state.Players[state.ActivePlayerIndex].Hand)-1)
	for i, c := range state.Players[state.ActivePlayerIndex].Hand {
		if i == cardIndexInHand {
			continue
		}
		hand = append(hand, c)
	}
	state.Players[state.ActivePlayerIndex].Hand = hand

	// could factor into a function since it's also used in copying abilities, but that would require scope stuff
	state.Stack = append(state.Stack,
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
	if len(state.Stack) == 0 {
		return
	}

	ability := state.Stack[len(state.Stack)-1]
	state.Stack = state.Stack[:len(state.Stack)-1]
	ability.BoundEffect()
}

// TODO: populates event listeners, moves the card to the waiting zone, etc.
// only done when the stack is empty (to prevent this from being used on copies; this should only be used on real cards)
func PostResolve(card *CardInstance, state *BoardState) {
	state.StackCard = nil
	state.Waiting = append(state.Waiting, card)
	// populate event listeners :)
}

func Draw(state *BoardState) {
	card := state.Deck[0] // assuming deck is never empty... maybe that's an incorrect assumption
	state.Deck = state.Deck[1:]
	state.Players[state.ActivePlayerIndex].Hand = append(state.Players[state.ActivePlayerIndex].Hand, card)
	// TODO: trigger draw abilities
}

// shuffles all hands into Waiting and puts it on the bottom of the deck
func PlayPhaseEnd(state *BoardState) {
	for i := range state.Players { // I don't assume the order matters
		state.Waiting = append(state.Waiting, state.Players[i].Hand...)
		state.Players[i].Hand = make([]*CardInstance, 0)
		// maybe can trigger some stuff here?
	}

	rand.Shuffle(len(state.Waiting), func(i, j int) {
		state.Waiting[i], state.Waiting[j] = state.Waiting[j], state.Waiting[i]
	}) // may need to change to be deterministic, idk

	state.Deck = append(state.Deck, state.Waiting...)
}

func PassPhaseBegin(state *BoardState) {
	// assumes all hands are empty, waiting is empty, stack is empty, and deck is full
	for range constants.CardsInHand {
		for i := range state.Players { // idk if first draw should rotate or what
			state.ActivePlayerIndex = i
			Draw(state)
		}
	}
	// leaves the active player index on the last player; change if desired
}
