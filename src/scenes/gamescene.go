package scenes

import (
	"image"

	"github.com/LWDaniels/Card-Game/src/archetypes/factory"
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/LWDaniels/Card-Game/src/constants"
	"github.com/LWDaniels/Card-Game/src/logic"
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
	Hand        []*donburi.Entry
	HoveredZone *donburi.Entry
	HeldCard    *donburi.Entry
	State       *logic.BoardState
}

func NewGameScene() *GameScene {
	g := &GameScene{donburi.NewWorld(), make([]*donburi.Entry, 0), nil, nil, logic.NewBoardState()}
	// for range startingCards {
	// 	c := factory.CreateCard(g.World, math.NewVec2(float64(constants.WorldWidth()/2),
	// 		float64(constants.WorldHeight()/2)))
	// 	g.Hand = append(g.Hand, c)
	// }

	// prob want a method for this
	factory.CreateZone(g.World, math.NewVec2(10, 10), image.Pt(constants.WorldWidth()-20, 100))
	factory.CreateZone(g.World, math.NewVec2(10, 120), image.Pt(100, 300))
	factory.CreateZone(g.World, math.NewVec2(float64(constants.WorldWidth()-110), 120), image.Pt(100, 300))

	return g
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

func (g *GameScene) PassCard(card *donburi.Entry) {
	if g.State.ActivePlayerIndex != 0 || !g.State.Phase.Is(logic.PhasePass) { // maybe need to keep track of index?
		return
	}

	acc := make([]*donburi.Entry, 0)
	for _, e := range g.Hand {
		if card.Id() == e.Id() {
			continue
		}
		acc = append(acc, e)
	}
	g.Hand = acc
	transform.RemoveRecursive(card)
}

var cardQuery = donburi.NewQuery(filter.Contains(components.Card))

func (g *GameScene) ManageHand() {
	newCards := make([]*logic.CardInstance, 0)
	// detecting new cards; makes use of the fact that cards are added to the back of the hand only
	// would be easier to trigger off of draw, but drawing and logic are pretty much completely separated rn (not sure if that's great)
	for i := len(g.State.Players[0].Hand) - 1; i >= 0; i-- { // may need to change off 0-index
		c := g.State.Players[0].Hand[i]
		found := false
		for _, entry := range g.Hand {
			if components.Card.Get(entry).Instance.Id == c.Id {
				found = true
				break
			}
		}
		if found {
			break
		}
		// not found, so we can are sure that c is a new card
		newCards = append(newCards, c)
	}
	for i, c := range newCards {
		card := factory.CreateCard(g.World, SlotPos(len(g.Hand)+i, len(g.Hand)+len(newCards)), c)
		g.Hand = append(g.Hand, card)
	}

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
