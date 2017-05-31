package musicplay

import "testing"

func TestMP3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("TestMP3")
		}
	}()
	Play("e:\\dzg.mp3", "mp3")
}

func TestWAV(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("TestWAV")
		}
	}()
}
