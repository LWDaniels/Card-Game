package presets

import "github.com/LWDaniels/Card-Game/src/logic"

type Listing struct { // could instead do map[*logic.CardPreset]int
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
