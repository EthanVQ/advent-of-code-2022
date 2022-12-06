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

	star1 := solve_first_star(data)

	fmt.Println(star1)

	star2 := solve_second_star(data)

	fmt.Println(star2)
}

func solve_first_star(data string) int {
	for i := 0; i < len(data)-4; i++ {

		noDuplicates := true
		letters := strings.Split(data[i:i+4], "")

		for index, letter := range letters {

			copiedLetters := make([]string, 4)
			copy(copiedLetters, letters)

			// remove letter we are checking
			copiedLetters[index] = copiedLetters[len(copiedLetters)-1]
			copiedLetters = copiedLetters[:len(copiedLetters)-1]

			if strings.Contains(strings.Join(copiedLetters, ""), letter) {
				noDuplicates = false
			}
		}

		if noDuplicates {
			return i + 4
		}
	}

	return -1
}

func solve_second_star(data string) int {
	for i := 0; i < len(data)-14; i++ {

		noDuplicates := true
		letters := strings.Split(data[i:i+14], "")

		for index, letter := range letters {

			copiedLetters := make([]string, 14)
			copy(copiedLetters, letters)

			// remove letter we are checking
			copiedLetters[index] = copiedLetters[len(copiedLetters)-1]
			copiedLetters = copiedLetters[:len(copiedLetters)-1]

			if strings.Contains(strings.Join(copiedLetters, ""), letter) {
				noDuplicates = false
			}
		}

		if noDuplicates {
			return i + 14
		}
	}

	return -1
}

func read_puzzle_input(filename string) string {
	text := ""
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
