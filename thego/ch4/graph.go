package main

var graph = make(map[string]map[string]bool)

func main() {

}

func addEdge(from, to string) {
	// gv, ok := graph[from] // subscripting a map
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
