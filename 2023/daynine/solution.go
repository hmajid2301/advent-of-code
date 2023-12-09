package daynine

import (
	"strconv"
	"strings"
)

func CalcaulteNextValue(sequences string) int {
	lines := strings.Split(sequences, "\n")

	sum := 0
	for _, line := range lines {
		sequence := StringToIntSlice(line)
		sum += nextValue(sequence)
	}
	return sum

}

func CalcaultePrevValue(sequences string) int {
	lines := strings.Split(sequences, "\n")

	sum := 0
	for _, line := range lines {
		sequence := StringToIntSlice(line)
		sum += previousValue(sequence)
	}
	return sum

}
func previousValue(sequence []int) int {
	if isZeroSlice(sequence) {
		return 0
	}

	diffList := diffs(sequence)
	previousDiff := previousValue(diffList)
	return sequence[0] - previousDiff
}

func nextValue(sequence []int) int {
	if isZeroSlice(sequence) {
		return 0
	}

	diffList := diffs(sequence)
	nextDiff := nextValue(diffList)
	return sequence[len(sequence)-1] + nextDiff
}

func diffs(sequence []int) []int {
	diffs := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		diffs[i] = sequence[i+1] - sequence[i]
	}
	return diffs
}

func isZeroSlice(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}

func StringToIntSlice(s string) []int {
	var t2 = []int{}

	for _, i := range strings.Fields(s) {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}
