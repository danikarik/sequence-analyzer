package sequence_test

import (
	"testing"

	"github.com/danikarik/sequence"
)

func TestAnalyze(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    []int
		Expected sequence.Direction
	}{
		{
			Name:     "Increasing",
			Input:    []int{1, 2, 3, 4, 5},
			Expected: sequence.Increasing,
		},
		{
			Name:     "Decreasing",
			Input:    []int{5, 4, 3, 2, 1, 0, -1, -2, -3},
			Expected: sequence.Decreasing,
		},
		{
			Name:     "Stable",
			Input:    []int{1, 2, 3, 2, 1, 4, 5},
			Expected: sequence.Stable,
		},
		{
			Name:     "Empty",
			Input:    []int{},
			Expected: sequence.Stable,
		},
		{
			Name:     "Single",
			Input:    []int{},
			Expected: sequence.Stable,
		},
		{
			Name:     "EqualWhenIncreasing",
			Input:    []int{1, 2, 2},
			Expected: sequence.Stable,
		},
		{
			Name:     "EqualWhenDecreasing",
			Input:    []int{5, 4, 4},
			Expected: sequence.Stable,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := sequence.Analyze(tc.Input)
			if tc.Expected != res {
				t.Errorf("expected %s got %s", tc.Expected, res)
			}
		})
	}
}
