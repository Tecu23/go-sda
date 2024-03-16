// Package array contains common algorithms on arrays
package array

import (
	"container/heap"
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

/*
  - Return the minimal length of a subarray whose sum is greater than or equal to target.
    return 0 if not found.

-   @params {int} target -> the target we are searching for
-   @params {int[]} nums -> the array we are searching in

-   @return {int} -> the count of all unique characters in all substrings
*/
func minimumSizeSubarraySum(target int, nums []int) int {
	currLen, minLen := 0, len(nums)
	found, sum := false, 0
	i, j := 0, 0

	for j < len(nums) || sum >= target {
		if sum < target {
			sum += nums[j]
			j++
			currLen++
		} else {
			if currLen <= minLen {
				found = true
				minLen = currLen
			}
			currLen--
			sum -= nums[i]
			i++
		}
	}

	if !found {
		return 0
	}
	return minLen
}

/*
  - Return the length of the longest substring containing the same letter you can after replacing
    any letter k amount of times

    ex: ABACABA , k = 2 => 5, because AB AAAAA

-   @params {string} s -> the string we are searching in
-   @params {int} k -> the numer of replacements we are allowed to make

-   @return {int} -> the length of the substring
*/
func longestRepeatingCharacter(s string, k int) int {
	count := make(map[byte]int)
	res, maxCount, start := 0, 0, 0

	for end := range s {
		count[s[end]]++

		if maxCount < count[s[end]] {
			maxCount = count[s[end]]
		}

		if end-start+1-maxCount > k {
			count[s[start]]--
			start++
		}

		if res < end-start+1 {
			res = end - start + 1
		}
	}
	return res
}

/*
  - Given a string, find the length of the longest substring without repeating characters

-   @params {string} s -> the string we are searching in

-   @return {int} -> the length of the substring
*/
func longestSubstringWithoutRepeatingCharacters(s string) int {
	count := make(map[byte]int)

	start, end := 0, 0
	maxC, currC := 0, 0

	for end < len(s) {
		val := count[s[end]]
		if val == 0 {
			count[s[end]]++
			currC++
			end++
		} else {
			for count[s[end]] > 0 {
				count[s[start]]--
				currC--
				start++
			}
		}

		if currC > maxC {
			maxC = currC
		}

	}

	return maxC
}

/*
  - You are given an array of integers nums, there is a sliding window of size k which

is moving from the very left of the array to the very right. You can only see the k numbers
in the window. Each time the sliding window moves right by one position.

-   @params {[]int} nums -> the arrays we are using the window for
-   @params {int} k -> the size of the sliding window

-   @return {[]int} -> the max sliding windows array
*/
type pair struct {
	index, value int
}

type pairHeap []pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}

func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func slidingWindowMaximum(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)

	i, j, idx := 0, 0, 0
	pq := &pairHeap{}
	heap.Init(pq)

	for j < len(nums) {
		pair := pair{j, nums[j]}
		heap.Push(pq, pair)
		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			for (*pq)[0].index < i {
				heap.Pop(pq)
			}
			first := (*pq)[0]
			res[idx] = first.value
			if first.index < i+1 {
				heap.Pop(pq)
			}
			i++
			j++
			idx++
		}
	}

	return res
}

/*
  - Given two strings s and t of lengths m and n respectively, return the minimum window substring

    of s such that every character in t (including duplicates) is included in the window.
    If there is no such substring, return the empty string "".

-   @params {string} s -> the string in which we are searching for substrings
-   @params {string} t -> the string which characters have to be in the substring

-   @return {string} -> the substring of s which contains every letter of t
*/
func minimumWindowSubstring(s, t string) string {
	if len(t) > len(s) {
		return ""
	}

	count := make([]int, 60)

	for c := range t {
		count[t[c]-65]++
	}

	i, j := 0, 0
	countChar := len(t)
	minLen, startIndex := int(^uint(0)>>1), 0

	for j < len(s) {
		if count[s[j]-65] > 0 {
			countChar--
		}

		count[s[j]-65]--
		j++

		for countChar == 0 {
			if (j - i) < minLen {
				startIndex = i
				minLen = j - i
			}

			if count[s[i]-65] == 0 {
				countChar++
			}

			count[s[i]-65]++
			i++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
