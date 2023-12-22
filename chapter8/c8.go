package chapter8

import "fmt"

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

func multiplyWithoutSign(big, small int) int {
	if small > big {
		return multiplyWithoutSign(small, big)
	}

	switch small {
	case 0:
		return 0
	case 1:
		return big
	case 2:
		return big << 1
	case 3:
		return (big << 1) + big
	}

	if small&1 == 1 {
		return big + multiplyWithoutSign(big<<1, small>>1)
	}
	return multiplyWithoutSign(big<<1, small>>1)
}

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func paintFill(screen [][]int, point [2]int, color int) {
	x, y := point[0], point[1]
	currColor := screen[y][x]
	screen[y][x] = color

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(screen[0]) && newY >= 0 && newY < len(screen) && screen[newY][newX] == currColor {
			paintFill(screen, [2]int{newX, newY}, color)
		}
	}
}

func parenWays(s string, boolean bool) int {
	trueMap := make(map[string]int)
	falseMap := make(map[string]int)
	trueMap["1"] = 1
	falseMap["0"] = 1

	res := recParenWays(s, boolean, trueMap, falseMap)
	fmt.Println(trueMap, falseMap)
	return res
}

func recParenWays(s string, boolean bool, waysTrue map[string]int, waysFalse map[string]int) int {
	if boolean {
		if _, ok := waysTrue[s]; ok {
			return waysTrue[s]
		}
	} else {
		if _, ok := waysFalse[s]; ok {
			return waysFalse[s]
		}
	}

	ways := 0
	for i := 1; i < len(s); i += 2 {
		left := s[:i]
		right := s[i+1:]
		op := string(s[i])
		v := opSwitch(left, right, op, boolean, waysTrue, waysFalse)
		if boolean {
			waysTrue[s] = v
		} else {
			waysFalse[s] = v
		}
		ways += v
	}
	return ways
}

func opSwitch(left string, right string, op string, boolean bool, waysTrue map[string]int, waysFalse map[string]int) int {
	fmt.Println(left, right, op, boolean)
	switch op {
	case "^":
		if boolean {
			return recParenWays(left, false, waysTrue, waysFalse)*recParenWays(right, true, waysTrue, waysFalse) +
				recParenWays(left, true, waysTrue, waysFalse)*recParenWays(right, false, waysTrue, waysFalse)
		} else {
			return recParenWays(left, true, waysTrue, waysFalse)*recParenWays(right, true, waysTrue, waysFalse) +
				recParenWays(left, false, waysTrue, waysFalse)*recParenWays(right, false, waysTrue, waysFalse)
		}
	case "&":
		return opSwitch(left, right, "|", !boolean, waysTrue, waysFalse)
	case "|":
		if boolean {
			return recParenWays(left, true, waysTrue, waysFalse)*recParenWays(right, true, waysTrue, waysFalse) +
				recParenWays(left, true, waysTrue, waysFalse)*recParenWays(right, false, waysTrue, waysFalse) +
				recParenWays(left, false, waysTrue, waysFalse)*recParenWays(right, true, waysTrue, waysFalse)
		} else {
			return recParenWays(left, false, waysTrue, waysFalse) * recParenWays(right, false, waysTrue, waysFalse)
		}
	}
	panic("not implemented")
}
