package main

import (
	"fmt"
	"math"
	"time"
)

const (
	inp = 3004953
)

func josephus(i int) int {
	return (i-int(math.Exp2(math.Floor(math.Log2(float64(i))))))*2 + 1
}

func josephusAcross(i int) int {
	return i - int(math.Pow(3, math.Floor(math.Log(float64(i))/math.Log(3))))
}

func main() {
	start := time.Now()

	fmt.Println(josephus(inp))
	fmt.Println(josephusAcross(inp))

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
