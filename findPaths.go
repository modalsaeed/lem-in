package main

import (
	"fmt"
	"strings"
)

func findPaths(colony Colony) [][]string {
	start := colony.startRoom.name
	end := colony.endRoom.name
	allPaths := [][]string{}

	DFS(colony, start, end, map[string]bool{}, []string{}, &allPaths)

	for i, path := range allPaths {
		fmt.Println("Path", i+1, ":", strings.Join(path, " -> "))
	}
	return allPaths
}

func DFS(colony Colony, current, end string, visited map[string]bool, currentPath []string, allPaths *[][]string) {

	for _, completedPath := range *allPaths {
		if contains(completedPath[1:], currentPath) {
			return
		}
	}

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
