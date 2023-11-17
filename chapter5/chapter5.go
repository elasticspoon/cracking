package chapter5

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
