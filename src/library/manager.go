package library

import (
	"errors"
	"fmt"
)

type Music struct {
	Name      string
	Artist    string
	Location  string
	Musictype string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{musics: make([]Music, 0)}
}

func (mm *MusicManager) Len() int {
	return len(mm.musics)
}

func (mm *MusicManager) Get(index int) (m *Music, err error) {
	if index < 0 || index >= mm.Len() {
		err = errors.New("The music not exists.")
		return
	}

	return &mm.musics[index], nil
}

func (mm *MusicManager) Display() {
	for index, music := range mm.musics {
		fmt.Println(index+1, music.Name, music.Artist, music.Location, music.Musictype)
	}
	return

}

func (mm *MusicManager) Find(name string) (m *Music, err error) {
	if mm.Len() == 0 {
		err = errors.New("The music library is empty.")
		return
	}
	for index, music := range mm.musics {
		if music.Name == name {
			return &mm.musics[index], nil
		}
	}
	return nil, errors.New("The music not exists.")
}

func (mm *MusicManager) Add(music *Music) error {
	mm.musics = append(mm.musics, *music)

	return nil
}

func (mm *MusicManager) RemoveByName(name string) (m *Music, err error) {
	for index, value := range mm.musics {
		if value.Name == name {
			m = &mm.musics[index]
			if mm.Len() == 1 {
				mm.musics = make([]Music, 0)
				return m, nil
			}
			if index == 0 {
				mm.musics = mm.musics[1:]
			} else if index == mm.Len()-1 {
				mm.musics = mm.musics[:index]
			} else {
				mm.musics = append(mm.musics[0:index], mm.musics[index+1:]...)
			}
			return m, nil
		}
	}
	return nil, errors.New("The music " + name + " not exists")
}
