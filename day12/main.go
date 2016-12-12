package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	file_name = "input.txt"
	cpy       = iota
	inc
	dec
	jnz
)

type (
	Instruction struct {
		op         int
		register   rune
		isVal      bool
		val        int
		isRegister bool
	}

	Executor struct {
		memory map[rune]int
	}
)

func (e *Executor) execute(i Instruction) int {
	switch i.op {
	case cpy:
		if i.isRegister {
			e.memory[i.register] = e.memory[rune(i.val)]
		} else {
			e.memory[i.register] = i.val
		}
	case inc:
		e.memory[i.register]++
	case dec:
		e.memory[i.register]--
	case jnz:
		if (i.isVal && int(i.register) != 0) || e.memory[i.register] != 0 {
			return i.val
		}
	}
	return 1
}

func (e *Executor) run(wg *sync.WaitGroup, instructs []Instruction) {
	defer wg.Done()
	i := 0
	for i < len(instructs) {
		i += e.execute(instructs[i])
	}
}

func NewExecutor() *Executor {
	return &Executor{
		memory: map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
			'd': 0,
		},
	}
}

func NewExecutor2() *Executor {
	return &Executor{
		memory: map[rune]int{
			'a': 0,
			'b': 0,
			'c': 1,
			'd': 0,
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
					op:         cpy,
					register:   rune(s[2][0]),
					val:        num,
					isRegister: false,
				}
			} else {
				instruct = Instruction{
					op:         cpy,
					register:   rune(s[2][0]),
					val:        int(s[1][0]),
					isRegister: true,
				}
			}

		case "inc":
			instruct = Instruction{
				op:       inc,
				register: rune(s[1][0]),
			}

		case "dec":
			instruct = Instruction{
				op:       dec,
				register: rune(s[1][0]),
			}

		case "jnz":
			num, _ := strconv.Atoi(s[2])
			val, err := strconv.Atoi(s[1])
			if err == nil {
				instruct = Instruction{
					op:       jnz,
					register: rune(val),
					isVal:    true,
					val:      num,
				}
			} else {
				instruct = Instruction{
					op:       jnz,
					register: rune(s[1][0]),
					isVal:    false,
					val:      num,
				}
			}
		}
		k = append(k, instruct)
	}

	e := NewExecutor()
	e2 := NewExecutor2()

	var wg sync.WaitGroup

	wg.Add(2)
	go e.run(&wg, k)
	go e2.run(&wg, k)

	wg.Wait()
	fmt.Println(e.memory['a'])
	fmt.Println(e2.memory['a'])

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
