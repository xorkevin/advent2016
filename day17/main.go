package main

import (
	"container/heap"
	"crypto/md5"
	"fmt"
	"time"
)

type (
	Vec2 struct {
		x, y int
	}

	Node struct {
		Vec2
		dir, cost, heuristic int
		parent               *Node
	}

	Nodelist []*Node

	Searcher struct {
		target  *Node
		path    *Node
		openset *Nodelist
	}
)

// Map Layout

func hashToHex(i byte) int {
	switch {
	case i >= '0' && i <= '9':
		return int(i - '0')
	case i == 'a':
		return 10
	case i == 'b':
		return 11
	case i == 'c':
		return 12
	case i == 'd':
		return 13
	case i == 'e':
		return 14
	case i == 'f':
		return 15
	default:
		return 16
	}
}

func checkDoorOpen(n *Node, passcode string) (up, down, left, right bool) {
	k := passcode + n.path()
	hash := fmt.Sprintf("%x", md5.Sum([]byte(k)))
	u := hashToHex(hash[0])
	d := hashToHex(hash[1])
	l := hashToHex(hash[2])
	r := hashToHex(hash[3])
	return u > 10, d > 10, l > 10, r > 10
}

// Node

func (n *Node) path() string {
	k := ""
	for i := n; i != nil; i = i.parent {
		switch i.dir {
		case 1:
			k = "U" + k
		case 2:
			k = "D" + k
		case 3:
			k = "L" + k
		case 4:
			k = "R" + k
		}
	}
	return k
}

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

func (n *Node) nextNodes(target *Node, z string) *Nodelist {
	k := &Nodelist{}
	north := &Node{
		Vec2: Vec2{
			x: n.x,
			y: n.y - 1,
		},
		dir:    1,
		cost:   n.cost + 1,
		parent: n,
	}
	north.calcHeuristic(target)
	east := &Node{
		Vec2: Vec2{
			x: n.x + 1,
			y: n.y,
		},
		dir:    4,
		cost:   n.cost + 1,
		parent: n,
	}
	east.calcHeuristic(target)
	south := &Node{
		Vec2: Vec2{
			x: n.x,
			y: n.y + 1,
		},
		dir:    2,
		cost:   n.cost + 1,
		parent: n,
	}
	south.calcHeuristic(target)
	west := &Node{
		Vec2: Vec2{
			x: n.x - 1,
			y: n.y,
		},
		dir:    3,
		cost:   n.cost + 1,
		parent: n,
	}
	west.calcHeuristic(target)

	no, so, we, ea := checkDoorOpen(n, z)

	if no && north.x >= 0 && north.y >= 0 && north.x <= 3 && north.y <= 3 {
		heap.Push(k, north)
	}
	if ea && east.x >= 0 && east.y >= 0 && east.x <= 3 && east.y <= 3 {
		heap.Push(k, east)
	}
	if so && south.x >= 0 && south.y >= 0 && south.x <= 3 && south.y <= 3 {
		heap.Push(k, south)
	}
	if we && west.x >= 0 && west.y >= 0 && west.x <= 3 && west.y <= 3 {
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

func (n *Nodelist) add(node *Node) {
	heap.Push(n, node)
}

func (n *Nodelist) pop() *Node {
	return heap.Pop(n).(*Node)
}

// Search

func NewSearcher(init, target *Node) *Searcher {
	k := &Searcher{
		target:  target,
		openset: &Nodelist{},
	}
	k.openset.add(init)
	return k
}

func (s *Searcher) search(z string, cap int, first bool) bool {
	if s.openset.len() < 1 {
		if s.path != nil {
			fmt.Println("cost: ", s.path.cost)
		} else {
			fmt.Println("fail")
		}
		return false
	}

	current := s.openset.pop()
	if s.target.isEqual(current) {
		s.path = current
		return !first
	}

	for _, i := range *current.nextNodes(s.target, z) {
		s.openset.add(i)
	}

	return true
}

const (
	inp = "hhhxzeay"
)

func main() {
	start := time.Now()

	init := &Node{
		Vec2:   Vec2{0, 0},
		cost:   0,
		parent: nil,
	}

	target := &Node{
		Vec2:   Vec2{3, 3},
		cost:   0,
		parent: nil,
	}

	s := NewSearcher(init, target)
	for s.search(inp, 4096, true) {
	}
	fmt.Println("path: ", s.path.path())
	for s.search(inp, 4096, false) {
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
