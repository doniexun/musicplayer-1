package main

import (
	"bufio"
	"fmt"
	"library"
	"musicplay"
	"os"
	"strings"
)

var prompt = func() {
	fmt.Println(`
Enter following commands to control the player:
lib list                                  --view the existing music lib
lib add <name> <artist> <location> <type> --add a music to the music lib
lib remove <name>                         --remove the specified music from the music lib
play <name>                               --play the specified music
quit/exit                                 --exit the program`)
}

var lib *library.MusicManager

func main() {
	prompt()

	lib = library.NewMusicManager()

	rd := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command->")

		rawLine, _, _ := rd.ReadLine()

		line := string(rawLine)

		if line == "" {
			continue
		}

		if line == "quit" || line == "exit" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("unsupported commands:", tokens[0])
		}
		prompt()
	}
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		lib.Display()
	case "add":
		if len(tokens) == 6 {
			music := &library.Music{Name: tokens[2], Artist: tokens[3], Location: tokens[4], Musictype: tokens[5]}
			err := lib.Add(music)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("usage: lib add <name> <artist> <location> <type>")
		}
	case "remove":
		if len(tokens) != 3 {
			fmt.Println("usage: lib remove <name>")
			return
		}

		_, err := lib.RemoveByName(tokens[2])
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unsupported lib command", tokens[1])
	}
}

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("usage: play <name>")
		return
	}

	music, err := lib.Find(tokens[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	musicplay.Play(music.Location, music.Musictype)
}
