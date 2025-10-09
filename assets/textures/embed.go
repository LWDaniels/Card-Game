package textures

import (
	_ "embed"
)

type TextureKey = int

// when adding a texture, need to:
// 1) add the texture to assets/textures (pref as png)
// 2) add a const/fake-enum of the texture id
// 3) add a go:embed of the texture bytes
// 4) add the enum and the bytes to the map
// could maybe use stringify to reduce the # of steps but w/e

// StartRead and StopRead are a hacky way of allowing looping over the enums; don't try to load them
const (
	// DO NOT USE
	StartRead TextureKey = iota
	Gopher
	BlackLotus
	AncestralRecall
	Leveler
	Sting
	Border
	// DO NOT USE
	StopRead
)

// seems kind of expensive per sprite but what do I know
var (
	//go:embed gopher.png
	gopher_png []byte
	//go:embed black-lotus.jpg
	blackLotus_jpg []byte
	//go:embed border.png
	border_png []byte
	//go:embed ancestral-recall.png
	ancestralRecall_png []byte
	//go:embed leveler.png
	leveler_png []byte
	//go:embed sting.png
	sting_png []byte
)

var TextureBytes = map[TextureKey][]byte{
	Gopher:          gopher_png,
	BlackLotus:      blackLotus_jpg,
	Border:          border_png,
	AncestralRecall: ancestralRecall_png,
	Leveler:         leveler_png,
	Sting:           sting_png,
}
