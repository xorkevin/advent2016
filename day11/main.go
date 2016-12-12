package main

import (
	"fmt"
	"math"
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

	Index struct {
		id   int
		chip bool
	}
)

func NewIndex(id int, chip bool) Index {
	return Index{id, chip}
}

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

func (s *State) copy() State {
	newPairs := make([]Pair, len(s.pairs))
	copy(newPairs, s.pairs)
	return State{
		cost:      s.cost,
		heuristic: s.heuristic,
		elevator:  s.elevator,
		pairs:     newPairs,
	}
}

func (s *State) inc_cost() {
	s.cost++
}

func (s *State) calc_heuristic(target *State) {
	k := 0
	for n, i := range s.pairs {
		chip := target.pairs[n].chip
		gen := target.pairs[n].gen
		k += int(math.Abs(float64(chip-i.chip))) + int(math.Abs(float64(gen-i.gen)))
	}
	s.heuristic = k
}

func (s *State) elevator_up(limit int) bool {
	k := s.elevator + 1
	if k >= limit {
		return false
	}
	s.elevator = k
	return true
}

func (s *State) elevator_down(limit int) bool {
	k := s.elevator - 1
	if k < limit {
		return false
	}
	s.elevator = k
	return true
}

func (s *State) up(limit, pairId int, chip bool) (State, bool) {
	newState := s.copy()
	pair, success := s.pairs[pairId].up(limit, chip)
	if success {
		newState.pairs[pairId] = pair
		return newState, true
	} else {
		return State{}, false
	}
}

func (s *State) down(limit, pairId int, chip bool) (State, bool) {
	newState := s.copy()
	pair, success := s.pairs[pairId].down(limit, chip)
	if success {
		newState.pairs[pairId] = pair
		return newState, true
	} else {
		return State{}, false
	}
}

func (s *State) nextStates(target *State) []State {
	states := []State{}

	indicies := []Index{}
	for n, i := range s.pairs {
		if i.gen == s.elevator {
			indicies = append(indicies, NewIndex(n, false))
		}
		if i.chip == s.elevator {
			indicies = append(indicies, NewIndex(n, true))
		}
	}

	for i := 0; i < len(indicies); i++ {
		k := indicies[i]
		next_up, success1 := s.up(4, k.id, k.chip)
		if success1 {
			next_up.elevator_up(4)
			next_up.inc_cost()
			next_up.calc_heuristic(target)
			states = append(states, next_up)
		}
		next_down, success2 := s.down(0, k.id, k.chip)
		if success2 {
			next_down.elevator_down(0)
			next_down.inc_cost()
			next_down.calc_heuristic(target)
			states = append(states, next_down)
		}
		for j := i + 1; j < len(indicies); j++ {
			l := indicies[j]
			if success1 {
				next, _ := next_up.up(4, l.id, l.chip)
				next.calc_heuristic(target)
				states = append(states, next)
			}
			if success2 {
				next, _ := next_up.down(0, l.id, l.chip)
				next.calc_heuristic(target)
				states = append(states, next)
			}
		}
	}
	return states
}

var (
	init_state = State{
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

	target_state = State{
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
)

func main() {
	start := time.Now()

	init_state.calc_heuristic(&target_state)

	nextStates := init_state.nextStates(&target_state)

	for _, i := range nextStates {
		fmt.Println(i)
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
