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

	DrawColony(colony)
	paths := findPaths(colony)
	if len(paths) < 1 {
		fmt.Println("No possible Paths")
		return
	}
	simulateAnts(colony.ants, paths, colony.endRoom.name)
}
