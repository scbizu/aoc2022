package main

import (
	"container/list"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/queue"
)

var (
	m monkeys
	// 最小公倍数
	leastCommonMultiple int = 1
)

type monkeys struct {
	list []monkey
}

type monkey struct {
	inspectTime int
	data        *queue.Queue[int]
	op          func(n int)
}

func newMonkey(vals []int, op func(n int)) monkey {
	s := &queue.Queue[int]{
		List: list.New(),
	}
	s.PushN(vals...)
	return monkey{
		data: s,
		op:   op,
	}
}

func init() {
	mks := 8
	m.list = make([]monkey, mks)
}

var template = `Monkey (.*):
  Starting items: (.*)
  Operation: new = old (.*)
  Test: divisible by (.*)
    If true: throw to monkey (.*)
    If false: throw to monkey (.*)`

type monkeyTemplate struct {
	id            int
	startingItems string
	op            string
	div           int
	trueIndex     int
	falseIndex    int
}

func templateMonkey(mt monkeyTemplate, fn func(worryLevel int) int) monkey {
	items := strings.Split(mt.startingItems, ",")
	var vals []int
	for _, item := range items {
		var val int
		fmt.Sscanf(item, "%d", &val)
		vals = append(vals, val)
	}
	m := newMonkey(vals, func(old int) {
		var new int
		parts := strings.Split(mt.op, " ")
		var opval int
		if parts[1] == "old" {
			opval = old
		} else {
			fmt.Sscanf(parts[1], "%d", &opval)
		}
		switch parts[0] {
		case "*":
			new = (old * opval)
		case "+":
			new = (old + opval)
		case "/":
			new = (old / opval)
		case "-":
			new = (old - opval)
		}
		new = fn(new)
		if new%mt.div == 0 {
			m.list[mt.trueIndex].data.Push(new)
		} else {
			m.list[mt.falseIndex].data.Push(new)
		}
	})
	return m
}

func main() {
	if err := input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", func(blocks []string) error {
		for _, block := range blocks {
			mt := &monkeyTemplate{}
			// parse block to template use regexp
			re := regexp.MustCompile(template).FindAllStringSubmatch(block, -1)
			fmt.Sscanf(re[0][1], "%d", &mt.id)
			mt.startingItems = re[0][2]
			mt.op = re[0][3]
			fmt.Sscanf(re[0][4], "%d", &mt.div)
			fmt.Sscanf(re[0][5], "%d", &mt.trueIndex)
			fmt.Sscanf(re[0][6], "%d", &mt.falseIndex)
			m.list[mt.id] = templateMonkey(*mt, func(worryLevel int) int {
				return worryLevel / 3
			})
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var rounds int

	for {
		rounds++
		for i := 0; i < len(m.list); i++ {
			if m.list[i].data.Len() == 0 {
				continue
			}
			// pop all
			for m.list[i].data.Len() > 0 {
				m.list[i].op(m.list[i].data.Pop())
				m.list[i].inspectTime++
			}
		}
		if rounds == 20 {
			break
		}
	}
	for i := 0; i < len(m.list); i++ {
		println(m.list[i].data.String())
	}
	println("p1:")
	for i := 0; i < len(m.list); i++ {
		fmt.Printf("monkey %d, inspect: %d\n", i, m.list[i].inspectTime)
	}

	if err := input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", func(blocks []string) error {
		for _, block := range blocks {
			mt := &monkeyTemplate{}
			// parse block to template use regexp
			re := regexp.MustCompile(template).FindAllStringSubmatch(block, -1)
			fmt.Sscanf(re[0][1], "%d", &mt.id)
			mt.startingItems = re[0][2]
			mt.op = re[0][3]
			fmt.Sscanf(re[0][4], "%d", &mt.div)
			leastCommonMultiple *= mt.div
			fmt.Sscanf(re[0][5], "%d", &mt.trueIndex)
			fmt.Sscanf(re[0][6], "%d", &mt.falseIndex)
			m.list[mt.id] = templateMonkey(*mt, func(worryLevel int) int {
				return worryLevel % leastCommonMultiple
			})
		}
		return nil
	}); err != nil {
		panic(err)
	}

	rounds = 0

	for {
		rounds++
		for i := 0; i < len(m.list); i++ {
			if m.list[i].data.Len() == 0 {
				continue
			}
			// pop all
			for m.list[i].data.Len() > 0 {
				m.list[i].op(m.list[i].data.Pop())
				m.list[i].inspectTime++
			}
		}
		if rounds == 10000 {
			break
		}
	}

	println("p2:")
	maxTwo := [2]int{0, 0}
	for i := 0; i < len(m.list); i++ {
		if m.list[i].inspectTime > maxTwo[0] {
			maxTwo[1] = maxTwo[0]
			maxTwo[0] = m.list[i].inspectTime
		} else if m.list[i].inspectTime > maxTwo[1] {
			maxTwo[1] = m.list[i].inspectTime
		}
		fmt.Printf("monkey %d, inspect: %d\n", i, m.list[i].inspectTime)
	}
	fmt.Fprintf(os.Stdout, "p2: res: %d\n", maxTwo[0]*maxTwo[1])
}
