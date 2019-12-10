package main

import (
	"fmt"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	var i, max int64 = 0, 3124 //Quinary numeral system (base 5) largest number with max 3124 is 44444
	//5*5*5*5*5 (5^5) = 3125
	for ; i <= max; i++ {
		a, b, c, d, e := splitQuinaryFromNumber(i)
		fmt.Println(a, b, c, d, e)
	}
}

func splitQuinaryFromNumber(number int64) (int64, int64, int64, int64, int64) {
	base := 5
	quinaryString := strconv.FormatInt(number, base)
	fmt.Println(quinaryString)
	length := len(quinaryString)

	var a, b, c, d, e int64 = 0, 0, 0, 0, 0
	var err error

	if length >= 1 {
		e, err = strconv.ParseInt(quinaryString[length-1:length], base, 64)
		util.PanicOnError(err)
	}

	if length >= 2 {
		num := quinaryString[length-2 : length-1]
		d, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 3 {
		num := quinaryString[length-3 : length-2]
		c, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 4 {
		num := quinaryString[length-4 : length-3]
		b, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 5 {
		num := quinaryString[length-5 : length-4]
		a, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	return a, b, c, d, e
}
