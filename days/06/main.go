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
	tree := make(map[string][]string)
	for _, s := range result {
		split := strings.Split(s, ")")
		left := split[0]
		right := split[1]
		if tree[left] == nil {
			tree[left] = make([]string, 0)
		}

		tree[left] = append(tree[left], right)

		for k, v := range tree {
			if util.ContainsStr(v, left) {
				tree[k] = append(tree[k], right)
			}
		}
	}

	//okay it is not an algorithm. it is just bruteforce
	changed := true
	for changed {
		changed = addRemainder(tree)
	}

	sum := 0
	for _, v := range tree {
		sum += len(v)
	}

	fmt.Println(sum)
	util.Elapsed()
}

func addRemainder(tree map[string][]string) bool {
	ret := false
	for k, v := range tree {
		var newValues = make([]string, len(v))
		copy(newValues, v)
		for _, x1 := range newValues {
			if tree[x1] != nil {
				for _, x2 := range tree[x1] {
					if !util.ContainsStr(newValues, x2) {
						ret = true
						newValues = append(newValues, x2)
					}
				}
			}
		}

		if ret {
			tree[k] = newValues
		}
	}

	return ret
}
