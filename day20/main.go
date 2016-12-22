package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type (
	Range struct {
		begin, end int
	}

	Rangelist struct {
		list []Range
	}
)

func (r Range) combine(other Range) Range {
	b := r.begin
	if other.begin < b {
		b = other.begin
	}
	e := r.end
	if other.end > e {
		e = other.end
	}

	return Range{b, e}
}

func (r Range) intersect(other Range) bool {
	return r.end+1 >= other.begin
}

func NewRangelist() *Rangelist {
	return &Rangelist{
		list: []Range{},
	}
}

func (r *Rangelist) addRange(low, high int) {
	r.list = append(r.list, Range{low, high})
}

func (r Rangelist) Len() int {
	return len(r.list)
}
func (r Rangelist) Less(i, j int) bool {
	a := r.list[i].begin
	b := r.list[j].begin
	if a == b {
		return r.list[i].end < r.list[j].end
	} else {
		return a < b
	}
}
func (r Rangelist) Swap(i, j int) {
	r.list[i], r.list[j] = r.list[j], r.list[i]
}

const (
	file_name = "input.txt"
	max       = 4294967295
)

func main() {
	start := time.Now()

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := NewRangelist()

	s := bufio.NewScanner(f)
	for s.Scan() {
		k := strings.Split(s.Text(), "-")
		low, _ := strconv.Atoi(k[0])
		high, _ := strconv.Atoi(k[1])
		r.addRange(low, high)
	}

	sort.Sort(r)

	for i := 1; i < len(r.list); i++ {
		if r.list[i-1].intersect(r.list[i]) {
			r.list = append(append(r.list[:i-1], r.list[i-1].combine(r.list[i])), r.list[i+1:]...)
			i--
		}
	}

	lowest := 0

	for _, i := range r.list {
		if lowest >= i.begin && lowest <= i.end {
			lowest = i.end + 1
		}
	}

	fmt.Println("lowest: ", lowest)

	count := 0
	for i := 1; i < len(r.list); i++ {
		count += r.list[i].begin - r.list[i-1].end - 1
	}

	fmt.Println("count: ", count)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
