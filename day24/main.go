package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

var maze = grid.NewVecMatrix[rune]()

type blizzardDirection uint8

const (
	Up blizzardDirection = iota
	Down
	Left
	Right
)

func (bd blizzardDirection) Rune() rune {
	switch bd {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	}
	panic("invalid blizzard direction")
}

type blizzard struct {
	grid.Vec
	direction blizzardDirection
}

func (b blizzard) move() blizzard {
	switch b.direction {
	case Up:
		vec := grid.Vec{X: b.X, Y: b.Y - 1}
		if vec.Y == 0 {
			vec.Y = r - 2
		}
		return blizzard{Vec: vec, direction: b.direction}
	case Down:
		vec := grid.Vec{X: b.X, Y: b.Y + 1}
		if vec.Y == r-1 {
			vec.Y = 1
		}
		return blizzard{Vec: vec, direction: b.direction}
	case Left:
		vec := grid.Vec{X: b.X - 1, Y: b.Y}
		if vec.X == 0 {
			vec.X = col - 2
		}
		return blizzard{Vec: vec, direction: b.direction}
	case Right:
		vec := grid.Vec{X: b.X + 1, Y: b.Y}
		if vec.X == col-1 {
			vec.X = 1
		}
		return blizzard{Vec: vec, direction: b.direction}
	}
	panic("invalid blizzard direction")
}

type mazeState struct {
	maze      grid.VecMatrix[rune]
	blizzards *set.Set[blizzard]
	entry     grid.Vec
	dest      grid.Vec
	myLoc     grid.Vec
	steps     int
}

func (ms mazeState) getMyNextLocations(reverse bool) []grid.Vec {
	current := ms.myLoc
	possible := set.New(
		grid.Vec{X: current.X + 1, Y: current.Y},
		grid.Vec{X: current.X, Y: current.Y + 1},
		grid.Vec{X: current.X, Y: current.Y - 1},
		grid.Vec{X: current.X - 1, Y: current.Y},
	)
	if reverse {
		possible = set.New(
			grid.Vec{X: current.X - 1, Y: current.Y},
			grid.Vec{X: current.X, Y: current.Y - 1},
			grid.Vec{X: current.X, Y: current.Y + 1},
			grid.Vec{X: current.X + 1, Y: current.Y},
		)
	}
	var locs []grid.Vec
	possible.Each(func(vec grid.Vec) bool {
		if c, ok := ms.maze.Get(vec); ok && c == '.' {
			locs = append(locs, vec)
		}
		return true
	})
	return locs
}

func (ms mazeState) reset(nb *set.Set[blizzard]) grid.VecMatrix[rune] {
	newMaze := grid.NewVecMatrix[rune]()
	ms.maze.ForEach(func(v grid.Vec, r rune) {
		if r == '^' || r == 'v' || r == '<' || r == '>' {
			newMaze.Add(v, '.')
		} else {
			newMaze.Add(v, r)
		}
	})
	nb.Each(func(b blizzard) bool {
		newMaze.Add(b.Vec, b.direction.Rune())
		return true
	})
	return newMaze
}

func (ms mazeState) moveUntil(reverse bool) mazeState {
	for {
		newState, ok := ms.nextState(reverse)
		if ok {
			if newState.myLoc == newState.dest {
				return newState
			}
			ms = newState
			// ms.maze.Print(os.Stdout, "%c")
			// fmt.Printf("myLoc: %s\n", ms.myLoc)
			// fmt.Printf("steps: %d\n", ms.steps)
		} else {
			panic("no next state")
		}
	}
}

type pathState struct {
	loc   grid.Vec
	steps int
}

var (
	states   = []mazeState{}
	stateMap = make(map[pathState]struct{})
)

func (ms mazeState) nextState(reverse bool) (mazeState, bool) {
	// blizzard's turn
	newBlizzards := set.New[blizzard]()

	ms.blizzards.Each(func(item blizzard) bool {
		newBlizzards.Add(item.move())
		return true
	})
	ms.maze = ms.reset(newBlizzards)
	// my turn
	locs := ms.getMyNextLocations(reverse)
	// move to next location
	for _, loc := range locs {
		newState := mazeState{
			maze:      ms.maze,
			blizzards: newBlizzards,
			entry:     ms.entry,
			dest:      ms.dest,
			myLoc:     loc,
			steps:     ms.steps + 1,
		}
		ps := pathState{
			loc:   newState.myLoc,
			steps: newState.steps,
		}
		if _, ok := stateMap[ps]; !ok {
			states = append(states, newState)
			stateMap[ps] = struct{}{}
		}
	}

	if c, ok := ms.maze.Get(ms.myLoc); ok && c == '.' {
		// or still wait
		state := mazeState{
			maze:      ms.maze,
			blizzards: newBlizzards,
			entry:     ms.entry,
			dest:      ms.dest,
			myLoc:     ms.myLoc,
			steps:     ms.steps + 1,
		}
		ps := pathState{
			loc:   state.myLoc,
			steps: state.steps,
		}
		if _, ok := stateMap[ps]; !ok {
			states = append(states, state)
			stateMap[ps] = struct{}{}
		}
	}

	if len(states) == 0 {
		return mazeState{}, false
	}

	next := states[0]
	states = states[1:]
	return next, true
}

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	initBlizzards := set.New[blizzard]()
	var start, dest grid.Vec
	maze.ForEach(func(v grid.Vec, rn rune) {
		if v.Y == 0 && rn == '.' {
			start = v
		}
		if v.Y == r-1 && rn == '.' {
			dest = v
		}
		switch rn {
		case '^':
			initBlizzards.Add(blizzard{Vec: v, direction: Up})
		case 'v':
			initBlizzards.Add(blizzard{Vec: v, direction: Down})
		case '<':
			initBlizzards.Add(blizzard{Vec: v, direction: Left})
		case '>':
			initBlizzards.Add(blizzard{Vec: v, direction: Right})
		}
	})
	// maze.Print(os.Stdout, "%c")
	state := mazeState{
		maze:      maze,
		blizzards: initBlizzards,
		entry:     start,
		dest:      dest,
		myLoc:     start,
		steps:     0,
	}

	ps := pathState{
		loc:   state.myLoc,
		steps: state.steps,
	}
	states = append(states, state)
	stateMap[ps] = struct{}{}

	state = state.moveUntil(false)

	step := state.steps

	fmt.Fprintf(os.Stdout, "part1: steps: %d\n", step)

	states = []mazeState{}
	stateMap = make(map[pathState]struct{})

	state = mazeState{
		maze:      state.maze,
		blizzards: state.blizzards,
		entry:     dest,
		dest:      start,
		myLoc:     dest,
		steps:     0,
	}
	states = append(states, state)
	stateMap[ps] = struct{}{}
	state = state.moveUntil(true)
	back := state.steps

	states = []mazeState{}
	stateMap = make(map[pathState]struct{})

	state = mazeState{
		maze:      state.maze,
		blizzards: state.blizzards,
		entry:     start,
		dest:      dest,
		myLoc:     start,
		steps:     0,
	}
	states = append(states, state)
	stateMap[ps] = struct{}{}
	state = state.moveUntil(false)
	rt := state.steps

	step += back + rt

	fmt.Fprintf(os.Stdout, "part2: steps: %d\n", step)
}

var (
	r   int
	col int
)

func handler(line string) error {
	if r == 0 {
		col = len(line)
	}
	for index, e := range line {
		maze.Add(grid.Vec{
			X: index,
			Y: r,
		}, e)
	}
	r++
	return nil
}
