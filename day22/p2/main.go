package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

const (
	// 50 * 50 (per line 4 blocks)
	width         = 50
	height        = 50
	perLineBlocks = 4
)

type faceDirection struct {
	index int
	di    byte
	trans func(grid.Vec) (grid.Vec, bool)
}

// for my input
var neighborBoard = map[int]map[byte]faceDirection{
	1: {
		'R': {
			2, 'R', func(v grid.Vec) (grid.Vec, bool) {
				return v, false
			},
		},
		'L': {
			8, 'R', func(v grid.Vec) (grid.Vec, bool) {
				if v.X < width {
					v.X = 0
					v.Y = 2*height + (height - 1 - v.Y)
					return v, true
				}
				return v, false
			},
		},
		'U': {
			12, 'R', func(v grid.Vec) (grid.Vec, bool) {
				if v.Y < 0 {
					v.Y = v.X - width + 3*height
					v.X = 0
					return v, true
				}
				return v, false
			},
		},
		'D': {
			5, 'D', func(v grid.Vec) (grid.Vec, bool) {
				return v, false
			},
		},
	},
	2: {
		'R': {9, 'L', func(v grid.Vec) (grid.Vec, bool) {
			if v.X > 3*width-1 {
				v.X = 2*width - 1
				v.Y = 2*height + (height - 1 - v.Y)
				return v, true
			}
			return v, false
		}},
		'L': {1, 'D', func(v grid.Vec) (grid.Vec, bool) {
			return v, false
		}},
		'U': {12, 'U', func(v grid.Vec) (grid.Vec, bool) {
			if v.Y < 0 {
				v.Y = 4*height - 1
				v.X = v.X - width*2
				return v, true
			}
			return v, false
		}},
		'D': {5, 'L', func(v grid.Vec) (grid.Vec, bool) {
			if v.Y > height-1 {
				v.Y = height + v.X - 2*width
				v.X = 2*width - 1
				return v, true
			}
			return v, false
		}},
	},
	5: {
		'U': {1, 'U', func(v grid.Vec) (grid.Vec, bool) {
			return v, false
		}},
		'D': {9, 'D', func(v grid.Vec) (grid.Vec, bool) {
			return v, false
		}},
		'L': {8, 'D', func(v grid.Vec) (grid.Vec, bool) {
			if v.X < width {
				v.X = v.Y - height
				v.Y = 2 * height
				return v, true
			}
			return v, false
		}},
		'R': {
			2, 'U', func(v grid.Vec) (grid.Vec, bool) {
				if v.X > 2*width-1 {
					v.X = v.Y - height + 2*width
					v.Y = height - 1
					return v, true
				}
				return v, false
			},
		},
	},
	8: {
		'R': {
			9, 'R', func(v grid.Vec) (grid.Vec, bool) {
				return v, false
			},
		},
		'L': {1, 'R', func(v grid.Vec) (grid.Vec, bool) {
			if v.X < 0 {
				v.X = width
				v.Y = 3*height - 1 - v.Y
				return v, true
			}
			return v, false
		}},
		'U': {
			5, 'R', func(v grid.Vec) (grid.Vec, bool) {
				if v.Y < 2*height {
					v.Y = v.X + height
					v.X = width
					return v, true
				}
				return v, false
			},
		},
		'D': {
			12, 'D', func(v grid.Vec) (grid.Vec, bool) {
				return v, false
			},
		},
	},
	9: {
		'L': {8, 'L', func(v grid.Vec) (grid.Vec, bool) {
			return v, false
		}},
		'R': {2, 'L', func(v grid.Vec) (grid.Vec, bool) {
			if v.X > 2*width-1 {
				v.X = 3*width - 1
				v.Y = 3*height - 1 - v.Y
				return v, true
			}
			return v, false
		}},
		'U': {
			5, 'U', func(v grid.Vec) (grid.Vec, bool) {
				return v, false
			},
		},
		'D': {
			12, 'L', func(v grid.Vec) (grid.Vec, bool) {
				if v.Y > 3*height-1 {
					v.Y = 3*height + v.X - width
					v.X = width - 1
					return v, true
				}
				return v, false
			},
		},
	},
	12: {
		'U': {8, 'U', func(v grid.Vec) (grid.Vec, bool) {
			return v, false
		}},
		'R': {9, 'U', func(v grid.Vec) (grid.Vec, bool) {
			if v.X > width-1 {
				v.X = width + v.Y - 3*height
				v.Y = 3*height - 1
				return v, true
			}
			return v, false
		}},
		'L': {1, 'D', func(v grid.Vec) (grid.Vec, bool) {
			if v.X < 0 {
				v.X = v.Y - 3*height + width
				v.Y = 0
				return v, true
			}
			return v, false
		}},
		'D': {2, 'D', func(v grid.Vec) (grid.Vec, bool) {
			if v.Y > 4*height-1 {
				v.X = 2*width + v.X
				v.Y = 0
				return v, true
			}
			return v, false
		}},
	},
}

// test data
// var neighborBoard = map[int]map[byte]faceDirection{
// 	2: {
// 		'R': {
// 			11, 'L',
// 			func(v grid.Vec) (grid.Vec, bool) {
// 				if v.X > 3*width-1 {
// 					v.X = 4*width - 1
// 					v.Y = 3*height - 1 - v.Y
// 					return v, true
// 				}
// 				return v, false
// 			},
// 		},
// 		'L': {5, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.X < 2*width {
// 				v.X = (1*width - 1) + v.Y
// 				v.Y = 1*height - 1
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'U': {4, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y < 0 {
// 				v.X = v.X - (2*width - 1)
// 				v.Y = 1*height - 1
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'D': {6, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 	},
// 	4: {
// 		'R': {5, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'L': {11, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.X < 0 {
// 				v.X = 4*width - (v.Y - (1*height - 1))
// 				v.Y = 3*height - 1
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'U': {2, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y < height {
// 				v.Y = 0
// 				v.X = (2*width - 1) + v.X
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'D': {10, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y > 2*height-1 {
// 				v.Y = 3*height - 1
// 				v.X = v.X + (2*width - 1)
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 	},
// 	5: {
// 		'R': {6, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'L': {4, 'L', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'U': {2, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y < height {
// 				v.Y = v.X - width
// 				v.X = 2 * width
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'D': {10, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y > 2*height-1 {
// 				v.Y = 3*height - (v.X - (1*width - 1))
// 				v.X = 2 * width
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 	},
// 	6: {
// 		'R': {11, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.X > 3*width-1 {
// 				// 16 - (5- 3)
// 				v.X = 4*width - (v.Y - (height - 1))
// 				v.Y = 2 * height
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'L': {5, 'L', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'U': {2, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'D': {10, 'D', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 	},
// 	10: {
// 		'R': {11, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'L': {5, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.X < 2*width {
// 				v.X = 2*width - (v.Y - (2*height - 1))
// 				v.Y = 2*height - 1
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'U': {6, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'D': {4, 'U', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y > 3*height-1 {
// 				v.Y = 2*height - 1
// 				v.X = width - (v.X - (2*width - 1))
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 	},
// 	11: {
// 		'R': {2, 'L', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.X > 4*width {
// 				v.X = 3*width - 1
// 				v.Y = 3*height - v.Y
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'L': {10, 'L', func(v grid.Vec) (grid.Vec, bool) {
// 			return v, false
// 		}},
// 		'U': {6, 'L', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y < 2*height {
// 				v.Y = height + (3*width - 1 - v.X)
// 				v.X = 3*width - 1
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 		'D': {4, 'R', func(v grid.Vec) (grid.Vec, bool) {
// 			if v.Y > 3*height-1 {
// 				v.Y = 3*height - (v.X - (2*width - 1))
// 				v.X = 0
// 				return v, true
// 			}
// 			return v, false
// 		}},
// 	},
// }

type loc struct {
	grid.Vec
	face byte
}

func (l loc) String() string {
	return fmt.Sprintf("(%d,%d,%c)", l.X, l.Y, l.face)
}

func (l loc) Next() loc {
	index := l.Y/height*perLineBlocks + l.X/width
	fmt.Printf("at index: %d\n", index)
	vec := l.Add(direction(l.face, 1))
	if nb, ok := neighborBoard[index]; ok {
		if fd, ok := nb[l.face]; ok {
			di := l.face
			vec2, ok := fd.trans(vec)
			if ok {
				fmt.Printf("face: %c,trans: %s => %s\n", di, vec, vec2)
				di = fd.di
			}
			return loc{
				Vec:  vec2,
				face: di,
			}
		} else {
			panic("invalid face")
		}
	}
	return loc{
		Vec:  vec,
		face: l.face,
	}
}

func direction(face byte, step int) grid.Vec {
	switch face {
	case 'R':
		return grid.Vec{X: step}
	case 'L':
		return grid.Vec{X: -step}
	case 'U':
		return grid.Vec{Y: -step}
	case 'D':
		return grid.Vec{Y: step}
	}
	panic("invalid face")
}

func main() {
	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", handler)
	current := loc{}
	t, s := parsePath(path)
	for i := 0; i < len(t); i++ {
		current.face = turn(current.face, t[i])
		vec := current.Next()
		remStep := s[i]
		for remStep > 0 {
			fmt.Printf("next: %s\n", vec)
			if v, ok := board[grid.Vec{X: vec.X, Y: vec.Y}]; ok && v == '#' {
				break
			}
			if v, ok := board[grid.Vec{X: vec.X, Y: vec.Y}]; ok && v == '.' {
				current = vec
				remStep--
				fmt.Printf("current: %s,rem: %d\n", current, remStep)
			}
			vec = vec.Next()
		}
	}
	fmt.Fprintf(os.Stdout, "current: %s\n", current)
	fmt.Fprintf(os.Stdout, "password: %d\n", 1000*(current.Y+1)+4*(current.X+1)+faceNum(current.face))
}

func faceNum(face byte) int {
	switch face {
	case 'R':
		return 0
	case 'L':
		return 2
	case 'U':
		return 3
	case 'D':
		return 1
	}
	panic("invalid face")
}

var (
	board map[grid.Vec]byte = make(map[grid.Vec]byte)
	path  string
)

func handler(lines []string) error {
	if len(lines) == 0 {
		return nil
	}
	if len(lines) != 2 {
		panic("invalid input")
	}
	bl := strings.Split(lines[0], "\n")
	board = make(map[grid.Vec]byte)
	for i, line := range bl {
		for j, c := range line {
			board[grid.Vec{X: j, Y: i}] = byte(c)
		}
	}
	path = strings.TrimSuffix(lines[1], "\n")
	return nil
}

func parsePath(path string) ([]byte, []int) {
	di := []byte{'R'}
	var steps []int

	for len(path) > 0 {
		var step int
		for i, c := range path {
			if c >= '0' && c <= '9' {
				step = step*10 + int(c-'0')
				if i == len(path)-1 {
					path = path[i+1:]
					break
				}
			} else {
				path = path[i:]
				break
			}
		}
		steps = append(steps, step)
		if len(path) > 0 {
			di = append(di, path[0])
			path = path[1:]
		}
	}

	if len(di) != len(steps) {
		fmt.Printf("di: %v, steps: %v\n", di, steps)
		panic("invalid path")
	}

	return di, steps
}

func turn(raw byte, to byte) byte {
	switch raw {
	case 'R':
		if to == 'L' {
			return 'U'
		}
		return 'D'
	case 'L':
		if to == 'L' {
			return 'D'
		}
		return 'U'
	case 'U':
		if to == 'L' {
			return 'L'
		}
		return 'R'
	case 'D':
		if to == 'L' {
			return 'R'
		}
		return 'L'
	}
	return to
}
