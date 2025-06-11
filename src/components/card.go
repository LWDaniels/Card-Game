package components

import "github.com/yohamta/donburi"

type CardData struct {
	Child *donburi.Entry // ref to entry storing Sprite, Transform, and Interactable data
	// add some effect type too
}

var Card = donburi.NewComponentType[CardData]()

func InitCard(e *donburi.Entry, child *donburi.Entry) {
	c := Card.Get(e)
	c.Child = child
}
