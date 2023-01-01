package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

const (
	human = "humn"
)

var errHasHuman = errors.New("has human")

var fn = func(m *monkey) error {
	if m.name == human {
		return errHasHuman
	}
	return nil
}

func main() {
	ctx := context.Background()
	input.NewTXTFile("input.txt").ReadByLine(ctx, handler)
	answer, _ := cal(cacheMonkeys["root"])
	fmt.Fprintf(os.Stdout, "p1: res: %d\n", answer)

	// reset
	cacheMonkeys = map[string]*monkey{}
	input.NewTXTFile("input.txt").ReadByLine(ctx, handler)
	parts, _ := extractMonkey(cacheMonkeys["root"].op)
	var v int
	var bad *monkey
	v1, err := cal(cacheMonkeys[parts[0]], fn)
	if err == nil {
		v = v1
		bad = cacheMonkeys[parts[1]]
	}
	v2, err := cal(cacheMonkeys[parts[1]], fn)
	if err == nil {
		v = v2
		bad = cacheMonkeys[parts[0]]
	}
	bad.number = &v
	n := revCal(bad)

	fmt.Fprintf(os.Stdout, "p2: res: %d\n", n)
}

var cacheMonkeys = map[string]*monkey{}

type monkey struct {
	name   string
	op     string
	number *int
}

func revCal(m *monkey) int {
	if m.name == human {
		return *m.number
	}
	if m.op == "" {
		panic("not possible")
	}
	parts, op := extractMonkey(m.op)
	switch op {
	case '+':
		v1, err := cal(cacheMonkeys[parts[0]], fn)
		if err == nil {
			n := *m.number - v1
			cacheMonkeys[parts[1]].number = &n
			return revCal(cacheMonkeys[parts[1]])
		}
		v2, err := cal(cacheMonkeys[parts[1]], fn)
		if err == nil {
			n := *m.number - v2
			cacheMonkeys[parts[0]].number = &n
			return revCal(cacheMonkeys[parts[0]])
		}
	case '*':
		v1, err := cal(cacheMonkeys[parts[0]], fn)
		if err == nil {
			n := *m.number / v1
			cacheMonkeys[parts[1]].number = &n
			return revCal(cacheMonkeys[parts[1]])
		}
		v2, err := cal(cacheMonkeys[parts[1]], fn)
		if err == nil {
			n := *m.number / v2
			cacheMonkeys[parts[0]].number = &n
			return revCal(cacheMonkeys[parts[0]])
		}
	case '-':
		v1, err := cal(cacheMonkeys[parts[0]], fn)
		if err == nil {
			n := v1 - *m.number
			cacheMonkeys[parts[1]].number = &n
			return revCal(cacheMonkeys[parts[1]])
		}
		v2, err := cal(cacheMonkeys[parts[1]], fn)
		if err == nil {
			n := v2 + *m.number
			cacheMonkeys[parts[0]].number = &n
			return revCal(cacheMonkeys[parts[0]])
		}
	case '/':
		v1, err := cal(cacheMonkeys[parts[0]], fn)
		if err == nil {
			n := v1 / *m.number
			cacheMonkeys[parts[1]].number = &n
			return revCal(cacheMonkeys[parts[1]])
		}
		v2, err := cal(cacheMonkeys[parts[1]], fn)
		if err == nil {
			n := v2 * *m.number
			cacheMonkeys[parts[0]].number = &n
			return revCal(cacheMonkeys[parts[0]])
		}
	}
	panic(fmt.Sprintf("can not rev cal %+v", m))
}

func cal(m *monkey, fns ...func(m *monkey) error) (int, error) {
	for _, fn := range fns {
		if err := fn(m); err != nil {
			return 0, err
		}
	}
	if m.number != nil {
		return *m.number, nil
	}
	parts, op := extractMonkey(m.op)
	switch op {
	case '+':
		v1, err := cal(cacheMonkeys[parts[0]], fns...)
		if err != nil {
			return 0, err
		}
		v2, err := cal(cacheMonkeys[parts[1]], fns...)
		if err != nil {
			return 0, err
		}
		v := v1 + v2
		cacheMonkeys[m.name].number = &v
		return v, nil
	case '*':
		v1, err := cal(cacheMonkeys[parts[0]], fns...)
		if err != nil {
			return 0, err
		}
		v2, err := cal(cacheMonkeys[parts[1]], fns...)
		if err != nil {
			return 0, err
		}
		v := v1 * v2
		cacheMonkeys[m.name].number = &v
		return v, nil

	case '/':
		v1, err := cal(cacheMonkeys[parts[0]], fns...)
		if err != nil {
			return 0, err
		}
		v2, err := cal(cacheMonkeys[parts[1]], fns...)
		if err != nil {
			return 0, err
		}
		v := v1 / v2
		cacheMonkeys[m.name].number = &v
		return v, nil
	case '-':
		v1, err := cal(cacheMonkeys[parts[0]], fns...)
		if err != nil {
			return 0, err
		}
		v2, err := cal(cacheMonkeys[parts[1]], fns...)
		if err != nil {
			return 0, err
		}
		v := v1 - v2
		cacheMonkeys[m.name].number = &v
		return v, nil
	}
	panic(fmt.Sprintf("can not rev cal %+v", op))
}

func extractMonkey(raw string) ([2]string, rune) {
	opIndex := strings.IndexFunc(raw, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})
	return [2]string{strings.TrimSpace(raw[:opIndex]), strings.TrimSpace(raw[opIndex+1:])}, rune(raw[opIndex])
}

func handler(line string) error {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		return errors.New("invalid input")
	}
	name := parts[0]
	number, err := strconv.Atoi(parts[1])
	if err != nil {
		cacheMonkeys[name] = &monkey{name: name, op: parts[1]}
		return nil
	}
	cacheMonkeys[name] = &monkey{name: name, number: &number}
	return nil
}
