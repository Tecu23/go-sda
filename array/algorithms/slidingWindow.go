// Package arrayalgorithms contains common algorithms on arrays
package arrayalgorithms

/*
  - Find a slice whose length is equal to k and has the maximum average value

-   @params {[]int} nums -> the array we are seaching in
-   @params {int} k -> the length of a continuous subarray with the maximum average

-   @return {float64} -> the maximum average
*/
func maximumAverageSubarray(nums []int, k int) float64 {
	maxSum := 0
	l := len(nums)

	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	window := maxSum

	start := k

	for start < l {
		window += nums[start] - nums[start-k]
		if window > maxSum {
			maxSum = window
		}
		start++

	}

	return float64(maxSum) / float64(k)
}
