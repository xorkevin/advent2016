package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
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

func main() {
	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	filestring := strings.TrimSpace(string(file))

	directions := strings.Split(filestring, "\n")

	digits := make([]int, len(directions))

	currentButton := &Vec2{1, 1}
	for n, i := range directions {
		for _, j := range i {
			switch j {
			case 'U':
				currentButton.up()
			case 'D':
				currentButton.down()
			case 'L':
				currentButton.left()
			case 'R':
				currentButton.right()
			}
		}
		digits[n] = currentButton.toDigit()
	}

	fmt.Println(digits)
}
