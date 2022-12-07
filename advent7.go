package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type File struct {
	name     string
	_type    string
	size     int
	children []*File
	parent   *File
}

func bubbleFileSizeUp(child File) {
	size := child.size
	current := child.parent

	for current != nil {
		current.size += size
		current = current.parent
	}
}

func dfsForOne(current File) int {
	result := 0

	if current.size <= 100000 && current._type == "dir" {
		result += current.size
	}

	for _, child := range current.children {
		result += dfsForOne(*child)
	}

	return result
}

func calculateTotalSizesOfDirsLowerThan100k(root File) int {
	current := root
	return dfsForOne(current)
}

func dfsForTwo(current File, totalSizeNeeded int, result *File) File {
	if current.size >= totalSizeNeeded && current._type == "dir" && current.size < result.size {
		result = &current
	}

	for _, child := range current.children {
		childResult := dfsForTwo(*child, totalSizeNeeded, result)

		if childResult.size < result.size {
			result = &childResult
		}
	}

	return *result
}

func calculateLowestPossibleDeletionForUpdate(root File, totalSizeNeeded int) {
	current := root
	result := dfsForTwo(current, totalSizeNeeded, &root)
	fmt.Println("Second Result", result)
}

func one(input string) {
	root := File{
		name:   "/",
		_type:  "dir",
		size:   0,
		parent: nil,
	}

	current := &root

	for _, command := range strings.Split(input, "\n$ ") {
		parts := strings.Split(command, "\n")

		if parts[0] == "ls" {
			for i := 1; i < len(parts); i++ {
				data := strings.Split(parts[i], " ")

				if data[0] == "dir" {
					dir := File{
						name:   data[1],
						_type:  "dir",
						size:   0, // initialized at 0
						parent: current,
					}

					current.children = append(current.children, &dir)
				} else {
					size, _ := strconv.Atoi(data[0])
					file := File{
						name:   data[1],
						_type:  "file",
						size:   size,
						parent: current,
					}

					current.children = append(current.children, &file)
					bubbleFileSizeUp(file)
				}

			}
		} else {
			cdparts := strings.Split(parts[0], " ")

			if cdparts[1] == ".." {
				current = current.parent
			} else {
				for _, child := range current.children {
					if child.name == cdparts[1] {
						current = child
						break
					}
				}
			}
		}
	}

	resultOne := calculateTotalSizesOfDirsLowerThan100k(root)
	fmt.Println("First result", resultOne)

	totalSizeNeeded := 30000000 - (70000000 - root.size)
	fmt.Println("totalSizeNeeded", totalSizeNeeded)

	calculateLowestPossibleDeletionForUpdate(root, totalSizeNeeded)
}

func main() {
	content, err := ioutil.ReadFile("7.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
}
