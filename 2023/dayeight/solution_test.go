package dayeight_test

import (
	"testing"

	"github.com/hmajid2301/advent_of_code/2023/dayeight"
	"gotest.tools/assert"
)

func TestCalculateStartToEnd(t *testing.T) {
	t.Run("Should return steps to z", func(t *testing.T) {
		treeStr := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

		value := dayeight.CalculateStartToEnd(treeStr)
		assert.Equal(t, value, 2)
	})

	t.Run("Should return steps to z again", func(t *testing.T) {
		treeStr := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

		value := dayeight.CalculateStartToEnd(treeStr)
		assert.Equal(t, value, 6)
	})
}
