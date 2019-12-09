package main

import (
	"fmt"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")
	util.Start()
	result := util.OpenFile("06.txt")
	//result := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	backward := make(map[string]string)

	for _, s := range result {
		split := strings.Split(s, ")")
		left := split[0]
		right := split[1]

		backward[right] = left
	}

	sum := 0
	for k := range backward {
		sum += goBackwardsByKey(backward, k)
	}

	fmt.Println(sum)
	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	you := goBackwardsWithDistance(backward, "YOU")
	san := goBackwardsWithDistance(backward, "SAN")
	min := 0
	for k, youV := range you {
		sanV, ok := san[k]
		if ok {
			if min == 0 || sanV+youV < min {
				min = sanV + youV
			}
		}
	}

	fmt.Println(min - 2) //minus 2, because of: Between the objects they are orbiting - not between YOU and SAN.
	util.Elapsed()
}

func goBackwardsByKey(backward map[string]string, start string) int {
	distance := 0

	v, ok := backward[start]
	for ok {
		distance++
		v, ok = backward[v]
	}

	return distance
}

func goBackwardsWithDistance(backward map[string]string, start string) map[string]int {
	result := make(map[string]int)
	distance := 0
	v, ok := backward[start]
	for ok {
		distance++
		result[v] = distance
		v, ok = backward[v]
	}

	return result
}
