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

		value := daysix.CalculateRaceProduct(results)
		assert.Equal(t, value, 288)
	})

	t.Run("Should return product of races real input", func(t *testing.T) {
		results := `Time:        55     82     64     90
Distance:   246   1441   1012   1111`

		value := daysix.CalculateRaceProduct(results)
		assert.Equal(t, value, 608902)
	})

	t.Run("Should return product of one race", func(t *testing.T) {
		results := `Time:      7  15   30
Distance:  9  40  200`

		value := daysix.CalculateOneRaceProduct(results)
		assert.Equal(t, value, 71503)
	})

	t.Run("Should return product of races real input", func(t *testing.T) {
		results := `Time:        55     82     64     90
	Distance:   246   1441   1012   1111`

		value := daysix.CalculateOneRaceProduct(results)
		assert.Equal(t, value, 46173809)
	})

}
