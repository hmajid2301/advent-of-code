package dayone

import (
	"strconv"
	"strings"
)

func CalculateCalibrations(calibrations string) int {
	lines := strings.Split(calibrations, "\n")

	var total int
	for _, line := range lines {
		chars := strings.Split(line, "")

		digits := []int{}
		for _, char := range chars {
			if v, err := strconv.Atoi(char); err == nil {
				digits = append(digits, v)
			}
		}

		total += digits[0]*10 + digits[len(digits)-1]
	}
	return total
}

func CalculateCalibrationsWithStrings(calibrations string) int {
	type LookUp struct {
		stringValue  string
		integerValue int
	}

	numbers := map[string][]LookUp{
		"o": {{integerValue: 1, stringValue: "one"}},
		"t": {{integerValue: 2, stringValue: "two"}, {integerValue: 3, stringValue: "three"}},
		"f": {{integerValue: 4, stringValue: "four"}, {integerValue: 5, stringValue: "five"}},
		"s": {{integerValue: 6, stringValue: "six"}, {integerValue: 7, stringValue: "seven"}},
		"e": {{integerValue: 8, stringValue: "eight"}},
		"n": {{integerValue: 9, stringValue: "nine"}},
	}

	var total int
	lines := strings.Split(calibrations, "\n")
	for _, line := range lines {
		i := 0
		digits := []int{}

		// TODO: simplify this mess
		for i < len(line) {
			char := string(line[i])
			if lookups, ok := numbers[char]; ok {
				for _, lookup := range lookups {
					ok := strings.HasPrefix(line[i:], lookup.stringValue)
					if ok {
						digits = append(digits, lookup.integerValue)
						i += len(lookup.stringValue) - 2
						break
					}
				}
			} else if v, err := strconv.Atoi(char); err == nil {
				digits = append(digits, v)
			}
			i++
		}

		total += digits[0]*10 + digits[len(digits)-1]
	}
	return total
}
