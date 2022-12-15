# Notes

* Neighbor is a vertex which is accessible through an edge (direction matters).
* Adjacency list is a map with vertices as keys and lists of neighbors as
  values. If a vertex has no neighbors, then its value is an empty list.
* DFS uses a stack.
* BFS uses a queue
* A cycle -> a path which can lead to the same source node:
	A->B->A or A->B->C->A, etc.
* Acyclic graph -> a graph without cycles

## DFS

### Time complexity

Let V is the number of vertices and E is the number of edges.
Every vertex which is accessible from the source will be added to the set of
seen nodes exactly once. Add operation has O(1) time complexity.
Every vertex which is accessible from the source will be pushed to the stack
exactly once. Push operation has O(1) time complexity.
Every vertex which is accessible from the source will be popped from the stack
exactly once. Pop operation has O(1) time complexity.
Every neighbor of every vertex which is accessible from the source will be
checked for the presence in the set of seen nodes. The number of neighbors for
a particular vertex is equal to the number of edges from it. Searching in a
set is O(1) operation.
Resulting time complexity is O(V+E).

In the worst case there is an edge from every vertex to every other vertex.
There are V - 1 edges from every vertex. Total number of edges is V * (V - 1) or
V<sup>2</sup> - V.
Time complexity is O(V<sup>2</sup>)

### Space Complexity

Before ending the algorithm application a set of seen vertices will contain
every vertex which is reachable from the source.
Space complexity is O(V).

## BFS

Generally the same thinking as for DFS.

### Time complexity

Time complexity: O(V+E)

### Space complexity
Space complexity: O(V)
