package main

import (
	"fmt"
	"math"
	"time"
)

var (
	input = []bool{
		false, false, true, true, true, true, false, true, true, true, true, true, false, true, false, false, false,
	}
)

func order(position, curveLength int) int {
	return int(math.Max(math.Log2(float64(position+1)/float64(curveLength+1)), 0))
}

func invOrder(ord, curveLength int) int {
	return int(math.Exp2(float64(ord))*float64(curveLength+1)) - 2
}

func dragonCurve(position int, curve []bool) bool {
	curveLength := len(curve)
	ord := order(position, curveLength)
	ordprev := order(position-1, curveLength)

	if ord == 0 {
		if position < curveLength {
			return curve[position]
		} else if position == curveLength {
			return false
		} else {
			return !curve[2*curveLength-position]
		}
	} else if ord-ordprev == 1 {
		return false
	} else {
		return !dragonCurve(invOrder(ord+1, curveLength)-position, curve)
	}
}

func main() {
	start := time.Now()

	for i := 0; i < 50; i++ {
		fmt.Println(i, dragonCurve(i, input))
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
