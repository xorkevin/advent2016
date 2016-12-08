package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	file_name = "input.txt"
)

func parseLine(screen *[6][50]bool, line string) {
	m := strings.Fields(line)
	if m[0] == "rect" {
		dims := strings.Split(m[1], "x")
		a, _ := strconv.Atoi(dims[0])
		b, _ := strconv.Atoi(dims[1])

		for i := 0; i < b; i++ {
			for j := 0; j < a; j++ {
				screen[i][j] = true
			}
		}
	} else if m[0] == "rotate" {
		num, _ := strconv.Atoi(strings.Split(m[2], "=")[1])
		count, _ := strconv.Atoi(m[4])

		if m[1] == "row" {
			newscreen := [50]bool{}
			for i := 0; i < 50; i++ {
				newscreen[i] = screen[num][(i-count+50)%50]
			}
			for i := 0; i < 50; i++ {
				screen[num][i] = newscreen[i]
			}
		} else if m[1] == "column" {
			newscreen := [6]bool{}
			for i := 0; i < 6; i++ {
				newscreen[i] = screen[(i-count+6)%6][num]
			}
			for i := 0; i < 6; i++ {
				screen[i][num] = newscreen[i]
			}
		}
	}
}

func main() {
	start := time.Now()

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	screen := [6][50]bool{}

	for s.Scan() {
		parseLine(&screen, s.Text())
	}

	k := 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if screen[i][j] {
				fmt.Print("#")
				k++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("total on: ", k)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
