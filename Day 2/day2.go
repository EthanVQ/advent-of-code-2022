package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	data := read_puzzle_input(filename)

	// print_puzzle_input(data)

	star1 := solve_star1(data)

	fmt.Println(star1)

	star2 := solve_star2(data)

	fmt.Println(star2)
}

func solve_star1(data [][]string) int {
	scoreDict := map[string]int{
		"AX": 1 + 3,
		"AY": 2 + 6,
		"AZ": 3 + 0,
		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,
		"CX": 1 + 6,
		"CY": 2 + 0,
		"CZ": 3 + 3,
	}

	totalScore := 0
	for _, round := range data {
		combination := round[0] + round[1]
		totalScore += scoreDict[combination]
	}

	return totalScore
}

func solve_star2(data [][]string) int {
	scoreDict := map[string]int{
		"AX": 3 + 0,
		"AY": 1 + 3,
		"AZ": 2 + 6,
		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,
		"CX": 2 + 0,
		"CY": 3 + 3,
		"CZ": 1 + 6,
	}

	totalScore := 0
	for _, round := range data {
		combination := round[0] + round[1]
		totalScore += scoreDict[combination]
	}

	return totalScore
}

func read_puzzle_input(filename string) [][]string {
	data := [][]string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		text := scanner.Text()
		text = strings.Trim(text, " ")
		splitText := strings.Split(text, " ")
		data = append(data, splitText)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func print_puzzle_input(data [][]string) {
	for index, element := range data {
		fmt.Printf("Item %s\n", index)
		for _, element2 := range element {
			fmt.Println(element2)
		}
	}
}

type winConditions struct {
	AX int
	AY int
	AZ int
	BX int
	BY int
	BZ int
	CX int
	CY int
	CZ int
}
