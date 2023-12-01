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

		num := fmt.Sprintf("%s%s", first, last)
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		fmt.Println("NUm", num)

		value += i
	}
	return value
}
