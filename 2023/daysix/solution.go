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

		maxNum := getMax(timeNum, minDistance)
		minNum := getMin(timeNum, minDistance)
		if total == 0 {
			total = maxNum - minNum + 1
		} else {
			total *= maxNum - minNum + 1
		}
	}

	return total
}

func getMin(timeNum, minDistance int) int {
	var minNum int
	rejectedOptions := map[int]struct{}{
		timeNum: {},
	}

	timeHoldingButton := int(math.Floor((float64(timeNum) + 1) / 2))
	for len(rejectedOptions) < (timeNum+1)/2 {
		distanceCovered := (timeNum - timeHoldingButton) * timeHoldingButton

		_, ok1 := rejectedOptions[timeHoldingButton]
		if ok1 {
			break
		}

		if distanceCovered > minDistance {
			if minNum == 0 || minNum > timeHoldingButton {
				minNum = timeHoldingButton
			}
			_, ok1 := rejectedOptions[minNum-1]
			if ok1 {
				break
			}

			timeHoldingButton = int(math.Floor((float64(0) + float64(timeHoldingButton)) / 2))
		} else {
			rejectedOptions[timeHoldingButton] = struct{}{}
			timeHoldingButton = int(math.Floor((float64(minNum) + float64(timeHoldingButton)) / 2))
		}
	}
	return minNum
}

func getMax(timeNum, minDistance int) int {
	var maxNum int
	rejectedOptions := map[int]struct{}{
		timeNum: {},
	}

	timeHoldingButton := int(math.Floor((float64(timeNum) + 1) / 2))
	for len(rejectedOptions) < (timeNum+1)/2 {
		distanceCovered := (timeNum - timeHoldingButton) * timeHoldingButton

		_, ok1 := rejectedOptions[timeHoldingButton]
		if ok1 {
			break
		}

		if distanceCovered > minDistance {
			if maxNum == 0 || maxNum < timeHoldingButton {
				maxNum = timeHoldingButton
			}
			_, ok1 := rejectedOptions[maxNum+1]
			if ok1 {
				break
			}

			timeHoldingButton = int(math.Floor((float64(timeNum) + float64(timeHoldingButton)) / 2))
		} else {
			rejectedOptions[timeHoldingButton] = struct{}{}
			timeHoldingButton = int(math.Floor((float64(maxNum) + float64(timeHoldingButton)) / 2))
		}
	}
	return maxNum
}
