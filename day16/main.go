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

	memo = map[int]bool{}
)

func order(position, curveLength int) int {
	return int(math.Max(math.Log2(float64(position+1)/float64(curveLength+1)), 0))
}

func invOrder(ord, curveLength int) int {
	return int(math.Exp2(float64(ord))*float64(curveLength+1)) - 2
}

func dragonCurveVal(position int, curve []bool) bool {
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

func dragonCurve(position int, curve []bool) bool {
	if val, ok := memo[position]; ok {
		return val
	} else {
		k := dragonCurveVal(position, curve)
		memo[position] = k
		return k
	}
}

func checksumVal(position, factor int, curve []bool) bool {
	if factor < 2 {
		return dragonCurve(position*2, curve) == dragonCurve(position*2+1, curve)
	} else {
		return checksumVal(position*2, factor-1, curve) == checksumVal(position*2+1, factor-1, curve)
	}
}

func checksum(size int, curve []bool) []bool {
	k := 0
	finalSize := size
	for finalSize%2 == 0 {
		k++
		finalSize /= 2
	}

	arr := []bool{}
	for i := 0; i < finalSize; i++ {
		arr = append(arr, checksumVal(i, k, curve))
	}

	return arr
}

func main() {
	start := time.Now()

	printrate := 8

	fmt.Println("checksum")

	for n, i := range checksum(272, input) {
		if i {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if (n+1)%printrate == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Println("checksum2")

	for n, i := range checksum(35651584, input) {
		if i {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if (n+1)%printrate == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
