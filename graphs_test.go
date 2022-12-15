package main

import (
	"fmt"
	"sort"
	"strings"
)

func ExampleDFS() {
	DFS("a", g)
	// Output: adijclkbfheg
}

func ExampleBFS() {
	BFS("a", g)
	// Output: abcdefklighj
}

func ExampleDFSR() {
	DFSR("a", g)
	// Output: abegfhckldij
}

func ExampleHasPathDFS_true() {
	var rsl bool
	rsl = HasPathDFS("a", "l", g)
	fmt.Println(rsl)
	// Output: true
}

func ExampleHasPathDFS_false() {
	var rsl bool
	rsl = HasPathDFS("a", "m", g)
	fmt.Println(rsl)
	// Output: false
}

func ExampleHasPathBFS_true() {
	var rsl bool
	rsl = HasPathBFS("a", "l", g)
	fmt.Println(rsl)
	// Output: true
}

func ExampleHasPathBFS_false() {
	var rsl bool
	rsl = HasPathBFS("a", "m", g)
	fmt.Println(rsl)
	// Output: false
}

type stringPairs [][2]string

func (sp stringPairs) Len() int {
	return len(sp)
}

func (sp stringPairs) Less(i, j int) bool {
	return sp[i][0] < sp[j][0]
}

func (sp stringPairs) Swap(i, j int) {
	sp[i], sp[j] = sp[j], sp[i]
}

func ExampleFromEdgeListToAdjacencyList_directed() {
	var graph AdjacencyList
	edges := EdgeList{
		{"a", "b"},
		{"a", "c"},
		{"a", "d"},
		{"d", "c"},
	}
	graph = FromEdgeListToAdjacencyList(edges, true)
	// AdjacencyList is an alias to a map which has no order
	// As we need consistent printing, we convert a map into a list of string pairs
	// 1-st element of each pair is a vertex
	// 2-nd element is a string of all neighbors for the vertex
	var rsl [][2]string
	for node, neighbors := range graph {
		rsl = append(rsl, [2]string{node, strings.Join(neighbors, "")})
	}
	sort.Sort(stringPairs(rsl))
	fmt.Println(rsl)
	// Output: [[a bcd] [d c]]
}

func ExampleFromEdgeListToAdjacencyList_undirected() {
	var graph AdjacencyList
	edges := EdgeList{
		{"a", "b"},
		{"a", "c"},
		{"a", "d"},
		{"d", "c"},
	}
	graph = FromEdgeListToAdjacencyList(edges, false)
	// AdjacencyList is an alias to a map which has no order
	// As we need consistent printing, we convert a map into a list of string pairs
	// 1-st element of each pair is a vertex
	// 2-nd element is a string of all neighbors for the vertex
	var rsl [][2]string
	for node, neighbors := range graph {
		rsl = append(rsl, [2]string{node, strings.Join(neighbors, "")})
	}
	sort.Sort(stringPairs(rsl))
	fmt.Println(rsl)
	// Output: [[a bcd] [b a] [c ad] [d ac]]
}
