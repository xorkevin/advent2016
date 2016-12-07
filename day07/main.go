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

func hasAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func hasABA_BAB(s1, s2 string) bool {
	for i := 0; i < len(s1)-2; i++ {
		if s1[i] == s1[i+2] && s1[i] != s1[i+1] {
			a := s1[i]
			b := s1[i+1]
			for j := 0; j < len(s2)-2; j++ {
				if s2[j] == s2[j+2] && s2[j+1] == a && s2[j] == b {
					return true
				}
			}
		}
	}

	return false
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

	supportsTLS := 0

	supportsSSL := 0

	for s.Scan() {
		m := strings.Split(s.Text(), "[")

		// part 1

		outsideABBA := false
		insideABBA := false

		if hasAbba(m[0]) {
			outsideABBA = true
		}

		for _, i := range m[1:] {
			k := strings.Split(i, "]")
			if hasAbba(k[0]) {
				insideABBA = true
			}
			if hasAbba(k[1]) {
				outsideABBA = true
			}
			if insideABBA && outsideABBA {
				break
			}
		}

		if !insideABBA && outsideABBA {
			supportsTLS++
		}

		// part 2

		outside := []string{}
		inside := []string{}

		outside = append(outside, m[0])

		for _, i := range m[1:] {
			k := strings.Split(i, "]")
			inside = append(inside, k[0])
			outside = append(outside, k[1])
		}

		success := false

		for _, i := range outside {
			for _, j := range inside {
				if hasABA_BAB(i, j) {
					success = true
				}
			}
			if success {
				break
			}
		}

		if success {
			supportsSSL++
		}
	}

	fmt.Println("supports TLS: ", supportsTLS)
	fmt.Println("supports SSL: ", supportsSSL)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
