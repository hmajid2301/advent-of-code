package dayseven

import (
	"strconv"
	"strings"
)

type HandType int

const (
	FiveKind HandType = iota
	FourKind
	FullHouse
	ThreeKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	Cards map[string]int
	Bid   int
	Hand  []int
	Type  HandType
}

func CalculateWinningTotal(camelCards string) int {
	var total int

	cardValueMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	hands := map[HandType][]Hand{}
	cardLines := strings.Split(camelCards, "\n")
	for _, cardBid := range cardLines {
		line := strings.Split(cardBid, " ")
		cards := line[0]
		bid := line[1]

		bidNum, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}

		hand := Hand{
			Bid: bidNum,
		}

		handValue := []int{}
		cardMap := map[string]int{}
		for _, card := range strings.Split(cards, "") {
			val, ok := cardValueMap[card]
			if !ok {
				panic("card not in map somehow")
			}
			handValue = append(handValue, val)
			cardMap[card]++
		}

		hand.Hand = handValue
		hand.Cards = cardMap
		if len(cardMap) == 1 {
			hand.Type = FiveKind
		} else if len(cardMap) == 2 {
			for k := range cardMap {
				val := cardMap[k]
				if val == 1 || val == 4 {
					hand.Type = FourKind
				} else {
					hand.Type = FullHouse
				}
				break
			}
		} else if len(cardMap) == 3 {
			for k := range cardMap {
				val := cardMap[k]
				hand.Type = TwoPair
				if val == 3 {
					hand.Type = ThreeKind
					break
				}
			}
			hand.Type = ThreeKind
		} else if len(cardMap) == 4 {
			hand.Type = OnePair
		} else {
			hand.Type = HighCard
		}

		hands[hand.Type] = append(hands[hand.Type], hand)
	}

	return total
}
