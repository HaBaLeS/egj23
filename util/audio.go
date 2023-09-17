package util

import (
	"github.com/habales/egj23/assets"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var audioContext *audio.Context = audio.NewContext(sampleRate)
var audioPlayer *audio.Player

const (
	sampleRate = 48000
)

func LoadAudio(file string) *audio.Player {

	f, e := assets.Data.Open(file)
	if e != nil {
		panic(e)
	}
	s, err := mp3.DecodeWithoutResampling(f)
	if e != nil {
		panic(err)
	}

	p, err := audioContext.NewPlayer(s)

	return p
}
