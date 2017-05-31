package musicplay

import "fmt"
import "time"

type WAVPlay struct {
	progress int
}

func (wav *WAVPlay) Play(location string) {
	fmt.Println("Playing wav music", location)
	wav.progress = 0
	for wav.progress <= 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		wav.progress += 10
	}
	fmt.Println("\nFinished playing.")
}
