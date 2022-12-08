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

	star1 := solve_first_star(data)

	fmt.Println(star1)

	star2 := solve_second_star(data)

	fmt.Println(star2)
}

func solve_first_star(data [][]int) int {
	height := len(data)
	width := len(data[0])

	// calculate edges first
	visibleEdges := height*2 + width*2 - 4

	visibleInside := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			visible := check_if_tree_is_visbile(data, i, j)
			if visible {
				visibleInside++
			}
		}
	}

	return visibleEdges + visibleInside
}

func solve_second_star(data [][]int) int {
	height := len(data)
	width := len(data[0])

	maxViewingScore := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			score := calculate_viewing_score(data, i, j)
			if score > maxViewingScore {
				maxViewingScore = score
			}
		}
	}

	return maxViewingScore
}

func check_if_tree_is_visbile(data [][]int, i int, j int) bool {
	treeHeight := data[i][j]
	height := len(data)
	width := len(data[0])
	visibleAngles := 4

	// look left
	for left := j - 1; left >= 0; left-- {
		if treeHeight <= data[i][left] {
			visibleAngles--
			break
		}
	}

	// look right
	for right := j + 1; right < width; right++ {
		if treeHeight <= data[i][right] {
			visibleAngles--
			break
		}
	}

	// look up
	for up := i - 1; up >= 0; up-- {
		if treeHeight <= data[up][j] {
			visibleAngles--
			break
		}
	}

	// look down
	for down := i + 1; down < height; down++ {
		if treeHeight <= data[down][j] {
			visibleAngles--
			break
		}
	}

	if visibleAngles > 0 {
		return true
	} else {
		return false
	}
}

func calculate_viewing_score(data [][]int, i int, j int) int {
	treeHeight := data[i][j]
	height := len(data)
	width := len(data[0])

	// look left
	leftCounter := 0
	for left := j - 1; left >= 0; left-- {
		leftCounter++
		if treeHeight <= data[i][left] {
			break
		}
	}

	// look right
	rightCounter := 0
	for right := j + 1; right < width; right++ {
		rightCounter++
		if treeHeight <= data[i][right] {
			break
		}
	}

	// look up
	upCounter := 0
	for up := i - 1; up >= 0; up-- {
		upCounter++
		if treeHeight <= data[up][j] {
			break
		}
	}

	// look down
	downCounter := 0
	for down := i + 1; down < height; down++ {
		downCounter++
		if treeHeight <= data[down][j] {
			break
		}
	}

	return leftCounter * rightCounter * upCounter * downCounter
}

func read_puzzle_input(filename string) [][]int {
	data := [][]int{}
	text := ""
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = scanner.Text()
		stringHeights := strings.Split(text, "")
		intHeights := []int{}

		for _, string := range stringHeights {
			int, _ := strconv.Atoi(string)
			intHeights = append(intHeights, int)
		}
		data = append(data, intHeights)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print_matrix(data)
	return data
}

func print_matrix(data [][]int) {
	for _, row := range data {
		row_string := ""
		for _, col := range row {
			row_string = row_string + strconv.Itoa(col) + " "
		}
		fmt.Println(row_string)
	}
}
