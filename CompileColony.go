package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Colony struct {
	ants      int
	startRoom Room
	endRoom   Room
	rooms     []Room
	paths     []Path
}
type Room struct {
	name   string
	xCoord int
	yCoord int
}

type Path struct {
	room1 string
	room2 string
}

func CompileColony(filename string) (Colony, error) {
	Colony := Colony{}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("error opening file")
		return Colony, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	first := true
	for scanner.Scan() {

		if first {
			line := scanner.Text()
			ants, err := strconv.Atoi(line)
			if err != nil || ants < 1 {
				fmt.Println("error: invalid number of ants")
				return Colony, err
			}
			first = false
			Colony.ants = ants
			continue
		}

		line := scanner.Text()

		if line == "##start" {

			scanner.Scan()
			lines := strings.Split(scanner.Text(), " ")

			if len(lines) != 3 {

				fmt.Println("error: invalid start room")
				err := errors.New("error: invalid start room")
				return Colony, err
			}

			x, err := strconv.Atoi(lines[1])
			if err != nil {
				fmt.Println("error: invalid x coordinate")
				return Colony, err
			}

			y, err := strconv.Atoi(lines[2])
			if err != nil {
				fmt.Println("error: invalid x coordinate")
				return Colony, err
			}

			Colony.startRoom = Room{lines[0], x, y}
			Colony.rooms = append(Colony.rooms, Colony.startRoom)

		} else if line == "##end" {

			scanner.Scan()
			lines := strings.Split(scanner.Text(), " ")

			if len(lines) != 3 {
				fmt.Println("error: invalid end room")
				err := errors.New("error: invalid end room")
				return Colony, err
			}

			x, err := strconv.Atoi(lines[1])
			if err != nil {
				fmt.Println("error: invalid x coordinate")
				return Colony, err
			}

			y, err := strconv.Atoi(lines[2])
			if err != nil {
				fmt.Println("error: invalid x coordinate")
				return Colony, err
			}

			Colony.endRoom = Room{lines[0], x, y}
			Colony.rooms = append(Colony.rooms, Colony.endRoom)

		} else if line == "" || line[0] == '#' {
			continue
		} else {

			if strings.Contains(line, "-") {
				lines := strings.Split(line, "-")

				if len(lines) != 2 {
					fmt.Println("error: invalid path")
					err := errors.New("error: invalid path")
					return Colony, err
				}

				Colony.paths = append(Colony.paths, Path{lines[0], lines[1]})

			} else {
				lines := strings.Split(line, " ")

				if len(lines) != 3 {
					fmt.Println("error: invalid room")
					err := errors.New("error: invalid room")
					return Colony, err
				}

				x, err := strconv.Atoi(lines[1])
				if err != nil {
					fmt.Println("error: invalid x coordinate")
					return Colony, err
				}

				y, err := strconv.Atoi(lines[2])
				if err != nil {
					fmt.Println("error: invalid x coordinate")
					return Colony, err
				}

				Colony.rooms = append(Colony.rooms, Room{lines[0], x, y})
			}

		}

	}
	return Colony, err
}
