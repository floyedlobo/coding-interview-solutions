package main

//https://leetcode.com/problems/is-subsequence/

//O(n)
func IsValidSubsequence(array []int, sequence []int) bool {

	subArrIdx := 0

	for _, num := range array {
		if sequence[subArrIdx] == num {
			subArrIdx++
		}
	}
	if subArrIdx == len(sequence) {
		return true
	} else {
		return false
	}
}
