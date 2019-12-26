package main

import (
	"fmt"
	"math"

	"github.com/madimaa/aoc2019/lib/array2d"

	util "github.com/madimaa/aoc2019/lib"
)

type vector struct {
	x, y                 int
	magnitude, direction float64
}

func createVector(x, y int, magnitude, direction float64) *vector {
	return &vector{x: x, y: y, magnitude: magnitude, direction: direction}
}

func (vec *vector) getCoordinates() (int, int) {
	return vec.x, vec.y
}

func (vec *vector) getVector() (int, int, float64, float64) {
	return vec.x, vec.y, vec.magnitude, vec.direction
}

var directions map[float64][]*vector

func main() {
	directions = make(map[float64][]*vector)
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("10.txt")

	array := array2d.Create(len(result[0]), len(result))
	for y, s := range result {
		for x := 0; x < len(s); x++ {
			array.Put(x, y, string(s[x]))
		}
	}

	maxX, maxY := array.GetSize()
	detected := 0
	var detX, detY int
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if array.Get(x, y) == "#" {
				current := detectableAsteroids(array, x, y, false)
				if current > detected {
					detX = x
					detY = y
					detected = current
				}
			}
		}
	}

	fmt.Println(detX, detY, detected)
	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	detectableAsteroids(array, detX, detY, true)
	vaporizer(array, detX, detY)

	util.Elapsed()
}

func detectableAsteroids(array *array2d.Array2D, originalX, originalY int, fillDirections bool) int {
	list := make([]float64, 0)
	maxX, maxY := array.GetSize()
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if array.Get(x, y) == "#" {
				if originalX == x && originalY == y {
					continue
				}

				direction := math.Atan2(float64(originalY-y), float64(x-originalX))
				if !util.ContainsFloat(list, direction) {
					list = append(list, direction)
				}

				if fillDirections {
					if directions[direction] == nil {
						directions[direction] = make([]*vector, 0)
					}

					magnitude := math.Sqrt(math.Pow(float64(originalY-y), 2.0) + math.Pow(float64(x-originalX), 2.0))
					vec := createVector(x, y, magnitude, direction)
					directions[direction] = append(directions[direction], vec)
				}
			}
		}
	}

	return len(list)
}

func vaporizer(array *array2d.Array2D, laserX, laserY int) {
	destroyed := 0
	vaporizerDirection := math.Atan2(1, 0)
	for len(directions) > 0 {
		content := directions[vaporizerDirection]
		if content != nil {
			distance := 0.0
			index := 0
			var asteroid *vector
			for i := range content {
				vec := content[i]
				if distance > vec.magnitude || distance == 0.0 {
					index = i
					distance = vec.magnitude
					asteroid = vec
				}
			}

			content[index] = content[len(content)-1]
			content[len(content)-1] = nil
			content = content[:len(content)-1]

			destroyed++

			if len(content) == 0 {
				delete(directions, vaporizerDirection)
			} else {
				directions[vaporizerDirection] = content
			}

			if destroyed == 200 {
				fmt.Println(asteroid.x*100 + asteroid.y)
			}

			vaporizerDirection = findNextDirection(vaporizerDirection)
		}
	}
}

func findNextDirection(direction float64) float64 {
	closest := -42.0 //magic numbers
	closestDistance := 42.0
	for key := range directions {
		val := direction - key

		if closestDistance > val && direction > key {
			closestDistance = val
			closest = key
		}
	}

	if closest == -42.0 && closestDistance == 42.0 {
		for key := range directions {
			if key > closest || closest == -42.0 {
				closest = key
			}
		}
	}

	return closest
}
