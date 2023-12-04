package dayfour

import (
	"strconv"
	"strings"
)

// It would be more efficient to go through the string just once
// But probably easier to follow if we split the string
func CalculateCardWorth(cards string) int {
	var total int

	lines := strings.Split(cards, "\n")
	for _, line := range lines {
		var score int
		card := strings.Split(line, "|")
		winningNumbersList := card[0]
		nums := strings.Split(winningNumbersList, ":")[1]

		winningsNumbers := map[int]struct{}{}
		for _, num := range strings.Split(nums, " ") {
			if v, err := strconv.Atoi(num); err == nil {
				winningsNumbers[v] = struct{}{}
			}
		}

		for _, num := range strings.Split(card[1], " ") {
			if v, err := strconv.Atoi(num); err == nil {
				if _, ok := winningsNumbers[v]; ok {
					if score != 0 {
						score *= 2
					} else {
						score = 1
					}
				}
			}
		}
		total += score
	}

	return total
}

// TODO: refactor common code
func CalculateTotalScratchCards(cards string) int {

	lines := strings.Split(cards, "\n")
	cardScoreMap := map[int]int{}
	for gameNum, line := range lines {
		var score int
		card := strings.Split(line, "|")
		winningNumbersList := card[0]
		nums := strings.Split(winningNumbersList, ":")[1]

		winningsNumbers := map[int]struct{}{}
		for _, num := range strings.Split(nums, " ") {
			if v, err := strconv.Atoi(num); err == nil {
				winningsNumbers[v] = struct{}{}
			}
		}

		for _, num := range strings.Split(card[1], " ") {
			if v, err := strconv.Atoi(num); err == nil {
				if _, ok := winningsNumbers[v]; ok {
					score += 1
				}
			}
		}

		cardScoreMap[gameNum] = score
	}

	var total int

	games := map[int]int{}
	for game := range cardScoreMap {
		games[game] = 1
	}

	for j := 0; j < len(games); j++ {
		score := cardScoreMap[j]
		for x := 0; x < games[j]; x++ {
			for i := 0; i < score; i++ {
				games[j+i+1] += 1

			}
		}
	}

	for _, value := range games {
		total += value
	}

	return total
}
