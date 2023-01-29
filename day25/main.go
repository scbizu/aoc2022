package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/math"
)

var snafus []string

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	var total int
	for _, snafu := range snafus {
		total += convertSnafu(snafu)
	}
	finalSnafu := convertToSnafu(total)
	fmt.Fprintf(os.Stdout, "p1: snafu: %s\n", finalSnafu)
}

func convertToSnafu(number int) string {
	var result string
	for number > 0 {
		digit := number % 5
		var debt int
		switch digit {
		case 0:
			result = "0" + result
		case 1:
			result = "1" + result
		case 2:
			result = "2" + result
		case 3:
			result = "=" + result
			debt += 2
		case 4:
			result = "-" + result
			debt += 1
		}
		number = (number + debt) / 5
	}
	return result
}

func convertSnafu(snafu string) int {
	var result int
	for index, char := range snafu {
		var digit int
		switch char {
		case '-':
			digit = -1
		case '=':
			digit = -2
		default:
			digit = int(char - '0')
		}
		result += digit * math.Power(5, (len(snafu)-index-1))
	}
	return result
}

func handler(line string) error {
	snafus = append(snafus, line)
	return nil
}
