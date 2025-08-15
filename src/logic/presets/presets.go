package presets

import "github.com/LWDaniels/Card-Game/src/logic"

var (
	Dagger = logic.CardPreset{
		Name:           "Dagger",
		Text:           "Take {level} damage.",
		RequiresTarget: false,
		Effects: map[logic.Trigger]logic.Effect{
			logic.TriggerResolve: func(state *logic.BoardState, casterIndex int, originalCard *logic.CardInstance) {
				state.Players[casterIndex].Health -= originalCard.Level
			},
		},
	}
	Upgrade = logic.CardPreset{
		Name:           "Upgrade",
		Text:           "Upgrade the next card {level} times. If you can't, copy its ability instead (at this card's level).",
		RequiresTarget: false,
		Effects: map[logic.Trigger]logic.Effect{
			logic.TriggerNextPlay: func(state *logic.BoardState, casterIndex int, originalCard *logic.CardInstance) {
				effectedCard := state.StackCard
				for range originalCard.Level {
					if !effectedCard.Upgrade() {
						state.Stack = append(state.Stack,
							logic.Ability{Trigger: logic.TriggerResolve,
								BoundEffect: func() {
									effect, ok := effectedCard.Preset.Effects[logic.TriggerResolve]
									if ok {
										effect(state, state.ActivePlayerIndex, effectedCard)
									}
									// no post-resolve since this is just a copy, not a real card
								},
							})
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
			logic.TriggerDraw: func(state *logic.BoardState, activePlayerIndex int, originalCard *logic.CardInstance) {
				// TODO: effect to draw from deck (should be easy, but should have it as a function to make things easy)
				// state.Players[activePlayerIndex].Hand
			},
		},
	}
)
