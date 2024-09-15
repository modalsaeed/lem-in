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
	fmt.Println(colony.ants)
	fmt.Println(colony.startRoom)
	fmt.Println(colony.endRoom)
	fmt.Println(colony.rooms)
	fmt.Println(colony.paths)
}
