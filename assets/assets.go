// currently using an embedding system
// unfortunately this may not scale for all assets so maybe expand to runtime loading in the future
package assets

/*
Notes on asset structure

Embed allows for portable, quickly loading (in theory),
and easy to use (since we have syntax highlighting
and compile-time checks) assets.
However, it also balloons the file size of the binary/exe
and probably increases compilation time.

So, if I add a ton of assets I will prob want to change to runtime loading.
This however would come at the cost of portability (esp for web)
*/

import (
	"bytes"
	"image"
	_ "image/png" // to enable png decoding
	"log"

	"github.com/LWDaniels/Card-Game/assets/textures"
	eb "github.com/hajimehoshi/ebiten/v2"
)

var loadedTextures = make(map[textures.TextureKey]*eb.Image)

// returns a texture; loads it if it isn't loaded already
func GetTexture(key textures.TextureKey) *eb.Image {
	tex, texLoaded := loadedTextures[key]
	if texLoaded {
		return tex
	}

	LoadTexture(key)
	return loadedTextures[key]
}

func LoadTexture(key textures.TextureKey) {
	_, texLoaded := loadedTextures[key]
	if texLoaded {
		return
	}
	b, exists := textures.TextureBytes[key]
	if !exists {
		// already loaded (so we dumped it),
		// forgot to add to the map,
		// or key == StartRead/StopRead

		// stringify would be helpful here
		log.Fatalf("key invalid; key integer value = %d\n"+
			"Check if key is in TextureBytes map or if it was disposed early.", key)
		return
	}

	tex, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err) // hopefully this doesn't happen often lol
		// seems like it can only happen if the data isn't in proper png format (in which case we prob fail earlier)
	}

	loadedTextures[key] = eb.NewImageFromImage(tex)
	delete(textures.TextureBytes, key) // hopefully this gets rid of the ref so GC can handle it
}

func LoadAll() {
	// add non-texture stuff later
	for key := textures.StartRead + 1; key < textures.StopRead; key++ {
		LoadTexture(key)
	}
}

// not convinced this is needed bc of GC; however,
// GC doesn't seem to work properly with ebiten's
// graphics card usage, so it's good to be safe
func UnloadTextures() {
	// will deallocate even textures with existing pointers elsewhere
	for k, v := range loadedTextures {
		v.Deallocate()
		delete(loadedTextures, k)
	}
}
