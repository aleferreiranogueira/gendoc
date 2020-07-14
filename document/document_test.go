package document

import (
	"fmt"
	"testing"
)

func TestSeedBase(t *testing.T) {
	base := seedBase()

	if len(base) != 9 {
		t.Errorf("Base document number inbalid, got %v, wanted 9", len(base))
	}
}

func TestCalculateWeight(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			210,
		},
		{
			[]int{1, 1, 1, 4, 4, 4, 7, 7, 7},
			162,
		},
		{
			[]int{7, 2, 5, 6, 1, 6, 8, 9, 2},
			269,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			result := calculateWeight(tt.input)
			if result != tt.want {
				t.Errorf("Total sum for identifier invalid, got %v, wanted %v", result, tt.want)
			}
		})
	}
}
