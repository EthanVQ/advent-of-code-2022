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

	boxes, commands := read_puzzle_input(filename)

	star1 := solve_first_star(boxes, commands)

	fmt.Println(star1)

	boxes2, commands2 := read_puzzle_input(filename)

	star2 := solve_second_star(boxes2, commands2)

	fmt.Println(star2)
}

func solve_first_star(boxes map[int][]string, commands [][]int) string {
	for _, command := range commands {
		count := 0
		for _, item := range boxes[command[1]] {
			if count >= command[0] {
				break
			} else {
				boxes[command[2]] = append([]string{item}, boxes[command[2]]...)
				boxes[command[1]] = boxes[command[1]][1:]
				count++
			}
		}
	}

	topOfStacks := ""
	for i := 1; i <= len(boxes); i++ {
		topOfStacks = topOfStacks + boxes[i][0]
	}

	return topOfStacks
}

func solve_second_star(boxes map[int][]string, commands [][]int) string {
	for _, command := range commands {
		copyiedList := make([]string, len(boxes[command[1]][:command[0]]))
		copy(copyiedList, boxes[command[1]][:command[0]])
		boxes[command[2]] = append(copyiedList, boxes[command[2]]...)
		boxes[command[1]] = boxes[command[1]][command[0]:]
	}

	topOfStacks := ""
	for i := 1; i <= len(boxes); i++ {
		topOfStacks = topOfStacks + boxes[i][0]
	}

	return topOfStacks
}

func read_puzzle_input(filename string) (map[int][]string, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	boxes := []string{}
	commands := []string{}
	isCommands := false
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			isCommands = true
			boxes = boxes[:len(boxes)-1]
			continue
		}

		if !isCommands {
			boxes = append(boxes, text)
		} else {
			commands = append(commands, text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	boxesData := read_boxes_data(boxes)
	commandsData := read_commands_data(commands)

	return boxesData, commandsData
}

func read_boxes_data(boxes []string) map[int][]string {
	data := map[int][]string{}
	for _, line := range boxes {
		for i := 1; i < len(line); i += 4 {

			stack := ((i - 1) / 4) + 1
			character := line[i : i+1]

			if character == " " {
				continue
			}

			if _, ok := data[stack]; ok {
				data[stack] = append(data[stack], character)
			} else {
				data[stack] = []string{character}
			}
		}
	}
	return data
}

func read_commands_data(commands []string) [][]int {
	data := [][]int{}
	for _, line := range commands {
		moveInd := strings.Index(line, "move")
		fromInd := strings.Index(line, "from")
		toInd := strings.Index(line, "to")

		moveNum, _ := strconv.Atoi(line[moveInd+5 : fromInd-1])
		fromNum, _ := strconv.Atoi(line[fromInd+5 : toInd-1])
		toNum, _ := strconv.Atoi(line[toInd+3:])

		command := []int{moveNum, fromNum, toNum}
		data = append(data, command)
	}
	return data
}

func print_boxes(boxes map[int][]string) {
	for i := 1; i <= len(boxes); i++ {
		fmt.Println("stack: " + strconv.Itoa(i))
		for _, item := range boxes[i] {
			fmt.Println(item)
		}
	}
}
