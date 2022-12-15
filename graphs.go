package main

import (
	"fmt"

	"github.com/savelyevvv/go-collections/queue"
	"github.com/savelyevvv/go-collections/stack"
)

func DFS(src string, graph map[string][]string) {
	s := stack.New[string]()
	s.Push(src)
	seen := map[string]struct{}{src: {}}
	for !s.IsEmpty() {
		node := s.Pop()
		fmt.Print(node)
		for _, neighbor := range graph[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = struct{}{}
				s.Push(neighbor)
			}
		}
	}
	fmt.Println()
}

func BFS(src string, graph map[string][]string) {
	q := queue.New[string]()
	q.Add(src)
	seen := map[string]struct{}{src: {}}
	for !q.IsEmpty() {
		node := q.Remove()
		for _, neighbor := range graph[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = struct{}{}
				q.Add(neighbor)
			}
		}
		fmt.Print(node)
	}
	fmt.Println()
}

func DFSR(src string, graph map[string][]string) {
	fmt.Print(src)
	for _, neighbor := range graph[src] {
		DFSR(neighbor, graph)
	}
}

func HasPathDFS(src, dst string, graph map[string][]string) bool {
	var rsl bool
	s := stack.New[string]()
	s.Push(src)
	seen := map[string]struct{}{src: {}}
	for !s.IsEmpty() {
		node := s.Pop()
		if node == dst {
			rsl = true
			break
		}
		for _, neighbor := range graph[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = struct{}{}
				s.Push(neighbor)
			}
		}
	}
	return rsl
}

func HasPathBFS(src, dst string, graph map[string][]string) bool {
	var rsl bool
	q := queue.New[string]()
	q.Add(src)
	seen := map[string]struct{}{src: {}}
	for !q.IsEmpty() {
		node := q.Remove()
		if node == dst {
			rsl = true
			break
		}
		for _, neighbor := range graph[node] {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = struct{}{}
				q.Add(neighbor)
			}
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
