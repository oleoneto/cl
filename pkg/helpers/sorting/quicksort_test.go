package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expectation []int
	}{
		{
			name:        "1234a",
			input:       []int{1, 2, 3, 4},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234b",
			input:       []int{2, 3, 4, 1},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234c",
			input:       []int{2, 1, 3, 4},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234d",
			input:       []int{4, 3, 2, 1},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "repeated tokens - 1",
			input:       []int{1, 9, 4, 2, 3, 5, 4},
			expectation: []int{1, 2, 3, 4, 4, 5, 9},
		},
		{
			name:        "repeated tokens - 2",
			input:       []int{1, 9, 4, 2, 9, 3, 5, 4},
			expectation: []int{1, 2, 3, 4, 4, 5, 9, 9},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectation, QuickSort(test.input))
	}
}
