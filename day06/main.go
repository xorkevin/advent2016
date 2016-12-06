package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

	m := [8]string{}

	for s.Scan() {
		s := s.Text()
		for n, i := range s {
			m[n] += string(i)
		}
	}

	final := ""
	final2 := ""

	for _, i := range m {
		count := 0
		count2 := 999
		str := ""
		str2 := ""
		for j := 'a'; j < 'z'+1; j++ {
			ch := string(j)
			k := strings.Count(i, ch)
			if k > count {
				count = k
				str = ch
			}
			if k < count2 {
				count2 = k
				str2 = ch
			}
		}

		final += str
		final2 += str2
	}

	fmt.Println("most common chars per column: ", final)
	fmt.Println("least common chars per column: ", final2)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
