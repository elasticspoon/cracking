package chapter8

func numJumps(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}

	s1, s2, s3 := 1, 2, 4
	for i := 4; i <= n; i++ {
		temp := s1 + s2 + s3
		s3, s2, s1 = temp, s3, s2
	}

	return s3
}

func magicIndex(arr []int) int {
	for l, r := 0, len(arr)-1; l <= r; {
		m := (l + r) / 2
		if arr[m] == m {
			return m
		} else if arr[m] < m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func magicIndexDups(arr []int) int {
	return magicIndexDupsRec(arr, 0, len(arr)-1)
}

func magicIndexDupsRec(arr []int, l int, r int) int {
	if r < l {
		return -1
	}

	mid := (l + r) / 2

	if arr[mid] == mid {
		return mid
	}

	leftI := min(mid-1, arr[mid])
	if left := magicIndexDupsRec(arr, l, leftI); left >= 0 {
		return left
	}

	rightI := max(mid+1, arr[mid])
	if right := magicIndexDupsRec(arr, rightI, r); right >= 0 {
		return right
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func multiplyWithoutSign(a, b int) int {
	if b > a {
		return multiplyWithoutSign(b, a)
	}

	switch b {
	case 0:
		return 0
	case 1:
		return a
	case 2:
		return a << 1
	case 3:
		return (a << 1) + a
	default:
		left := b >> 1
		return multiplyWithoutSign(a<<1, left) + multiplyWithoutSign(a, b-(left<<1))
	}
}
