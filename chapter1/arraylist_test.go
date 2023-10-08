package ctci

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	tests := []struct {
		additions   []int
		expectedArr []int
		expectedLen int
		expectedCap int
	}{
		{[]int{1}, []int{1}, 1, 1},
		{[]int{1, 2}, []int{1, 2}, 2, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3, 4},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5, 8},
	}

	for _, tt := range tests {
		arrayList := NewArrayList[int]()

		for _, v := range tt.additions {
			arrayList.Append(v)
		}
		if arrayList.Length != tt.expectedLen {
			t.Errorf("expected length to be %d got %d", tt.expectedLen, arrayList.Length)
		}
		if arrayList.Capacity != tt.expectedCap {
			t.Errorf("expected length to be %d got %d", tt.expectedCap, arrayList.Capacity)
		}
		for i, v := range tt.expectedArr {
			if arrayList.At(i) != v {
				t.Errorf("expected arr[%d] to be %d got %d", i, v, arrayList.At(i))
			}
		}
	}
}
