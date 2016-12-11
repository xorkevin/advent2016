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
		items    []Item
	}
)

func NewState() *State {
	return &State{
		cost:     0,
		elevator: 0,
		items: []Item{
			Item{
				chip:    false,
				element: "s",
				floor:   0,
			},
			Item{
				chip:    true,
				element: "s",
				floor:   0,
			},
			Item{
				chip:    false,
				element: "p",
				floor:   0,
			},
			Item{
				chip:    true,
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
				chip:    true,
				element: "r",
				floor:   1,
			},
			Item{
				chip:    false,
				element: "c",
				floor:   1,
			},
			Item{
				chip:    true,
				element: "c",
				floor:   1,
			},
			Item{
				chip:    true,
				element: "t",
				floor:   2,
			},
		},
	}
}

func (s *State) nextStates() []State {
	k := []State{}
	for _, i := range s.items {

	}
	return k
}

func main() {
	start := time.Now()

	k := NewState()

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
