package daysix_test

import (
	"testing"

	"github.com/hmajid2301/advent_of_code/2023/daysix"
	"github.com/stretchr/testify/assert"
)

func TestRaceProduct(t *testing.T) {
	t.Run("Should return product of races", func(t *testing.T) {
		results := `Time:      7  15   30
Distance:  9  40  200`

		value := daysix.CalculatePartNumSum(results)
		assert.Equal(t, value, 288)
	})
}
