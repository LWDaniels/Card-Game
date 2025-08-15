package components

import "github.com/yohamta/donburi"

// only for visible cards rn, not the actual card effects rn
type CardData struct {
	Child *donburi.Entry // ref to entry storing Sprite, Transform, and Interactable data
}

var Card = donburi.NewComponentType[CardData]()

func InitCard(e *donburi.Entry, child *donburi.Entry) {
	c := Card.Get(e)
	c.Child = child
}
