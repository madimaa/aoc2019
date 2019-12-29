package main

import (
	"fmt"
	"math"

	util "github.com/madimaa/aoc2019/lib"
)

type vector struct {
	x, y, z int
}

type moon struct {
	position vector
	velocity vector
}

func createMoon(x, y, z int) *moon {
	return &moon{position: vector{x, y, z}, velocity: vector{0, 0, 0}}
}

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("12.txt")
	moons := make([]*moon, 0)
	for _, item := range result {
		var x, y, z int
		fmt.Sscanf(item, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, createMoon(x, y, z))
	}

	steps := 10
	for i := 0; i < steps; i++ {
		for _, moon := range moons {
			applyGravity(moon, moons)
		}

		for _, moon := range moons {
			applyVelocity(moon)
		}
	}

	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += calculateMoonEnergy(moon)
		fmt.Println(fmt.Sprintf("pos=<x= %d, y= %d, z= %d>, vel=<x= %d, y= %d, z= %d>", moon.position.x, moon.position.y, moon.position.z, moon.velocity.x, moon.velocity.y, moon.velocity.z))
	}

	fmt.Println("Total energy", totalEnergy)

	util.Elapsed()
}

func applyGravity(moon *moon, moons []*moon) {
	for _, other := range moons {
		if other != moon {
			if other.position.x > moon.position.x {
				moon.velocity.x++
			} else if other.position.x < moon.position.x {
				moon.velocity.x--
			}

			if other.position.y > moon.position.y {
				moon.velocity.y++
			} else if other.position.y < moon.position.y {
				moon.velocity.y--
			}

			if other.position.z > moon.position.z {
				moon.velocity.z++
			} else if other.position.z < moon.position.z {
				moon.velocity.z--
			}
		}
	}
}

func applyVelocity(moon *moon) {
	moon.position.x += moon.velocity.x
	moon.position.y += moon.velocity.y
	moon.position.z += moon.velocity.z
}

func calculateMoonEnergy(moon *moon) int {
	//potential energy
	potential := int(math.Abs(float64(moon.position.x)) + math.Abs(float64(moon.position.y)) + math.Abs(float64(moon.position.z)))

	//kinetic energy
	kinetic := int(math.Abs(float64(moon.velocity.x)) + math.Abs(float64(moon.velocity.y)) + math.Abs(float64(moon.velocity.z)))

	return potential * kinetic
}

func part2() {
	//I want to do very similar things, but i don't want to ruin the first solution with part 2 codes
	fmt.Println("Part 2")
	util.Start()

	result := util.OpenFile("12.txt")
	moons := make([]*moon, 0)
	initialPos := make([]vector, 0)
	var stepX, stepY, stepZ int64
	for _, item := range result {
		var x, y, z int
		fmt.Sscanf(item, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, createMoon(x, y, z))
		initialPos = append(initialPos, vector{x, y, z})
	}

	var step int64 = 0
	for {
		if step != 0 {
			var x, y, z bool = true, true, true
			for i, moon := range moons {
				if x && stepX == 0 && (moon.position.x != initialPos[i].x || moon.velocity.x != 0) {
					x = false
				}

				if y && stepY == 0 && (moon.position.y != initialPos[i].y || moon.velocity.y != 0) {
					y = false
				}

				if z && stepZ == 0 && (moon.position.z != initialPos[i].z || moon.velocity.z != 0) {
					z = false
				}
			}

			if stepX == 0 && x {
				stepX = step
			}

			if stepY == 0 && y {
				stepY = step
			}

			if stepZ == 0 && z {
				stepZ = step
			}

			if stepX != 0 && stepY != 0 && stepZ != 0 {
				break
			}
		}

		for _, moon := range moons {
			applyGravity(moon, moons)
		}

		for _, moon := range moons {
			applyVelocity(moon)
		}

		step++
	}

	fmt.Println(stepX, stepY, stepZ)
	fmt.Println("Steps:", util.Lcm(stepX, stepY, stepZ))

	util.Elapsed()
}
