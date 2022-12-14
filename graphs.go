package main

import (
	"fmt"

	"github.com/savelyevvv/go-collections/queue"
	"github.com/savelyevvv/go-collections/stack"
)

func DFS(start string, graph map[string][]string) {
	s := stack.New[string]()
	s.Push(start)
	for !s.IsEmpty() {
		node := s.Pop()
		for _, neighbor := range graph[node] {
			s.Push(neighbor)
		}
		fmt.Print(node)
	}
	fmt.Println()
}

func DFSR(source string, graph map[string][]string) {
	fmt.Print(source)
	for _, neighbor := range graph[source] {
		DFSR(neighbor, graph)
	}
}

func BFS(source string, graph map[string][]string) {
	q := queue.New[string]()
	q.Add(source)
	for !q.IsEmpty() {
		node := q.Remove()
		for _, neighbor := range graph[node] {
			q.Add(neighbor)
		}
		fmt.Print(node)
	}
	fmt.Println()
}

func HasPathDFS(src, dst string, graph map[string][]string) bool {
	var rsl bool
	stack := stack.New[string]()
	stack.Push(src)
	for !stack.IsEmpty() {
		node := stack.Pop()
		if node == dst {
			rsl = true
			break
		}
		for _, neighbor := range graph[node] {
			stack.Push(neighbor)
		}
	}
	return rsl
}

func HasPathBFS(src, dst string, graph map[string][]string) bool {
	var rsl bool
	queue := queue.New[string]()
	queue.Add(src)
	for !queue.IsEmpty() {
		node := queue.Remove()
		if node == dst {
			rsl = true
			break
		}
		for _, neighbor := range graph[node] {
			queue.Add(neighbor)
		}
	}
	return rsl
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
	"m": {"l"},
}
