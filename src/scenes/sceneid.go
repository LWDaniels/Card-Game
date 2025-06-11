package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(*ebiten.Image)
	// maybe add an Init() error one too
}

type SceneID uint

const (
	MainMenuSceneID SceneID = iota // default scene
	GameSceneID
)

var scenes = map[SceneID]Scene{
	MainMenuSceneID: NewMainMenuScene(),
	GameSceneID:     NewGameScene(),
}

var nextScene SceneID = 0

func SetNextScene(id SceneID) {
	nextScene = id
}

// this may be the current scene :)
func NextScene() Scene {
	return scenes[nextScene]
}
