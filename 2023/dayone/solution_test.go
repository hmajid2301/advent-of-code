package dayone_test

import (
	"testing"

	"github.com/hmajid2301/advent_of_code/2023/dayone"
	"github.com/stretchr/testify/assert"
)

func TestCalculateCaibrations(t *testing.T) {
	t.Run("Should return calibration values", func(t *testing.T) {
		calibrations := `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`

		value := dayone.CalculateCalibrations(calibrations)
		assert.Equal(t, value, 142)
	})
}
