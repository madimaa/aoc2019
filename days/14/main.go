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

var ore int

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("sample3.txt")
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

	calculateFuel(chemicals)
	util.Elapsed()
}

func splitChemical(input string) (int, string) {
	output := strings.Split(strings.TrimSpace(input), " ")
	quantity, err := strconv.Atoi(output[0])
	util.PanicOnError(err)

	return quantity, output[1]
}

func calculateFuel(chemicals recipes) {
	requiredChemicals := make(map[string]int)
	inputChemicals := chemicals.product["FUEL"]
	for _, item := range inputChemicals {
		requiredChemicals[item.name] = item.quantity
	}

	calculateRequiredQuantities(chemicals, requiredChemicals)
	fmt.Println(ore)
}

func calculateRequiredQuantities(chemicals recipes, required map[string]int) {
	requiredChemicals := make(map[string]int)
	for k, v := range required {
		if _, ok := requiredChemicals[k]; ok {
			requiredChemicals[k] += v
			continue
		}

		if k == "ORE" {
			if _, ok := requiredChemicals["ORE"]; ok {
				requiredChemicals["ORE"] += v
			} else {
				requiredChemicals["ORE"] = v
			}

			continue
		}

		inputChemicals := chemicals.product[k]
		for _, item := range inputChemicals {
			quantity, name := item.quantity, item.name
			value := int(math.Ceil(float64(v)/float64(chemicals.getQuantity(k))) * float64(quantity))
			//value := int(math.Ceil(float64(v*quantity) / float64(chemicals.getQuantity(k))))
			if _, ok := requiredChemicals[name]; ok {
				requiredChemicals[name] += value
			} else {
				requiredChemicals[name] = value
			}
		}
	}

	_, ok := requiredChemicals["ORE"]
	if len(requiredChemicals) > 1 || !ok {
		for k, v := range requiredChemicals {
			if k != "ORE" {
				quantity := chemicals.getQuantity(k)
				remaining := v % quantity
				if remaining != 0 {
					requiredChemicals[k] += quantity - remaining
				}
			}
		}

		calculateRequiredQuantities(chemicals, requiredChemicals)
	} else {
		ore = requiredChemicals["ORE"]
	}
}
