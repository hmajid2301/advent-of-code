package daythree_test

import (
	"testing"

	"github.com/hmajid2301/advent_of_code/2023/daytwo"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePossibleGames(t *testing.T) {
	t.Run("Should return sum of part numbers", func(t *testing.T) {
		results := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		value := daytwo.CalculatePossibleGames(results)
		assert.Equal(t, value, 4361)
	})
}
