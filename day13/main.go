package main

import (
	"container/heap"
	"fmt"
	"time"
)

type (
	Vec2 struct {
		x, y int
	}

	Node struct {
		Vec2
		cost, heuristic int
		parent          *Node
	}

	Nodelist []*Node
)

// Map Layout

func hashFunc(x, y, z int) int {
	return x*x + 3*x + 2*x*y + y + y*y + z
}

func oddOnes(x int) bool {
	k := false
	for x > 0 {
		if x%2 != 0 {
			k = !k
		}
		x /= 2
	}
	return k
}

func isWall(vec Vec2, z int) bool {
	return oddOnes(hashFunc(vec.x, vec.y, z))
}

// Node

func (n *Node) isEqual(other *Node) bool {
	return n.x == other.x && n.y == other.y
}

func (n *Node) isLess(other *Node) bool {
	i := n.cost + n.heuristic
	j := other.cost + other.heuristic
	if i == j {
		return n.cost < other.cost
	}
	return i < j
}

func (n *Node) nextNodes() Nodelist {
	k := Nodelist{}
	k = append(k)
	return k
}

// Node Priority Queue

func (n Nodelist) Len() int {
	return len(n)
}
func (n Nodelist) Less(i, j int) bool {
	return n[i].cost+n[i].heuristic < n[j].cost+n[j].heuristic
}
func (n Nodelist) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n *Nodelist) Push(x interface{}) {
	*n = append(*n, x.(*Node))
}
func (n *Nodelist) Pop() interface{} {
	nodes := *n
	size := len(nodes)
	item := nodes[size-1]
	*n = nodes[0 : size-1]
	return item
}

func (n *Nodelist) add(node *Node) {
	unique := true
	for j, i := range *n {
		if i.isEqual(node) {
			if node.isLess(i) {
				heap.Remove(n, j)
			} else {
				unique = false
			}
		}
	}
	if unique {
		heap.Push(n, node)
	}
}

func (n *Nodelist) pop() *Node {
	return heap.Pop(n).(*Node)
}

// Search

const (
	inp = 1362
)

func main() {
	start := time.Now()

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
