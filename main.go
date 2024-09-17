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

	if err != nil {
		return
	}

	fmt.Println("Colony Compiled!")

	DrawColony(colony)

	fmt.Println("Colony Drawn!")

	paths := findPaths(colony)
	if len(paths) < 1 {
		fmt.Println("No possible Paths")
		return
	}
	fmt.Println("Paths Found!")

	simulateAnts(colony.ants, paths, colony.endRoom.name)
	simulateAnts(colony.ants, [][]string{{"richard", "erlich", "jimYoung", "peter"}}, colony.endRoom.name)
}
