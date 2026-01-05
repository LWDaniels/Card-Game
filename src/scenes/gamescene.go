package scenes

import (
	"image"

	"github.com/LWDaniels/Card-Game/src/archetypes/factory"
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/logic"
	"github.com/LWDaniels/Card-Game/src/logic/presets"
	"github.com/LWDaniels/Card-Game/src/logic/structures"
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
	World       donburi.World
	Hand        []*donburi.Entry // should not be updated directly; instead use Sync(...) below
	HoveredZone *donburi.Entry
	HeldCard    *donburi.Entry
	State       *logic.BoardState
}

func NewGameScene() *GameScene {
	g := &GameScene{donburi.NewWorld(), make([]*donburi.Entry, 0), nil, nil, logic.NewBoardState(GenerateDeck())}

	// prob want a method for this
	factory.CreateZone(g.World, math.NewVec2(10, 10), image.Pt(constants.WorldWidth()-20, 100))
	factory.CreateZone(g.World, math.NewVec2(10, 120), image.Pt(100, 300))
	factory.CreateZone(g.World, math.NewVec2(float64(constants.WorldWidth()-110), 120), image.Pt(100, 300))

	// could start the game after a prompt if needed
	g.State.Transition(logic.EventStart)
	return g
}

func GenerateDeck() structures.Stack[*logic.CardInstance] {
	deck := structures.Stack[*logic.CardInstance]{}
	for _, item := range presets.DeckList {
		t := logic.TargetNone
		if item.Card.RequiresTarget {
			t = logic.TargetLeft
		}
		for range item.Count {
			deck.PushBack(logic.NewInstance(item.Card, t))
			// oscillate target directions
			if t == logic.TargetNone {
				continue
			} else if t == logic.TargetLeft {
				t = logic.TargetRight
			} else if t == logic.TargetRight {
				t = logic.TargetLeft
			}
		}
	}
	return deck
}

var zoneQuery = donburi.NewQuery(filter.Contains(tags.Zone))

func (g *GameScene) ManageZone() {
	g.HoveredZone = nil
	zoneQuery.Each(g.World, func(e *donburi.Entry) {
		interactable := components.Interactable.Get(e)
		if interactable.Hovered {
			g.HoveredZone = e
		}
	})
}

func (g *GameScene) PassCard(card *donburi.Entry) { // TODO: pass in target player index
	if !g.State.Phase.Is(logic.PhasePass) {
		return
	}

	c := components.Card.Get(card)
	logic.PassCard(g.State, 0, 1 /*TODO: target properly*/, c.Instance)
	g.Hand = g.CardsToEntries(g.Hand, g.State.Players[0].Hand)
	transform.RemoveRecursive(card) // can remove this once I update CardsToEntries
}

var cardQuery = donburi.NewQuery(filter.Contains(components.Card))

/*
Returns a slice of entries representing [cards].
Reuses entries from [prevEntries] whenever possible
(can just use nil if you want a brand new hand; be wary of duplicates ofc)
*/
func (g *GameScene) CardsToEntries(prevEntries []*donburi.Entry, cards []*logic.CardInstance) []*donburi.Entry {
	// not super efficient; may be better to make events to update card by card? or maybe to make hands a map...
	newEntries := make([]*donburi.Entry, len(cards))

	for i, cardInstance := range cards {
		found := false
		for _, entry := range prevEntries {
			if components.Card.Get(entry).Instance.Id == cardInstance.Id {
				found = true
				newEntries[i] = entry
				break
			}
		}
		if !found {
			newEntries[i] = factory.CreateCard(g.World, SlotPos(i, len(cards)), cardInstance)
		}
	}
	// TODO: remove entries that arent in the new list from transform
	return newEntries
}

func (g *GameScene) ManageHand() {
	g.Hand = g.CardsToEntries(g.Hand, g.State.Players[0].Hand)

	// gather held card
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		if g.HeldCard != nil && g.HoveredZone != nil {
			g.PassCard(g.HeldCard)
		}
		g.HeldCard = nil
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		cardQuery.Each(g.World, func(e *donburi.Entry) {
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
	procedures.TriggerInteractables(g.World)
	g.ManageZone()
	g.ManageHand()
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	procedures.DrawSprites(g.World, screen)
}
