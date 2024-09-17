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

	for _, path := range allPaths {
		if isUnique(path, selectedPaths) {
			selectedPaths = append(selectedPaths, path)
		}
	}
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
