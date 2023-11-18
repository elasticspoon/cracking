package chapter5

import "fmt"

func InsertBitsBetween(n, m uint32, i, j int) uint32 {
	var mask uint32
	mask = (1 << (j + 1)) - 1
	mask &= ^((1 << i) - 1)
	n &= ^mask
	return n | (m << i)
}

func FloatDecToBinary(n float64) string {
	if n >= 1 || n < 0 {
		return "ERROR"
	}
	var result uint32

	for i := 1; n > 0 && i < 33; i++ {
		tmp := 1 << i
		if v := 1 / float64(tmp); n >= v {
			result = result<<1 + 1
			n -= v
		} else {
			result <<= 1
		}
	}

	if n != 0 {
		return "ERROR"
	} else {
		return fmt.Sprintf("0.%b", result)
	}
}

func FlipBit(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var curr, prev, maxLen int
	prev = -1

	for n > 0 {
		if n%2 == 1 {
			curr++
			maxLen = max(maxLen, curr)
		} else {
			prev = curr
			curr = 0
		}

		if n >= 0 {
			maxLen = max(maxLen, curr+prev+1)
		}

		n /= 2
	}

	return maxLen
}

func nextLarger(n int) int {
	length := 0
	numBits := 0
	neededBitPos := -1

	for i := n; i > 0; i /= 2 {
		if i&1 == 1 {
			numBits++
		}
		if (i^0b10)&0b11 == 0b11 && neededBitPos == -1 {
			neededBitPos = length
		}
		length++
	}

	if numBits == 0 {
		return -1
	}

	if neededBitPos == length-1 {
		return 1<<length + 1<<(numBits-1) - 1
	}

	return n - (1 << neededBitPos) + (1 << (neededBitPos + 1))
}

func nextSmaller(n int) int {
	length := 0

	for i := n; i > 0; i /= 2 {
		length++
		if (i^0b01)&0b11 == 0b11 {
			return n - (1 << length) + (1 << (length - 1))
		}
	}

	return -1
}

func bookNextLarger(n int) int {
	var c0, c1 int
	c := n

	for (c&1) == 0 && (c != 0) {
		c0++
		c >>= 1
	}

	for (c & 1) == 1 {
		c1++
		c >>= 1
	}

	if c0+c1 == 31 || c0+c1 == 0 {
		return -1
	}

	p := c0 + c1
	n |= (1 << p)
	n &= ^((1 << p) - 1)
	n |= (1 << (c1 - 1)) - 1
	return n
}

func BitsToFlip(a, b int) int {
	count := 0
	for c := a ^ b; c != 0; c >>= 1 {
		count += c & 1
	}
	return count
}

func SwapOddEvenBits(n int) int {
	evenMask := 0
	for i := 0; i < 63; i += 2 {
		evenMask |= 1 << i
	}
	oddMask := evenMask << 1

	evens := n & evenMask
	odds := n & oddMask

	return (evens << 1) | (odds >> 1)
}
