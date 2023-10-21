package chapter3

import (
	"ctci/builtins"
	"strconv"
)

type Calculator struct {
	nums *builtins.Stack
	ops  *builtins.Stack
}

func NewCalculator() *Calculator {
	return &Calculator{
		nums: builtins.NewStack(),
		ops:  builtins.NewStack(),
	}
}

func (c *Calculator) execute(v1 float64, v2 float64, op byte) float64 {
	switch op {
	case '+':
		return v1 + v2
	case '-':
		return v1 - v2
	case '*':
		return v1 * v2
	case '/':
		return v1 / v2
	default:
		panic("Invalid operator")
	}
}

func (c *Calculator) Calculate(in string) float64 {
	var pos int

	for pos < len(in) {
		if !c.isOp(in[pos]) {
			pos = c.processInt(in, pos)
		} else {
			op := in[pos]
			c.ops.Push(in[pos])
			pos++

			if op == '*' || op == '/' {
				pos = c.processInt(in, pos)
				v2 := (c.nums.Pop()).(float64)
				v1 := (c.nums.Pop()).(float64)
				c.nums.Push(c.execute(v1, v2, op))
				c.ops.Pop()
			}
		}
	}

	for c.ops.Len() > 0 {
		v2 := (c.nums.Pop()).(float64)
		v1 := (c.nums.Pop()).(float64)
		op := (c.ops.Pop()).(byte)

		c.nums.Push(c.execute(v1, v2, op))
	}
	return (c.nums.Pop()).(float64)
}

func (c *Calculator) isOp(v byte) bool {
	return v == '+' || v == '-' || v == '*' || v == '/'
}

func (c *Calculator) processInt(in string, pos int) int {
	startPos := pos
	for pos < len(in) && !c.isOp(in[pos]) {
		pos++
	}

	i, err := strconv.Atoi(in[startPos:pos])
	if err != nil {
		panic(err)
	}
	c.nums.Push(float64(i))
	return pos
}
