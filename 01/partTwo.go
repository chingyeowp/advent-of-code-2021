package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
    input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err) 
	}

	content := strings.Split(string(input), "\n") 

	var window []int

	for i, _ := range content {
		windowSum := computeWindowSum(3, i, content)
		window = append(window, windowSum)
	}

	sum := 0
	prevNum := -1

	for _, windowSum := range window {
		if prevNum != -1 && windowSum > prevNum {
			sum++
		}
		prevNum = windowSum
	}
	fmt.Println(sum)
}

func computeWindowSum(windowSize int, curr int, input []string) int {
	sum := 0

	for i := 0; i < windowSize; i++ {
		if curr < len(input) {
			convertedInt, err := strconv.Atoi(input[curr]) 
			if err == nil {
				sum = sum + convertedInt
			}
			curr++
		}
	}

	return sum
}