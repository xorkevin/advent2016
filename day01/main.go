package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	file_name = "input.txt"
)

type (
	Vec2 struct {
		x int
		y int
	}
)

func (v Vec2) distance() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)))
}

func (v Vec2) asString() string {
	return fmt.Sprintf("(%d,%d)", v.x, v.y)
}

func main() {
	start := time.Now()

	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	filestring := strings.TrimSpace(string(file))

	// filestring = "R2, R2, R2"
	// filestring = "R5, L5, R5, R3"

	directions := strings.Split(filestring, ", ")

	history := make(map[Vec2]bool)

	x := 0
	y := 0
	orientation := 0

	history[Vec2{x, y}] = true

	checktwice := true

	firsttwice := 0

	for _, i := range directions {
		prev := Vec2{x, y}

		switch i[0] {
		case 'R':
			orientation += 3
		case 'L':
			orientation += 1
		}
		orientation %= 4
		m, err := strconv.Atoi(i[1:])
		if err != nil {
			fmt.Println(err)
			return
		}
		switch orientation {
		case 0:
			y += m
		case 1:
			x -= m
		case 2:
			y -= m
		case 3:
			x += m
		}

		if checktwice {
			for i := 1; i <= m; i++ {
				pos := prev
				switch orientation {
				case 0:
					pos = Vec2{pos.x, pos.y + i}
				case 1:
					pos = Vec2{pos.x - i, pos.y}
				case 2:
					pos = Vec2{pos.x, pos.y - i}
				case 3:
					pos = Vec2{pos.x + i, pos.y}
				}

				if history[pos] {
					checktwice = false
					firsttwice = pos.distance()
					break
				} else {
					history[pos] = true
				}
			}
		}
	}

	dist := Vec2{x, y}.distance()

	fmt.Println(fmt.Sprintf("final position: %s", Vec2{x, y}.asString()))
	fmt.Println(fmt.Sprintf("final distance: %d", dist))

	fmt.Println(fmt.Sprintf("first visited twice distance: %d", firsttwice))

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
