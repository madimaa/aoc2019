package main

import (
	"container/list"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/intcode"
)

const (
	north   = 1
	south   = 2
	west    = 3
	east    = 4
	wall    = "#"
	path    = "."
	deadEnd = "x"
	oxygen  = "O"
)

type point struct {
	x, y int
}

var oxygenSystem point

func (vec *point) move(direction int) {
	switch direction {
	case north:
		vec.y++
	case south:
		vec.y--
	case west:
		vec.x--
	case east:
		vec.x++
	}
}

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("15.txt")
	content := strings.Split(result[0], ",")
	intcodeArr := make([]int, 10000)
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	part1(intcodeArr)
	util.Elapsed()
	fmt.Println("Part 2")
	util.Start()
	part2(intcodeArr)
	util.Elapsed()
}

func part1(intcodeArr []int) {
	computer := intcode.CreateComputer(intcodeArr)
	var output int
	halt := false

	journeyMap := make(map[string]string)
	position := point{0, 0}
	track := list.New()
	track.PushBack(position)
	direction := north

	handleTrack := func() {
		track.Remove(track.Back())
		track.Remove(track.Back())
	}

	for !halt {
		computer.AddInput(direction)
		_, output, halt = computer.Computer()
		switch output {
		case 0:
			//The droid hit a wall, position has not changed.
			var x, y int = position.x, position.y
			switch direction {
			case north:
				y++
			case south:
				y--
			case west:
				x--
			case east:
				x++
			}

			key := fmt.Sprintf("%d %d", x, y)
			journeyMap[key] = wall
			direction = newDirection(journeyMap, position, direction, handleTrack)
		case 1:
			//The droid has moved in the requested direction.
			position.move(direction)
			track.PushBack(position)
			key := fmt.Sprintf("%d %d", position.x, position.y)
			journeyMap[key] = path
			direction = newDirection(journeyMap, position, direction, handleTrack)
		case 2:
			//The droid has moved in the requested direction. The location is the oxygen system.
			position.move(direction)
			track.PushBack(position)
			halt = true
		}
	}

	oxygenSystem = position
	fmt.Println(fmt.Sprintf("Oxygen system's position: %d %d", position.x, position.y))
	fmt.Println(track.Len() - 1) //track starts with 0.0 it is not a step but the starting position
}

func part2(intcodeArr []int) {
	computer := intcode.CreateComputer(intcodeArr)
	var output int
	halt := false

	journeyMap := make(map[string]string)
	position := point{0, 0}
	direction := north

	for !halt {
		computer.AddInput(direction)
		_, output, halt = computer.Computer()
		switch output {
		case 0:
			//The droid hit a wall, position has not changed.
			var x, y int = position.x, position.y
			switch direction {
			case north:
				y++
			case south:
				y--
			case west:
				x--
			case east:
				x++
			}

			key := fmt.Sprintf("%d %d", x, y)
			journeyMap[key] = wall
		case 1:
			//The droid has moved in the requested direction.
			position.move(direction)
			key := fmt.Sprintf("%d %d", position.x, position.y)
			journeyMap[key] = path
		case 2:
			//The droid has moved in the requested direction. The location is the oxygen system.
			position.move(direction)
			key := fmt.Sprintf("%d %d", position.x, position.y)
			journeyMap[key] = oxygen
			//halt = true
		}

		direction = newDirection(journeyMap, position, direction, nil)
		if direction == 0 {
			halt = true
		}
	}

	printImage(journeyMap)

	fmt.Printf("Oxygen spread in %d minutes\n", spreadOxygen(journeyMap))
}

func newDirection(journeyMap map[string]string, position point, direction int, handleTrack func()) int {
	x, y := position.x, position.y
	if checkCoordinates(journeyMap, x, y+1, false) {
		return north
	} else if checkCoordinates(journeyMap, x, y-1, false) {
		return south
	} else if checkCoordinates(journeyMap, x-1, y, false) {
		return west
	} else if checkCoordinates(journeyMap, x+1, y, false) {
		return east
	}

	if handleTrack != nil {
		handleTrack()
	}

	if checkCoordinates(journeyMap, x, y+1, true) {
		key := fmt.Sprintf("%d %d", x, y)
		journeyMap[key] = deadEnd
		return north
	} else if checkCoordinates(journeyMap, x, y-1, true) {
		key := fmt.Sprintf("%d %d", x, y)
		journeyMap[key] = deadEnd
		return south
	} else if checkCoordinates(journeyMap, x-1, y, true) {
		key := fmt.Sprintf("%d %d", x, y)
		journeyMap[key] = deadEnd
		return west
	} else if checkCoordinates(journeyMap, x+1, y, true) {
		key := fmt.Sprintf("%d %d", x, y)
		journeyMap[key] = deadEnd
		return east
	}

	return 0
}

func checkCoordinates(journeyMap map[string]string, x, y int, canGoBack bool) bool {
	key := fmt.Sprintf("%d %d", x, y)
	v, ok := journeyMap[key]

	if !ok {
		return true
	} else if v == wall || v == deadEnd {
		return false
	} else if (v == path || v == oxygen) && canGoBack {
		return true
	}

	return false
}

func spreadOxygen(journeyMap map[string]string) int {
	key := fmt.Sprintf("%d %d", oxygenSystem.x, oxygenSystem.y)
	journeyMap[key] = oxygen

	minutes := 0
	spread := true
	for spread {
		spread = false
		oxygenCoords := findOxygens(journeyMap)
		for _, coordinates := range oxygenCoords {
			northCoord := fmt.Sprintf("%d %d", coordinates.x, coordinates.y+1)
			southCoord := fmt.Sprintf("%d %d", coordinates.x, coordinates.y-1)
			westCoord := fmt.Sprintf("%d %d", coordinates.x-1, coordinates.y)
			eastCoord := fmt.Sprintf("%d %d", coordinates.x+1, coordinates.y)

			if journeyMap[northCoord] != wall && journeyMap[northCoord] != oxygen {
				journeyMap[northCoord] = oxygen
				spread = true
			}

			if journeyMap[southCoord] != wall && journeyMap[southCoord] != oxygen {
				journeyMap[southCoord] = oxygen
				spread = true
			}

			if journeyMap[westCoord] != wall && journeyMap[westCoord] != oxygen {
				journeyMap[westCoord] = oxygen
				spread = true
			}

			if journeyMap[eastCoord] != wall && journeyMap[eastCoord] != oxygen {
				journeyMap[eastCoord] = oxygen
				spread = true
			}
		}

		if spread {
			minutes++
		}
	}

	return minutes
}

func findOxygens(journeyMap map[string]string) []point {
	coords := make([]point, 0)
	for k, v := range journeyMap {
		if v == oxygen {
			var x, y int
			fmt.Sscanf(k, "%d %d", &x, &y)

			//check adjacent
			northCoord := fmt.Sprintf("%d %d", x, y+1)
			southCoord := fmt.Sprintf("%d %d", x, y-1)
			westCoord := fmt.Sprintf("%d %d", x-1, y)
			eastCoord := fmt.Sprintf("%d %d", x+1, y)

			ok := false
			if journeyMap[northCoord] != wall && journeyMap[northCoord] != oxygen {
				ok = true
			} else if journeyMap[southCoord] != wall && journeyMap[southCoord] != oxygen {
				ok = true
			} else if journeyMap[westCoord] != wall && journeyMap[westCoord] != oxygen {
				ok = true
			} else if journeyMap[eastCoord] != wall && journeyMap[eastCoord] != oxygen {
				ok = true
			}

			if ok {
				coords = append(coords, point{x, y})
			}
		}
	}

	return coords
}

func printImage(pixels map[string]string) {
	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{100, 100}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for k, v := range pixels {
		var x, y int
		fmt.Sscanf(k, "%d %d", &x, &y)
		if x == oxygenSystem.x && y == oxygenSystem.y {
			v = oxygen
		}

		x += 50
		y += 50

		if x == 50 && y == 50 {
			img.Set(x, y, color.RGBA{0, 255, 0, 255})
			continue
		}

		switch v {
		case wall:
			img.Set(x, y, color.Black)
		case path, deadEnd:
			img.Set(x, y, color.White)
		case oxygen:
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	file, err := os.Create("15.png")
	util.PanicOnError(err)
	png.Encode(file, img)
}
