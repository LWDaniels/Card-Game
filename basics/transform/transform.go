package transform

import (
	"github.com/LWDaniels/Card-Game/basics/vec2"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	Pos   vec2.Vec2
	Scale vec2.Vec2
	// in radians, counterclockwise
	Rotation float32
	Parent   TransformContainer
	Children []TransformContainer
	Z        float32
	// need to set container if you want the containing object to be referenceable while traversing nodes
	Container any
}

type TransformContainer interface {
	Transform() Transform
	SetTransform(Transform)
	AddChild(TransformContainer)
	RemoveChild(TransformContainer)
	Orphan()
}

// embed a pointer to this and things should work :)
type TransformContainerDefault struct {
	Trans Transform
}

func (tci *TransformContainerDefault) Transform() Transform {
	return tci.Trans
}

func (tci *TransformContainerDefault) SetTransform(t Transform) {
	tci.Trans = t
}

// removes child's current Parent reference
func (tci *TransformContainerDefault) AddChild(child TransformContainer) {
	child.Orphan()
	tci.Trans.Children = append(tci.Trans.Children, child)
	ct := child.Transform()
	ct.Parent = tci
	child.SetTransform(ct)
}

func (tci *TransformContainerDefault) RemoveChild(child TransformContainer) {
	acc := make([]TransformContainer, 0)
	for _, c := range tci.Trans.Children {
		if c != child { // questionable equality check, prob need an ID
			acc = append(acc, c)
		}
	}
	tci.Trans.Children = acc
}

// may change global GeoM
func (tci *TransformContainerDefault) Orphan() {
	if tci.Trans.Parent == nil {
		return
	}
	tci.Trans.Parent.RemoveChild(tci)

	tci.Trans.Parent = nil
}

func NewTCD(scale vec2.Vec2, rotation float32, position vec2.Vec2, z float32) *TransformContainerDefault {
	return &TransformContainerDefault{NewTransform(scale, rotation, position, z)}
}

func NewTCDDefault() *TransformContainerDefault {
	return NewTCD(vec2.One(), 0, vec2.Zero(), 0)
}

func NewTransform(scale vec2.Vec2, rotation float32, position vec2.Vec2, z float32) Transform {
	return Transform{position, scale, rotation, nil, nil, z, nil}
}

func Identity() Transform {
	return NewTransform(vec2.One(), 0, vec2.Zero(), 0)
}

// global transform
func (t Transform) GlobalGeoM() eb.GeoM {
	if t.Parent == nil {
		return t.GeoM()
	}
	global := t.Parent.Transform().GlobalGeoM()
	global.Concat(t.GeoM())
	return global
}

// local transform
func (t Transform) GeoM() eb.GeoM {
	geoM := eb.GeoM{}
	sX, sY := t.Scale.F64()
	geoM.Scale(sX, sY)
	geoM.Rotate(float64(t.Rotation))
	x, y := t.Pos.F64()
	geoM.Translate(x, y)
	return geoM
}

// doesn't use scaled Z values
func (t Transform) GlobalZ() float32 {
	if t.Parent == nil {
		return t.Z
	}
	return t.Z + t.Parent.Transform().GlobalZ()
}
