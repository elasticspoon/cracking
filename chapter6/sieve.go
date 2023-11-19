package chapter6

import (
	"fmt"
	"math"
)

func ithPrime(n int) int {
	notPrime := make([]bool, 1000000)
	max := int(math.Sqrt(float64(len(notPrime))))

	if n > max {
		panic(fmt.Sprintf("this function is limited to primes < %v", max))
	}

	for i := 2; i < int(math.Sqrt(float64(len(notPrime)))); {
		for j := i + 1; j < len(notPrime); j++ {
			// for j := i + 1; j < len(notPrime); j++ {
			if !notPrime[j] && j%i == 0 {
				notPrime[j] = true
			}
		}
		i++
		for i < len(notPrime) && notPrime[i] {
			i++
		}
	}

	for i, count := 2, 0; i < len(notPrime); i++ {
		if !notPrime[i] {
			count++
			if count == n {
				return i
			}
		}
	}

	panic("should never get here")
}
