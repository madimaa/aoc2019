package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
)

type ingredient struct {
	name     string
	quantity int
}

type recipes struct {
	product  map[string][]ingredient
	quantity map[string]int
}

func initRecipes() recipes {
	return recipes{product: make(map[string][]ingredient), quantity: make(map[string]int)}
}

func (rcp *recipes) add(key string, quantity int, ingredients []ingredient) {
	rcp.product[key] = ingredients
	rcp.quantity[key] = quantity
}

func (rcp *recipes) getQuantity(key string) int {
	return rcp.quantity[key]
}

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("14.txt")
	chemicals := initRecipes()
	for _, item := range result {
		split := strings.Split(item, "=>")
		outputQuantity, outputName := splitChemical(split[1])
		ingredients := make([]ingredient, 0)
		leftSide := strings.Split(split[0], ",")
		for _, item := range leftSide {
			quantity, name := splitChemical(item)
			ingredients = append(ingredients, ingredient{name, quantity})
		}
		chemicals.add(outputName, outputQuantity, ingredients)
	}

	calculateOreRequirement(chemicals)

	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	ore := 1000000000000
	calculateFuel(chemicals, ore)

	util.Elapsed()
}

func splitChemical(input string) (int, string) {
	output := strings.Split(strings.TrimSpace(input), " ")
	quantity, err := strconv.Atoi(output[0])
	util.PanicOnError(err)

	return quantity, output[1]
}

func calculateOreRequirement(chemicals recipes) {
	remainingIngredients := make(map[string]int)
	ore := calc(chemicals, remainingIngredients, ingredient{"FUEL", 1})
	fmt.Println(ore)
}

func calculateFuel(chemicals recipes, initialOreQuantity int) {
	remainingIngredients := make(map[string]int)
	remainingIngredients["ORE"] = initialOreQuantity
	producedFuel := 0

	const fuelQuantity = 10000

	for {
		if !produceFuel(chemicals, remainingIngredients, fuelQuantity) {
			break
		}

		producedFuel++
	}

	producedFuel *= fuelQuantity

	for {
		if !produceFuel(chemicals, remainingIngredients, 1) {
			break
		}

		producedFuel++
	}

	fmt.Println("Produced amount of fuel:", producedFuel)
}

func produceFuel(chemicals recipes, remainingIngredients map[string]int, fuelQuantity int) bool {
	remainingInput := make(map[string]int)
	for k, v := range remainingIngredients {
		remainingInput[k] = v
	}

	ore := calc(chemicals, remainingIngredients, ingredient{"FUEL", fuelQuantity})
	if remainingIngredients["ORE"] >= ore {
		remainingIngredients["ORE"] -= ore
		return true
	}

	for k := range remainingIngredients {
		delete(remainingIngredients, k)
	}

	for k, v := range remainingInput {
		remainingIngredients[k] = v
	}

	return false
}

func calc(chemicals recipes, remainingIngredients map[string]int, input ingredient) int {
	requiredQuantity := input.quantity
	remaining, ok := remainingIngredients[input.name]

	ore := 0

	if ok && remaining >= requiredQuantity {
		remainingIngredients[input.name] -= requiredQuantity
		if remainingIngredients[input.name] == 0 {
			delete(remainingIngredients, input.name)
		}

		return 0
	} else if ok {
		delete(remainingIngredients, input.name)
		requiredQuantity -= remaining
	}

	outputQuantity := chemicals.getQuantity(input.name)
	if requiredQuantity%outputQuantity != 0 {
		remainder := outputQuantity - requiredQuantity%outputQuantity
		if _, ok := remainingIngredients[input.name]; ok {
			remainingIngredients[input.name] += remainder
		} else {
			remainingIngredients[input.name] = remainder
		}

		requiredQuantity += remainder
	}

	ingredients := chemicals.product[input.name]
	for _, item := range ingredients {
		reqQuantity := requiredQuantity / outputQuantity * item.quantity
		if item.name == "ORE" {
			ore += reqQuantity
		} else {
			ore += calc(chemicals, remainingIngredients, ingredient{item.name, reqQuantity})
		}
	}

	return ore
}
