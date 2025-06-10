// archetypes are a collection of components
package archetypes

import (
	"github.com/LWDaniels/Card-Game/src/archetypes/tags"
	"github.com/LWDaniels/Card-Game/src/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

var (
	Card   = newArchetype(tags.Card, components.Card, components.Sprite, transform.Transform)
	Button = newArchetype(tags.Button, components.Sprite, transform.Transform)
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
