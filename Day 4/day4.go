package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	data := read_puzzle_input(filename)

	// print_puzzle_input(data)

	star1 := solve_first_star(data)

	fmt.Println(star1)

	star2 := solve_second_star(data)

	fmt.Println(star2)
}

func solve_first_star(data [][][]int) int {
	containsCount := 0
	for _, pair := range data {
		if (pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1]) || (pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1]) {
			containsCount++
		}
	}
	return containsCount
}

func solve_second_star(data [][][]int) int {
	containsCount := 0
	for _, pair := range data {
		if pair[0][1] >= pair[1][0] && pair[1][1] >= pair[0][0] || pair[1][1] >= pair[0][0] && pair[0][1] >= pair[1][0] {
			containsCount++
		}
	}
	return containsCount
}

func print_puzzle_input(data [][][]int) {
	for index, pair := range data {
		fmt.Println(index)
		fmt.Println(strconv.Itoa(pair[0][0]) + "-" + strconv.Itoa(pair[0][1]))
		fmt.Println(strconv.Itoa(pair[1][0]) + "-" + strconv.Itoa(pair[1][1]))
	}
}

func read_puzzle_input(filename string) [][][]int {
	data := [][][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, " ")

		pairOfIdStrings := strings.Split(text, ",")

		pair := [][]int{}
		for _, idString := range pairOfIdStrings {
			hyphonIndex := strings.Index(idString, "-")
			lowerId, _ := strconv.Atoi(idString[:hyphonIndex])
			upperId, _ := strconv.Atoi(idString[hyphonIndex+1:])

			pair = append(pair, []int{lowerId, upperId})
		}
		data = append(data, pair)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
