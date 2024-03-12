package array

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

func TestFindAllAnagramsInAString(t *testing.T) {
	tests := []struct {
		id   string
		s    string
		p    string
		want []int
	}{
		{
			id:   "1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		}, {
			id:   "2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			result := findAllAnagramsInAString(tt.s, tt.p)

			assert.EqualSlices(t, result, tt.want)
		})
	}
}

func TestPermutationInString(t *testing.T) {
	tests := []struct {
		id   string
		s1   string
		s2   string
		want bool
	}{
		{
			id:   "1",
			s2:   "eidbaooo",
			s1:   "ab",
			want: true,
		}, {
			id:   "2",
			s2:   "eidboaoo",
			s1:   "ab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			result := permutationInString(tt.s1, tt.s2)
			assert.Equal(t, result, tt.want)
		})
	}
}

func TestCountUniqueCharactersOfAllSubstringsOfAGivenString(t *testing.T) {
	tests := []struct {
		id   string
		s    string
		want int
	}{
		{
			id:   "1",
			s:    "ABC",
			want: 10,
		}, {
			id:   "1",
			s:    "ABA",
			want: 8,
		}, {
			id:   "3",
			s:    "LEETCODE",
			want: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			result := countUniqueCharactersOfAllSubstringsOfAGivenString(tt.s)
			assert.Equal(t, result, tt.want)
		})
	}
}

func TestMinimumSizeSubarraySum(t *testing.T) {
	tests := []struct {
		id     string
		target int
		nums   []int
		want   int
	}{
		{
			id:     "1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		}, {
			id:     "2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		}, {
			id:     "3",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1},
			want:   0,
		}, {
			id:     "4",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			result := minimumSizeSubarraySum(tt.target, tt.nums)
			assert.Equal(t, result, tt.want)
		})
	}
}
