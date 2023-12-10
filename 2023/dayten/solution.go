package dayten

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func CalculateMaxSteps(input string) int {
	pipesTypes := map[string][]Coordinate{
		"|": {{0, -1}, {0, 1}},
		"J": {{0, -1}, {-1, 0}},
		"L": {{0, -1}, {1, 0}},
		"7": {{0, 1}, {-1, 0}},
		"F": {{0, 1}, {1, 0}},
		"-": {{-1, 0}, {1, 0}},
		".": {},
	}

	start := Coordinate{}
	var maze = make([][][]Coordinate, len(strings.Split(input, "\n")))
	for i := range maze {
		maze[i] = make([][]Coordinate, len(strings.Split(input, "\n")[0]))
	}

	var s = make([][]string, len(strings.Split(input, "\n")))
	for i := range maze {
		s[i] = make([]string, len(strings.Split(input, "\n")[0]))
	}
	for j, line := range strings.Split(input, "\n") {
		for i, char := range strings.Split(line, "") {
			maze[i][j] = pipesTypes[char]
			s[i][j] = char

			if char == "S" {
				start = Coordinate{i, j}
			}
		}
	}

	var loop int
	for _, direction := range []Coordinate{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		loop = findLoop(s, start, maze, direction)
		if loop > 0 {
			break
		}

	}
	return loop / 2
}

func findLoop(symbols [][]string, start Coordinate, maze [][][]Coordinate, direction Coordinate) int {
	loop := []Coordinate{start}
	location := Coordinate{start.x + direction.x, start.y + direction.y}

	fmt.Println()
	var symbol string
	for location != start {
		// TODO: this will break if you move to a symbol you shouldnt, as the first move
		loop = append(loop, location)
		if location.x < 0 || location.x >= len(maze) || location.y < 0 || location.y >= len(maze[0]) {
			return 0
		}

		pipe := maze[location.x][location.y]
		symbol = symbols[location.x][location.y]
		fmt.Print(symbol)
		fmt.Print(" ")
		fmt.Print(location)
		fmt.Print(" ")
		if len(pipe) == 0 {
			return 0
		}

		opposite := Coordinate{-direction.x, -direction.y}
		for _, dir := range pipe {
			if dir.x == opposite.x && dir.y == opposite.y {
				continue
			}
			direction = dir
			location = Coordinate{
				x: location.x + dir.x,
				y: location.y + dir.y,
			}
			break
		}
	}
	fmt.Println(loop)
	fmt.Println()
	return len(loop)
}
