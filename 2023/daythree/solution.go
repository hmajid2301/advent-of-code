package daythree

import (
	"strconv"
	"strings"
)

func CalculatePartNumSum(schematics string) int {
	symbols := map[string]struct{}{
		"#":  struct{}{},
		"+":  struct{}{},
		"$":  struct{}{},
		"*":  struct{}{},
		"-":  struct{}{},
		"=":  struct{}{},
		"/":  struct{}{},
		"\\": struct{}{},
		"%":  struct{}{},
	}

	type Coord struct {
		x int
		y int
	}

	var total int
	schematicsArray := multilineStringTo2DArray(schematics)
	for i, row := range schematicsArray {
		for j, col := range row {
			digits := []int{}
			if v, err := strconv.Atoi(col); err == nil {
				digits = append(digits, v)
			} else {
				digits = []int{}
			}

			var isSymbolAdjacent bool
			adjacentCoords := []Coord{
				Coord{x: i + 1, y: j},
				Coord{x: i - 1, y: j},
				Coord{x: i, y: j - 1},
				Coord{x: i, y: j + 1},
				Coord{x: i + 1, y: j + 1},
				Coord{x: i - 1, y: j + 1},
				Coord{x: i + 1, y: j - 1},
				Coord{x: i - 1, y: j - 1},
			}
			for _, coord := range adjacentCoords {
				if coord.x < 0 {
					continue
				}
				if coord.y > len(row) {
					continue
				}
				item := schematicsArray[coord.x][coord.y]
			}

		}
	}

	return total

}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func multilineStringTo2DArray(input string) [][]string {
	lines := strings.Split(input, "\n")

	var array [][]string

	for _, line := range lines {
		chars := strings.Split(line, "")
		array = append(array, chars)
	}

	return array
}
