package main

import "fmt"

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
