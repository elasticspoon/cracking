package ctci

import "strings"

type StringBuilder struct {
	content []string
}

func (sb *StringBuilder) Push(s string) {
	sb.content = append(sb.content, s)
}

func (sb *StringBuilder) String() string {
	return strings.Join(sb.content, "")
}
