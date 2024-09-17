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
	name           string
	xCoord         int
	yCoord         int
	connectedRooms []string
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
		line := scanner.Text()

		if line == "" {
			continue
		}

		if first {
			ants, err := strconv.Atoi(line)
			if err != nil || ants < 1 {
				fmt.Println("error: invalid number of ants")
				return Colony, err
			}
			first = false
			Colony.ants = ants
			continue
		}

		if line == "##start" {
			if Colony.startRoom.name != "" {
				fmt.Println("error: duplicate start room")
				err := errors.New("error: duplicate start room")
				return Colony, err
			}
			scanner.Scan()
			lines, flag := checkRoom(scanner.Text())

			if !flag {
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
				fmt.Println("error: invalid y coordinate")
				return Colony, err
			}

			Colony.startRoom = Room{lines[0], x, y, nil}
			Colony.rooms = append(Colony.rooms, Colony.startRoom)

		} else if line == "##end" {

			if Colony.endRoom.name != "" {
				fmt.Println("error: duplicate end room")
				err := errors.New("error: duplicate end room")
				return Colony, err
			}

			scanner.Scan()
			lines, flag := checkRoom(scanner.Text())

			if !flag {
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
				fmt.Println("error: invalid y coordinate")
				return Colony, err
			}

			Colony.endRoom = Room{lines[0], x, y, nil}
			Colony.rooms = append(Colony.rooms, Colony.endRoom)

		} else if line == "" || line[0] == '#' {
			continue
		} else {
			if strings.Contains(line, "-") {
				lines, flag := checkPath(line)

				if !flag {
					fmt.Println("error: invalid path")
					err := errors.New("error: invalid path")
					return Colony, err
				}

				if len(lines) != 2 {
					fmt.Println("error: invalid path")
					err := errors.New("error: invalid path")
					return Colony, err
				}

				if lines[0] == lines[1] {
					fmt.Println("error: invalid path")
					err := errors.New("error: invalid path")
					return Colony, err
				}

				Colony.paths = append(Colony.paths, Path{lines[0], lines[1]})

				for i := 0; i < len(Colony.rooms); i++ {
					if lines[0] == Colony.rooms[i].name {

						if Colony.rooms[i].name == Colony.startRoom.name {
							Colony.startRoom.connectedRooms = append(Colony.startRoom.connectedRooms, lines[1])
						} else if Colony.endRoom.name == Colony.rooms[i].name {
							Colony.endRoom.connectedRooms = append(Colony.endRoom.connectedRooms, lines[1])
						}

						Colony.rooms[i].connectedRooms = append(Colony.rooms[i].connectedRooms, lines[1])

					} else if lines[1] == Colony.rooms[i].name {

						if Colony.rooms[i].name == Colony.startRoom.name {
							Colony.startRoom.connectedRooms = append(Colony.startRoom.connectedRooms, lines[0])
						} else if Colony.endRoom.name == Colony.rooms[i].name {
							Colony.endRoom.connectedRooms = append(Colony.endRoom.connectedRooms, lines[0])
						}

						Colony.rooms[i].connectedRooms = append(Colony.rooms[i].connectedRooms, lines[0])
					}
				}

			} else {
				lines, flag := checkRoom(line)
				if !flag {
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
					fmt.Println("error: invalid y coordinate")
					return Colony, err
				}

				Colony.rooms = append(Colony.rooms, Room{lines[0], x, y, nil})
			}
		}

	}

	if Colony.startRoom.name == "" || Colony.endRoom.name == "" {
		fmt.Println("error: no start or end room")
		err := errors.New("error: no start or end room")
		return Colony, err
	}

	return Colony, err
}
