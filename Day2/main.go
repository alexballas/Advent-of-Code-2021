package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	input = "input_2"
)

type Operation struct {
	Op    string
	Value int
}

func main() {
	operations := make([]Operation, 0)

	f, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		line := strings.Split(buf.Text(), " ")
		number, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln(err)
		}

		operations = append(operations, Operation{line[0], number})
	}

	fmt.Println(part1(operations))

	fmt.Println(part2(operations))
}

func part1(o []Operation) int {
	allData := make(map[string]int)

	for _, op := range o {
		processOp_part1(op.Op, op.Value, allData)
	}
	return allData["depth"] * allData["forward"]
}

func processOp_part1(s string, i int, allData map[string]int) {
	switch s {
	case "forward":
		allData[s] = allData[s] + i
	case "down":
		allData["depth"] = allData["depth"] + i
	case "up":
		allData["depth"] = allData["depth"] - i
	}
}

func part2(o []Operation) int {
	allData := make(map[string]int)

	for _, op := range o {
		processOp_part2(op.Op, op.Value, allData)
	}
	return allData["depth"] * allData["forward"]
}

func processOp_part2(s string, i int, allData map[string]int) {
	switch s {
	case "forward":
		allData["depth"] = i*allData["aim"] + allData["depth"]
		allData["forward"] = allData["forward"] + i
	case "down":
		allData["aim"] = allData["aim"] + i
	case "up":
		allData["aim"] = allData["aim"] - i
	}
}
