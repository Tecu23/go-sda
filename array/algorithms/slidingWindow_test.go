package arrayalgorithms

import (
	"testing"

	assert "github.com/Tecu23/go-sda/internal/helpers"
)

func TestMaximumAverageSubarray(t *testing.T) {
	tests := []struct {
		id   string
		nums []int
		k    int
		want float64
	}{
		{
			id:   "1",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		}, {
			id:   "2",
			nums: []int{5},
			k:    1,
			want: 5.0,
		}, {
			id:   "3",
			nums: []int{10, 2, -15, -26, 5, 3, -2, 7, 4, -6, 20, 5},
			k:    6,
			want: 4.666666666666667,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			avg := maximumAverageSubarray(tt.nums, tt.k)

			assert.Equal(t, avg, tt.want)
		})
	}
}
