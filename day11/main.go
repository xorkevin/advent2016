package main

import (
	"fmt"
	"time"
)

type (
	Pair struct {
		gen  int
		chip int
	}

	State struct {
		cost      int
		heuristic int
		elevator  int
		pairs     []Pair
	}
)

func NewPair(gen, chip int) Pair {
	return Pair{
		gen:  gen,
		chip: chip,
	}
}

func (p *Pair) up(limit int, chip bool) (Pair, bool) {
	if chip {
		k := p.chip + 1
		if k >= limit {
			return NewPair(0, 0), false
		}
		return NewPair(p.gen, k), true
	} else {
		k := p.gen + 1
		if k >= limit {
			return NewPair(0, 0), false
		}
		return NewPair(k, p.chip), true
	}
}

func (p *Pair) down(limit int, chip bool) (Pair, bool) {
	if chip {
		k := p.chip - 1
		if k < limit {
			return NewPair(0, 0), false
		}
		return NewPair(p.gen, k), true
	} else {
		k := p.gen - 1
		if k < limit {
			return NewPair(0, 0), false
		}
		return NewPair(k, p.chip), true
	}
}

func (s *State) up(limit, pairId int, chip bool) (State, bool) {
	newState := *s
	pair, success := s.pairs[pairId].up(limit, chip)
	if success {
		newState.pairs[pairId] = pair
		return newState, true
	} else {
		return State{}, false
	}
}

func (s *State) down(limit, pairId int, chip bool) (State, bool) {
	newState := *s
	pair, success := s.pairs[pairId].down(limit, chip)
	if success {
		newState.pairs[pairId] = pair
		return newState, true
	} else {
		return State{}, false
	}
}

func main() {
	start := time.Now()

	init_state := State{
		cost:      0,
		heuristic: 0,
		elevator:  0,
		pairs: []Pair{
			NewPair(0, 0),
			NewPair(0, 0),
			NewPair(1, 1),
			NewPair(1, 1),
			NewPair(1, 2),
		},
	}

	target_state := State{
		cost:      0,
		heuristic: 0,
		elevator:  3,
		pairs: []Pair{
			NewPair(3, 3),
			NewPair(3, 3),
			NewPair(3, 3),
			NewPair(3, 3),
			NewPair(3, 3),
		},
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
