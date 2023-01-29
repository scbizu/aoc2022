package math

func Max(ints ...int) int {
	max := ints[0]
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func TriangleSequence(size int) []int {
	tri := make([]int, size)
	for i := 0; i < size; i++ {
		tri[i] = i * (i + 1) / 2
	}
	return tri
}

func Power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
