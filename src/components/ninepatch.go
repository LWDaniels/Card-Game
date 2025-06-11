package components

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type NinePatchData struct {
	children []*donburi.Entry
	im       *ebiten.Image
}

// requires transform atm
var NinePatch = donburi.NewComponentType[NinePatchData]()

// x, y in [0,2]
func subImageFromIndex(Image *ebiten.Image, xIndex, yIndex int) *ebiten.Image {
	thirdX, thirdY := Image.Bounds().Dx()/3, Image.Bounds().Dy()/3
	return Image.SubImage(image.Rect(xIndex*thirdX, yIndex*thirdX, (xIndex+1)*thirdX, (yIndex+1)*thirdY)).(*ebiten.Image)
}

// divides Image into 3x3 chunks before stretching it based on Size
// may be suspect if doing a ninepatch made of ninepatch sprites and/or sprite sheets (which may be an issue bc of ebiten batching?)
func InitNinePatch(e *donburi.Entry, Image *ebiten.Image, Size image.Point) {
	np := NinePatch.Get(e)
	np.im = Image
	np.children = make([]*donburi.Entry, 9)
	for i := range 9 {
		xIndex := i % 3
		yIndex := i / 3
		np.children[i] = e.World.Entry(e.World.Create(Sprite, transform.Transform))
		InitSprite(np.children[i], subImageFromIndex(Image, xIndex, yIndex))
		transform.AppendChild(e, np.children[i], false)
	}
	np.StretchTo(Size)
}

func (np *NinePatchData) StretchTo(size image.Point) {
	imSizeX, imSizeY := np.im.Bounds().Dx(), np.im.Bounds().Dy()
	for i, child := range np.children {
		xIndex := i % 3
		yIndex := i / 3
		t := transform.Transform.Get(child)
		x := float64(0)
		y := float64(0)
		xScale := float64(1)
		yScale := float64(1)

		// prob not ideal lol
		if xIndex == 2 {
			x = float64(size.X - imSizeX/3)
		}
		if yIndex == 2 {
			y = float64(size.Y - imSizeY/3)
		}

		if xIndex == 1 {
			x = float64(xIndex * imSizeX / 3)

			newSize := size.X - 2*imSizeX/3
			xScale = float64(newSize) / float64(imSizeX/3)
		}
		if yIndex == 1 {
			y = float64(yIndex * imSizeY / 3)

			newSize := size.Y - 2*imSizeY/3
			yScale = float64(newSize) / float64(imSizeY/3)
		}

		t.LocalPosition = math.NewVec2(x, y)
		t.LocalScale = math.NewVec2(xScale, yScale)
	}
}

func (np *NinePatchData) SetTint(tint ebiten.ColorScale) {
	for _, child := range np.children {
		Sprite.Get(child).Tint = tint
	}
}
