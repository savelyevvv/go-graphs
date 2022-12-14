package main

import (
	"github.com/savelyevvv/go-collections/queue"
	"github.com/savelyevvv/go-collections/stack"
)

func DFS(list map[string][]string) {
	s := stack.NewStack[int]()
	_ = s
}

func BFS(list map[string][]string) {
	s := queue.NewQueue[int]()
	_ = s
}

func main() {
	g := map[string][]string{
		"a": {"b", "c", "d"},
		"b": {"e", "f"},
		"c": {"k", "l"},
		"d": {"i"},
		"e": {"g"},
		"f": {"h"},
		"g": {},
		"h": {},
		"i": {"j"},
		"j": {},
		"k": {},
		"l": {},
	}
	_ = g
}
