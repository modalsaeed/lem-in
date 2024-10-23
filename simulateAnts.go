package main

import (
	"fmt"
	"sort"
)

func simulateAnts(ants int, paths [][]string, end string) {
	optimizedAllocation := optimizePaths(paths, ants)

	occupiedRooms := make(map[string]bool)
	occupiedPaths := make(map[int]bool)
	if end == "" {
		return
	}
	antPaths := make([][]string, ants)
	antSteps := make([]int, ants)
	antNumbers := make([]int, ants)
	nextAntNumber := 1

	antsAssigned := 0
	for pathIndex, antCount := range optimizedAllocation {
		for i := 0; i < antCount; i++ {
			if antsAssigned == ants {
				break
			}
			antPaths[antsAssigned] = paths[pathIndex]
			antSteps[antsAssigned] = 1
			antsAssigned++
		}
	}
	steps := 0
	flag := true
	for flag {
		for ant, path := range antPaths {
			if antSteps[ant] == -1 {
				continue
			}
			isDirectPath := len(path) == 2
			pathIndex := -1

			for i, p := range paths {
				if compareSlice(p, path) {
					pathIndex = i
					break
				}
			}
			if occupiedPaths[pathIndex] {
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
				if isDirectPath {
					occupiedPaths[pathIndex] = true
				}
				if antNumbers[ant] == 0 {
					antNumbers[ant] = nextAntNumber
					nextAntNumber++
				}
				fmt.Printf("L%d-%s ", antNumbers[ant], path[step])
				antSteps[ant]++
			}

			if path[step] == end {
				antSteps[ant] = -1
			}
		}
		steps++
		reset(occupiedPaths)
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

func optimizePaths(paths [][]string, antCount int) []int {
	// Sort paths by length
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Initialize all ants on the shortest path
	allocation := make([]int, len(paths))
	allocation[0] = antCount

	// Iteratively optimize
	for {
		improved := false
		for i := 0; i < len(paths)-1; i++ {
			if moveAntImproves(paths, allocation, i, i+1) {
				allocation[i]--
				allocation[i+1]++
				improved = true
			}
		}
		if !improved {
			break
		}
	}

	return allocation
}

func moveAntImproves(paths [][]string, allocation []int, fromPath, toPath int) bool {
	currentMax := max(len(paths[fromPath])+allocation[fromPath], len(paths[toPath])+allocation[toPath])
	newMax := max(len(paths[fromPath])+allocation[fromPath]-1, len(paths[toPath])+allocation[toPath]+1)
	return newMax < currentMax
}

func compareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reset(m map[int]bool) {
	for k := range m {
		m[k] = false
	}
}
