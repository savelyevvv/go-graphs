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
}

func DFSR(node string, list map[string][]string) {
	fmt.Print(node)
	for _, neighbor := range list[node] {
		DFSR(neighbor, list)
	}
}

func BFS(start string, list map[string][]string) {
	q := queue.NewQueue[string]()
	q.Add(start)
	for !q.IsEmpty() {
		node := q.Remove()
		for _, neighbor := range list[node] {
			q.Add(neighbor)
		}
		fmt.Print(node)
	}
	fmt.Println()
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
