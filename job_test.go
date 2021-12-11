package queue

import (
	"fmt"
	"testing"
)

func TestJobResult(t *testing.T) {
	tests := []struct {
		actual   JobResult
		expected string
	}{
		{
			actual:   Fail,
			expected: "Fail",
		},
		{
			actual:   Success,
			expected: "Success",
		},
		{
			actual:   Skip,
			expected: "Skip",
		},
		{
			actual:   100,
			expected: "JobResult(100)",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("JobResult-%d", i), func(t *testing.T) {
			actual := test.actual.String()
			if actual != test.expected {
				t.Errorf("expected string of %s, got %s", test.expected, actual)
			}
		})
	}
}
