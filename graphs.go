package main

import (
	"fmt"

	"github.com/savelyevvv/go-collections/queue"
	"github.com/savelyevvv/go-collections/stack"
)

type EdgeList [][2]string
type AdjacencyList map[string][]string

func DFS(src string, graph AdjacencyList) {
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

func BFS(src string, graph AdjacencyList) {
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

func DFSR(src string, graph AdjacencyList) {
	fmt.Print(src)
	for _, neighbor := range graph[src] {
		DFSR(neighbor, graph)
	}
}

func HasPathDFS(src, dst string, graph AdjacencyList) bool {
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

func HasPathBFS(src, dst string, graph AdjacencyList) bool {
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

var g = AdjacencyList{
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

var edges = EdgeList{
	{"a", "b"},
	{"b", "e"},
	{"b", "f"},
	{"e", "g"},
	{"g", "f"},
	{"f", "h"},
	{"a", "c"},
	{"c", "k"},
	{"c", "l"},
	{"m", "l"},
	{"m", "k"},
	{"k", "d"},
	{"k", "i"},
	{"k", "j"},
	{"a", "d"},
	{"d", "i"},
	{"i", "j"},
}
