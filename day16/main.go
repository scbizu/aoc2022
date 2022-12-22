package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/math"
)

// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
// Valve HH has flow rate=22; tunnel leads to valve GG
var regexpExpr = `Valve (?P<name>\w+) has flow rate=(?P<rate>\d+); tunnel(s?) lead(s?) to valve(s?) (?P<children>.+)`

var (
	total int
	times int = 1
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	for ii, v := range valves {
		for i, e := range v.edge {
			valves[ii].rate[i] = rateMap[e] - rateMap[v.name]
		}
	}
	// p1
	fmt.Fprintf(os.Stdout, "p1: %d", visitFrom(valves["AA"], 30, []string{}))

	stateSet = make(map[string]int)
	// p2
	// my released pressure + elephant's released pressure
	fmt.Fprintf(os.Stdout, "p2: %d", visit2From(valves["AA"], 26, []string{}, 1))
}

var stateSet = make(map[string]int)

func stringSliceContains(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func stringSliceRemove(s []string, v string) []string {
	for i, vv := range s {
		if vv == v {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func visitFrom(v *Valve, times int, opened []string) int {
	if times <= 0 {
		return 0
	}
	vvState := &state{start: v.name, times: times, opened: opened}
	if _, ok := stateSet[vvState.String()]; ok {
		return stateSet[vvState.String()]
	}

	var release int
	if v.selfRate > 0 && !stringSliceContains(opened, v.name) {
		opened = append(opened, v.name)
		sort.Slice(opened, func(i, j int) bool {
			return rateMap[opened[i]] < rateMap[opened[j]]
		})
		r := (times-1)*v.selfRate + visitFrom(v, times-1, opened)
		// other paths can be better ?
		opened = stringSliceRemove(opened, v.name)
		release = r
	}

	for _, e := range v.edge {
		release = math.Max(release, visitFrom(valves[e], times-1, opened))
	}
	stateSet[vvState.String()] = release
	return release
}

func visit2From(v *Valve, times int, opened []string, role int) int {
	if times == 0 {
		if role == 0 {
			return 0
		} else {
			// plus elephant's best released pressure
			return visit2From(valves["AA"], 26, opened, role-1)
		}
	}
	vvState := &p2state{start: v.name, times: times, opened: opened, role: role}
	if _, ok := stateSet[vvState.String()]; ok {
		return stateSet[vvState.String()]
	}

	var release int
	if v.selfRate > 0 && !stringSliceContains(opened, v.name) {
		opened = append(opened, v.name)
		sort.Slice(opened, func(i, j int) bool {
			return rateMap[opened[i]] < rateMap[opened[j]]
		})
		r := (times-1)*v.selfRate + visit2From(v, times-1, opened, role)
		// other paths can be better ?
		opened = stringSliceRemove(opened, v.name)
		release = r
	}

	for _, e := range v.edge {
		release = math.Max(release, visit2From(valves[e], times-1, opened, role))
	}
	stateSet[vvState.String()] = release
	return release
}

type state struct {
	start  string
	times  int
	opened []string
}

type p2state struct {
	start  string
	times  int
	opened []string
	role   int
}

func (s *p2state) String() string {
	return fmt.Sprintf("%d:%s:%d:%v", s.role, s.start, s.times, s.opened)
}

func (s *state) String() string {
	return fmt.Sprintf("%s:%d:%v", s.start, s.times, s.opened)
}

func handler(line string) error {
	r := regexp.MustCompile(regexpExpr)
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}
	expr := new(Expr)
	for i, name := range r.SubexpNames() {
		switch name {
		case "name":
			expr.Name = matches[i]
		case "rate":
			expr.Rate = input.Atoi(matches[i])
		case "children":
			expr.Tunnels = strings.Split(matches[i], ", ")
		}
		rateMap[expr.Name] = expr.Rate
	}

	Init(expr)
	return nil
}

type Expr struct {
	Name    string
	Rate    int
	Tunnels []string
}

type Valve struct {
	name     string
	selfRate int
	rate     []int
	edge     []string
}

var rateMap = make(map[string]int)

var valves = make(map[string]*Valve)

func Init(e *Expr) {
	v := &Valve{
		name:     e.Name,
		selfRate: e.Rate,
		edge:     make([]string, len(e.Tunnels)),
		rate:     make([]int, len(e.Tunnels)),
	}
	copy(v.edge, e.Tunnels)
	valves[v.name] = v
}
