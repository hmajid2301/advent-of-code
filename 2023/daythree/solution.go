package daythree

import (
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Coord struct {
	x int
	y int
}

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

	var total int
	schematicsArray := multilineStringTo2DArray(schematics)

	for i, row := range schematicsArray {
		digits := []int{}
		var isSymbolAdjacent bool
		for j, col := range row {
			v, err := strconv.Atoi(col)
			if err == nil {
				digits = append(digits, v)
				isSymbolAdjacent = isSymbolAdjacentToCoord(Coord{x: i, y: j}, schematicsArray, row, symbols)
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

func isSymbolAdjacentToCoord(currentCoord Coord, schematicsArray [][]string, row []string, symbols map[string]struct{}) bool {
	i := currentCoord.x
	j := currentCoord.y

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
			return true
		}
	}
	return false
}

func CalculateGearRatio(schematics string) int {
	var total int
	schematicsArray := multilineStringTo2DArray(schematics)

	potentialGears := map[Coord][]int{}
	for i, row := range schematicsArray {
		digits := []int{}
		adjacentAsterisk := map[Coord]struct{}{}

		for j, col := range row {
			v, err := strconv.Atoi(col)
			if err == nil {
				digits = append(digits, v)
				newPossibleGears := getPotentialGearCoords(Coord{x: i, y: j}, schematicsArray, row)
				adjacentAsterisk = maps.Copy(adjacentAsterisk, newPossibleGears)
			}
			if err != nil || j+1 == len(row) {
				for asteriskCoord := range adjacentAsterisk {
					num := sliceToInt(digits)
					potentialGears[asteriskCoord] = append(potentialGears[asteriskCoord], num)

				}
				digits = []int{}
				adjacentAsterisk = map[Coord]struct{}{}
			}
		}
	}

	for _, potential := range potentialGears {
		if len(potential) == 2 {
			total += potential[0] * potential[1]
		}
	}

	return total
}

func getPotentialGearCoords(currentCoord Coord, schematicsArray [][]string, row []string) map[Coord]struct{} {
	i := currentCoord.x
	j := currentCoord.y
	adjacentSymbol := map[Coord]struct{}{}

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

	return adjacentSymbol
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
