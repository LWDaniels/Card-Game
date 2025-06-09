package game

import (
	"math"
	"sort"

	"github.com/LWDaniels/Card-Game/basics/transform"
	"github.com/LWDaniels/Card-Game/basics/vec2"
	"github.com/LWDaniels/Card-Game/constants"
	"github.com/LWDaniels/Card-Game/game/card"
	"github.com/LWDaniels/Card-Game/game/sprite"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*transform.TransformContainerImplementer
}

func NewGame() *Game {
	g := Game{transform.NewTCIDefault()}

	startX, stopX := float32(75), float32(constants.WorldWidth()-75)
	arcMinY, arcMaxY := float32(constants.WorldHeight()-30), float32(constants.WorldHeight()-100)
	startRot, endRot := float32(6.28*-.07), float32(6.28*.07)
	nCards := int(5)
	for i := range nCards {
		interp := float32(i) / float32(nCards-1)
		c := card.NewCard()

		// an arc of an ellipsoid
		midN := float32(nCards) / 2
		dX := (float32(i) + .5 - midN) // idk exactly why the .5 helps but w/e
		cardY := (arcMaxY-arcMinY)*float32(math.Sqrt(float64(midN*midN-dX*dX)))/midN + arcMinY
		cardX := startX + (stopX-startX)*interp // worth noting that X is distributed along a line and not an arc
		c.Trans.Pos = vec2.NewVec2(cardX, cardY)
		c.Trans.Rotation = startRot + (endRot-startRot)*interp
		g.AddChild(c)
	}

	return &g
}

// note that this is a fixed update
func (g *Game) Update() error {
	g.HandleInput()

	// will need to propagate updates to transforms or something
	return nil
}

func (g *Game) HandleInput() {
	// if !inpututil.IsMouseButtonJustPressed(eb.MouseButton0) {
	// 	return
	// }

	// // will need to change for mobile
	// mouseXInt, mouseYInt := eb.CursorPosition()
	// mousePos := vec2.Vec2{X: float32(mouseXInt), Y: float32(mouseYInt)}

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

	sort.Slice(drawables, func(a, b int) bool { return drawables[a].Z >= drawables[b].Z })

	for n := range drawables {
		op.GeoM = drawables[n].GeoM
		screen.DrawImage(drawables[n].Texture, op)
	}
}

// determines render target size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WorldWidth(), constants.WorldHeight()
}
