package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/math"
)

var blueprintPattern = `Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.`

type blueprint struct {
	index             int
	oreCostOre        int
	clayCostOre       int
	obsidianCostOre   int
	obsidianCostClay  int
	geodeCostOre      int
	geodeCostObsidian int
}

type bot int

const (
	oreBot bot = iota
	clayBot
	obsidianBot
	geodeBot
)

var blueprints []blueprint

type state struct {
	bp            blueprint
	remMinutes    int
	ore           int
	clay          int
	obsidian      int
	geode         int
	oreRobot      int
	clayRobot     int
	obsidianRobot int
	geodeRobot    int
	wantMake      bot
}

var maxGeodeCollection = make(map[int]int)

func decide(lastState state) {
	if lastState.remMinutes == 0 {
		maxGeodeCollection[lastState.bp.index] = math.Max(maxGeodeCollection[lastState.bp.index], lastState.geode)
		return
	}
	// robot overflow:
	// if the ore robot is more than the max ore cost of the current blueprint, then we should prune the branch
	if lastState.wantMake == oreBot &&
		lastState.oreRobot >= math.Max(lastState.bp.oreCostOre, lastState.bp.clayCostOre,
			lastState.bp.obsidianCostOre, lastState.bp.geodeCostOre) {
		return
	}
	// if the clay robot is more than the max clay cost of the current blueprint , then we should prune the branch
	if lastState.wantMake == clayBot &&
		lastState.clayRobot >= lastState.bp.obsidianCostClay {
		return
	}
	// if the obsidian robot is more than the max obsidian cost of the current blueprint, then we should prune the branch
	if lastState.wantMake == obsidianBot {
		if lastState.obsidianRobot >= lastState.bp.geodeCostObsidian {
			return
		}
		// the necessary resource is not claimed
		if lastState.clay == 0 {
			return
		}

	}
	// the necessary resource is not claimed
	if lastState.wantMake == geodeBot {
		if lastState.obsidian == 0 {
			return
		}
	}
	// assume that we can finally collect one(or more than one which depends on the robot numbers) geode per minute
	// then the total could collected geode is : geode + geodeRobot * remMinutes + triangle(remMinutes)
	//                                                                             <- if we can craft as much as possible geode robot before the deadline comes
	if lastState.geode+lastState.geodeRobot*lastState.remMinutes+math.TriangleSequence(33)[lastState.remMinutes] <= maxGeodeCollection[lastState.bp.index] {
		return
	}
	// fmt.Printf("state: %+v\n", lastState)
	// crafting:
	// always craft ore robot if possible
	if lastState.ore >= lastState.bp.oreCostOre && lastState.wantMake == oreBot {
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            lastState.bp,
				remMinutes:    lastState.remMinutes - 1,
				ore:           lastState.ore + lastState.oreRobot - lastState.bp.oreCostOre,
				clay:          lastState.clay + lastState.clayRobot,
				obsidian:      lastState.obsidian + lastState.obsidianRobot,
				geode:         lastState.geode + lastState.geodeRobot,
				oreRobot:      lastState.oreRobot + 1,
				clayRobot:     lastState.clayRobot,
				obsidianRobot: lastState.obsidianRobot,
				geodeRobot:    lastState.geodeRobot,
				wantMake:      bot(i),
			})
		}
		return
	}
	// always craft clay robot if possible
	if lastState.ore >= lastState.bp.clayCostOre && lastState.wantMake == clayBot {
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            lastState.bp,
				remMinutes:    lastState.remMinutes - 1,
				ore:           lastState.ore + lastState.oreRobot - lastState.bp.clayCostOre,
				clay:          lastState.clay + lastState.clayRobot,
				obsidian:      lastState.obsidian + lastState.obsidianRobot,
				geode:         lastState.geode + lastState.geodeRobot,
				oreRobot:      lastState.oreRobot,
				clayRobot:     lastState.clayRobot + 1,
				obsidianRobot: lastState.obsidianRobot,
				geodeRobot:    lastState.geodeRobot,
				wantMake:      bot(i),
			})
		}
		return
	}
	// always craft obsidian robot if possible
	if lastState.ore >= lastState.bp.obsidianCostOre &&
		lastState.clay >= lastState.bp.obsidianCostClay &&
		lastState.wantMake == obsidianBot {
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            lastState.bp,
				remMinutes:    lastState.remMinutes - 1,
				ore:           lastState.ore + lastState.oreRobot - lastState.bp.obsidianCostOre,
				clay:          lastState.clay + lastState.clayRobot - lastState.bp.obsidianCostClay,
				obsidian:      lastState.obsidian + lastState.obsidianRobot,
				geode:         lastState.geode + lastState.geodeRobot,
				oreRobot:      lastState.oreRobot,
				clayRobot:     lastState.clayRobot,
				obsidianRobot: lastState.obsidianRobot + 1,
				geodeRobot:    lastState.geodeRobot,
				wantMake:      bot(i),
			})
		}
		return
	}
	// always craft geode robot if possible
	if lastState.ore >= lastState.bp.geodeCostOre &&
		lastState.obsidian >= lastState.bp.geodeCostObsidian &&
		lastState.wantMake == geodeBot {
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            lastState.bp,
				remMinutes:    lastState.remMinutes - 1,
				ore:           lastState.ore + lastState.oreRobot - lastState.bp.geodeCostOre,
				clay:          lastState.clay + lastState.clayRobot,
				obsidian:      lastState.obsidian + lastState.obsidianRobot - lastState.bp.geodeCostObsidian,
				geode:         lastState.geode + lastState.geodeRobot,
				oreRobot:      lastState.oreRobot,
				clayRobot:     lastState.clayRobot,
				obsidianRobot: lastState.obsidianRobot,
				geodeRobot:    lastState.geodeRobot + 1,
				wantMake:      bot(i),
			})
		}
		return
	}
	// do nothing , just always collect
	// and till we can craft the wantMake robot
	newState := state{
		bp:            lastState.bp,
		remMinutes:    lastState.remMinutes - 1,
		ore:           lastState.ore + lastState.oreRobot,
		clay:          lastState.clay + lastState.clayRobot,
		obsidian:      lastState.obsidian + lastState.obsidianRobot,
		geode:         lastState.geode + lastState.geodeRobot,
		oreRobot:      lastState.oreRobot,
		clayRobot:     lastState.clayRobot,
		obsidianRobot: lastState.obsidianRobot,
		geodeRobot:    lastState.geodeRobot,
		wantMake:      lastState.wantMake,
	}
	decide(newState)
}

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	var sum int
	for _, bp := range blueprints {
		maxGeodeCollection[bp.index] = 0
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            bp,
				remMinutes:    24,
				ore:           0,
				clay:          0,
				obsidian:      0,
				geode:         0,
				oreRobot:      1,
				clayRobot:     0,
				obsidianRobot: 0,
				geodeRobot:    0,
				wantMake:      bot(i),
			})
		}
		sum += int(maxGeodeCollection[bp.index] * bp.index)
	}
	fmt.Fprintf(os.Stdout, "p1: sum: %d\n", sum)

	multi := 1
	maxGeodeCollection = make(map[int]int)
	for _, bp := range blueprints[:3] {
		for i := 0; i < 4; i++ {
			decide(state{
				bp:            bp,
				remMinutes:    32,
				ore:           0,
				clay:          0,
				obsidian:      0,
				geode:         0,
				oreRobot:      1,
				clayRobot:     0,
				obsidianRobot: 0,
				geodeRobot:    0,
				wantMake:      bot(i),
			})
		}
		multi *= maxGeodeCollection[bp.index]
	}
	fmt.Fprintf(os.Stdout, "p2: multi: %d\n", multi)
}

func handler(line string) error {
	var bp blueprint
	if _, err := fmt.Sscanf(line, blueprintPattern,
		&bp.index, &bp.oreCostOre, &bp.clayCostOre, &bp.obsidianCostOre,
		&bp.obsidianCostClay, &bp.geodeCostOre, &bp.geodeCostObsidian,
	); err != nil {
		return err
	}
	blueprints = append(blueprints, bp)
	return nil
}
