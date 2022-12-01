package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func read_puzzle_input(filename string) [][]int {
	data := [][]int{}
	item := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		text := scanner.Text()

		if text == "" {

			data = append(data, item)

			item = []int{}
		} else {

			parsedInt, _ := strconv.Atoi(text)

			item = append(item, parsedInt)
		}

	}

	if len(item) > 0 {
		data = append(data, item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func solve_first_star(data [][]int) int {
	// calculate total calories for each elf
	var caloriesPerElf = make([]int, len(data))
	for index, elf := range data {
		totalCalories := 0
		for _, foodItem := range elf {
			totalCalories += foodItem
		}
		caloriesPerElf[index] = totalCalories
	}

	// find max total calories
	elfAndMaxCalories := []int{-1, 0}
	for index, elfCalories := range caloriesPerElf {
		if elfCalories > elfAndMaxCalories[1] {
			elfAndMaxCalories = []int{index, elfCalories}
		}
	}

	return elfAndMaxCalories[1]
}

func solve_second_star(data [][]int) int {
	// calculate total calories for each elf
	var caloriesPerElf = make([]int, len(data))
	for index, elf := range data {
		totalCalories := 0
		for _, foodItem := range elf {
			totalCalories += foodItem
		}
		caloriesPerElf[index] = totalCalories
	}

	// find top 3 elves
	star2 := 0
	for i := 1; i < 4; i++ {
		elfAndMaxCalories := get_max_calories_and_elf_from_list(caloriesPerElf)
		caloriesPerElf = remove_elf(caloriesPerElf, elfAndMaxCalories[0])
		star2 += elfAndMaxCalories[1]
	}

	return star2
}

func remove_elf(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func get_max_calories_and_elf_from_list(caloriesPerElf []int) []int {
	// find max total calories
	elfAndMaxCalories := []int{-1, 0}
	for index, elfCalories := range caloriesPerElf {
		if elfCalories > elfAndMaxCalories[1] {
			elfAndMaxCalories = []int{index, elfCalories}
		}
	}

	return elfAndMaxCalories
}

func print_puzzle_input(data [][]int) {
	for index, element := range data {
		fmt.Printf("Item %s\n", index)
		for _, element2 := range element {
			fmt.Println(element2)
		}
	}
}
