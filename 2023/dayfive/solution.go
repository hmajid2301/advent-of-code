package dayfive

import (
	"fmt"
	"strconv"
	"strings"
)

type CategoryMap struct {
	source      int
	destination int
	length      int
}

func CalculateMinimumLocation(seedsInfo string) int {
	lines := strings.Split(seedsInfo, "\n")

	seedToSoil, lineNum := getMap(lines, 3)
	soilToFert, lineNum := getMap(lines, lineNum+1)
	fertToWater, lineNum := getMap(lines, lineNum+1)
	waterToLight, lineNum := getMap(lines, lineNum+1)
	lightToTemp, lineNum := getMap(lines, lineNum+1)
	tempToHumidity, lineNum := getMap(lines, lineNum+1)
	humidityToLoc, _ := getMap(lines, lineNum+1)

	locationMaps := [][]CategoryMap{
		seedToSoil,
		soilToFert,
		fertToWater,
		waterToLight,
		lightToTemp,
		tempToHumidity,
		humidityToLoc,
	}

	seeds := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
	minLoc := -1
	for _, seed := range seeds {
		v, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}

		val := v
		for _, myMap := range locationMaps {
			val = getValFromMap(val, myMap)
		}

		if minLoc == -1 || minLoc > val {
			minLoc = val
		}
		fmt.Println()
	}
	return minLoc
}

func getValFromMap(source int, category []CategoryMap) int {
	for _, categoryMapping := range category {
		min := categoryMapping.source
		max := categoryMapping.source + categoryMapping.length
		if source >= min && source <= max {
			return categoryMapping.destination + source - min
		}
	}
	return source
}

func getMap(lines []string, startLineNum int) ([]CategoryMap, int) {
	myMakp := []CategoryMap{}
	lineNum := startLineNum
	for _, line := range lines[startLineNum:] {
		lineNum++
		if line == "" {
			break
		}
		nums := strings.Split(line, " ")
		digits := []int{}
		for _, num := range nums {
			if v, err := strconv.Atoi(num); err == nil {
				digits = append(digits, v)
			}
		}
		myMakp = append(myMakp, CategoryMap{source: digits[1], destination: digits[0], length: digits[2]})
	}

	return myMakp, lineNum
}
