package main

import (
	"fmt"
)

func simulateAnts(ants int, paths [][]string, end string) {
	antCount := 0
	occupiedRooms := make(map[string]bool)
	if end == "" {
		return
	}
	antPaths := make([][]string, ants)
	antSteps := make([]int, ants)
	for antCount < ants {
		for i := 0; i < len(paths); i++ {
			if antCount == ants {
				break
			}
			antPaths[antCount] = paths[i]
			antSteps[antCount] = 1
			antCount++
		}
	}
	steps := 0
	flag := true
	for flag {

		for ant, path := range antPaths {
			if antSteps[ant] == -1 {
				continue
			}
			step := antSteps[ant]
			if !occupiedRooms[path[step]] {
				if step-1 >= 0 {
					occupiedRooms[path[step-1]] = false
				}
				if path[step] != end {
					occupiedRooms[path[step]] = true
				}
				fmt.Printf("L%d-%s ", ant+1, path[step])
				antSteps[ant]++
			}

			if path[step] == end {
				antSteps[ant] = -1
			}

		}
		steps++
		fmt.Println()
		if sum(antSteps) == ants*-1 {
			flag = false
		}
	}
	fmt.Printf("Done in %d steps.\n", steps)
}

func sum(steps []int) int {
	total := 0
	for _, step := range steps {
		total += step
	}
	return total
}
