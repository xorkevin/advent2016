package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

type (
	Pair struct {
		gen  int
		chip int
	}

	Pairlist []Pair

	State struct {
		cost      int
		heuristic int
		elevator  int
		pairs     Pairlist
	}

	Index struct {
		id   int
		chip bool
	}

	Statelist struct {
		states []State
	}

	Searcher struct {
		cost       int
		target     State
		openlist   *Statelist
		closedlist *Statelist
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

func (p Pairlist) Len() int {
	return len(p)
}
func (p Pairlist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Pairlist) Less(i, j int) bool {
	k := p[i].gen
	l := p[j].gen
	if k == l {
		return p[i].chip < p[j].chip
	} else {
		return k < l
	}
}

func StatesEqual(s1 *State, s2 *State) bool {
	return reflect.DeepEqual(s1, s2)
}

func NewStatelist() *Statelist {
	return &Statelist{
		states: []State{},
	}
}

func (s *Statelist) add(other State) {
	unique := true
	for _, i := range s.states {
		if StatesEqual(&i, &other) {
			unique = false
			break
		}
	}
	if unique {
		s.states = append(s.states, other)
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

func (s *State) sort() {
	sort.Sort(s.pairs)
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

func (s *State) nextStates(target *State) *Statelist {
	states := NewStatelist()

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
		}
		next_down, success2 := s.down(0, k.id, k.chip)
		if success2 {
			next_down.elevator_down(0)
			next_down.inc_cost()
		}
		for j := i + 1; j < len(indicies); j++ {
			l := indicies[j]
			if success1 {
				next, _ := next_up.up(4, l.id, l.chip)
				next.sort()
				// next.calc_heuristic(target)
				states.add(next)
			}
			if success2 {
				next, _ := next_up.down(0, l.id, l.chip)
				next.sort()
				// next.calc_heuristic(target)
				states.add(next)
			}
		}
		if success1 {
			next_up.sort()
			// next_up.calc_heuristic(target)
			states.add(next_up)
		}
		if success2 {
			next_down.sort()
			// next_down.calc_heuristic(target)
			states.add(next_down)
		}
	}
	return states
}

func NewSearcher(init, target State) *Searcher {
	k := Searcher{
		cost:       0,
		target:     target,
		openlist:   NewStatelist(),
		closedlist: NewStatelist(),
	}
	k.openlist.add(init)
	return &k
}

func (s *Searcher) search() bool {

	return true
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

	nextStates := init_state.nextStates(&target_state)

	for _, i := range nextStates.states {
		fmt.Println(i)
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
