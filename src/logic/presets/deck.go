package presets

import "github.com/LWDaniels/Card-Game/src/logic"

type Listing struct {
	Card  *logic.CardPreset
	Count int
}

var DeckList = []Listing{
	{
		&Dagger, 4,
	},
	{
		&Upgrade, 4,
	},
	{
		&Seed, 4,
	},
}
