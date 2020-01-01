package main

import (
	"fmt"
	"math"
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

	layers := orderByLevel(chemicals)

	calculateFuel(chemicals, layers)
	util.Elapsed()
}

func orderByLevel(chemicals recipes) map[string]int {
	layers := make(map[string]int)

	//0th layer is FUEL
	layers["FUEL"] = 0
	actualLayer := make([]string, 1)
	actualLayer[0] = "FUEL"

	nextLayer := make([]string, 0)
	level := 1
	for {
		inputChemicals := chemicals.product[actualLayer[0]]
		for _, item := range inputChemicals {
			if v, ok := layers[item.name]; !ok || level > v {
				layers[item.name] = level
			}

			if !util.ContainsStr(nextLayer, item.name) {
				nextLayer = append(nextLayer, item.name)
			}
		}

		//remove the first element
		actualLayer = actualLayer[1:]

		if len(nextLayer) == 0 && len(actualLayer) == 0 {
			break
		}

		if len(actualLayer) == 0 {
			actualLayer = make([]string, len(nextLayer))
			copy(actualLayer, nextLayer)
			nextLayer = make([]string, 0)
			level++
		}
	}

	return layers
}

func splitChemical(input string) (int, string) {
	output := strings.Split(strings.TrimSpace(input), " ")
	quantity, err := strconv.Atoi(output[0])
	util.PanicOnError(err)

	return quantity, output[1]
}

func calculateFuel(chemicals recipes, layers map[string]int) {
	requiredChemicals := make(map[string]int)
	inputChemicals := chemicals.product["FUEL"]
	for _, item := range inputChemicals {
		requiredChemicals[item.name] = item.quantity
	}

	ore := calculateRequiredQuantities(chemicals, requiredChemicals, layers)
	fmt.Println(ore)
}

func calculateRequiredQuantities(chemicals recipes, required map[string]int, layers map[string]int) int {
	requiredChemicals := make(map[string]int)

	hadElement := true
	level := 0
	for hadElement {
		hadElement = false
		for input, value := range layers {
			if value == level && input != "ORE" {
				requiredQuantity := 1
				if _, ok := requiredChemicals[input]; ok {
					requiredQuantity = requiredChemicals[input]
					if requiredQuantity%chemicals.getQuantity(input) != 0 {
						requiredQuantity += chemicals.getQuantity(input) - requiredQuantity%chemicals.getQuantity(input)
					}
					delete(requiredChemicals, input)
				}

				ingredients := chemicals.product[input]
				outputQuantity := chemicals.getQuantity(input)
				for _, item := range ingredients {
					value := int(math.Ceil(float64(requiredQuantity)/float64(outputQuantity)) * float64(item.quantity))
					if _, ok := requiredChemicals[item.name]; ok {
						requiredChemicals[item.name] += value
					} else {
						requiredChemicals[item.name] = value
					}
				}

				hadElement = true
			}
		}

		level++
	}

	return requiredChemicals["ORE"]
}
