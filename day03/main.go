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

func isValidTriangle(k [3]int64) bool {
	a := k[0]
	b := k[1]
	c := k[2]
	return a+b > c && b+c > a && a+c > b
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

	possible := 0
	possibleVertical := 0

	j := 0
	triangles := [3][3]int64{}

	for s.Scan() {
		s := strings.Fields(s.Text())

		k := [3]int64{}
		for n, i := range s {
			k[n], _ = strconv.ParseInt(i, 10, 64)
		}

		if isValidTriangle(k) {
			possible++
		}

		triangles[0][j] = k[0]
		triangles[1][j] = k[1]
		triangles[2][j] = k[2]

		j++
		if j > 2 {
			j = 0

			for _, i := range triangles {
				if isValidTriangle(i) {
					possibleVertical++
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("horizontal triangles: %d", possible))
	fmt.Println(fmt.Sprintf("vertical triangles: %d", possibleVertical))

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
