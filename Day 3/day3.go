package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	filename := os.Args[1]

	data := read_puzzle_input_star1(filename)

	// print_puzzle_input(data)

	star1 := solve_first_star(data)

	fmt.Println(star1)

	data = read_puzzle_input_star2(filename)

	// print_puzzle_input(data)

	star2 := solve_second_star(data)

	fmt.Println(star2)
}

func solve_first_star(data [][]string) int {
	sumOfPrioities := 0
	for _, rucksack := range data {
		rucksackPriority := 0
		firstCompartment := rucksack[0]
		secondCompartment := rucksack[1]

		for _, item1 := range strings.Split(firstCompartment, "") {
			for _, item2 := range strings.Split(secondCompartment, "") {
				if item1 == item2 {
					rucksackPriority = get_priority(item2)
				}
			}
		}

		sumOfPrioities += rucksackPriority
	}
	return sumOfPrioities
}

func solve_second_star(data [][]string) int {
	sumOfPrioities := 0
	for _, group := range data {
		rucksackPriority := 0

		for _, item1 := range strings.Split(group[0], "") {
			for _, item2 := range strings.Split(group[1], "") {
				for _, item3 := range strings.Split(group[2], "") {
					if (item1 == item2) && (item1 == item3) {
						rucksackPriority = get_priority(item2)
					}
				}
			}
		}

		sumOfPrioities += rucksackPriority
	}

	return sumOfPrioities
}

func get_priority(letter string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	priority := 0

	priority += strings.Index(alphabet, strings.ToLower(letter))

	priority += 1 //offset

	r, _ := utf8.DecodeRuneInString(letter)
	if unicode.IsUpper(r) {
		priority += 26 // capital
	}

	return priority
}

func read_puzzle_input_star1(filename string) [][]string {
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

		firstCompartment := text[len(text)/2:]
		secondCompartment := text[:len(text)/2]

		item := []string{firstCompartment, secondCompartment}
		data = append(data, item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func read_puzzle_input_star2(filename string) [][]string {
	data := [][]string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	iterator := 0
	item := []string{}
	for scanner.Scan() {

		text := scanner.Text()
		text = strings.Trim(text, " ")
		item = append(item, text)
		iterator++

		if iterator > 2 {
			data = append(data, item)
			iterator = 0
			item = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func print_puzzle_input(data [][]string) {
	for index, element := range data {
		fmt.Println(index)
		for _, element2 := range element {
			fmt.Println(element2)
		}
	}
}
