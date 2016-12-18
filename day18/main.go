package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// - Its left and center tiles are traps, but its right tile is not.
// - Its center and right tiles are traps, but its left tile is not.
// - Only its left tile is a trap.
// - Only its right tile is a trap.

type (
	Vec2 struct {
		r, c int
	}

	Field struct {
		dim   Vec2
		table [][]bool
	}
)

func NewField(r, c int) *Field {
	return &Field{
		dim:   Vec2{r, c},
		table: make([][]bool, r),
	}
}

func (f *Field) get(r, c int) bool {
	if c >= 0 && c < f.dim.c && r >= 0 && r < f.dim.r {
		return f.table[r][c]
	}
	return false
}

func trapCondition(a, b, c, x, y, z bool) bool {
	return x == a && y == b && z == c
}

func (f *Field) deriveTrap(r, c int) bool {
	le := f.get(r-1, c-1)
	ce := f.get(r-1, c)
	re := f.get(r-1, c+1)
	return trapCondition(true, true, false, le, ce, re) || trapCondition(false, true, true, le, ce, re) || trapCondition(true, false, false, le, ce, re) || trapCondition(false, false, true, le, ce, re)
}

const (
	file_name = "input.txt"
	rows      = 400000
	// rows      = 40
)

func main() {
	start := time.Now()

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	l := make([]bool, 0)

	count := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, i := range s.Text() {
			l = append(l, i == '^')
			if i != '^' {
				count++
			}
		}
	}

	cols := len(l)
	field := NewField(rows, cols)
	field.table[0] = l

	for i := 1; i < rows; i++ {
		r := make([]bool, 0)
		for j := 0; j < cols; j++ {
			t := field.deriveTrap(i, j)
			r = append(r, t)
			if !t {
				count++
			}
		}
		field.table[i] = r
	}

	fmt.Println("safe tiles: ", count)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
