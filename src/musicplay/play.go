package musicplay

import "fmt"

type Player interface {
	Play(location string)
}

func Play(location string, musicType string) {
	var play Player
	switch musicType {
	case "mp3":
		play = &MP3Play{}
	case "wav":
		play = &WAVPlay{}
	default:
		fmt.Println("unsupported music type.")
		return
	}
	play.Play(location)
}
