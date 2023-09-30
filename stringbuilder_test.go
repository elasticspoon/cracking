package ctci

import (
	"testing"
)

func TestStringBuilder(t *testing.T) {
	tests := []struct {
		additions []string
		expected  string
	}{
		{[]string{"a"}, "a"},
		{[]string{"a", "b"}, "ab"},
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "b", "c", "d", "e"}, "abcde"},
		{[]string{"hello", "world"}, "helloworld"},
	}
	for _, tt := range tests {
		sb := StringBuilder{make([]string, 0)}

		for _, a := range tt.additions {
			sb.Push(a)
		}

		actual := sb.String()
		if actual != tt.expected {
			t.Errorf("expected %s got %s", tt.expected, actual)
		}
	}
}
