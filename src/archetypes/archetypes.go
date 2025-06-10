// archetypes are a collection of components
package archetypes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// can also do unions over these which may be nice (not 100% sure how much I should encourage that tho)
// can just use appends/... on Arch.components for now

var (
	Card   = newArchetype(tags.Card, components.Card, components.Sprite, transform.Transform, components.Interactable)
	Button = newArchetype(tags.Button, components.Sprite, transform.Transform, components.Interactable)
)

// prob want to add a child *archetype or something similar to allow for nesting
type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
	}
}

func (a *archetype) Spawn(world donburi.World, extraComponents ...donburi.IComponentType) *donburi.Entry {
	e := world.Entry(world.Create(
		append(a.components, extraComponents...)...,
	))
	return e
}
