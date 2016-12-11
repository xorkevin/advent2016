package main

import (
	"fmt"
	"time"
)

type (
	Item struct {
		chip    bool
		element string
		floor   int
	}

	State struct {
		cost     int
		elevator int
		gens     []Item
		chips    []Item
	}

	Index struct {
		index int
		chip  bool
	}

	Search struct {
		cost   int
		target State
		open   []State
		closed []State
	}
)

func (i *Item) up(limit int) bool {
	if i.floor+1 < limit {
		i.floor += 1
		return true
	} else {
		return false
	}
}

func (i *Item) down(limit int) bool {
	if i.floor-1 >= limit {
		i.floor -= 1
		return true
	} else {
		return false
	}
}

func NewState() State {
	return State{
		cost:     0,
		elevator: 0,
		gens: []Item{
			Item{
				chip:    false,
				element: "s",
				floor:   0,
			},
			Item{
				chip:    false,
				element: "p",
				floor:   0,
			},
			Item{
				chip:    false,
				element: "t",
				floor:   1,
			},
			Item{
				chip:    false,
				element: "r",
				floor:   1,
			},
			Item{
				chip:    false,
				element: "c",
				floor:   1,
			},
		},
		chips: []Item{
			Item{
				chip:    true,
				element: "s",
				floor:   0,
			},
			Item{
				chip:    true,
				element: "p",
				floor:   0,
			},
			Item{
				chip:    true,
				element: "t",
				floor:   2,
			},
			Item{
				chip:    true,
				element: "r",
				floor:   1,
			},
			Item{
				chip:    true,
				element: "c",
				floor:   1,
			},
		},
	}
}

func NewTargetState() State {
	return State{
		cost:     0,
		elevator: 4,
		gens: []Item{
			Item{
				chip:    false,
				element: "s",
				floor:   4,
			},
			Item{
				chip:    false,
				element: "p",
				floor:   4,
			},
			Item{
				chip:    false,
				element: "t",
				floor:   4,
			},
			Item{
				chip:    false,
				element: "r",
				floor:   4,
			},
			Item{
				chip:    false,
				element: "c",
				floor:   4,
			},
		},
		chips: []Item{
			Item{
				chip:    true,
				element: "s",
				floor:   4,
			},
			Item{
				chip:    true,
				element: "p",
				floor:   4,
			},
			Item{
				chip:    true,
				element: "t",
				floor:   4,
			},
			Item{
				chip:    true,
				element: "r",
				floor:   4,
			},
			Item{
				chip:    true,
				element: "c",
				floor:   4,
			},
		},
	}
}

func EqualStates(a, b State) bool {
	if a.elevator != b.elevator {
		return false
	}

	for i, _ := range a.chips {
		if a.chips[i] != b.chips[i] || a.gens[i] != b.gens[i] {
			return false
		}
	}

	return true
}

func (s *State) copy() State {
	a := []Item{}
	b := []Item{}
	copy(a, s.gens)
	copy(b, s.chips)
	return State{
		cost:     s.cost,
		elevator: s.elevator,
		gens:     a,
		chips:    b,
	}
}

func (s *State) elevator_up(limit int) bool {
	if s.elevator+1 < limit {
		s.elevator++
		return true
	} else {
		return false
	}
}

func (s *State) elevator_down(limit int) bool {
	if s.elevator-1 >= limit {
		s.elevator--
		return true
	} else {
		return false
	}
}

func (s *State) up(chip bool, element, limit int) bool {
	if chip {
		return s.chips[element].up(limit)
	} else {
		return s.gens[element].up(limit)
	}
}

func (s *State) down(chip bool, element, limit int) bool {
	if chip {
		return s.chips[element].down(limit)
	} else {
		return s.gens[element].down(limit)
	}
}

func (s *State) incCost() {
	s.cost++
}

func (s *State) nextStates() []State {
	k := []State{}
	indicies := []Index{}
	for i := 0; i < len(s.chips); i++ {
		if s.elevator == s.chips[i].floor {
			indicies = append(indicies, Index{i, true})
		}
	}
	for i := 0; i < len(s.gens); i++ {
		if s.elevator == s.gens[i].floor {
			indicies = append(indicies, Index{i, false})
		}
	}

	for _, i := range indicies {
		j := s.copy()
		if j.elevator_up(4) && j.up(i.chip, i.index, 4) {
			j.incCost()
			k = append(k, j)
		}
		j = s.copy()
		if j.elevator_down(0) && j.down(i.chip, i.index, 0) {
			j.incCost()
			k = append(k, j)
		}
	}
	for n, i := range indicies {
		for l := n + 1; l < len(indicies); l++ {
			j := s.copy()
			if j.elevator_up(4) && j.up(i.chip, i.index, 4) && j.up(indicies[n].chip, indicies[n].index, 4) {
				j.incCost()
				k = append(k, j)
			}
			j = s.copy()
			if j.elevator_down(0) && j.down(i.chip, i.index, 0) && j.down(indicies[n].chip, indicies[n].index, 0) {
				j.incCost()
				k = append(k, j)
			}
		}
	}

	return k
}

func NewSearch(init_state, target_state State) *Search {
	return &Search{
		target: target_state,
		open:   []State{init_state},
		closed: []State{},
	}
}

func (s *Search) addClosed(states []State) {
state_loop:
	for _, i := range states {
		for _, j := range s.closed {
			if EqualStates(i, j) {
				continue state_loop
			}
		}
		s.closed = append(s.closed, i)
	}
}

func (s *Search) addOpen(states []State) {
state_loop:
	for _, i := range states {
		for _, j := range s.closed {
			if EqualStates(i, j) {
				continue state_loop
			}
		}
		for _, j := range s.open {
			if EqualStates(i, j) {
				continue state_loop
			}
		}
		s.open = append(s.open, i)
	}
}

func (s *Search) search() bool {
	fmt.Println("\n\nrun")
	fmt.Println("\n\n", s.closed)
	fmt.Println("\n\n", s.open)
	if len(s.open) < 1 {
		return false
	}
	current := s.open[0]
	s.open = s.open[1:]
	if EqualStates(s.target, current) {
		s.cost = current.cost
		return false
	} else {
		s.addClosed([]State{current})
		s.addOpen(current.nextStates())
		return true
	}
}

func main() {
	start := time.Now()

	k := NewSearch(NewState(), NewTargetState())

	for k.search() {
	}

	fmt.Println("Cost: ", k.cost)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
