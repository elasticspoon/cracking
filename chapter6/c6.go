package chapter6

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
