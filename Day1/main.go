package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	input     = "input_1"
	allNumber = make([]int, 0)
)

func main() {
	loadData()
	fmt.Printf("Solution 1: %d\nSolution 2: %d\n", part1(), part2())
}

func loadData() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		number, err := strconv.Atoi(buf.Text())
		if err != nil {
			log.Fatalln(err)
		}
		allNumber = append(allNumber, number)
	}
}

func part1() int {

	previous := 0
	found := 0

	for counter := range allNumber {
		if allNumber[counter] > previous && previous > 0 {
			found++
		}
		previous = allNumber[counter]
	}

	return found
}

func part2() int {

	previous := 0
	found := 0

	for counter := 0; counter < len(allNumber)-2; counter++ {
		a := allNumber[counter]
		b := allNumber[counter+1]
		c := allNumber[counter+2]

		num := a + b + c

		if num > previous && previous > 0 {
			found++
		}

		previous = num
	}

	return found
}
