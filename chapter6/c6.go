package chapter6

import "math/rand"

func CalcLockers(n int) int {
	lockers := make([]bool, n)
	for i := 1; i <= n; i++ {
		for j := i - 1; j < len(lockers); j += i {
			lockers[j] = !lockers[j]
		}
	}
	count := 0
	for _, v := range lockers {
		if v {
			count++
		}
	}
	return count
}

func calcBGRatio(families int) float64 {
	var boys, girls int

	for i := 0; i < families; i++ {
		for {
			if rand.Intn(2) == 0 {
				girls++
				break
			}
			boys++
		}
	}

	return float64(boys) / float64(girls)
}
