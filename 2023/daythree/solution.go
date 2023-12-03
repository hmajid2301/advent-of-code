package daythree

import (
	"strconv"
	"strings"
)

func CalculatePartNumSum(schematics string) int {
	symbols := map[string]struct{}{
		"#":  {},
		"+":  {},
		"$":  {},
		"*":  {},
		"-":  {},
		"=":  {},
		"/":  {},
		"\\": {},
		"%":  {},
		"&":  {},
		"@":  {},
	}

	type Coord struct {
		x int
		y int
	}

	var total int
	schematicsArray := multilineStringTo2DArray(schematics)

	for i, row := range schematicsArray {
		digits := []int{}
		var isSymbolAdjacent bool
		for j, col := range row {
			v, err := strconv.Atoi(col)
			if err == nil {
				digits = append(digits, v)
				adjacentCoords := []Coord{
					{x: i + 1, y: j},
					{x: i - 1, y: j},
					{x: i, y: j - 1},
					{x: i, y: j + 1},
					{x: i + 1, y: j + 1},
					{x: i - 1, y: j + 1},
					{x: i + 1, y: j - 1},
					{x: i - 1, y: j - 1},
				}
				for _, coord := range adjacentCoords {
					if coord.x < 0 || coord.y < 0 {
						continue
					}
					if coord.y+1 > len(schematicsArray) || coord.x+1 > len(row) {
						continue
					}

					item := schematicsArray[coord.x][coord.y]
					if _, ok := symbols[item]; ok {
						isSymbolAdjacent = true
						break
					}
				}
			}
			if err != nil || j+1 == len(row) {
				if isSymbolAdjacent {
					total += sliceToInt(digits)
				}
				digits = []int{}
				isSymbolAdjacent = false
			}

		}
	}

	return total

}
func CalculateGearRatio(schematics string) int {
	type Coord struct {
		x int
		y int
	}

	var total int
	schematicsArray := multilineStringTo2DArray(schematics)

	a := map[Coord][]int{}
	for i, row := range schematicsArray {
		digits := []int{}
		adjacentSymbol := map[Coord]struct{}{}

		for j, col := range row {
			v, err := strconv.Atoi(col)
			if err == nil {
				digits = append(digits, v)
				adjacentCoords := []Coord{
					{x: i + 1, y: j},
					{x: i - 1, y: j},
					{x: i, y: j - 1},
					{x: i, y: j + 1},
					{x: i + 1, y: j + 1},
					{x: i - 1, y: j + 1},
					{x: i + 1, y: j - 1},
					{x: i - 1, y: j - 1},
				}
				for _, coord := range adjacentCoords {
					if coord.x < 0 || coord.y < 0 {
						continue
					}
					if coord.y+1 > len(schematicsArray) || coord.x+1 > len(row) {
						continue
					}

					item := schematicsArray[coord.x][coord.y]
					if item == "*" {
						adjacentSymbol[coord] = struct{}{}
					}
				}
			}
			if err != nil || j+1 == len(row) {
				for s := range adjacentSymbol {
					num := sliceToInt(digits)
					a[s] = append(a[s], num)

				}
				digits = []int{}
				adjacentSymbol = map[Coord]struct{}{}
			}
		}
	}

	for _, potential := range a {
		if len(potential) == 2 {
			total += potential[0] * potential[1]
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
