package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	Permutor struct {
		a    []rune
		c    []int
		n, i int
	}
)

func NewPermutor(a []rune) *Permutor {
	return &Permutor{
		a: a,
		c: make([]int, len(a)),
		n: len(a),
		i: 0,
	}
}

// Heap's algorithm
// procedure generate(n : integer, A : array of any):
//   c : array of int
//
//   for i := 0; i < n; i += 1 do
//       c[i] := 0
//   end for
//
//   output(A)
//
//   i := 0;
//   while i < n do
//       if  c[i] < i then
//           if i is even then
//               swap(A[0], A[i])
//           else
//               swap(A[c[i]], A[i])
//           end if
//           output(A)
//           c[i] += 1
//           i := 0
//       else
//           c[i] := 0
//           i += 1
//       end if
//   end while

func (p *Permutor) permute() []rune {
	n := p.n
	i := p.i
	if i < n {
		if p.c[i] < i {
			if i%2 == 0 {
				p.a[0], p.a[i] = p.a[i], p.a[0]
			} else {
				p.a[p.c[i]], p.a[i] = p.a[i], p.a[p.c[i]]
			}
			p.c[i]++
			p.i = 0
			return p.a[:]
		} else {
			p.c[i] = 0
			p.i++
			return p.permute()
		}
	} else {
		return nil
	}
}

func op(opstring, text string) string {
	s := strings.Fields(opstring)
	switch s[0] {
	case "swap":
		if s[1] == "position" {
			x, _ := strconv.Atoi(s[2])
			y, _ := strconv.Atoi(s[5])
			t := []rune(text)
			t[x], t[y] = t[y], t[x]
			return string(t)
		} else {
			x := strings.Index(text, s[2])
			y := strings.Index(text, s[5])
			t := []rune(text)
			t[x], t[y] = t[y], t[x]
			return string(t)
		}
	case "rotate":
		if s[3] == "steps" || s[3] == "step" {
			x, _ := strconv.Atoi(s[2])
			x %= len(text)
			if x > 0 {
				if s[1] == "left" {
					return text[x:] + text[:x]
				} else {
					x = len(text) - x
					return text[x:] + text[:x]
				}
			} else {
				return text
			}
		} else {
			x := strings.Index(text, s[6])
			if x > 3 {
				x += 1
			}
			x += 1
			x %= len(text)
			x = len(text) - x
			return text[x:] + text[:x]
		}
	case "reverse":
		x, _ := strconv.Atoi(s[2])
		y, _ := strconv.Atoi(s[4])
		t := []rune(text)
		for i, j := x, y; i < j; i, j = i+1, j-1 {
			t[i], t[j] = t[j], t[i]
		}
		return string(t)
	case "move":
		x, _ := strconv.Atoi(s[2])
		y, _ := strconv.Atoi(s[5])

		t := text
		char := string(t[x])
		t = t[:x] + t[x+1:]
		return t[:y] + char + t[y:]
	default:
		return text
	}
}

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

	text := "abcdefgh"

	m := []string{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		text = op(t, text)
		m = append(m, t)
	}

	fmt.Println(text)

	target := "fbgdceah"

	p := NewPermutor([]rune("abcdefgh"))

	k := p.permute()

	for k != nil {
		j := string(k)
		t := j
		for _, i := range m {
			t = op(i, t)
		}
		if t == target {
			fmt.Println(j)
			break
		}
		k = p.permute()
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
