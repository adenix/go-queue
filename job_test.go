package queue

import (
	"fmt"
	"testing"
)

func TestJobResult(t *testing.T) {
	cases := []struct {
		result   JobResult
		expected string
	}{
		{
			result:   Fail,
			expected: "fail",
		},
		{
			result:   Success,
			expected: "success",
		},
		{
			result:   Skip,
			expected: "skip",
		},
		{
			result:   100,
			expected: "unknown",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("JobResult-%d", i), func(t *testing.T) {
			actual := c.result.String()
			if actual != c.expected {
				t.Errorf("expected string of %s, got %s", c.expected, actual)
			}
		})
	}
}
