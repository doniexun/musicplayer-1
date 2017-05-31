package library

import "testing"

func TestManager1(t *testing.T) {
	mm := NewMusicManager()

	if mm.Len() != 0 {
		t.Error("mm:=NewMusicManager()")
	}

	m := &Music{"大中国", "周杰伦", "e:\\dzg.mp3", "mp3"}

	err := mm.Add(m)
	if err != nil {
		t.Error("err := mm.Add(m)")
	}

	if mm.Len() != 1 {
		t.Error("err := mm.Add(m)")
	}

	music, err := mm.Find("大中国")
	if err != nil {
		t.Error("music,err:=mm.Find(\"大中国\")")
	}
	if music.Name != "大中国" {
		t.Error("music,err:=mm.Find(\"大中国\")")
	}

	music, err = mm.Remove(0)
	if err != nil {
		t.Error("music,err=mm.Remove(0)")
	}
	if music.Name != "大中国" {
		t.Error("music,err=mm.Remove(0)")
	}

	if mm.Len() != 0 {
		t.Error("music,err=mm.Remove(0)")
	}

}
