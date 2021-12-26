package main

import (
	"fmt"
)

type graph map[string][]string

func addNode(g graph, u, v string) {
	g[u] = append(g[u], v)
}

func DFS(g graph, u string) {
	_, ok := g[u]
	if !ok {
		return
	}

	str := ""
	discover := map[string]struct{}{}

	stack := []string{}
	stack = append(stack, u)
	for len(stack) != 0 {
		// Pop the stack; mark node as discovered
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		discover[node] = struct{}{}
		str += node + " "

		// Loop through out-neighbor of node
		for _, v := range g[node] {
			// Check if node has been discovered
			if _, discovered := discover[v]; discovered {
				// Skip this node
				continue
			} else {
				stack = append(stack, v)
			}
		}
	}
	fmt.Println("DFS Traversal:", str)
}

func main() {
	g := graph{}
	addNode(g, "a", "b")
	addNode(g, "a", "c")
	addNode(g, "c", "d")
	addNode(g, "c", "e")
	addNode(g, "c", "a")
	addNode(g, "c", "c")
	DFS(g, "a")
}
