package dayone

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateCalibrations(calibrations string) int {
	lines := strings.Split(calibrations, "\n")

	var value int
	for _, line := range lines {
		chars := strings.Split(line, "")

		var first string
		var last string
		for _, char := range chars {
			if _, err := strconv.Atoi(char); err == nil {
				if first == "" {
					first = char
				}
				last = char
			}
		}

		parsedValue := fmt.Sprintf("%s%s", first, last)

		num, err := strconv.Atoi(parsedValue)
		if err != nil {
			panic(err)
		}

		value += num
	}
	return value
}

func CalculateCalibrationsWithStrings(calibrations string) int {
	lines := strings.Split(calibrations, "\n")

	type LookUp struct {
		integerValue string
		stringValue  string
	}

	numbers := map[string][]LookUp{
		"o": {{integerValue: "1", stringValue: "one"}},
		"t": {{integerValue: "2", stringValue: "two"}, {integerValue: "3", stringValue: "three"}},
		"f": {{integerValue: "4", stringValue: "four"}, {integerValue: "5", stringValue: "five"}},
		"s": {{integerValue: "6", stringValue: "six"}, {integerValue: "7", stringValue: "seven"}},
		"e": {{integerValue: "8", stringValue: "eight"}},
		"n": {{integerValue: "9", stringValue: "nine"}},
	}

	var value int
	for _, line := range lines {

		var first string
		var last string

		i := 0
		for i < len(line) {

			char := string(line[i])
			if lookups, ok := numbers[char]; ok {
				for _, lookup := range lookups {
					ok := strings.HasPrefix(line[i:], lookup.stringValue)
					if ok {
						if first == "" {
							first = lookup.integerValue
						}
						last = lookup.integerValue
						i += len(lookup.stringValue) - 2
						break
					}
				}
			} else if _, err := strconv.Atoi(char); err == nil {
				if first == "" {
					first = char
				}
				last = char
			}
			i++
		}

		parsedValue := fmt.Sprintf("%s%s", first, last)

		num, err := strconv.Atoi(parsedValue)
		if err != nil {
			panic(err)
		}

		value += num
	}
	return value
}
