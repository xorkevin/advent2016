package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	file_name = "input.txt"
)

func main() {
	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	filestring := strings.TrimSpace(string(file))

	possible := 0
	possibleVertical := 0
	j := 0
	triangles := [3][3]int{}
	for _, i := range strings.Split(filestring, "\n") {
		s := strings.Fields(i)
		a, _ := strconv.Atoi(string(s[0]))
		b, _ := strconv.Atoi(string(s[1]))
		c, _ := strconv.Atoi(string(s[2]))

		if a+b > c && b+c > a && a+c > b {
			possible++
		}

		triangles[0][j] = a
		triangles[1][j] = b
		triangles[2][j] = c
		j++

		if j > 2 {
			j = 0

			for _, i := range triangles {
				a := i[0]
				b := i[1]
				c := i[2]
				if a+b > c && b+c > a && a+c > b {
					possibleVertical++
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("horizontal triangles: %d", possible))
	fmt.Println(fmt.Sprintf("vertical triangles: %d", possibleVertical))
}
