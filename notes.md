# Notes

* Neighbor -> a node which is accessible through an edge (direction matters)
* Adjacency list -> a map with nodes as keys and lists of neighbors as values
   If a node has no neighbors, then the value for it is an empty list
* DFS uses a stack
* BFS uses a queue
* A cycle -> a path which can lead to the same source node:
A->B->A or A->B->C->A, etc.
* Acyclic graph -> a graph without cycles
