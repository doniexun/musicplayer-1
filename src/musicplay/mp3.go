package musicplay

import "fmt"

import "time"

type MP3Play struct {
	progress int
}

func (mp3 *MP3Play) Play(location string) {
	fmt.Println("Playing mp3 music", location)
	mp3.progress = 0
	for mp3.progress <= 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		mp3.progress += 10
	}
	fmt.Println("\nFinished playing.")
}
