package dayeight

import "strings"

func CalculateStepsInRoute(tree string) int {
	lines := strings.Split(tree, "\n")
	treeMap := map[string][2]string{}
	for _, line := range lines[2:] {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]

		treeMap[node] = [2]string{left, right}
	}

	var dir = map[string]int{"L": 0, "R": 1}

	path := lines[0]
	var steps int
	atNode := "AAA"
	for atNode != "ZZZ" {
		for _, route := range strings.Split(path, "") {
			atNode = treeMap[atNode][dir[route]]
			steps++
			if atNode == "ZZZ" {
				break
			}
		}
	}
	return steps
}

func CalculateStepsToAllZ(tree string) int {
	lines := strings.Split(tree, "\n")
	treeMap := map[string][2]string{}
	for _, line := range lines[2:] {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]

		treeMap[node] = [2]string{left, right}
	}

	var dir = map[string]int{"L": 0, "R": 1}

	path := lines[0]
	var steps int
	atNode := "AAA"
	for atNode != "ZZZ" {
		for _, route := range strings.Split(path, "") {
			atNode = treeMap[atNode][dir[route]]
			steps++
			if atNode == "ZZZ" {
				break
			}
		}
	}
	return steps
}
