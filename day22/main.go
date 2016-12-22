package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	Node struct {
		x, y, size, free, used int
	}
)

var (
	m = []Node{}
)

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

	s := bufio.NewScanner(f)
	s.Scan()
	s.Scan()
	for s.Scan() {
		s := strings.Fields(s.Text())
		xy := strings.Split(s[0], "-")
		x, _ := strconv.Atoi(xy[1][1:])
		y, _ := strconv.Atoi(xy[2][1:])
		size, _ := strconv.Atoi(s[1][:len(s[1])-1])
		used, _ := strconv.Atoi(s[2][:len(s[2])-1])
		free, _ := strconv.Atoi(s[3][:len(s[3])-1])

		m = append(m, Node{x, y, size, free, used})
	}

	count := 0

	for k, i := range m {
		for j := k + 1; j < len(m); j++ {
			if i.used > 0 && i.used <= m[j].free || m[j].used > 0 && m[j].used <= i.free {
				count++
			}
		}
	}

	fmt.Println(count)

	grid := [25][37]rune{}

	for _, i := range m {
		symbol := '.'
		if i.used > 200 {
			symbol = '#'
		} else if i.used < 1 {
			symbol = 'O'
		}
		grid[i.y][i.x] = symbol
	}

	fmt.Print("   ")
	for i := 0; i < 37; i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println()

	for n, i := range grid {
		fmt.Printf("%3d", n)
		for _, j := range i {
			fmt.Printf("  %c", j)
		}
		fmt.Println()
	}

	fmt.Println("\ncounting moves visually: ", 22+18+22+5*35+1)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
