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
