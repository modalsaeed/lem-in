package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

const (
	imageWidth  = 800
	imageHeight = 600
	roomSize    = 20
)

func DrawColony(colony Colony) {
	colonyImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(colonyImage, colonyImage.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// Find min and max coordinates
	minX, minY, maxX, maxY := findMinMaxCoordinates(colony.rooms)

	for _, room := range colony.rooms {
		drawRoom(colonyImage, room, minX, minY, maxX, maxY, colony)
	}

	for _, path := range colony.paths {
		drawPath(colonyImage, colony.rooms, path, minX, minY, maxX, maxY)
	}

	file, err := os.Create("colony.png")
	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer file.Close()
	png.Encode(file, colonyImage)
}

func findMinMaxCoordinates(rooms []Room) (minX, minY, maxX, maxY int) {
	minX, minY = math.MaxInt32, math.MaxInt32
	maxX, maxY = math.MinInt32, math.MinInt32

	for _, room := range rooms {
		minX = min(minX, room.xCoord)
		minY = min(minY, room.yCoord)
		maxX = max(maxX, room.xCoord)
		maxY = max(maxY, room.yCoord)
	}
	return
}

func scaleCoordinate(value, minValue, maxValue, dimension int) int {
	if maxValue == minValue {
		return dimension / 2
	}
	return (value-minValue)*(dimension-roomSize*2)/(maxValue-minValue) + roomSize
}

func drawRoom(colonyImage *image.RGBA, room Room, minX, minY, maxX, maxY int, colony Colony) {
	x := scaleCoordinate(room.xCoord, minX, maxX, imageWidth)
	y := scaleCoordinate(room.yCoord, minY, maxY, imageHeight)

	roomColor := color.RGBA{0, 0, 255, 255}
	if room.name == colony.startRoom.name {
		roomColor = color.RGBA{0, 255, 0, 255}
	}
	if room.name == colony.endRoom.name {
		roomColor = color.RGBA{255, 165, 0, 255}
	}
	draw.Draw(colonyImage, image.Rect(x-roomSize/2, y-roomSize/2, x+roomSize/2, y+roomSize/2), &image.Uniform{roomColor}, image.Point{}, draw.Src)
}

func drawPath(colonyImage *image.RGBA, rooms []Room, path Path, minX, minY, maxX, maxY int) {
	var from, to Room

	for _, room := range rooms {
		if room.name == path.room1 {
			from = room
		}
		if room.name == path.room2 {
			to = room
		}
	}

	fromX := scaleCoordinate(from.xCoord, minX, maxX, imageWidth)
	fromY := scaleCoordinate(from.yCoord, minY, maxY, imageHeight)
	toX := scaleCoordinate(to.xCoord, minX, maxX, imageWidth)
	toY := scaleCoordinate(to.yCoord, minY, maxY, imageHeight)

	pathColor := color.RGBA{255, 0, 0, 255}
	drawLine(colonyImage, fromX, fromY, toX, toY, pathColor)
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.RGBA) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
