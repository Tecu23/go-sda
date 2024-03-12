// Package arrayalgorithms contains common algorithms on arrays
package array

import (
	"fmt"
)

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

/*
  - Return an array of all the start indices of p's anagrams in s. (Can return in any order)

-   @params {string} s -> the string we are searching anagrams in
-   @params {string} p -> the string we are searching for

-   @return {int[]} -> the positions of all the anagrams found in s
*/
func findAllAnagramsInAString(s string, p string) []int {
	result := make([]int, 0, 20)

	if len(s) < len(p) {
		return result
	}

	freqS := make([]int, 26)
	freqP := make([]int, 26)

	for i := 0; i < len(p); i++ {
		freqS[int(s[i]-'a')]++
		freqP[int(p[i]-'a')]++
	}

	start := 0
	end := len(p)

	if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
		result = append(result, start)
	}

	for end < len(s) {
		freqS[int(s[start]-'a')]--
		freqS[int(s[end]-'a')]++

		if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
			result = append(result, start+1)
		}

		start++
		end++
	}

	return result
}

/*
- Find whether string s2 contains a permutation of s1

-   @params {string} s1 -> the string which permutations we are searching for
-   @params {string} s2 -> the string we are searching in

-   @return {bool} -> whether s2 contains a permutation of s1
*/
func permutationInString(s1, s2 string) bool {
	l, freqS1 := 0, [26]int{}
	for i := range s1 {
		freqS1[s1[i]-97]++
	}

	for r := range s2 {
		freqS1[s2[r]-97]--
		if freqS1 == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			freqS1[s2[l]-97]++
			l++
		}
	}
	return false
}

/*
- Find all the unique chracters in all the substrings of a string

-   @params {string} s -> the string we are counting in

-   @return {int} -> the count of all unique characters in all substrings
*/
func countUniqueCharactersOfAllSubstringsOfAGivenString(s string) int {
	/*

	   Ex. LEETCODE

	   The final "pos" array will be for chars "A", "B", "C", .....
	   [
	       [-1 8] [-1 8] [-1 4 8] [-1 6 8] [-1 1 2 7 8] [-1 8]
	       [-1 8] [-1 8] [-1 8] [-1 8] [-1 8] [-1 0 8] [-1 8]
	       [-1 8] [-1 5 8] [-1 8] [-1 8] [-1 8] [-1 8] [-1 3 8]
	       [-1 8] [-1 8] [-1 8] [-1 8] [-1 8] [-1 8]
	   ]

	   so lets take C for example: [-1 4 8] -> C is at position 4 in "LEETCODE"

	   - the unique contributions of C will be all the substring that only contain one C and
	   it will be the cross product of the number of substrings that end in C and the number
	   of subtrings that start with C:

	       L E E T C           C O D E
	         E E T C           C O D
	           E T C           C O
	             T C           C
	               C

	   - because all substring that end with C could be appended to the substring that start with C
	    and vice-versa

	    => The C adds to the sum 4 * 5 => 20
	        L E E T C O D E
	        L E E T C O D
	        L E E T C O
	        L E E T C
	          E E T C O D E
	          E E T C O D
	          E E T C O
	          E E T C
	            E T C O D E
	            E T C O D
	            E T C O
	            E T C
	              T C O D E
	              T C O D
	              T C O
	              T C
	                C O D E
	                C O D
	                C O
	                C
	*/
	sum := 0
	n := len(s)
	pos := make([][]int, 26)

	for i := 0; i < 26; i++ {
		pos[i] = append(pos[i], -1)
	}

	for i := range s {
		pos[s[i]-'A'] = append(pos[s[i]-'A'], i)
	}

	for i := 0; i < 26; i++ {
		pos[i] = append(pos[i], n)
	}

	for i := 0; i < 26; i++ {
		for k := 1; k < len(pos[i])-1; k++ {
			sum += (pos[i][k] - pos[i][k-1]) * (pos[i][k+1] - pos[i][k])
		}
	}

	return sum
}
