package loader

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const SampleRate = 44100

type ResourceLoader struct {
}

//go:embed misc/*
var embedResource embed.FS

func (r *ResourceLoader) GetSprite(filename string) *ebiten.Image {
	path := "misc/"

	img, _, err := ebitenutil.NewImageFromFileSystem(embedResource, path+filename)
	if err != nil {
		panic(err)
	}

	return img
}

func (r *ResourceLoader) GetSound(name string, audioContext *audio.Context) *audio.Player {
	path := "misc/"

	filename := SoundsFilenames[name]
	file, err := embedResource.Open(path + filename)
	if err != nil {
		panic(err)
	}

	decoded, err := mp3.DecodeWithSampleRate(SampleRate, file)
	if err != nil {
		panic(err)
	}

	player, err := audioContext.NewPlayer(decoded)
	if err != nil {
		panic(fmt.Errorf("Could not create player for " + filename + ". " + err.Error()))
	}

	return player
}

func (r *ResourceLoader) GetSoundList() []string {
	res := make([]string, 0, len(SoundsFilenames))
	for name := range SoundsFilenames {
		res = append(res, name)
	}

	return res
}
