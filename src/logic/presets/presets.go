package presets

import "github.com/LWDaniels/Card-Game/src/logic"

var (
	Dagger = logic.CardPreset{
		Name:           "Dagger",
		Text:           "Take {level} damage.",
		RequiresTarget: false,
		Effects: map[logic.Trigger]logic.Effect{
			logic.TriggerResolve: func(state *logic.BoardState, activePlayerIndex int, originalCard *logic.CardInstance, triggerParameters ...any) {
				state.Players[activePlayerIndex].Health -= originalCard.Level
			},
		},
	}
	Upgrade = logic.CardPreset{
		Name:           "Upgrade",
		Text:           "Upgrade the next card {level} times. If you can't, copy its ability instead (at this card's level).",
		RequiresTarget: false,
		Effects: map[logic.Trigger]logic.Effect{
			logic.TriggerNextPlay: func(state *logic.BoardState, activePlayerIndex int, originalCard *logic.CardInstance, triggerParameters ...any) {
				effectedCard := triggerParameters[0].(*logic.CardInstance)
				for range originalCard.Level {
					if !effectedCard.Upgrade() {
						// TODO: copy the card's effects (will need to be done with some on-play function that populates trigger event handlers)
						return
					}
				}
			},
		},
	}
	Seed = logic.CardPreset{
		Name:           "Seed",
		Text:           "Start your next hand with {level} more cards.",
		RequiresTarget: false,
		Effects: map[logic.Trigger]logic.Effect{
			logic.TriggerDraw: func(state *logic.BoardState, activePlayerIndex int, originalCard *logic.CardInstance, triggerParameters ...any) {
				// TODO: effect to draw from deck (should be easy, but should have it as a function to make things easy)
				// state.Players[activePlayerIndex].Hand
			},
		},
	}
)
