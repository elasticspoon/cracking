package ctci

func IsPermutation(s1, s2 string) bool {
	sMap := make(map[byte]int)

	if len(s1) != len(s2) {
		return false
	} else if s1 == s2 {
		return true
	}

	for i := 0; i < len(s1); i++ {
		sMap[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		if val, ok := sMap[s2[i]]; ok {
			if val == 0 {
				return false
			} else {
				sMap[s2[i]]--
			}
		} else {
			return false
		}
	}
	return true
}

func IsPermutationPalindrome(s string) bool {
	hash := make(map[byte]int)
	numOdd := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		hash[s[i]]++
		if hash[s[i]]%2 == 1 {
			numOdd++
		} else {
			numOdd--
		}
	}

	return numOdd <= 1
}

func RotateMatrix(array [][]int) [][]int {
	n := len(array)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			rotato(&array, i, j)
		}
	}
	return array
}

func rotato(array *[][]int, x int, y int) {
	arr := *array
	n := len(arr) - 1
	nx, ny := n-x, n-y

	temp := arr[y][nx]
	arr[y][nx] = arr[x][y]
	arr[nx][ny], temp = temp, arr[nx][ny]
	arr[ny][x], temp = temp, arr[ny][x]
	arr[x][y] = temp
}
