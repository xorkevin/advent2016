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

	Searcher struct {
		target    *Node
		path      *Node
		openset   *Nodelist
		closedset *Nodelist
	}
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

func (n *Node) calcHeuristic(target *Node) {
	i := target.x - n.x
	if i*-1 > i {
		i *= -1
	}
	j := target.y - n.y
	if j*-1 > j {
		j *= -1
	}
	n.heuristic = i + j
}

func (n *Node) nextNodes(target *Node, z int) *Nodelist {
	k := &Nodelist{}
	north := &Node{
		Vec2: Vec2{
			x: n.x,
			y: n.y - 1,
		},
		cost:   n.cost + 1,
		parent: n,
	}
	north.calcHeuristic(target)
	east := &Node{
		Vec2: Vec2{
			x: n.x + 1,
			y: n.y,
		},
		cost:   n.cost + 1,
		parent: n,
	}
	east.calcHeuristic(target)
	south := &Node{
		Vec2: Vec2{
			x: n.x,
			y: n.y + 1,
		},
		cost:   n.cost + 1,
		parent: n,
	}
	south.calcHeuristic(target)
	west := &Node{
		Vec2: Vec2{
			x: n.x - 1,
			y: n.y,
		},
		cost:   n.cost + 1,
		parent: n,
	}
	west.calcHeuristic(target)

	if !isWall(north.Vec2, z) && north.x >= 0 && north.y >= 0 {
		heap.Push(k, north)
	}
	if !isWall(east.Vec2, z) && east.x >= 0 && east.y >= 0 {
		heap.Push(k, east)
	}
	if !isWall(south.Vec2, z) && south.x >= 0 && south.y >= 0 {
		heap.Push(k, south)
	}
	if !isWall(west.Vec2, z) && west.x >= 0 && west.y >= 0 {
		heap.Push(k, west)
	}

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

func (n *Nodelist) len() int {
	return len(*n)
}

func (n *Nodelist) hasNode(node *Node) int {
	for j, i := range *n {
		if i.isEqual(node) {
			return j
		}
	}
	return -1
}

func (n *Nodelist) add(node *Node) {
	if j := n.hasNode(node); j > -1 {
		if node.isLess((*n)[j]) {
			heap.Remove(n, j)
			heap.Push(n, node)
		}
	} else {
		heap.Push(n, node)
	}
}

func (n *Nodelist) pop() *Node {
	return heap.Pop(n).(*Node)
}

// Search

func NewSearcher(init, target *Node) *Searcher {
	k := &Searcher{
		target:    target,
		openset:   &Nodelist{},
		closedset: &Nodelist{},
	}
	k.openset.add(init)
	return k
}

func (s *Searcher) search(z, cap int) bool {
	if s.openset.len() < 1 {
		return false
	}

	current := s.openset.pop()
	if s.target.isEqual(current) {
		s.path = current
		return false
	}

	s.closedset.add(current)

	for _, i := range *current.nextNodes(s.target, z) {
		if s.closedset.hasNode(i) < 0 && i.cost <= cap {
			s.openset.add(i)
		}
	}

	return true
}

const (
	inp = 1362
)

func main() {
	start := time.Now()

	init := &Node{
		Vec2:   Vec2{1, 1},
		cost:   0,
		parent: nil,
	}

	target := &Node{
		Vec2:   Vec2{31, 39},
		cost:   0,
		parent: nil,
	}

	s := NewSearcher(init, target)
	for s.search(inp, 100) {
	}
	fmt.Println("cost: ", s.path.cost)
	fmt.Println("searched: ", s.closedset.len())

	s2 := NewSearcher(init, target)
	for s2.search(inp, 50) {
	}
	fmt.Println("nodes within cost 50: ", s2.closedset.len())

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
