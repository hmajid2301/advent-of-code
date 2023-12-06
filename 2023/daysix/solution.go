package daysix

import (
	"math"
	"strconv"
	"strings"
)

func CalculateRaceSum(races string) int {
	var total int

	timeLine := strings.Split(races, "\n")
	times := strings.Split(timeLine[0], "Time: ")
	distances := strings.Split(strings.Split(timeLine[1], "Distance: ")[1], " ")

	for i, time := range strings.Split(times[1], "") {
		timeNum, err := strconv.Atoi(time)
		if err != nil {
			panic(err)
		}
		minDistance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}

		var max int
		var min int
		rejectedOptions := map[int]struct{}{
			0:       {},
			timeNum: {},
		}

		timeHoldingButton := int(math.Floor((float64(timeNum) + 1) / 2))
		for len(rejectedOptions) != timeNum+1 {
			distanceCovered := (timeNum - timeHoldingButton) * timeHoldingButton

			if min > max {
				break
			}

			if distanceCovered > minDistance {
				if max == 0 || max < timeHoldingButton {
					max = timeHoldingButton

				}
				if min == 0 || min > timeHoldingButton {
					min = timeHoldingButton
				}

				_, ok := rejectedOptions[min-1]
				_, ok1 := rejectedOptions[max+1]
				if ok && ok1 {
					break
				}
			} else {
				rejectedOptions[timeHoldingButton] = struct{}{}
			}

			timeHoldingButton = int(math.Floor((float64(timeNum) + float64(timeHoldingButton)) / 2))
		}

		total *= max - min + 1
	}

	return total
}
