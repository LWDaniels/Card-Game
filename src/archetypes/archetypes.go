// archetypes are a collection of components
package archetypes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// can also do unions over these which may be nice (not 100% sure how much I should encourage that tho)

var (
	Card   = newArchetype(tags.Card, components.Card, components.Sprite, transform.Transform, components.Interactable)
	Button = newArchetype(tags.Button, components.Sprite, transform.Transform, components.Interactable)
)

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
