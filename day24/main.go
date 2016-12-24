package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

const (
	square_wall = iota
	square_path
)

type (
	Square struct {
		state int
	}

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

	Permutor struct {
		a    []int
		c    []int
		n, i int
	}
)

// Map Layout

func isWall(vec Vec2, m [][]Square) bool {
	return m[vec.y][vec.x].state == square_wall
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

func (n *Node) nextNodes(target *Node, z [][]Square) *Nodelist {
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

func (s *Searcher) search(z [][]Square, cap int) bool {
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

func NewPermutor(a []int) *Permutor {
	return &Permutor{
		a: a,
		c: make([]int, len(a)),
		n: len(a),
		i: 0,
	}
}

func (p *Permutor) permute() []int {
	n := p.n
	i := p.i
	if i < n {
		if p.c[i] < i {
			if i%2 == 0 {
				p.a[0], p.a[i] = p.a[i], p.a[0]
			} else {
				p.a[p.c[i]], p.a[i] = p.a[i], p.a[p.c[i]]
			}
			p.c[i]++
			p.i = 0
			return p.a[:]
		} else {
			p.c[i] = 0
			p.i++
			return p.permute()
		}
	} else {
		return nil
	}
}

const (
	file_name = "input.txt"
)

func main() {
	start := time.Now()

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	board := [][]Square{}

	nodes := map[int]*Node{}

	s := bufio.NewScanner(f)
	for row := 0; s.Scan(); row++ {
		line := []Square{}
		for col, i := range s.Text() {
			j := square_path
			if i == '#' {
				j = square_wall
			} else if i >= '0' && i <= '9' {
				nodes[int(i-'0')] = &Node{Vec2: Vec2{col, row}}
			}

			line = append(line, Square{
				state: j,
			})
		}
		board = append(board, line)
	}

	salesman := map[int]map[int]int{}

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			s := NewSearcher(nodes[i], nodes[j])

			for s.search(board, 500) {
			}

			if _, ok := salesman[i]; !ok {
				salesman[i] = map[int]int{}
			}
			if _, ok := salesman[j]; !ok {
				salesman[j] = map[int]int{}
			}

			salesman[i][j] = s.path.cost
			salesman[j][i] = s.path.cost
		}
	}

	min := 999999
	p := NewPermutor([]int{1, 2, 3, 4, 5, 6, 7})
	for i := p.permute(); i != nil; i = p.permute() {
		sum := 0
		prev := 0
		for _, j := range i {
			sum += salesman[prev][j]
			prev = j
		}
		sum += salesman[prev][0] // toggle for part 2
		if sum < min {
			min = sum
		}
	}

	fmt.Println(min)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
