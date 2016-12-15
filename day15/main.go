package main

import (
	"fmt"
	"time"
)

type (
	Disc struct {
		size, i int
	}
)

func (d *Disc) tick(count int) {
	d.i = (d.i + count) % d.size
}

func (d *Disc) peek(count int) int {
	return (d.i + count) % d.size
}

var (
	d = []*Disc{
		&Disc{5, 2},
		&Disc{13, 7},
		&Disc{17, 10},
		&Disc{3, 2},
		&Disc{19, 9},
		&Disc{7, 0},
		&Disc{11, 0},
	}
)

func checkAlignment(discs []*Disc) bool {
	for n, i := range discs {
		if i.peek(n+1) != 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	t := 0

	forward := (2*d[0].size - d[0].i - 1) % d[0].size
	t += forward
	for _, i := range d {
		i.tick(forward)
	}
	forward = d[0].size

	for !checkAlignment(d) {
		t += forward
		for _, i := range d {
			i.tick(forward)
		}
	}

	fmt.Println(t)
	for n, i := range d {
		fmt.Println(n, i)
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
