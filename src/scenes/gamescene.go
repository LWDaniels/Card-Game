package scenes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/factory"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/procedures"
	"github.com/LWDaniels/Card-Game/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type GameScene struct {
	World    donburi.World
	Hand     []*donburi.Entry
	HeldCard *donburi.Entry
}

const startingCards = int(4)

func NewGameScene() *GameScene {
	g := &GameScene{donburi.NewWorld(), make([]*donburi.Entry, 0), nil}
	for range startingCards {
		c := factory.CreateCard(g.World, math.NewVec2(float64(constants.WorldWidth()/2),
			float64(constants.WorldHeight()/2)))
		g.Hand = append(g.Hand, c)
	}
	for i, card := range g.Hand {
		t := transform.GetTransform(card)
		t.LocalPosition = SlotPos(i, startingCards)
	}

	return g
}

var cardQuery = donburi.NewQuery(filter.Contains(components.Card))

func (g *GameScene) ManageHand() {
	// gather held card
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		g.HeldCard = nil
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		cardQuery.Each(g.World, func(e *donburi.Entry) {
			// this is annoying so maybe just have a pointer to the interactable in the card lol
			child, _ := transform.FindChildWithComponent(e, components.Interactable)
			interactable := components.Interactable.Get(child)
			if interactable.Hovered {
				g.HeldCard = e
			}
		})
	}

	// move it to cursor
	if g.HeldCard != nil {
		mouseX, mouseY := ebiten.CursorPosition()
		mousePos := math.NewVec2(float64(mouseX), float64(mouseY))
		transform.GetTransform(g.HeldCard).LocalPosition = mousePos
	}

	// move cards to hand slots
	nCards := len(g.Hand)
	for i, card := range g.Hand {
		if g.HeldCard != nil && g.HeldCard.Id() == card.Id() {
			return
		}

		t := transform.GetTransform(card)
		t.LocalPosition = utils.ExpDecayVec2(t.LocalPosition, SlotPos(i, nCards), 15)
	}
}

var leftMost = math.NewVec2(50, float64(constants.WorldHeight())-100)
var rightMost = math.NewVec2(float64(constants.WorldWidth())-50, float64(constants.WorldHeight())-100)

func SlotPos(cardIndex int, numCards int) math.Vec2 {
	return utils.LerpVec2(leftMost, rightMost, (float64(cardIndex)+.5)/float64(numCards))
}

func (g *GameScene) Update() error {
	g.ManageHand()
	procedures.TriggerInteractables(g.World)
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	procedures.DrawSprites(g.World, screen)
}
