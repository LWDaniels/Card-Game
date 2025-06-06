package game

import (
	"sort"

	"github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/constants"
	"github.com/LWDaniels/Card-Game/game/card"
	"github.com/LWDaniels/Card-Game/game/sprite"
	"github.com/LWDaniels/Card-Game/game/transform"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*transform.TransformContainerImplementer
}

func NewGame() *Game {
	g := Game{transform.NewTCIDefault()}
	// g.AddChild(card.NewCard())
	c := card.NewCard()
	c.Trans.Rotation += 3.14 * .5
	c.Trans.Pos = vec2.NewVec2(float32(constants.WorldWidth()/2), float32(constants.WorldHeight()/2))
	g.AddChild(c)

	return &g
}

// note that this is a fixed update
func (g *Game) Update() error {
	g.HandleInput()

	// will need to propagate updates to transforms or something

	return nil
}

func (g *Game) HandleInput() {
	if !inpututil.IsMouseButtonJustPressed(eb.MouseButton0) {
		return
	}

	// will need to change for mobile
	// mouseXInt, mouseYInt := eb.CursorPosition()
	// mousePos := vec2.Vec2{X: float32(mouseXInt), Y: float32(mouseYInt)}
	// for n := range g.items {
	// 	switch g.items[n].(type) {
	// 	// add more types if desired
	// 	case *card.Card:
	// 		c := g.items[n].(*card.Card)
	// 		if c.InBounds(mousePos) {
	// 			// shift right by 10
	// 			c.SetPos(vec2.Sum(c.Pos(), vec2.FromF(10)))
	// 		}
	// 	}
	// }
}

type Drawable struct {
	Texture *eb.Image
	GeoM    eb.GeoM
	Z       float32
}

func collectDrawables(tc transform.TransformContainer, parentGeoM eb.GeoM, parentZ float32) []Drawable {
	// can simplify this part greatly (cut down most parameters) but I'm too lazy
	acc := make([]Drawable, 0)
	if sp, ok := tc.(*sprite.Sprite); ok {
		acc = append(acc, Drawable{sp.Texture, parentGeoM, parentZ})
	}
	for _, child := range tc.Transform().Children {
		childTransform := child.Transform()
		cg := childTransform.GeoM()
		cg.Concat(parentGeoM)
		parentZ += childTransform.Z
		acc = append(acc, collectDrawables(child, cg, parentZ)...)
	}

	return acc
}

func (g *Game) Draw(screen *eb.Image) {
	op := &eb.DrawImageOptions{}
	op.Filter = eb.FilterLinear

	drawables := collectDrawables(g, op.GeoM, g.Trans.Z)

	sort.Slice(drawables, func(a, b int) bool { return drawables[a].Z < drawables[b].Z })

	for n := range drawables {
		op.GeoM = drawables[n].GeoM
		screen.DrawImage(drawables[n].Texture, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WorldWidth(), constants.WorldHeight()
}
