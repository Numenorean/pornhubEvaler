package pornhubEvaler

import (
	"math"
	"strconv"
)

func leastFactor(n int) int {
	if n == 0 {
		return 0
	}
	if n%1 == 1 {
		return 1
	}
	if n%2 == 0 {
		return 2
	}
	if n%3 == 0 {
		return 3
	}
	if n%5 == 0 {
		return 5
	}
	m := int(math.Sqrt(float64(n)))
	for i := 7; i <= m; i += 30 {
		if n%i == 0 {
			return i
		}
		if n%(i+4) == 0 {
			return i + 4
		}
		if n%(i+6) == 0 {
			return i + 6
		}
		if n%(i+10) == 0 {
			return i + 10
		}
		if n%(i+12) == 0 {
			return i + 12
		}
		if n%(i+16) == 0 {
			return i + 16
		}
		if n%(i+22) == 0 {
			return i + 22
		}
		if n%(i+24) == 0 {
			return i + 24
		}
	}
	return n
}

func sliceStringToUint8(input [][]string) []uint8 {
	result := make([]uint8, 0, len(input))
	for _, a := range input {
		i, _ := strconv.Atoi(a[1])
		result = append(result, uint8(i))
	}
	return result
}

func sliceStringToInt(input [][]string) [][]int {
	result := make([][]int, 0, len(input))
	for _, a := range input {
		i, _ := strconv.Atoi(a[1])
		i2, _ := strconv.Atoi(a[2])
		result = append(result, []int{i, i2})
	}
	return result
}
