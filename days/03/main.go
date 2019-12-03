package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")
	util.Start()
	result := util.OpenFile("03.txt")
	firstWire := strings.Split(result[0], ",")
	secondWire := strings.Split(result[1], ",")
	//firstWire := strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ",")
	//secondWire := strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")

	size := 50000
	panel := util.CreateArray2D(size, size)
	var centralX, centralY int = size / 2, size / 2
	panel.PutArray2D(centralX, centralY, "o")

	cross := make(map[int]int)
	drawCable(centralX, centralY, *panel, firstWire, "1", cross)
	drawCable(centralX, centralY, *panel, secondWire, "2", cross)
	// for i := 0; i < 30; i++ {
	// 	for j := 0; j < 30; j++ {
	// 		if panel.GetArray2D(j, i) == "" {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf(panel.GetArray2D(j, i))
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// os.Exit(1)
	var distance float64 = 0
	for k, v := range cross {
		manhattan := math.Abs(float64(centralX-k)) + math.Abs(float64(centralY-v))
		if manhattan < distance || distance == 0 {
			distance = manhattan
		}
	}

	fmt.Println(distance)
	util.Elapsed()
	fmt.Println("Part 2")
}

func drawCable(centralX, centralY int, panel util.Array2D, wire []string, wireNumber string, cross map[int]int) {
	var posX, posY int = centralX, centralY
	for _, s := range wire {
		length, err := strconv.Atoi(s[1:len(s)])
		util.LogOnError(err)
		posX, posY = draw(panel, posX, posY, length, s[0], wireNumber, cross)
	}
}

func draw(panel util.Array2D, startX, startY, length int, direction byte, wireNumber string, cross map[int]int) (int, int) {
	switch direction {
	case 'U':
		for i := 1; i <= length; i++ {
			startY++
			setPanel(panel, startX, startY, wireNumber, cross)
		}
	case 'R':
		for i := 1; i <= length; i++ {
			startX++
			setPanel(panel, startX, startY, wireNumber, cross)
		}
	case 'D':
		for i := 1; i <= length; i++ {
			startY--
			setPanel(panel, startX, startY, wireNumber, cross)
		}
	case 'L':
		for i := 1; i <= length; i++ {
			startX--
			setPanel(panel, startX, startY, wireNumber, cross)
		}
	}

	return startX, startY
}

func setPanel(panel util.Array2D, x, y int, wireNumber string, cross map[int]int) {
	content := panel.GetArray2D(x, y)
	if content == "" || content == wireNumber {
		panel.PutArray2D(x, y, wireNumber)
	} else {
		cross[x] = y
		panel.PutArray2D(x, y, "X")
	}
}
