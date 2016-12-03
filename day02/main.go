package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

const (
	file_name = "input.txt"
)

var (
	numpad = map[Vec2Keypad]rune{
		Vec2Keypad{0, -2}:  '1',
		Vec2Keypad{-1, -1}: '2',
		Vec2Keypad{0, -1}:  '3',
		Vec2Keypad{1, -1}:  '4',
		Vec2Keypad{-2, 0}:  '5',
		Vec2Keypad{-1, 0}:  '6',
		Vec2Keypad{0, 0}:   '7',
		Vec2Keypad{1, 0}:   '8',
		Vec2Keypad{2, 0}:   '9',
		Vec2Keypad{-1, 1}:  'A',
		Vec2Keypad{0, 1}:   'B',
		Vec2Keypad{1, 1}:   'C',
		Vec2Keypad{0, 2}:   'D',
	}
)

type (
	Vec2 struct {
		x int
		y int
	}

	Vec2Keypad struct {
		x int
		y int
	}
)

func (v *Vec2) up() {
	v.y = int(math.Max(0.0, float64(v.y-1)))
}

func (v *Vec2) down() {
	v.y = int(math.Min(2.0, float64(v.y+1)))
}

func (v *Vec2) left() {
	v.x = int(math.Max(0.0, float64(v.x-1)))
}

func (v *Vec2) right() {
	v.x = int(math.Min(2.0, float64(v.x+1)))
}

func (v *Vec2) toDigit() int {
	return v.y*3 + v.x + 1
}

func (v Vec2Keypad) distance() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)))
}

func (v *Vec2Keypad) set(x, y int) {
	k := Vec2Keypad{x, y}
	if k.distance() < 3 {
		v.x = x
		v.y = y
	}
}

func (v *Vec2Keypad) up() {
	v.set(v.x, v.y-1)
}

func (v *Vec2Keypad) down() {
	v.set(v.x, v.y+1)
}

func (v *Vec2Keypad) left() {
	v.set(v.x-1, v.y)
}

func (v *Vec2Keypad) right() {
	v.set(v.x+1, v.y)
}

func (v *Vec2Keypad) toDigit() rune {
	return numpad[*v]
}

func main() {
	start := time.Now()

	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	filestring := strings.TrimSpace(string(file))

	directions := strings.Split(filestring, "\n")

	digits := make([]int, len(directions))
	digits2 := make([]rune, len(directions))

	currentButton := &Vec2{1, 1}
	currentButton2 := &Vec2Keypad{-2, 0}
	for n, i := range directions {
		for _, j := range i {
			switch j {
			case 'U':
				currentButton.up()
				currentButton2.up()
			case 'D':
				currentButton.down()
				currentButton2.down()
			case 'L':
				currentButton.left()
				currentButton2.left()
			case 'R':
				currentButton.right()
				currentButton2.right()
			}
		}
		digits[n] = currentButton.toDigit()
		digits2[n] = currentButton2.toDigit()
	}

	fmt.Println("square keypad: ", digits)
	fmt.Println("larger keypad: ", string(digits2))

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
