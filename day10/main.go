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

type (
	Bot struct {
		id       int
		low      int
		high     int
		low_bot  bool
		high_bot bool
		carry    []int
	}

	Inp struct {
		val int
		bot int
	}

	Factory struct {
		bots  map[int]*Bot
		outs  map[int]int
		input []Inp
	}
)

func (b *Bot) feed(val int) {
	b.carry = append(b.carry, val)
}

func (b *Bot) simulate() (low, high int, hasTwo bool) {
	if len(b.carry) < 2 {
		return 0, 0, false
	}

	o1 := b.carry[0]
	o2 := b.carry[1]
	b.carry = b.carry[2:]

	if o1 < o2 {
		return o1, o2, true
	} else {
		return o2, o1, true
	}
}

func NewFactory() *Factory {
	return &Factory{
		bots:  map[int]*Bot{},
		outs:  map[int]int{},
		input: []Inp{},
	}
}

func (f *Factory) feed(val, bot int) {
	f.input = append(f.input, Inp{
		val: val,
		bot: bot,
	})
}

func (f *Factory) execFeed() {
	for _, i := range f.input {
		f.bots[i.bot].feed(i.val)
	}
}

func (f *Factory) simulate() bool {
	changed := false

	for _, val := range f.bots {
		low, high, hasTwo := val.simulate()

		if low == 17 && high == 61 {
			fmt.Println("bot id: ", val.id)
		}

		if hasTwo {
			changed = true

			if val.low_bot {
				f.bots[val.low].feed(low)
			} else {
				f.outs[val.low] = low
			}

			if val.high_bot {
				f.bots[val.high].feed(high)
			} else {
				f.outs[val.high] = high
			}
		}
	}

	return changed
}

func (f *Factory) initBot(id, low, high int, low_bot, high_bot bool) {
	f.bots[id] = &Bot{
		id:       id,
		low:      low,
		high:     high,
		low_bot:  low_bot,
		high_bot: high_bot,
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

	factory := NewFactory()

	for s.Scan() {
		s := strings.Fields(s.Text())
		if s[0] == "value" {
			val, _ := strconv.Atoi(s[1])
			bot, _ := strconv.Atoi(s[5])
			factory.feed(val, bot)
		} else if s[0] == "bot" {
			bot, _ := strconv.Atoi(s[1])
			low, _ := strconv.Atoi(s[6])
			high, _ := strconv.Atoi(s[11])
			low_bot := s[5] == "bot"
			high_bot := s[10] == "bot"
			factory.initBot(bot, low, high, low_bot, high_bot)
		}
	}

	factory.execFeed()

	for factory.simulate() {
	}

	k := factory.outs[0] * factory.outs[1] * factory.outs[2]

	fmt.Println(k)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
