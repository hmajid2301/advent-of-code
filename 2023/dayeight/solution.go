package dayeight

import "strings"

type Node struct {
	left  *Node
	right *Node
	data  string
}

func CalculateStartToEnd(tree string) int {
	var steps int
	path := tree[0]

	nodes := []Node{}
	for _, line := range strings.Split(tree[2:], "\n") {
		data := strings.Split(line, "=")
		nodeName := data[0]
		nextNodes := data[1]

	}
	return steps
}
