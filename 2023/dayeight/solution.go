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
	startNodes := []string{}
	for _, line := range lines[2:] {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]

		treeMap[node] = [2]string{left, right}
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}

	var dir = map[string]int{"L": 0, "R": 1}

	stepsTaken := []int{}
	path := lines[0]
	for _, node := range startNodes {
		var steps int
		atNode := node
		for !strings.HasSuffix(atNode, "Z") {
			for _, route := range strings.Split(path, "") {
				atNode = treeMap[atNode][dir[route]]
				steps++
			}
		}
		stepsTaken = append(stepsTaken, steps)
	}
	return lcmN(stepsTaken)
}

func lcmN(n []int) int {
	if len(n) == 2 {
		return lcm(n[0], n[1])
	}
	return lcm(n[0], lcmN(n[1:]))
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
