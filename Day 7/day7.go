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

func solve_first_star(data []string) int {
	fs := &Node{Filesize: 0, Folder: true}
	fs.Add("/", 0, true)

	recurse(0, data, 0, fs, []string{"/"})

	// print_file_system(fs, 0)

	scores := map[string]int{}

	_ = calculate_scores(fs, 0, []string{"/"}, scores)

	// print_file_system(fs, 0)

	star1 := find_scores_under_x(100000, scores)

	return star1
}

func solve_second_star(data []string) int {
	fs := &Node{Filesize: 0, Folder: true}
	fs.Add("/", 0, true)

	recurse(0, data, 0, fs, []string{"/"})

	// print_file_system(fs, 0)

	scores := map[string]int{}

	_ = calculate_scores(fs, 0, []string{"/"}, scores)

	min_deleted_space := 30000000 - (70000000 - scores["/"])

	star2 := find_closest_file(min_deleted_space, scores)

	return star2
}

func recurse(level int, data []string, dataIndex int, fs *Node, location []string) {

	if dataIndex >= len(data) {
		return
	}

	command := data[dataIndex]
	splitCommands := strings.Split(command, " ")
	copiedLocation := make([]string, len(location))
	copy(copiedLocation, location)

	if splitCommands[0] == "$" {
		// cd
		if splitCommands[1] == "cd" {
			if splitCommands[2] == ".." {
				recurse(level-1, data, dataIndex+1, fs, copiedLocation[:len(copiedLocation)-1])
			} else if splitCommands[2] == "/" {
				recurse(0, data, dataIndex+1, fs, []string{"/"})
			} else {
				fs.Get(location).Add(splitCommands[2], 0, true)
				recurse(level+1, data, dataIndex+1, fs, append(copiedLocation, splitCommands[2]))
			}

		}

		// ls
		if splitCommands[1] == "ls" {

			iterator := dataIndex + 1
			for true {
				if iterator >= len(data) {
					return
				}

				item := data[iterator]

				if strings.Contains(item, "$") {
					break
				}

				pair := strings.Split(item, " ")
				if pair[0] == "dir" {
					fs.Get(location).Add(pair[1], 0, true)
				} else {
					val, _ := strconv.ParseInt(pair[0], 10, 64)
					fs.Get(location).Add(pair[1], val, false)
				}

				iterator++
			}
			recurse(level, data, iterator, fs, copiedLocation)
		}
	} else {
		fmt.Println("Something went wrong, command not expected. Got: " + command)
	}
	return
}

func print_file_system(n *Node, level int) {
	if len(n.Subdir) <= 0 {
		return
	}
	for key, a := range n.Subdir {
		fileSizeString := ""
		if a.Filesize > 0 {
			fileSizeString = " , " + strconv.FormatInt(a.Filesize, 10)
		}
		fmt.Println(strings.Repeat(" ", level*2) + "- " + key + fileSizeString)
		print_file_system(n.Get([]string{key}), level+1)
	}
}

func calculate_scores(fs *Node, level int, location []string, scores map[string]int) int64 {

	var totalFileSize int64
	totalFileSize = 0

	if fs.Get(location).Folder {
		for key, _ := range fs.Get(location).Subdir {
			totalFileSize += calculate_scores(fs, level+1, append(location, key), scores)
		}
		scores[strings.Join(location, "")] = int(totalFileSize)
	} else {
		totalFileSize = fs.Get(location).Filesize
	}

	fs.Set(totalFileSize, location)
	return totalFileSize
}

func find_scores_under_x(limit int, scores map[string]int) int {
	sum := 0
	for _, score := range scores {
		if score <= 100000 {
			sum += score
		}
	}
	return sum
}

func find_closest_file(limit int, scores map[string]int) int {
	closest_score := 70000000
	for _, score := range scores {
		dif := score - limit
		if dif >= 0 {
			if score < closest_score {
				closest_score = score
			}
		}
	}
	return closest_score
}

func read_puzzle_input(filename string) []string {
	data := []string{}
	text := ""
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = scanner.Text()
		data = append(data, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

type Node struct {
	Subdir   map[string]*Node
	Filesize int64
	Folder   bool
}

func (n *Node) Add(key string, v int64, f bool) {
	if n.Subdir == nil {
		n.Subdir = map[string]*Node{}
	}
	n.Subdir[key] = &Node{Filesize: v, Folder: f}
}

func (n *Node) Get(keys []string) *Node {
	for _, key := range keys {
		n = n.Subdir[key]
	}
	return n
}

func (n *Node) Set(v int64, keys []string) {
	n = n.Get(keys)
	n.Filesize = v
}
