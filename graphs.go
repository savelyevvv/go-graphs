package main

import (
	"fmt"

	"github.com/savelyevvv/go-collections/queue"
	"github.com/savelyevvv/go-collections/stack"
)

func DFS(start string, list map[string][]string) {
	s := stack.NewStack[string]()
	s.Push(start)
	for !s.IsEmpty() {
		node := s.Pop()
		for _, neighbor := range list[node] {
			s.Push(neighbor)
		}
		fmt.Print(node)
	}
	fmt.Println()
	_ = s
}

func BFS(list map[string][]string) {
	s := queue.NewQueue[string]()
	_ = s
}

var g = map[string][]string{
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

func main() {
	DFS("a", g)
}
