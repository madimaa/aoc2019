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

	size := 1000
	var centralX, centralY int = size / 2, size / 2

	firstMap := drawCable(centralX, centralY, firstWire)
	secondMap := drawCable(centralX, centralY, secondWire)

	cross := make(map[int]int)
	for k, v := range firstMap {
		secondVals := secondMap[k]
		if secondVals != nil {
			for _, val := range v {
				if contains(secondVals, val) {
					cross[k] = val
				}
			}
		}
	}

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
	util.Start()

	firstDistance := calculateDistance(cross, centralX, centralY, size, firstWire)
	secondDistance := calculateDistance(cross, centralX, centralY, size, secondWire)

	shortestDistance := 0
	for k, v := range cross {
		key := size*k + v
		if shortestDistance == 0 {
			shortestDistance = firstDistance[key] + secondDistance[key]
		} else {
			actual := firstDistance[key] + secondDistance[key]
			if actual < shortestDistance {
				shortestDistance = actual
			}
		}
	}

	fmt.Println(shortestDistance)
	util.Elapsed()
}

func drawCable(centralX, centralY int, wire []string) map[int][]int {
	var posX, posY int = centralX, centralY

	result := make(map[int][]int)
	for _, s := range wire {
		length, err := strconv.Atoi(s[1:len(s)])
		util.LogOnError(err)
		posX, posY = draw(result, posX, posY, length, s[0])
	}

	return result
}

func draw(output map[int][]int, startX, startY, length int, direction byte) (int, int) {
	switch direction {
	case 'U':
		for i := 1; i <= length; i++ {
			startY++
			setMap(output, startX, startY)
		}
	case 'R':
		for i := 1; i <= length; i++ {
			startX++
			setMap(output, startX, startY)
		}
	case 'D':
		for i := 1; i <= length; i++ {
			startY--
			setMap(output, startX, startY)
		}
	case 'L':
		for i := 1; i <= length; i++ {
			startX--
			setMap(output, startX, startY)
		}
	}

	return startX, startY
}

func setMap(output map[int][]int, x, y int) {
	if output[x] == nil {
		output[x] = make([]int, 0)
	}

	output[x] = append(output[x], y)
}

func calculateDistance(cross map[int]int, centralX, centralY, size int, wire []string) map[int]int {
	var posX, posY int = centralX, centralY

	result := make(map[int]int)
	distance := 0
	for _, s := range wire {
		length, err := strconv.Atoi(s[1:len(s)])
		util.LogOnError(err)
		posX, posY, distance = distanceInOneDireciton(posX, posY, length, size, distance, s[0], cross, result)
	}

	return result
}

func distanceInOneDireciton(startX, startY, length, size, distance int, direction byte, cross, result map[int]int) (int, int, int) {
	i := 1
	switch direction {
	case 'U':
		for ; i <= length; i++ {
			startY++
			checkCrossing(result, cross, startX, startY, size, distance, i)
		}
	case 'R':
		for ; i <= length; i++ {
			startX++
			checkCrossing(result, cross, startX, startY, size, distance, i)
		}
	case 'D':
		for ; i <= length; i++ {
			startY--
			checkCrossing(result, cross, startX, startY, size, distance, i)
		}
	case 'L':
		for ; i <= length; i++ {
			startX--
			checkCrossing(result, cross, startX, startY, size, distance, i)
		}
	}

	return startX, startY, distance + i - 1 //the last i++ will cause bad distance calculation
}

func checkCrossing(result, cross map[int]int, x, y, size, distance, i int) {
	if cross[x] == y {
		result[x*size+y] = distance + i
	}
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}
