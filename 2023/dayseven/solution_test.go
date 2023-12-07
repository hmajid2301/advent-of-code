package dayseven_test

import (
	"testing"

	"github.com/hmajid2301/advent_of_code/2023/dayseven"
	"github.com/stretchr/testify/assert"
)

func TestTotalWinnings(t *testing.T) {
	t.Run("Should return product of card winnings", func(t *testing.T) {
		results := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

		value := dayseven.CalculateRaceProduct(results)
		assert.Equal(t, value, 6440)
	})
}
