package main

import (
	"sort"
)

func findPaths(colony Colony) [][]string {
	start := colony.startRoom.name
	end := colony.endRoom.name
	allPaths := [][]string{}
	selectedPaths := [][]string{}

	DFS(colony, start, end, map[string]bool{}, []string{}, &allPaths)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	pathSets:=[][][]string{}

	for _, path := range allPaths {
		for i,set:=range pathSets{
			if isUnique(path, set) {
				pathSets[i]=append(pathSets[i], path)
				break
			}
		}
			pathSets=append(pathSets, [][]string{path})
	}
	selectedPaths=selectOptimalSet(pathSets)
	return selectedPaths
}

func DFS(colony Colony, current, end string, visited map[string]bool, currentPath []string, allPaths *[][]string) {

	currentPath = append(currentPath, current)

	if current == end {
		pathCopy := make([]string, len(currentPath))
		copy(pathCopy, currentPath)
		*allPaths = append(*allPaths, pathCopy)
		return
	}

	if current != end && current != colony.startRoom.name {
		visited[current] = true
	}

	var currentRoom Room
	for _, room := range colony.rooms {
		if room.name == current {
			currentRoom = room
			break
		}
	}

	for _, connectedRoom := range currentRoom.connectedRooms {
		if (!visited[connectedRoom] || connectedRoom == end) && connectedRoom != colony.startRoom.name {
			localVisited := make(map[string]bool)
			for k, v := range visited {
				localVisited[k] = v
			}
			DFS(colony, connectedRoom, end, localVisited, currentPath, allPaths)
		}
	}
}

func contains(completedPath []string, currentPath []string) bool {
	for _, completedRoom := range completedPath {
		for _, currentRoom := range currentPath {
			if completedRoom == currentRoom {
				return true
			}
		}
	}
	return false
}

func isUnique(path []string, selectedPaths [][]string) bool {
	for _, selectedPath := range selectedPaths {
		if contains(path[1:len(path)-1], selectedPath[1:len(selectedPath)-1]) {
			return false
		}
	}
	return true
}

func selectOptimalSet(pathSets [][][]string) [][]string {
    var bestSet [][]string
    bestScore := float64(0)

    for _, set := range pathSets {
        // Calculate average path length
        totalLength := 0
        for _, path := range set {
            totalLength += len(path)
        }
        avgLength := float64(totalLength) / float64(len(set))
        
        // Score = number of paths / average length
        // This favors more paths with shorter lengths
        score := float64(len(set)) / avgLength
        
        if score > bestScore {
            bestScore = score
            bestSet = set
        }
    }
    
    return bestSet
}

