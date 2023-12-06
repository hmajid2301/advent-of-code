package daysix

import (
	"strconv"
	"strings"
)

func CalculateRaceProduct(races string) int {
	var total int

	timeLine := strings.Split(races, "\n")
	times := strings.Split(timeLine[0], "Time: ")
	distances := strings.Fields(strings.Trim(strings.Split(timeLine[1], "Distance: ")[1], " "))

	for i, time := range strings.Fields(times[1]) {

		timeNum, err := strconv.Atoi(strings.Trim(time, " "))
		if err != nil {
			panic(err)
		}
		minDistance, err := strconv.Atoi(strings.Trim(distances[i], " "))
		if err != nil {
			panic(err)
		}

		for j := 1; j < timeNum-1; j++ {
			distance := (timeNum - j) * j

			if distance > minDistance {
				if total == 0 {
					total = timeNum - 2*j + 1
				} else {
					total *= (timeNum - 2*j + 1)
				}

				break
			}
		}
	}

	return total
}

func CalculateOneRaceProduct(races string) int {
	var total int

	timeLine := strings.Split(races, "\n")
	times := strings.Split(timeLine[0], "Time: ")
	distances := strings.Split(timeLine[1], "Distance: ")

	timeNum, err := strconv.Atoi(strings.ReplaceAll(times[1], " ", ""))
	if err != nil {
		panic(err)
	}
	minDistance, err := strconv.Atoi(strings.ReplaceAll(distances[1], " ", ""))
	if err != nil {
		panic(err)
	}

	for j := 1; j < timeNum-1; j++ {
		distance := (timeNum - j) * j
		if distance > minDistance {
			total = timeNum - 2*j + 1
			break
		}
	}

	return total
}
