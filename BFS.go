package main

import (
	"container/list"
	"fmt"
)

func BFS(colony Colony) {
	roomNames := []string{}
	for _, room := range colony.rooms {
		roomNames = append(roomNames, room.name)
	}
	distances:=make(map[string]int)
	predescessors:=make(map[string]string)
	start:=colony.startRoom.name
	end:=colony.endRoom.name


	for _, room := range colony.rooms {
		distances[room.name] = -1
		predescessors[room.name] = ""
	}
	distances[start] = 0

	queue:= list.New()
	queue.PushBack(start)
	visited:=make(map[string]bool)
	visited[start] = true

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
