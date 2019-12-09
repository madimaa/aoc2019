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
	//result := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	backward := make(map[string]string)

	for _, s := range result {
		split := strings.Split(s, ")")
		left := split[0]
		right := split[1]

		backward[right] = left
	}

	sum := 0
	for k := range backward {
		sum += goBackwards(backward, k)
	}

	fmt.Println(sum)
	util.Elapsed()
}

func goBackwards(backward map[string]string, start string) int {
	distance := 0

	v, ok := backward[start]
	for ok {
		distance++
		v, ok = backward[v]
	}

	return distance
}
