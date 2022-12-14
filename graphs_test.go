package main

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
