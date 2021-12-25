package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	sum := 0
	prevNum := -1

	for scanner.Scan() { 
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		if prevNum != -1 && num > prevNum {
			sum++
		}

		prevNum = num
	}

	fmt.Println(sum)
}