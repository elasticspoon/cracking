package ctci

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"abc", "abc", true},
		{"abc", "cba", true},
		{"abc", "bca", true},
		{"abc", "cab", true},
		{"abc", "bac", true},
		{"abc", "acb", true},
		{"abc", "xyz", false},
		{"abc", "xy", false},
	}

	for _, c := range cases {
		val := IsPermutation(c.s1, c.s2)
		if val != c.want {
			t.Errorf("IsPermutation(%q, %q) == %t, want %t", c.s1, c.s2, val, c.want)
		}
	}
}

func TestIsPermutationPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abba", true},
		{"abcba", true},
		{"ab", false},
		{"abc", false},
		{"abab", true},
		{"ababc", true},
		{"taco cat", true},
		{"ra ce car", true},
	}

	for _, c := range cases {
		val := IsPermutationPalindrome(c.s)
		if val != c.want {
			t.Errorf("IsPermutationPalindrome(%q, %q) == %t, want %t", c.s, c.s, val, c.want)
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	cases := []struct {
		in, want [][]int
	}{
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 0, 1, 2}, {3, 4, 5, 6}}, [][]int{{3, 9, 5, 1}, {4, 0, 6, 2}, {5, 1, 7, 3}, {6, 2, 8, 4}}},
	}

	for _, c := range cases {
		val := RotateMatrix(c.in)
		if !compareMatrix(val, c.want) {
			t.Errorf("RotateMatrix(%v) == %v, want %v", c.in, val, c.want)
		}
	}
}

func compareMatrix(a, b [][]int) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
