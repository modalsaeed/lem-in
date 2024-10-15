package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 1 {
		fmt.Println("wrong number of arguments")
		return
	}

	filename := Args[0]

	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("wrong file type")
		return
	}

	colony, err := CompileColony(filename)
	fmt.Println()

	if err != nil {
		return
	}
	if isValidColony(colony) {
		fmt.Println("Colony is Invalid")
		return
	}
	//fmt.Println("Colony Compiled!")

	DrawColony(colony)

	//fmt.Println("Colony Drawn!")

	paths := findPaths(colony)

	if len(paths) < 1 {
		//fmt.Println("No possible Paths")
		return
	} else {
		//fmt.Println("Paths Found!")
	}

	simulateAnts(colony.ants, paths, colony.endRoom.name)
}

func isValidColony(c Colony) bool {
	if c.ants == 0 {
		return true
	}
	if c.endRoom.name == "" {
		return true
	}
	if c.startRoom.name == "" {
		return true
	}
	if c.rooms == nil {
		return true
	}
	if c.paths == nil {
		return true
	}
	return false
}
