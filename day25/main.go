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
	// file_name = "test.txt"
	file_name = "input.txt"
	cpy       = iota
	inc
	dec
	jnz
	tgl
	out
)

type (
	Instruction struct {
		op          int
		arg1        int
		isRegister1 bool
		arg2        int
		isRegister2 bool
	}

	Executor struct {
		memory map[rune]int
	}
)

func (e *Executor) execute(i Instruction, pos int, instructs []Instruction) int {
	switch i.op {
	case cpy:
		if i.isRegister2 {
			if i.isRegister1 {
				e.memory[rune(i.arg2)] = e.memory[rune(i.arg1)]
			} else {
				e.memory[rune(i.arg2)] = i.arg1
			}
		}
	case inc:
		e.memory[rune(i.arg1)]++
	case dec:
		e.memory[rune(i.arg1)]--
	case jnz:
		if (!i.isRegister1 && i.arg1 != 0) || e.memory[rune(i.arg1)] != 0 {
			if i.isRegister2 {
				return e.memory[rune(i.arg2)]
			} else {
				return i.arg2
			}
		}
	case tgl:
		k := e.memory[rune(i.arg1)] + pos
		if k < len(instructs) {
			switch instructs[k].op {
			case inc:
				instructs[k].op = dec
			case dec, tgl:
				instructs[k].op = inc
			case jnz:
				instructs[k].op = cpy
			case cpy:
				instructs[k].op = jnz
			}
		}
	case out:
		e.memory['e'] = e.memory[rune(i.arg1)]
		// fmt.Println("out: ", e.memory['e'])
		e.memory['f'] = 1
		return 1
	}
	e.memory['f'] = 0
	return 1
}

func (e *Executor) run(limit int, instructs []Instruction) bool {
	num := e.memory['a']
	flip := false
	prev := 1
	// for i, j := 0, 0; j < limit; j++ {
	for i, j := 0, 0; true; j++ {
		// for k := 'a'; k < 'g'; k++ {
		// 	fmt.Print(e.memory[k], " ")
		// }
		// fmt.Println()
		if i >= len(instructs) {
			return false
		}
		ins := i
		i += e.execute(instructs[i], i, instructs)

		k := ""
		switch instructs[ins].op {
		case inc:
			k = "inc"
		case dec:
			k = "dec"
		case cpy:
			k = "cpy"
		case jnz:
			k = "jnz"
		case tgl:
			k = "tgl"
		case out:
			k = "out"
		}
		fmt.Println(j+1, ins+1, k, "|", e.memory['a'], e.memory['b'], e.memory['c'], e.memory['d'])

		if e.memory['f'] == 1 {
			flip = true
			fmt.Println(e.memory['e'])
			if e.memory['e'] == prev {
				fmt.Println(num, "failed")
				return false
			}
			prev = e.memory['e']
		}
	}

	return flip
}

func NewExecutor(a int) *Executor {
	return &Executor{
		memory: map[rune]int{
			'a': a,
			'b': 0,
			'c': 0,
			'd': 0,
			'e': 0,
			'f': 0,
		},
	}
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

	k := []Instruction{}

	for s.Scan() {
		s := strings.Fields(s.Text())
		var instruct Instruction
		switch s[0] {
		case "cpy":
			num, err := strconv.Atoi(s[1])
			if err == nil {
				instruct = Instruction{
					op:          cpy,
					arg1:        num,
					isRegister1: false,
					arg2:        int(s[2][0]),
					isRegister2: true,
				}
			} else {
				instruct = Instruction{
					op:          cpy,
					arg1:        int(s[1][0]),
					isRegister1: true,
					arg2:        int(s[2][0]),
					isRegister2: true,
				}
			}

		case "inc":
			instruct = Instruction{
				op:   inc,
				arg1: int(s[1][0]),
			}

		case "dec":
			instruct = Instruction{
				op:   dec,
				arg1: int(s[1][0]),
			}

		case "jnz":
			val1, err1 := strconv.Atoi(s[1])
			val2, err2 := strconv.Atoi(s[2])
			if err1 != nil {
				val1 = int(s[1][0])
			}
			if err2 != nil {
				val2 = int(s[2][0])
			}
			instruct = Instruction{
				op:          jnz,
				arg1:        val1,
				isRegister1: err1 != nil,
				arg2:        val2,
				isRegister2: err2 != nil,
			}

		case "tgl":
			instruct = Instruction{
				op:   tgl,
				arg1: int(s[1][0]),
			}
		case "out":
			instruct = Instruction{
				op:   out,
				arg1: int(s[1][0]),
			}
		}

		k = append(k, instruct)
	}

	// for i := 0; i < 9999; i++ {
	// 	e := NewExecutor(i)
	// 	if e.run(99999, k) {
	// 		fmt.Println("register: ", i)
	// 		break
	// 	}
	// }
	e := NewExecutor(182)
	e.run(999999, k)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
