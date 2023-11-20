package chapter6

import "testing"

func TestCalcLockers(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{100, 10},
		{1000, 31},
		{10000, 100},
	}
	for _, tt := range tests {
		if got := CalcLockers(tt.n); got != tt.want {
			t.Errorf("CalcLockers(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestCalcBGRatio(t *testing.T) {
	tests := []struct {
		families int
	}{
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
		{100000000},
	}

	for _, tt := range tests {
		got := calcBGRatio(tt.families)
		t.Logf("calcBGRatio(%v) = %v", tt.families, got)
	}
}
