package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func solve_first_star(data []Command) int {
	// initialise grid
	Hlocation := []float64{0, 0}
	Tlocation := []float64{0, 0}
	TlocationsVisited := [][]float64{[]float64{0, 0}}

	for _, command := range data {
		for k := 0; k < int(command.Distance); k++ {

			// move H first
			if command.Direction == "R" {
				Hlocation[0]++
			}
			if command.Direction == "L" {
				Hlocation[0]--
			}
			if command.Direction == "U" {
				Hlocation[1]++
			}
			if command.Direction == "D" {
				Hlocation[1]--
			}

			// move T base on H location

			xDif := Hlocation[0] - Tlocation[0]
			yDif := Hlocation[1] - Tlocation[1]

			// stay still scenario
			if math.Abs(xDif) <= 1 && math.Abs(yDif) <= 1 {
				continue
			}

			// move scenario
			if xDif > 0 {
				Tlocation[0]++
			}

			if xDif < 0 {
				Tlocation[0]--
			}

			if yDif > 0 {
				Tlocation[1]++
			}

			if yDif < 0 {
				Tlocation[1]--
			}

			// store T location if its new
			if location_doesnt_already_exist(TlocationsVisited, Tlocation) {
				copiedTlocation := make([]float64, len(Tlocation))
				copy(copiedTlocation, Tlocation)
				TlocationsVisited = append(TlocationsVisited, copiedTlocation)
			}
		}
	}

	return len(TlocationsVisited)
}

func solve_second_star(data []Command) int {
	// initialise grid
	locations := [][]float64{}
	for i := 0; i < 10; i++ {
		locations = append(locations, []float64{0, 0})
	}
	TlocationsVisited := [][]float64{[]float64{0, 0}}

	for _, command := range data {
		for k := 0; k < int(command.Distance); k++ {

			// move H first
			if command.Direction == "R" {
				locations[0][0]++
			}
			if command.Direction == "L" {
				locations[0][0]--
			}
			if command.Direction == "U" {
				locations[0][1]++
			}
			if command.Direction == "D" {
				locations[0][1]--
			}

			// iterate through and update tailing locations
			for i := 1; i < len(locations); i++ {
				locations[i] = return_tailing_location(locations[i-1], locations[i])
			}

			// store T location if its new
			Tlocation := locations[len(locations)-1]
			if location_doesnt_already_exist(TlocationsVisited, Tlocation) {
				copiedTlocation := make([]float64, len(Tlocation))
				copy(copiedTlocation, Tlocation)
				TlocationsVisited = append(TlocationsVisited, copiedTlocation)
			}
		}
	}

	return len(TlocationsVisited)
}

func return_tailing_location(destination []float64, location []float64) []float64 {

	xDif := destination[0] - location[0]
	yDif := destination[1] - location[1]

	// stay still scenario
	if math.Abs(xDif) <= 1 && math.Abs(yDif) <= 1 {
		return location
	}

	// move scenario
	if xDif > 0 {
		location[0]++
	}

	if xDif < 0 {
		location[0]--
	}

	if yDif > 0 {
		location[1]++
	}

	if yDif < 0 {
		location[1]--
	}

	return location
}

func location_doesnt_already_exist(data [][]float64, location []float64) bool {
	for _, existLocation := range data {
		if existLocation[0] == location[0] && existLocation[1] == location[1] {
			return false
		}
	}
	return true
}

func read_puzzle_input(filename string) []Command {
	data := []Command{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		items := strings.Split(text, " ")
		distance, _ := strconv.ParseFloat(items[1], 64)
		command := Command{Direction: items[0], Distance: distance}
		data = append(data, command)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print_matrix(data)
	return data
}

type Command struct {
	Direction string
	Distance  float64
}
