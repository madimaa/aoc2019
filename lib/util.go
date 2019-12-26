package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var timeNano int64

//LogOnError - check and log the error
func LogOnError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//PanicOnError - panic on error
func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

//OpenFile - Open file from path, return file content in string array/slice
func OpenFile(path string) []string {
	file, err := os.Open(path)
	PanicOnError(err)

	scanner := bufio.NewScanner(file)
	fileContent := make([]string, 0)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	LogOnError(scanner.Err())
	LogOnError(file.Close())

	return fileContent
}

//Start - set the start timer
func Start() {
	timeNano = time.Now().UnixNano()
}

//Elapsed - printf the elapsed time from Start
func Elapsed() {
	fmt.Printf("Runtime: %f\n", float64(time.Now().UnixNano()-timeNano)/float64(time.Second))
}

//ContainsInt - returns true if val exists in slice
func ContainsInt(slice []int, val int) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}

	return false
}

//ContainsStr - returns true if val exists in slice
func ContainsStr(slice []string, val string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}

	return false
}

//ContainsFloat - returns true if val exists in slice
func ContainsFloat(slice []float64, val float64) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}

	return false
}
