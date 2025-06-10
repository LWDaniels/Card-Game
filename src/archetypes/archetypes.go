// archetypes are a collection of components, possibly with another archetype as its child
package archetypes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// can also do unions over these which may be nice (not 100% sure how much I should encourage that tho)
// can just use appends/... on Arch.components for now
// using underscore prefix (eg _Name) for things that shouldn't be instantiated directly
// these cannot be tagged and you get them upon instantiation by simply getting the first child
// generally good practice to store a reference to the child in a parent component

var (
	_CardInteractable = newArchetype(components.Sprite, transform.Transform, components.Interactable)
	Card              = addChild(newArchetype(tags.Card, components.Card, transform.Transform), _CardInteractable)
	Button            = newArchetype(tags.Button, components.Sprite, transform.Transform, components.Interactable)
)

type archetype struct {
	components []donburi.IComponentType
	// if child == nil, there's no child; else, requires transform component in this and child (adds them if absent)
	// note that this is a single pointer instead of a slice because if you want multiple, just union them
	// dont have these children cycle lol
	child *archetype
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
		child:      nil,
	}
}

// returns parent for chaining purposes
func addChild(parent *archetype, child *archetype) *archetype {
	parent.child = child
	return parent
}

// spawns the archetype (with extra components added) and any archetype children (recursively)
// generally call factory.CreateX instead of this
func (a *archetype) Spawn(world donburi.World, extraComponents ...donburi.IComponentType) *donburi.Entry {
	e := world.Entry(world.Create(
		append(a.components, extraComponents...)...,
	))
	if a.child != nil {
		if !e.HasComponent(transform.Transform) {
			e.AddComponent(transform.Transform)
		}
		child := a.child.Spawn(world)
		if !child.HasComponent(transform.Transform) {
			child.AddComponent(transform.Transform)
		}
		transform.AppendChild(e, child, false)
	}
	return e
}
