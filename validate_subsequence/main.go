package main

import "fmt"

//https://leetcode.com/problems/is-subsequence/

func main() {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	subArr := []int{1, 6, -1, 10}

	if IsValidSubsequence(arr, subArr) {
		fmt.Println("Valid Subsequence")
	} else {
		fmt.Println("Not a Valid subsequence")
	}
}

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
