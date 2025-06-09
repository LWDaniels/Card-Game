package box

import (
	"github.com/LWDaniels/Card-Game/basics/transform"
	"github.com/LWDaniels/Card-Game/basics/vec2"
)

type BoxContainer interface {
	// includes/requires transform container component (may mess up multiple inheritance)
	transform.TransformContainer
	Size() vec2.Vec2
	SetSize(vec2.Vec2)
	InBounds(worldPos vec2.Vec2) bool
}

type BoxContainerDefault struct {
	transform.TransformContainerDefault
	size vec2.Vec2
}

func (bcd *BoxContainerDefault) Size() vec2.Vec2 {
	return bcd.size
}

func (bcd *BoxContainerDefault) SetSize(size vec2.Vec2) {
	bcd.size = size
}

func (bcd *BoxContainerDefault) InBounds(worldPos vec2.Vec2) bool {
	m := bcd.Transform().GeoM()
	m.Invert()
	localPos := vec2.MatMult(m, worldPos)
	return localPos.X >= 0 && localPos.Y >= 0 && localPos.X <= bcd.size.X && localPos.Y <= bcd.size.Y
}

func NewBCDDefault() *BoxContainerDefault {
	bcd := &BoxContainerDefault{}
	bcd.size = vec2.One()
	bcd.TransformContainerDefault = *transform.NewTCDDefault()
	return bcd
}
